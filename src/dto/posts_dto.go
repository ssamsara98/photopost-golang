package dto

import "mime/multipart"

type GetPostByIDParams struct {
	ID string `uri:"postId" binding:"required"`
}

type GetPostByUserIDParams struct {
	ID string `uri:"userId" binding:"required"`
}

type CreatePostReqDto struct {
	Caption  string   `form:"caption"`
	PhotoIds []string `form:"photoIds"`
}

type UploadPhotoReqDto struct {
	Image multipart.FileHeader `form:"image" binding:"required"`
}
