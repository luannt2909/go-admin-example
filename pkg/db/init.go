package db

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"go-reminder-bot/pkg/config"
	"go-reminder-bot/pkg/enum"
	"go-reminder-bot/pkg/reminder"
	"go-reminder-bot/pkg/user"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"time"
)

const dbFileName = "./tmp/go-reminder-bot.db"

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
func InitSQLiteDB() (*gorm.DB, error) {
	isExistedDB := false
	if fileExists(dbFileName) {
		fmt.Println("db is existed already")
		isExistedDB = true
	} else {
		fmt.Println("db is not exists, start create and migrate db")
		os.MkdirAll("./tmp", 0755)
		os.Create(dbFileName)
	}

	db, err := gorm.Open(sqlite.Open(dbFileName), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if !isExistedDB {
		_ = db.AutoMigrate(&reminder.Reminder{}, &user.User{})
		adminEmail := "admin@reminderbot.com"
		adminPwd := os.Getenv("DEFAULT_ADMIN_PASSWORD")
		if adminPwd == "" {
			adminPwd = "reminderbot"
		}
		db.Create(&user.User{
			Model: gorm.Model{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Email:    adminEmail,
			Password: adminPwd,
			Role:     enum.RoleAdmin,
			IsActive: true,
		})
		adminReminder := reminder.DefaultReminder
		adminReminder.CreatedBy = adminEmail
		db.Create(&adminReminder)

		guestEmail := "guest@reminderbot.com"
		guestPwd := os.Getenv("DEFAULT_GUEST_PASSWORD")
		if guestPwd == "" {
			guestPwd = "reminderbot"
		}
		db.Create(&user.User{
			Model: gorm.Model{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Email:    guestEmail,
			Password: guestPwd,
			Role:     enum.RoleGuest,
			IsActive: true,
		})
		guestReminder := reminder.DefaultReminder
		guestReminder.CreatedBy = guestEmail
		db.Create(&guestReminder)
	}
	return db, nil
}

func InitDatabase(cfg config.DBConfig) (db *gorm.DB, err error) {
	db, err = initDB(cfg)
	if err != nil {
		return
	}
	err = db.AutoMigrate(&user.User{}, &reminder.Reminder{})
	return
}

func initDB(cfg config.DBConfig) (db *gorm.DB, err error) {
	switch cfg.DBClient {
	case "mysql":
		db, err = gorm.Open(mysql.Open(cfg.DBConnectionURI), &gorm.Config{})
		if err != nil {
			return
		}
		return
	default:
		return InitSQLiteDB()
	}
}
