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

	email := req.FormValue("email")

	var mail string

	err = db.QueryRow("SELECT email FROM newsletter WHERE email=?", email).Scan(&mail)

    switch {
    	case err == sql.ErrNoRows:

        _, err = db.Exec("INSERT INTO newsletter(email) VALUES(?)", email)
        if err != nil {
            res.WriteHeader(http.StatusInternalServerError)    
            return
        }

         default:
	    	http.Redirect(res, req, "/", 301)

    }

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
	log.Fatal(http.ListenAndServe(":8080", nil))
}
