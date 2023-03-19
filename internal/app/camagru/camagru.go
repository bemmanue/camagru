package camagru

import (
	"database/sql"
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
	srv := newServer(store, sessionsStore)

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
