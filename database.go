package main

import (
"database/sql"
"fmt"
"os"
"github.com/joho/godotenv"
_ "github.com/lib/pq"
)

func connectDatabase() *sql.DB {
err := godotenv.Load()
if err != nil {
panic("Failed to load .env")
}
host := os.Getenv("DB_HOST")
port := os.Getenv("DB_PORT")
user := os.Getenv("DB_USER")
password := os.Getenv("DB_PASSWORD")
dbname := os.Getenv("DB_NAME")

connectionString := fmt.Sprintf(
	"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	host,
	port,
	user,
	password,
	dbname,
)

db, err := sql.Open("postgres", connectionString)
if err != nil {
	panic(err)
}

err = db.Ping()
if err != nil {
	panic(err)
}

return db
}
