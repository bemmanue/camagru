package camagru

import (
	"database/sql"
	"fmt"
	"github.com/bemmanue/camagru/internal/mail/smtp"
	"github.com/bemmanue/camagru/internal/store/sqlstore"
	"github.com/gin-contrib/sessions/cookie"
)

func Start(config *Config) error {
	db, err := newDB(*config.Database)
	if err != nil {
		return err
	}
	defer db.Close()

	store := sqlstore.New(db)

	sessionsStore := cookie.NewStore([]byte(config.SessionKey))

	mail := smtp.New(
		config.SMTP.From,
		config.SMTP.Password,
		config.SMTP.Host,
		config.SMTP.Port,
	)

	srv := newServer(store, sessionsStore, mail)

	return srv.router.Run(config.BindAddr)
}

func newDB(config DatabaseConfig) (*sql.DB, error) {
	databaseURL := fmt.Sprintf(
		"host=localhost port=%d user=%s dbname=%s sslmode=%s",
		config.Port,
		config.User,
		config.Name,
		config.SSLMode,
	)

	// for database container
	//databaseURL := fmt.Sprintf(
	//	"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
	//	config.Host,
	//	config.Port,
	//	config.User,
	//	config.Password,
	//	config.Name,
	//	config.SSLMode,
	//)
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
