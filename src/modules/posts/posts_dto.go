package posts

import "mime/multipart"

type GetPostByIdParams struct {
	ID string `uri:"id" binding:"required"`
}

type CreatePostReqDto struct {
	Caption string `form:"caption"`
}

type UploadPhotoReqDto struct {
	Image multipart.FileHeader `form:"image" binding:"required"`
}
