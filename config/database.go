package config

import(
	"fmt"
	"log"
	"os"
	
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_"github.com/jackc/pgx/v5/stdlib"
)
func ConnectDB() *sqlx.DB  {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn:= fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	db, err:= sqlx.Connect("pgx",dsn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

