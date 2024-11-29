package main

import (
	"os"
)

const ENV_DB = "DB"
const ENV_DB_DEF = "postgres://thex:thex1234@db/thex"

func main() {
	dbStr := os.Getenv(ENV_DB)
	if dbStr == "" {
		dbStr = ENV_DB_DEF
	}
	DBInit(dbStr)
}
