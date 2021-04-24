package infrastructure

import (
	"os"

	"github.com/go-pg/pg"
)

var pool_cache *pg.DB

func GetDatabasConnection() *pg.DB {
	if pool_cache == nil {
		return ConnectDatabase()
	}

	return pool_cache
}

func ConnectDatabase() *pg.DB {
	addr := "localhost:5432"
	if os.Getenv("APP_ENV") == "production" {
		addr = "hash-db:5432"
	}

	db := pg.Connect(&pg.Options{
		Addr:     addr,
		User:     "postgres",
		Password: "YOUR_PASSWORD_HERE",
		Database: "hash",
	})

	return db
}
