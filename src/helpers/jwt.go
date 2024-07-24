package helpers

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/ssamsara98/photopost-golang/src/constants"
	"github.com/ssamsara98/photopost-golang/src/lib"
	"github.com/ssamsara98/photopost-golang/src/models"
)

// JWTAuth service relating to authorization
type JWTAuth struct {
	env    *lib.Env
	logger *lib.Logger
}

type Claims struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Type     string `json:"type"`
	jwt.RegisteredClaims
}

func NewJWTAuth(
	env *lib.Env,
	logger *lib.Logger,

) *JWTAuth {
	return &JWTAuth{
		env,
		logger,
	}
}

// CreateToken creates jwt auth token
func (j JWTAuth) CreateToken(user *models.User, tokenType string) (string, error) {
	var secret string
	var duration time.Duration
	if tokenType == constants.TokenAccess {
		secret = j.env.JWTAccessSecret
		duration = j.env.AccessTokenDuration
	} else if tokenType == constants.TokenRefresh {
		secret = j.env.JWTRefreshSecret
		duration = j.env.RefreshTokenDuration
	}

	iat := time.Now()
	exp := iat.Add(duration)
	claims := &Claims{
		Name:     user.Name,
		Email:    user.Email,
		Username: user.Username,
		Type:     tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			IssuedAt:  jwt.NewNumericDate(iat),
			ExpiresAt: jwt.NewNumericDate(exp),
			Subject:   fmt.Sprint(user.ID),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		j.logger.Error("jwt validation failed: ", err)
		return "", fmt.Errorf("jwt validation failed: %s", err)
	}

	return tokenString, nil
}

// Authorize authorizes the generated token
func (j JWTAuth) VerifyToken(tokenString string, tokenType string) (*Claims, error) {
	var secret string
	if tokenType == constants.TokenAccess {
		secret = j.env.JWTAccessSecret
	} else if tokenType == constants.TokenRefresh {
		secret = j.env.JWTRefreshSecret
	}

	claims := new(Claims)
	var keyfunc jwt.Keyfunc = func(t *jwt.Token) (interface{}, error) { return []byte(secret), nil }
	token, err := jwt.ParseWithClaims(tokenString, claims, keyfunc)

	if token != nil && token.Valid {
		return claims, nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return nil, jwt.ErrTokenMalformed
		}
		if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return nil, jwt.ErrTokenExpired
		}
	}

	return nil, err
}
