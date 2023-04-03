package camagru

import (
	"database/sql"
	"github.com/bemmanue/camagru/internal/mail/smtp"
	"github.com/bemmanue/camagru/internal/store/sqlstore"
	"github.com/gin-contrib/sessions/cookie"
)

func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}

	defer db.Close()
	store := sqlstore.New(db)
	sessionsStore := cookie.NewStore([]byte("secret"))
	mail := smtp.New(
		"smtp.mail.ru:587",
		"olivia2804@mail.ru",
		"cSafGhJ4DceKJBZ8h76Y",
		"smtp.mail.ru",
	)
	srv := newServer(store, sessionsStore, mail)

	return srv.router.Run(config.BindAddr)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
