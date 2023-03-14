package main

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/bemmanue/camagru/internal/app/camagru"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"net/mail"
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

func validateRegistryData(login, email, password, passwordConfirm string) (bool, error) {
	// check login
	if len(login) < 6 {
		return false, errors.New("invalid login")
	}

	// check email
	_, err := mail.ParseAddress(email)
	if err != nil {
		return false, errors.New("invalid email")
	}

	// check password
	if len(password) < 6 {
		return false, errors.New("invalid password")
	}

	// check password confirm
	if passwordConfirm != password {
		return false, errors.New("wrong password confirm")
	}

	return true, nil
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

		login := r.Form.Get("login")
		email := r.Form.Get("email")
		password := r.Form.Get("password")
		passwordConfirm := r.Form.Get("password_confirm")

		_, err := validateRegistryData(login, email, password, passwordConfirm)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		query := fmt.Sprintf("insert into users(login, email, password) values ('%s', '%s', '%s');",
			login, email, password)

		_, err = db.Exec(query)
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

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "config/camagru.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := camagru.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatalln(err)
	}

	server := camagru.New(config)
	if err := server.Start(); err != nil {
		log.Fatalln(err)
	}

	//db = initDB()
	//defer db.Close()
}

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "camagru"
)

func initDB() *sql.DB {
	var err error

	driver := "postgres"
	data := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		host, port, user, dbname)

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
