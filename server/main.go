package main

import (
	"log"
	"net/http"

	"github.com/zak75/InsertEmail/server/db"
)

func insertEmail(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		res.WriteHeader(http.StatusMethodNotAllowed)
		res.Write([]byte("something went wrong"))
		return
	}

	req.ParseForm()

	email := req.Form["email"]
	log.Println("email:", email)

	err := db.InsertEmail(email[0])

	if err != nil {
		res.Write([]byte(err.Error()))
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Write([]byte("Thanks, your email has been added to the newsletter"))
}

func main() {
	http.HandleFunc("/", insertEmail)
	log.Println("server running....")
	log.Fatal(http.ListenAndServe(":8090", nil))
}
