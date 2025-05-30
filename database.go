package main

import (
	ct "context"
	"log"
	"os"

	sq "github.com/jackc/pgx/v5"
)

func getConnectionStr() (dsn string) {
	if dsn = os.Getenv("SUPABASE_DB_DSN"); dsn == "" {
		log.Fatal("Please set SUPABASE_DB_DSN environment variable")
	}
	return
}

func connectToDb(dsn string) (conn *sq.Conn, err error) {
	if conn, err = sq.Connect(ct.Background(), dsn); err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return
	}
	return
}

/*	var version string
	if err := conn.QueryRow(context.Background(), "SELECT version()").Scan(&version); err != nil {
		log.Fatalf("Query failed: %v", err)
	}

	log.Println("Connected to:", version)

*/
