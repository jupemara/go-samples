package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open(
		"mysql",
		fmt.Sprintf("%s@(127.0.0.1:3306)/", os.Getenv("CLOUDSQL_USER")),
	)
	if err != nil {
		log.Fatalf("could not connect mysql through cloudsql-proxy: %s", err)
	}
	m := http.NewServeMux()
	m.HandleFunc("/", handler)
	log.Println("starting the http server...")
	if err := http.ListenAndServe("0.0.0.0:8080", m); err != nil {
		log.Fatalf("could not start the http server: %s", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	rows, err := db.Query("SHOW DATABASES;")
	if err != nil {
		http.Error(w, fmt.Sprintf("database error: %s", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	buf := "DBs:\n"
	for rows.Next() {
		var v string
		if err := rows.Scan(&v); err != nil {
			http.Error(w, fmt.Sprintf("rows read error: %s", err), http.StatusInternalServerError)
			return
		}
		buf += fmt.Sprintf("- %s\n", v)
	}
	w.Write([]byte(buf))
}
