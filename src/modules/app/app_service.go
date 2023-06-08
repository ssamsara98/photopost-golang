package app

import (
	"errors"
	"go-photopost/src/entities"
	"go-photopost/src/helpers"
	"go-photopost/src/lib"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AppServiceInf interface {
	Register(body *RegisterUserDto) *entities.User
	Login(body *LoginUserDto) (*lib.Token, error)
}

type AppService struct {
	Log           *zap.Logger
	DB            *gorm.DB
	JWTAuthHelper *lib.JWTAuthHelper
}

func NewAppService(
	log *zap.Logger,
	db *gorm.DB,
	jwtAuthHelper *lib.JWTAuthHelper,
) AppServiceInf {
	return &AppService{
		log,
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
