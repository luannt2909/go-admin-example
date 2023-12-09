package user

import (
	"context"
	"errors"
	"fmt"
	"go-admin/pkg/enum"
	"gorm.io/gorm"
	"time"
)

var ErrNotFound = errors.New("user not found")

type User struct {
	gorm.Model
	Username  string        `json:"username" gorm:"username"`
	Role      enum.UserRole `json:"role" gorm:"role"`
	IsActive  bool          `json:"is_active" gorm:"is_active"`
	CreatedBy string        `json:"created_by" gorm:"created_by"`
}

func NewUser(username string) User {
	now := time.Now()
	return User{
		Model: gorm.Model{
			CreatedAt: now,
			UpdatedAt: now,
		},
		Username: username,
		Role:     enum.RoleAdmin,
		IsActive: true,
	}
}

type Storage interface {
	Create(ctx context.Context, user User) (User, error)
	GetList(ctx context.Context, p GetListParams) ([]User, int64, error)
	GetOne(ctx context.Context, id int64) (User, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, user User) error
	GetByUsername(ctx context.Context, username string) (User, error)
}

type storage struct {
	db *gorm.DB
}

func (t *storage) GetByUsername(ctx context.Context, email string) (user User, err error) {
	err = t.db.WithContext(ctx).Where(&User{
		Username: email,
	}).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = ErrNotFound
	}
	return
}

func (t *storage) Update(ctx context.Context, user User) (err error) {
	err = t.db.WithContext(ctx).Select("*").Updates(&user).Error
	return
}

func (t *storage) Delete(ctx context.Context, id int64) (err error) {
	err = t.db.WithContext(ctx).Delete(&User{}, id).Error
	return
}

func (t *storage) Create(ctx context.Context, user User) (result User, err error) {
	db := t.db.WithContext(ctx).Create(&user)
	return user, db.Error
}

func (t *storage) GetList(ctx context.Context, param GetListParams) (users []User, count int64, err error) {
	err = t.db.WithContext(ctx).Offset(param.Offset).
		Limit(param.Limit).
		Where(param.Filter).
		Order(fmt.Sprintf("%s %s", param.SortBy, param.SortType)).
		Find(&users).Count(&count).Error
	if err != nil {
		return
	}
	return
}

func (t *storage) GetOne(ctx context.Context, id int64) (user User, err error) {
	err = t.db.WithContext(ctx).First(&user, id).Error
	return
}

func NewStorage(db *gorm.DB) Storage {
	return &storage{db: db}
}
