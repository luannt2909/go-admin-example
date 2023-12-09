package di

import (
	"go-admin/pkg/db"
	"go-admin/pkg/token"
	"go-admin/pkg/user"
	"go-admin/server"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"os"
)

var Module = fx.Provide(
	provideDB,
	provideTokenizer,
	provideUserStorage,
	provideServer,
)

func provideDB() (*gorm.DB, error) {
	return db.InitDatabase()
}

func provideUserStorage(db *gorm.DB) user.Storage {
	return user.NewStorage(db)
}

func provideTokenizer() token.Tokenizer {
	secretKey := os.Getenv("JWT_SIGNING_KEY")
	return token.NewJwtTokenizer([]byte(secretKey))
}

func provideServer(tokenizer token.Tokenizer, userStorage user.Storage) server.Router {
	handler := server.NewHandler(tokenizer, userStorage)
	authMdw := server.AuthHandler(tokenizer, userStorage)
	return server.NewRouter(handler, authMdw)
}
