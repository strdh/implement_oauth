package config

import (
	"time"
	"os"
	"fmt"
	"database/sql"
	"exercise/gooauth/utils"
	"github.com/joho/godotenv"
)

func NewDB() *sql.DB {
	err1 := godotenv.Load(".env")
	utils.PanicIfError(err1)

	database := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DATABASE_USERNAME"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_PORT"), os.Getenv("DATABASE_NAME"))

	db, err2 := sql.Open("mysql", database)
	utils.PanicIfError(err2)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
