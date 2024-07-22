package services

import (
	"github.com/ssamsara98/photopost-golang/src/api/dto"
	"github.com/ssamsara98/photopost-golang/src/infrastructure"
	"github.com/ssamsara98/photopost-golang/src/lib"
	"github.com/ssamsara98/photopost-golang/src/models"

	"gorm.io/gorm"
)

type UsersService struct {
	logger          *lib.Logger
	db              *infrastructure.Database
	paginationScope *gorm.DB
}

func NewUsersService(
	logger *lib.Logger,
	db *infrastructure.Database,
) *UsersService {
	return &UsersService{
		logger: logger,
		db:     db,
	}
}

// PaginationScope
func (u UsersService) SetPaginationScope(scope func(*gorm.DB) *gorm.DB) *UsersService {
	u.paginationScope = u.db.WithTrx(u.db.Scopes(scope)).DB
	return &u
}

func (u UsersService) GetUserList() (*[]models.User, *int64, error) {
	var items []models.User
	var count int64

	err := u.db.WithTrx(u.paginationScope).Find(&items).Offset(-1).Limit(-1).Count(&count).Error
	if err != nil {
		return nil, nil, err
	}

	return &items, &count, nil
}
func (u UsersService) GetUserListCursor(cursor *int64) (*[]models.User, error) {
	var items []models.User

	var err error
	if *cursor != 0 {
		err = u.db.WithTrx(u.paginationScope).Where("id > ?", *cursor).Find(&items).Limit(-1).Error
	} else {
		err = u.db.WithTrx(u.paginationScope).Find(&items).Limit(-1).Error
	}

	if err != nil {
		return nil, err
	}

	return &items, nil
}

func (u UsersService) GetUserByID(uri *dto.GetUserByIDParams) (user models.User, err error) {
	return user, u.db.First(&user, "id = ?", uri.ID).Error
}
