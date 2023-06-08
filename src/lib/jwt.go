package lib

import (
	"errors"
	"go-photopost/src/entities"

	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Token struct {
	Type  string `json:"type"`
	Token string `json:"token"`
}

// JWTAuthHelper service relating to authorization
type JWTAuthHelper struct {
	Log *zap.Logger
	Env *Env
	DB  *gorm.DB
}

// NewJWTAuthHelper creates a new auth service
func NewJWTAuthHelper(
	log *zap.Logger,
	env *Env,
	db *gorm.DB,
) *JWTAuthHelper {
	return &JWTAuthHelper{
		log,
		env,
		db,
	}
}

// CreateToken creates jwt auth token
func (j JWTAuthHelper) CreateToken(user entities.User) *Token {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":      user.ID,
		"name":     user.Name,
		"email":    *user.Email,
		"username": *user.Username,
	})

	tokenString, err := token.SignedString([]byte(j.Env.JWTSecret))

	if err != nil {
		j.Log.Sugar().Infof("JWT validation failed: ", err)
	}

	return &Token{
		Type:  "Bearer",
		Token: tokenString,
	}
}

// Authorize authorizes the generated token
func (j JWTAuthHelper) Authorize(tokenString string) (*entities.User, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.Env.JWTSecret), nil
	})

	if token.Valid {
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return nil, errors.New("token claims error")
		}

		var user entities.User
		j.DB.Find(&entities.User{}, claims["sub"]).First(&user)

		return &user, nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return nil, errors.New("token malformed")
		}
		if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return nil, errors.New("token expired")
		}
	}

	return nil, errors.New("couldn't handle token")
}
