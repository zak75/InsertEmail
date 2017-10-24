package main

import "net/http"

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

//Global sql.DB pour accéder à la base de données partout
var db *sql.DB
var err error 

func signupPage(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "../client/index.html")
	//servir signup.html pour récuperer les requetes sur /signup
	if req.Method != "POST" {
		http.ServeFile(res, req, "../client/index.html")
		return
	}

	//on va d'abord s'assurer de la disponibilité des informations dans la bdd, et si c'est le cas, on crée l'utilisateur
	email := req.FormValue("email")

	var mail string

	err := db.QueryRow("SELECT email FROM users WHERE email=?", email).Scan(&mail)

	switch {
		//pseudo est disponible
	case err == sql.ErrNoRows:

        _, err = db.Exec("INSERT INTO users(email) VALUES(?)", email)
        if err != nil {
            http.Error(res, "Server error, unable to create your account.", 500)    
            return
        }

        res.Write([]byte("Account created!"))
        return
    case err != nil: 
        http.Error(res, "Server error, unable to create your account.", 500)    
        return

    default: 
        http.Redirect(res, req, "/", 301)
    }
}


//func homePage(res http.ResponseWriter, req *http.Request) {
//	http.ServeFile(res, req, "../client/index.html")

//}


func main() {

	//Créer un sql.DB et vérifier les erreurs 
	db, err = sql.Open("mysql", "root:<PASSWORD>@/<DBNAME>")
	if err != nil {
		panic(err.Error())
	}

    //sql.DB doit être maintenu jusqu'à la fin de la fonction, il prend fin avec le defer
	defer db.Close()

    //Tester la connection mysql
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
    
	http.HandleFunc("/", signupPage) //rajouter loginPage, homePage, signupPage
	http.ListenAndServe("80", nil)
}
