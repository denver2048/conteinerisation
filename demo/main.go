package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

func connectDB() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	password := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf(
		"host=%s port=5432 user=postgres password=%s dbname=postgres sslmode=disable",
		host, password,
	)

	return sql.Open("postgres", dsn)
}

func waitForDB() {
	var err error

	for i := 0; i < 10; i++ {
		db, err = connectDB()
		if err == nil {
			err = db.Ping()
			if err == nil {
				fmt.Println("Connected to DB")
				return
			}
		}

		fmt.Println("⏳ Waiting for DB...")
		time.Sleep(2 * time.Second)
	}

	panic("DB not available")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "Hello from %s 🚀\n", hostname)
}

func dbHandler(w http.ResponseWriter, r *http.Request) {
	var result int
	err := db.QueryRow("SELECT 1").Scan(&result)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprintf(w, "DB response: %d\n", result)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Fprintf(w, "OK")
}

func main() {
	waitForDB()

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/db", dbHandler)
	http.HandleFunc("/health", healthHandler)

	fmt.Println("App started on :8080")
	http.ListenAndServe(":8080", nil)
}