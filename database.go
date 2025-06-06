package main

import (
	ct "context"
	"log"
	"os"

	sq "github.com/jackc/pgx/v5"
)

func getConnectionStr() (connStr string) {
	const connVar = `SUPABASE_CONN_STR`
	if connStr = os.Getenv(connVar); connStr == "" {
		log.Fatalln("Please set", connVar, "environment variable")
		return
	}
	return
}

func connectToDb(connStr string) (conn *sq.Conn, e error) {
	if conn, e = sq.Connect(ct.Background(), connStr); e != nil {
		log.Fatalf("Failed to connect to the database: %v\n", e)
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
