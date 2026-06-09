package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var Db *pgx.Conn
var err error

func ConnectDB() {
	urlString := os.Getenv("DATABASE_URL")
	Db, err = pgx.Connect(context.Background(), urlString)
	if err != nil {
		panic(err)
		os.Exit(1)
	}
	fmt.Println("Server is connecting")
}
