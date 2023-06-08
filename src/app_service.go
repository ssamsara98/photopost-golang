package src

import (
	"errors"
	"go-photopost/src/entities"
	"go-photopost/src/helpers"
	"go-photopost/src/lib"

	"gorm.io/gorm"
)

// type AppServiceInf interface {
// 	Register(body *RegisterUserDto) *entities.User
// 	Login(body *LoginUserDto) (*lib.Token, error)
// }

type AppService struct {
	DB            *gorm.DB
	JWTAuthHelper *lib.JWTAuthHelper
}

func NewAppService(
	db *gorm.DB,
	jwtAuthHelper *lib.JWTAuthHelper,
) *AppService {
	return &AppService{
		db,
		jwtAuthHelper,
	}
}

func (app AppService) Register(body *RegisterUserDto) *entities.User {
	hashedPassword := helpers.HashPassword([]byte(body.Password))

	user := entities.User{
		Email:    &body.Email,
		Username: &body.Username,
		Password: string(hashedPassword),
		Name:     body.Name,
	}
	app.DB.Create(&user)

	return &user
}

func (app AppService) Login(body *LoginUserDto) (*lib.Token, error) {
	var user entities.User
	res := app.DB.Where("email = ? OR username = ?", body.UserSession, body.UserSession).First(&user)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("True")
	}

	err := helpers.CompareHash([]byte(user.Password), []byte(body.Password))
	if err != nil {
		return nil, errors.New("True")
	}

	// create token
	token := app.JWTAuthHelper.CreateToken(user)

	return token, nil
}
