package dto

import "mime/multipart"

type GetPostByIDParams struct {
	ID string `uri:"postId" binding:"required"`
}

type GetPostByUserIDParams struct {
	ID string `uri:"userId" binding:"required"`
}

type CreatePostDto struct {
	Caption  string   `form:"caption"`
	PhotoIds []string `form:"photoIds"`
}

type UploadPhotoDto struct {
	Image multipart.FileHeader `form:"image" binding:"required"`
}

type UpdatePostDto struct {
	Title   *string `form:"title"`
	Content *string `form:"content"`
}

type PublishPostDto struct {
	IsPublished *bool `form:"isPublished"`
}

type AddPostCommentDto struct {
	Content string `form:"content"`
}
