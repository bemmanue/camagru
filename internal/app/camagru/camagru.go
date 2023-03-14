package camagru

import (
	"fmt"
	"github.com/bemmanue/camagru/internal/store"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Camagru struct {
	config *Config
	logger *logrus.Logger
	router *http.ServeMux
	store  *store.Store
}

func New(config *Config) *Camagru {
	return &Camagru{
		config: config,
		logger: logrus.New(),
		router: http.NewServeMux(),
	}
}

func (c *Camagru) Start() error {
	if err := c.configureLogger(); err != nil {
		return err
	}

	c.configureRouter()

	if err := c.configureStore(); err != nil {
		return err
	}

	c.logger.Info("starting camagru server")

	return http.ListenAndServe(c.config.BindAddr, c.router)
}

func (c *Camagru) configureLogger() error {
	level, err := logrus.ParseLevel(c.config.LogLevel)
	if err != nil {
		return err
	}

	c.logger.SetLevel(level)

	return nil
}

func (c *Camagru) configureRouter() {
	// serve static files
	fs := http.FileServer(http.Dir("web"))
	handler := http.StripPrefix("/web/", fs)
	c.router.Handle("/web/", handler)

	c.router.HandleFunc("/", c.HandleIndex())
	//c.router.HandleFunc("/sign_in", signIn)
	//c.router.HandleFunc("/sign_up", signUp)
	//c.router.HandleFunc("/profile", profile)
	//c.router.HandleFunc("/feed", feed)
	//c.router.HandleFunc("/settings", settings)
	//c.router.HandleFunc("/confirm", confirm)
}

func (c *Camagru) HandleIndex() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello")
	}
}

func (c *Camagru) configureStore() error {
	st := store.New(c.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	c.store = st

	return nil
}
