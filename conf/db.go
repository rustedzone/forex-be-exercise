package conf

import (
	"database/sql"
	"fmt"

	"github.com/akbarkn/aknenv"
	_ "github.com/lib/pq"
)

func Db() *sql.DB {
	host := aknenv.GetEnv("DB_HOST")
	port := aknenv.GetEnv("DB_PORT")
	user := aknenv.GetEnv("DB_USER")
	name := aknenv.GetEnv("DB_NAME")
	password := aknenv.GetEnv("DB_PASSWORD")
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&timezone=Asia/Jakarta", user, password, host, port, name))
	if err != nil {
		panic(err.Error())
	}
	return db
}
