package camagru

import (
	"errors"
	"github.com/bemmanue/camagru/internal/model"
	"github.com/bemmanue/camagru/internal/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	errIncorrectEmailOrPassword = errors.New("incorrect email or password")
)

type server struct {
	router *gin.Engine
	store  store.Store
}

// newServer ...
func newServer(store store.Store) *server {
	s := &server{
		router: gin.Default(),
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

	s.router.Static("/web", "./web")

	s.router.GET("/", s.getIndex)
	s.router.GET("/sign_in", s.getSignIn)
	s.router.GET("/sign_up", s.getSignUp)
	s.router.GET("/profile", s.getProfile)

	s.router.POST("/sign_in", s.postSignIn)
	s.router.POST("/sign_up", s.postSignUp)
}

func (s *server) getIndex(c *gin.Context) {
	c.File("./web/templates/index.html")
}

func (s *server) getSignIn(c *gin.Context) {
	c.File("./web/templates/sign_in.html")
}

func (s *server) getSignUp(c *gin.Context) {
	c.File("./web/templates/sign_up.html")
}

func (s *server) postSignIn(c *gin.Context) {
	type request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var form request

	err := c.BindJSON(&form)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := s.store.User().FindByUsername(form.Username)
	if err != nil || !u.ComparePassword(form.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": errIncorrectEmailOrPassword})
		return
	}

	c.JSON(http.StatusOK, gin.H{"username": form.Username, "password": form.Password})
}

func (s *server) postSignUp(c *gin.Context) {
	type request struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var form request

	err := c.BindJSON(&form)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := &model.User{
		Username: form.Username,
		Email:    form.Email,
		Password: form.Password,
	}

	if err := s.store.User().Create(u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"username": form.Username, "email": form.Email, "password": form.Password})
}

func (s *server) getProfile(c *gin.Context) {
	c.File("./web/templates/profile.html")
}
