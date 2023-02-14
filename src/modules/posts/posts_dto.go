package posts

import "mime/multipart"

type GetPostByIdUri struct {
	ID string `uri:"id" binding:"required"`
}

type CreatePostDto struct {
	Caption string `form:"caption"`
}

type UploadPhotoDto struct {
	Image multipart.FileHeader `form:"image"`
}
