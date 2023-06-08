package users

import (
	"go-photopost/src/entities"
	"go-photopost/src/helpers"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UsersServiceV1Inf interface {
	CreateUser(body CreateUserReqDto) *entities.User
	GetUserList() []entities.User
	GetUser(uri *GetUserByIdParams) *entities.User
}

type UsersServiceV1 struct {
	Log *zap.Logger
	DB  *gorm.DB
}

func NewUsersServiceV1(
	log *zap.Logger,
	db *gorm.DB,
) UsersServiceV1Inf {
	return &UsersServiceV1{
		Log: log,
		DB:  db,
	}
}

func (us UsersServiceV1) CreateUser(body CreateUserReqDto) *entities.User {
	hashedPassword := helpers.HashPassword([]byte(body.Password))

	user := entities.User{
		Email:     &body.Email,
		Username:  &body.Username,
		Password:  string(hashedPassword),
		Name:      body.Name,
		Birthdate: nil,
	}
	us.DB.Create(&user)

	return &user
}

func (us UsersServiceV1) GetUserList() []entities.User {
	var users []entities.User
	us.DB.Find(&users)

	return users
}

func (us UsersServiceV1) GetUser(uri *GetUserByIdParams) *entities.User {
	var user entities.User
	us.DB.First(&user, uri.ID)

	return &user
}
