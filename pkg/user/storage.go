package user

import (
	"context"
	"errors"
	"fmt"
	"go-reminder-bot/pkg/enum"
	"go-reminder-bot/pkg/reminder"
	"gorm.io/gorm"
	"time"
)

var ErrNotFound = errors.New("user not found")

type User struct {
	gorm.Model
	Email     string              `json:"email" gorm:"email;size:256;uniqueIndex"`
	Password  string              `json:"-" gorm:"password"`
	Role      enum.UserRole       `json:"role" gorm:"role"`
	IsActive  bool                `json:"is_active" gorm:"is_active"`
	Reminders []reminder.Reminder `gorm:"foreignKey:CreatedBy;references:Email"`
}

func NewUser(email, password string) User {
	now := time.Now()
	return User{
		Model: gorm.Model{
			CreatedAt: now,
			UpdatedAt: now,
		},
		Email:    email,
		Password: password,
		Role:     enum.RoleGuest,
		IsActive: true,
	}
}

type Storage interface {
	Create(ctx context.Context, user User) (User, error)
	GetList(ctx context.Context, p GetListParams) ([]User, int64, error)
	GetOne(ctx context.Context, id int64) (User, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, user User) error
	GetByEmail(ctx context.Context, email string) (User, error)
	GetActiveUsers(ctx context.Context) ([]User, error)
	UpdateNewPassword(ctx context.Context, userID uint, newPassword string) error
}

type storage struct {
	db *gorm.DB
}

func (t *storage) GetByEmail(ctx context.Context, email string) (user User, err error) {
	err = t.db.WithContext(ctx).Where(&User{
		Email: email,
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

func (t *storage) UpdateNewPassword(ctx context.Context, userID uint, newPassword string) (err error) {
	err = t.db.WithContext(ctx).Model(&User{}).
		Where("id = ?", userID).
		Select("password").
		Updates(User{Password: newPassword}).Error
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
		Order(fmt.Sprintf("%s %s", param.SortBy, param.SortType)).
		Find(&users).Count(&count).Error
	if err != nil {
		return
	}
	return
}

func (t *storage) GetActiveUsers(ctx context.Context) (users []User, err error) {
	err = t.db.WithContext(ctx).
		Where("is_active", true).
		Preload("Reminders").
		Find(&users).Error
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
