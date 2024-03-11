package main

import (
	"context"
	"fmt"
	"goldencoderam/web-server/db"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Error loading .env file: %s", err)
    }

    connection, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Unable to connecot to database: %v\n", err)
        os.Exit(1)
    }
    defer connection.Close(context.Background())

    query := db.New(connection)
    author, err := query.GetAuthor(context.Background(), 1)
    if err != nil {
        fmt.Fprintf(os.Stderr, "GetAuthor failed: %v\n", err)
        os.Exit(1)
    }

    fmt.Printf(author.Name)
}
