package utils

import (
	ht "net/http"

	sq "github.com/jackc/pgx/v5"
)

type (
	Rows = sq.Rows           // SQL Rows
	RW   = ht.ResponseWriter // HTTP Response Writer
	Req  = ht.Request        // HTTP Request
)
