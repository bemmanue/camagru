package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

var db *sql.DB

func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/templates/index.html")
}

func signIn(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		http.ServeFile(w, r, "./web/templates/sign_in.html")
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		http.Redirect(w, r, "/profile", http.StatusFound)
	}
}

func signUp(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		http.ServeFile(w, r, "./web/templates/sign_up.html")
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		username := r.Form.Get("login")
		email := r.Form.Get("email")
		password := r.Form.Get("password")

		query := fmt.Sprintf("insert into users(username, email, password) values ('%s', '%s', '%s');",
			username, email, password)
		_, err := db.Exec(query)
		if err != nil {
			log.Fatalln(err)
		}

		http.Redirect(w, r, "/confirm", http.StatusSeeOther)
	}
}

func profile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/templates/profile.html")
}

func feed(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/templates/feed.html")
}

func settings(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/templates/settings.html")
}

func confirm(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/templates/confirm.html")
}

func main() {
	db = initDB()
	defer db.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("/", index)
	mux.HandleFunc("/sign_in", signIn)
	mux.HandleFunc("/sign_up", signUp)
	mux.HandleFunc("/profile", profile)
	mux.HandleFunc("/feed", feed)
	mux.HandleFunc("/settings", settings)
	mux.HandleFunc("/confirm", confirm)

	log.Fatalln(http.ListenAndServe("localhost:8888", mux))
}

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "camagru"
)

func initDB() *sql.DB {
	driver := "postgres"
	data := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		host, port, user, dbname)

	var err error

	db, err = sql.Open(driver, data)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
