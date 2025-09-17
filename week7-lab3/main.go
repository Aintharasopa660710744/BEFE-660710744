package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
func initDB() {
	var err error
	var db *sql.DB
	host := getEnv("DB_HOST", "")
	name := getEnv("DB_NAME", "")
	user := getEnv("DB_USER", "")
	password := getEnv("DB_PASSWORD", "")
	port := getEnv("DB_PORT", "")

	conSt := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, name)

	//fmt.Println("Connection String:", conSt)

	db, err = sql.Open("postgres", conSt)
	if err != nil {
		log.Fatal("Failed to open database")
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	log.Println("Connected to database successfully")
}
func main() {
	initDB()
}
