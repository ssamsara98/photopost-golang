package posts

import (
	"fmt"
	"go-photopost/src/entities"
	"log"

	"gorm.io/gorm"
)

type PostsServiceV1Inf interface {
	UploadPhoto(keypath *string) *entities.PostPhoto
	CreatePost(user *entities.User, createPostDto *CreatePostReqDto) *entities.Post
	GetPostList() []entities.Post
	GetPost(uri *GetPostByIdParams) *entities.Post
}

type PostsServiceV1 struct {
	Log *log.Logger
	DB  *gorm.DB
}

func NewPostsServiceV1(
	log *log.Logger,
	db *gorm.DB,
) *PostsServiceV1 {
	return &PostsServiceV1{
		Log: log,
		DB:  db,
	}
}

func (ps PostsServiceV1) UploadPhoto(keypath *string) *entities.PostPhoto {
	newPhoto := &entities.PostPhoto{
		Keypath: *keypath,
	}
	ps.DB.Create(newPhoto)

	return newPhoto
}

func (ps PostsServiceV1) CreatePost(user *entities.User, createPostDto *CreatePostReqDto) *entities.Post {
	fmt.Println(createPostDto.PhotoIds)

	var newPost *entities.Post

	ps.DB.Transaction(func(tx *gorm.DB) error {
		newPost = &entities.Post{
			AuthorID: user.ID,
			Caption:  createPostDto.Caption,
		}
		tx.Create(newPost)

		postPhotoJoins := make([]*entities.PostToPhoto, 0)
		for i := 0; i < len(createPostDto.PhotoIds); i++ {
			newPostPhoto := &entities.PostToPhoto{
				Position: uint(i),
				PostID:   newPost.ID,
				PhotoID:  createPostDto.PhotoIds[i],
			}

			postPhotoJoins = append(postPhotoJoins, newPostPhoto)
		}
		tx.Create(&postPhotoJoins)

		return nil
	})

	return newPost
}

func (ps PostsServiceV1) GetPostList() []entities.Post {
	var postList []entities.Post
	ps.DB.Preload("Photos", func(db *gorm.DB) *gorm.DB {
		return ps.DB.Order("post_to_photos.position ASC").Preload("Photo")
	}).Order("created_at DESC").Find(&postList)

	return postList
}

func (ps PostsServiceV1) GetPost(params *GetPostByIdParams) *entities.Post {
	var post entities.Post
	ps.DB.Preload("Photos", func(db *gorm.DB) *gorm.DB {
		return ps.DB.Order("post_to_photos.position ASC").Preload("Photo")
	}).Order("created_at DESC").First(&post, params.ID)

	return &post
}
