package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectToDB() (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)

	defer cancel()

	connectionString := "postgres://admin:password@localhost:5432/Portfolio"

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
