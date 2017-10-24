package main

import (
	"database/sql"
	"errors"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func doDBStuffHere() error {
	return errors.New("database layer not implemented")
}

func insertEmail(res http.ResponseWriter, req *http.Request) {
	db, err := sql.Open("mysql", "root:<PASSWORD>@/<DBNAME>")
	if err != nil {
		log.Printf("database connection error: %v", err)
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("something went wrong"))
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Printf("database connection error: %v", err)
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("something went wrong"))
		return
	}

	if req.Method != "POST" {
		res.WriteHeader(http.StatusMethodNotAllowed)
		res.Write([]byte("something went wrong"))
		return
	}

	err = doDBStuffHere()
	if err != nil {
		log.Printf("Server error, unable to create account: %v", err)
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("something went wrong"))
		return
	}
}

func main() {
	http.HandleFunc("/", insertEmail)
	log.Println("server running....")
	log.Fatal(http.ListenAndServe(":8090", nil))
}
