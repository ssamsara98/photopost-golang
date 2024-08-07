package dto

import (
	"mime/multipart"
)

type GetPostByIDParams struct {
	PostID string `uri:"postId" json:"postId" validate:"required"`
}

type GetPostByUserIDParams struct {
	UserID string `uri:"userId" json:"userId" validate:"required"`
}

type CreatePostDto struct {
	Caption  string   `json:"caption"`
	PhotoIDs []string `json:"photoIds"`
}

type UploadPhotoDto struct {
	Image *multipart.FileHeader `json:"image" form:"image" validate:"required"`
}

type UpdatePostDto struct {
	Title   *string `json:"title"`
	Content *string `json:"content"`
}

type PublishPostDto struct {
	IsPublished *bool `json:"isPublished"`
}

type AddPostCommentDto struct {
	Content string `json:"content"`
}
