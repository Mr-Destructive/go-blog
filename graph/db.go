package graph

import (
	"fmt"

	"github.com/go-pg/pg"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "goblog"
	DB_HOST     = "localhost"
	DB_PORT     = 5432
)

func Connect() *pg.DB {
	DB_URL := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	opt, err := pg.ParseURL(DB_URL)
	if err != nil {
		panic(err)
	}
	db := pg.Connect(opt)
	if _, DBStatus := db.Exec("SELECT 1"); DBStatus != nil {
		panic("PostgreSQL is down")
	}
	return db
	//dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	//db, err := sql.Open("postgres", dbinfo)
	//if _, DBStatus := db.Exec("SELECT 1"); DBStatus != nil {
	//	panic(err)
	//}
	//return db
}
