package services

import (
	"fmt"
	"photopost/infrastructure"
	"photopost/lib"
	"photopost/models"
	"photopost/src/dto"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PostsService struct {
	logger          lib.Logger
	db              infrastructure.Database
	paginationScope *gorm.DB
}

func NewPostsService(
	logger lib.Logger,
	db infrastructure.Database,
) *PostsService {
	return &PostsService{
		logger: logger,
		db:     db,
	}
}

// PaginationScope
func (p PostsService) SetPaginationScope(scope func(*gorm.DB) *gorm.DB) PostsService {
	p.paginationScope = p.db.WithTrx(p.db.Scopes(scope)).DB
	return p
}

func (p PostsService) UploadPhoto(keypath *string) *models.PostPhoto {
	newPhoto := &models.PostPhoto{
		Keypath: *keypath,
	}
	p.db.Create(newPhoto)

	return newPhoto
}

func (p PostsService) CreatePost(user *models.User, createPostDto *dto.CreatePostReqDto) *models.Post {
	fmt.Println(createPostDto.PhotoIds)

	var newPost *models.Post

	_ = p.db.Transaction(func(tx *gorm.DB) error {
		newPost = &models.Post{
			AuthorID: user.ID,
			Caption:  createPostDto.Caption,
		}
		tx.Create(newPost)

		postPhotoJoins := make([]*models.PostToPhoto, 0)
		for i := 0; i < len(createPostDto.PhotoIds); i++ {
			newPostPhoto := &models.PostToPhoto{
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

func (p PostsService) GetPostList() (response gin.H, err error) {
	var postList []models.Post
	var count int64

	err = p.db.WithTrx(p.paginationScope).Preload("Author").Preload("Photos", func(db *gorm.DB) *gorm.DB {
		return p.db.Order("post_to_photos.position ASC").Preload("Photo")
	}).Order("created_at DESC").Find(&postList).Offset(-1).Limit(-1).Count(&count).Error
	if err != nil {
		return nil, err
	}

	return gin.H{"result": postList, "count": count}, nil
}

func (p PostsService) GetPost(params *dto.GetPostByIDParams) *models.Post {
	var post models.Post
	p.db.Preload("Author").Preload("Photos", func(db *gorm.DB) *gorm.DB {
		return p.db.Order("post_to_photos.position ASC").Preload("Photo")
	}).Order("created_at DESC").First(&post, params.ID)

	return &post
}

func (p PostsService) GetMyPostList(user *models.User) (response gin.H, err error) {
	var postList []models.Post
	var count int64

	err = p.db.WithTrx(p.paginationScope).Preload("Photos", func(db *gorm.DB) *gorm.DB {
		return p.db.Order("post_to_photos.position ASC").Preload("Photo")
	}).Order("created_at DESC").Where("author_id = ?", user.ID).Find(&postList).Offset(-1).Limit(-1).Count(&count).Error

	if err != nil {
		return nil, err
	}

	return gin.H{"result": postList, "count": count}, nil
}

func (p PostsService) GetUserPostList(params *dto.GetPostByUserIDParams) (response gin.H, err error) {
	var postList []models.Post
	var count int64

	err = p.db.WithTrx(p.paginationScope).Preload("Photos", func(db *gorm.DB) *gorm.DB {
		return p.db.Order("post_to_photos.position ASC").Preload("Photo")
	}).Order("created_at DESC").Where("author_id = ?", params.ID).Where("is_published = ?", true).Find(&postList).Offset(-1).Limit(-1).Count(&count).Error

	if err != nil {
		return nil, err
	}

	return gin.H{"result": postList, "count": count}, nil
}
