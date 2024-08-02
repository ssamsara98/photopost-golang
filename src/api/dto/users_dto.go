package dto

type GetUserByIDParams struct {
	UserID string `uri:"userId" json:"userId" validate:"required"`
}
