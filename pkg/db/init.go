package db

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"go-admin/pkg/enum"
	"go-admin/pkg/user"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"time"
)

const dbFileName = "./tmp/go-admin.db"

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
		_ = db.AutoMigrate(&user.User{})
		db.Create(&user.User{
			Model: gorm.Model{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Username: "admin",
			Role:     enum.RoleAdmin,
			IsActive: true,
		})
		db.Create(&user.User{
			Model: gorm.Model{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Username: "guest",
			Role:     enum.RoleGuest,
			IsActive: true,
		})
	}
	return db, nil
}

func InitDatabase() (db *gorm.DB, err error) {
	db, err = InitSQLiteDB()
	if err != nil {
		return
	}
	err = db.AutoMigrate(&user.User{})
	return
}
