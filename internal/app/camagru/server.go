package camagru

import (
	"encoding/json"
	"errors"
	"github.com/bemmanue/camagru/internal/model"
	"github.com/bemmanue/camagru/internal/store"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
)

var (
	errIncorrectEmailOrPassword = errors.New("incorrect email or password")
)

type server struct {
	router *http.ServeMux
	logger *logrus.Logger
	store  store.Store
}

// newServer ...
func newServer(store store.Store) *server {
	s := &server{
		router: http.NewServeMux(),
		logger: logrus.New(),
		store:  store,
	}

	s.configureRouter()

	return s
}

// serveHTTP ...
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// configureRouter ...
func (s *server) configureRouter() {

	fs := http.FileServer(http.Dir("web/"))
	handler := http.StripPrefix("/web/", fs)
	s.router.Handle("/web/", handler)

	s.router.HandleFunc("/", s.handleIndex())
	s.router.HandleFunc("/sign_up", s.handleSignUp())
	s.router.HandleFunc("/confirm", s.handleConfirm())
	s.router.HandleFunc("/sign_in", s.handleSignIn())
}

func (s *server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/index.html")
	}
}

func (s *server) handleSignUp() http.HandlerFunc {
	type request struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			http.ServeFile(w, r, "./web/templates/sign_up.html")
		case http.MethodPost:
			req := &request{}
			if err := json.NewDecoder(r.Body).Decode(req); err != nil {
				s.error(w, r, http.StatusBadRequest, err)
				return
			}

			u := &model.User{
				Username: req.Username,
				Email:    req.Email,
				Password: req.Password,
			}

			if err := s.store.User().Create(u); err != nil {
				log.Println(err)
				s.error(w, r, http.StatusUnprocessableEntity, err)
				return
			}

			u.Sanitize()
			s.respond(w, r, http.StatusCreated, u)
		}
	}
}

func (s *server) handleSignIn() http.HandlerFunc {
	type request struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			http.ServeFile(w, r, "./web/templates/sign_in.html")
		case http.MethodPost:
			req := &request{}
			if err := json.NewDecoder(r.Body).Decode(req); err != nil {
				s.error(w, r, http.StatusBadRequest, err)
				return
			}

			u, err := s.store.User().FindByUsername(req.Username)
			if err != nil || !u.ComparePassword(req.Password) {
				log.Println(err, req.Password)
				s.error(w, r, http.StatusUnauthorized, errIncorrectEmailOrPassword)
				return
			}

			s.respond(w, r, http.StatusOK, nil)
			//http.Redirect(w, r, "/profile", http.StatusFound)
		}
	}
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func (s *server) handleConfirm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/confirm.html")
	}
}

//	func profile(w http.ResponseWriter, r *http.Request) {
//		http.ServeFile(w, r, "./web/templates/profile.html")
//	}
//
//	func feed(w http.ResponseWriter, r *http.Request) {
//		http.ServeFile(w, r, "./web/templates/feed.html")
//	}
//
//	func settings(w http.ResponseWriter, r *http.Request) {
//		http.ServeFile(w, r, "./web/templates/settings.html")
//	}
