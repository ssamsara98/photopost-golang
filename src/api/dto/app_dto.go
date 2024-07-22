package dto

import "time"

type RegisterUserDto struct {
	Email    string `form:"email" binding:"required"`
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	Name     string `form:"name" binding:"required"`
}

type LoginUserDto struct {
	UserSession string `form:"userSession" binding:"required"`
	Password    string `form:"password" binding:"required"`
}

type UpdateProfileDto struct {
	Name      string     `form:"name"`
	SexType   string     `form:"sexType"`
	Birthdate *time.Time `form:"birthdate" time_format:"2006-01-02"`
}

type RenewAccessTokenReqDto struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}
