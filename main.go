package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := newDB()
	if err != nil {
		log.Printf("db connect error -> %s", err)
		panic(err)
	}

	conn, err := db.Conn(context.Background())
	if err != nil {
		log.Printf("db connect error2 -> %s", err)
		panic(err)
	}
	defer conn.Close()

	res, err := conn.ExecContext(context.Background(), "INSERT INTO test_tables VALUES('id', 'name')")
	if err != nil {
		log.Printf("exec error -> %s", err)
		panic(err)
	}
	log.Println("insert success", res)
}

func newDB() (*sql.DB, error) {
	config := mysql.Config{
		DBName: "test",
		User:   "root",
		Passwd: os.Getenv("DB_PASS"),
		Addr:   fmt.Sprintf("%s:%s", os.Getenv("DB_DOMAIN"), os.Getenv("DB_PORT")),
		Net:    "tcp",
	}

	log.Println("config", config.FormatDSN())
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		return nil, err
	}
	return db, nil
}
