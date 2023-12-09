package di

import (
	"go-reminder-bot/admin/server"
	"go-reminder-bot/pkg/config"
	"go-reminder-bot/pkg/db"
	"go-reminder-bot/pkg/user"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

var Module = fx.Provide(
	provideConfig,
	provideDB,
	provideServer,
	provideUserStorage,
)

func provideConfig() (config.Config, error) {
	return config.LoadEnv()
}

func provideDB(cfg config.Config) (*gorm.DB, error) {
	return db.InitDatabase(cfg.DBConfig)
}

func provideUserStorage(db *gorm.DB) user.Storage {
	return user.NewStorage(db)
}

func provideServer(userStorage user.Storage) server.Server {
	handler := server.NewHandler(userStorage)
	return server.NewServer(handler)
}
