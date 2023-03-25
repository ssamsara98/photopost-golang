package users

type GetUserByIdParams struct {
	ID string `uri:"userId" binding:"required"`
}

type CreateUserReqDto struct {
	Email    string `form:"name"`
	Username string `form:"username"`
	Password string `form:"password"`
	Name     string `form:"name"`
}
