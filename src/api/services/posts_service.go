package services

import (
	"github.com/ssamsara98/photopost-golang/src/api/dto"
	"github.com/ssamsara98/photopost-golang/src/infrastructure"
	"github.com/ssamsara98/photopost-golang/src/lib"
	"github.com/ssamsara98/photopost-golang/src/models"

	"gorm.io/gorm"
)

type PostsService struct {
	logger          *lib.Logger
	db              *infrastructure.Database
	paginationScope *gorm.DB
}

func NewPostsService(
	logger *lib.Logger,
	db *infrastructure.Database,
) *PostsService {
	return &PostsService{
		logger: logger,
		db:     db,
	}
}

// WithTrx delegates transaction to repository database
func (p PostsService) WithTrx(trxHandle *gorm.DB) *PostsService {
	p.db = p.db.WithTrx(trxHandle)
	return &p
}

// PaginationScope
func (p PostsService) SetPaginationScope(scope func(*gorm.DB) *gorm.DB) *PostsService {
	p.paginationScope = p.db.WithTrx(p.db.Scopes(scope)).DB
	return &p
}

func (p PostsService) UploadPhoto(keypath *string) *models.PostPhoto {
	newPhoto := &models.PostPhoto{
		Keypath: *keypath,
	}
	p.db.Create(newPhoto)

	return newPhoto
}

func (p PostsService) CreatePost(user *models.User, body *dto.CreatePostDto) *models.Post {
	newPost := &models.Post{
		AuthorID: &user.ID,
		Caption:  body.Caption,
	}
	p.db.Create(newPost)

	postPhotoJoins := make([]*models.PostToPhoto, 0)
	for i := 0; i < len(body.PhotoIds); i++ {
		newPostPhoto := &models.PostToPhoto{
			Position: uint(i),
			PostID:   newPost.ID,
			PhotoID:  body.PhotoIds[i],
		}
		postPhotoJoins = append(postPhotoJoins, newPostPhoto)
	}
	p.db.Create(&postPhotoJoins)

	return newPost
}

func (p PostsService) GetPostList(cursor *int64) (*[]models.Post, error) {
	var items []models.Post

	var err error
	if *cursor != 0 {
		err = p.db.WithTrx(p.paginationScope).Preload("Author").Preload("Photos", func(db *gorm.DB) *gorm.DB {
			return p.db.Order("post_to_photos.position ASC").Preload("Photo")
		}).Order("id DESC").Where("id < ?", *cursor).Find(&items).Limit(-1).Error
	} else {
		err = p.db.WithTrx(p.paginationScope).Preload("Author").Preload("Photos", func(db *gorm.DB) *gorm.DB {
			return p.db.Order("post_to_photos.position ASC").Preload("Photo")
		}).Order("id DESC").Find(&items).Limit(-1).Error
	}

	if err != nil {
		return nil, err
	}

	return &items, nil
}

func (p PostsService) GetPostById(uri *dto.GetPostByIDParams) (post models.Post, err error) {
	return post, p.db.Preload("Author").Preload("Photos", func(db *gorm.DB) *gorm.DB {
		return p.db.Order("post_to_photos.position ASC").Preload("Photo")
	}).Order("created_at DESC").First(&post, uri.ID).Error
}

func (p PostsService) GetMyPostList(user *models.User) (items []models.Post, count *int64, err error) {
	var c int64

	err = p.db.WithTrx(p.paginationScope).Preload("Photos", func(db *gorm.DB) *gorm.DB {
		return p.db.Order("post_to_photos.position ASC").Preload("Photo")
	}).Order("created_at DESC").Order("id DESC").Where("author_id = ?", user.ID).Find(&items).Offset(-1).Limit(-1).Count(&c).Error
	if err != nil {
		return nil, nil, err
	}
	count = &c

	return items, count, nil
}

func (p PostsService) GetUserPostList(uri *dto.GetPostByUserIDParams) (items []models.Post, count *int64, err error) {
	var c int64

	err = p.db.WithTrx(p.paginationScope).Preload("Photos", func(db *gorm.DB) *gorm.DB {
		return p.db.Order("post_to_photos.position ASC").Preload("Photo")
	}).Order("created_at DESC").Where("author_id = ?", uri.ID).Where("is_published = ?", true).Find(&items).Offset(-1).Limit(-1).Count(&c).Error
	if err != nil {
		return nil, nil, err
	}
	count = &c

	return items, count, nil
}

// func (p PostsService) UpdatePost(user *models.User, uri *dto.GetPostByIDParams, body *dto.UpdatePostDto) {
// 	var post models.Post
// 	p.db.Where("id = ?", uri.ID).Where("author_id = ?", user.ID).First(&post)

// 	// if body.Title != nil {
// 	// 	post.Title = *body.Title
// 	// }
// 	// if body.Content != nil {
// 	// 	post.Content = *body.Content
// 	// }

// 	p.db.Save(&post)
// }

func (p PostsService) PublishPost(post *models.Post, uri *dto.GetPostByIDParams, body *dto.PublishPostDto) {
	if body.IsPublished != nil {
		post.IsPublished = *body.IsPublished
	}

	p.db.Save(&post)
}

func (p PostsService) DeletePost(post *models.Post, user *models.User, uri *dto.GetPostByIDParams) {
	p.db.Where("id = ?", uri.ID).Delete(&post)
}
