package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func ConnectToDB() (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)

	defer cancel()

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}
	connectionString := os.Getenv("DATABASE_URL")

	db, err := pgxpool.New(ctx, connectionString)
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return nil, err
	} else {
		fmt.Println("Successfully connected to the database!")
	}

	err = db.Ping(ctx)
	if err != nil {
		fmt.Println("Error pinging the database:", err)

	} else {
		fmt.Println("Successfully pinged the database!")
	}

	return db, nil
}
