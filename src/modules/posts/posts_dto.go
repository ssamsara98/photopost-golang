package posts

import "mime/multipart"

type GetPostByIdParams struct {
	ID string `uri:"id" binding:"required"`
}

type CreatePostReqDto struct {
	Caption  string   `form:"caption"`
	PhotoIds []string `form:"photoIds"`
}

type UploadPhotoReqDto struct {
	Image multipart.FileHeader `form:"image" binding:"required"`
}
