package infrastructure

import "github.com/go-pg/pg"

var pool_cache *pg.DB;

func GetDatabasConnection() *pg.DB {
	if(pool_cache == nil){
		return ConnectDatabase();
	}

	return pool_cache;
}

func ConnectDatabase() *pg.DB {
	db := pg.Connect(&pg.Options{
		Addr:     "localhost:5432",
		User:     "postgres",
		Password: "YOUR_PASSWORD_HERE",
		Database: "hash",
	})

	return db;
}