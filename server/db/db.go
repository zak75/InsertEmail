package db

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//InsertEmail ...
func InsertEmail(email string) error {
	log.Println("preparing to insert email")
	var mail string

	db, err := sql.Open("mysql", "root:<PASSWORD>@/<DBNAME>")

	if err != nil {
		log.Printf("%v", err)
		return errors.New("error connecting to database")
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Printf("%v", err)
		return errors.New("error connecting to database")
	}

	log.Println("database opened")

	err = db.QueryRow("SELECT email FROM newsletter WHERE email=?", email).Scan(&mail)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("no email found")
			_, err = db.Exec("INSERT INTO newsletter(email) VALUES(?)", email)
			if err != nil {
				log.Println("database error")
				log.Printf("%v", err)
				return errors.New("error adding email")
			}
		} else {
			return nil
		}
		return err
	}
	return nil
}
