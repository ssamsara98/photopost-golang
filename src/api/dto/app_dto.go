package dto

import "time"

type RegisterUserDto struct {
	Email    string `json:"email" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
}

type LoginUserDto struct {
	UserSession string `json:"userSession" validate:"required"`
	Password    string `json:"password" validate:"required"`
}

type UpdateProfileDto struct {
	Name      string     `json:"name"`
	SexType   string     `json:"sexType"`
	Birthdate *time.Time `json:"birthdate"`
}

type RenewAccessTokenReqDto struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}
