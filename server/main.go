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

func test(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("bonjour!"))
}

func insertEmail(res http.ResponseWriter, req *http.Request) {
	db, err := sql.Open("mysql", "root:<PASSWORD>@/<DBNAME>")
	if err != nil {
		log.Printf("database connection error: %v", err)
		res.Write([]byte("something went wrong"))
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Printf("database connection error: %v", err)
		res.Write([]byte("something went wrong"))
		res.WriteHeader(http.StatusInternalServerError)
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
	http.HandleFunc("/test", test)
	log.Println("server running....")
	log.Fatal(http.ListenAndServe(":8090", nil))
}
