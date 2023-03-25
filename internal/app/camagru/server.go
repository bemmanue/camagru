package camagru

import (
	"errors"
	"github.com/bemmanue/camagru/internal/model"
	"github.com/bemmanue/camagru/internal/store"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"path/filepath"
	"time"
)

const (
	sessionName = "camagru"
	imagesPath  = "data/"
)

var (
	errIncorrectEmailOrPassword = errors.New("incorrect email or password")
	errUnauthorized             = errors.New("unauthorized")
)

type server struct {
	router       *gin.Engine
	store        store.Store
	sessionStore sessions.Store
}

// newServer ...
func newServer(store store.Store, sessionStore sessions.Store) *server {
	s := &server{
		router:       gin.Default(),
		store:        store,
		sessionStore: sessionStore,
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
	s.router.MaxMultipartMemory = 8 << 20

	s.router.Use(sessions.Sessions(sessionName, s.sessionStore))
	s.router.LoadHTMLGlob("web/templates/*")

	s.router.Static("/web", "./web")
	s.router.Static("/data", "./data")

	s.router.GET("/", s.getIndex)
	s.router.GET("/sign_in", s.getSignIn)
	s.router.GET("/sign_up", s.getSignUp)

	s.router.POST("/sign_in", s.postSignIn)
	s.router.POST("/sign_up", s.postSignUp)

	authorized := s.router.Group("")
	authorized.Use(AuthenticateUser())
	{
		authorized.GET("/feed", s.getFeed)
		authorized.GET("/new", s.getNew)
		authorized.GET("/profile", s.getProfile)
		authorized.GET("/settings", s.getSettings)

		authorized.POST("/new", s.postNew)
		authorized.POST("/like", s.postLike)
	}
}

func AuthenticateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionID := session.Get("user_id")
		if sessionID == nil {
			c.HTML(http.StatusUnauthorized, "error.html", gin.H{"error": errUnauthorized.Error()})
			c.Abort()
		}

		c.Set("user_id", sessionID.(int))
		c.Next()
	}
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
		c.JSON(http.StatusUnauthorized, gin.H{"error": errIncorrectEmailOrPassword.Error()})
		return
	}

	session := sessions.Default(c)
	session.Set("user_id", u.ID)
	if err = session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user_id": session.Get("user_id")})
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

func (s *server) postNew(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "get form err: %s", err.Error())
		return
	}

	uu, err := uuid.NewUUID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
		return
	}

	name := uu.String()
	extension := filepath.Ext(file.Filename)
	path := imagesPath + name + extension

	if err := c.SaveUploadedFile(file, path); err != nil {
		c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
		return
	}

	userId, ok := c.Get("user_id")
	if ok == false {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "no user id"})
		return
	}

	i := &model.Image{
		Name:       name,
		Extension:  extension,
		Path:       path,
		UploadTime: time.Now(),
		UserID:     userId.(int),
	}

	if err := s.store.Image().Create(i); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
		return
	}

	c.Header("Location", path)
	c.JSON(http.StatusCreated, gin.H{"status": "uploaded"})
}

func (s *server) getFeed(c *gin.Context) {
	userId, ok := c.Get("user_id")
	if ok == false {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "no user id"})
		return
	}

	images, err := s.store.Image().GetPostData(userId.(int))

	//images, err := s.store.Image().SelectImages()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "feed.html", gin.H{"Images": images})
}

func (s *server) getNew(c *gin.Context) {
	c.File("./web/templates/new.html")
}

func (s *server) getProfile(c *gin.Context) {
	userId, ok := c.Get("user_id")
	if ok == false {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "no user id"})
		return
	}

	images, err := s.store.Image().SelectUserImages(userId.(int))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "profile.html", gin.H{"Images": images})
}

func (s *server) getSettings(c *gin.Context) {
	c.File("./web/templates/settings.html")
}

func (s *server) postLike(c *gin.Context) {
	type request struct {
		PostID int `json:"post_id"`
	}

	var form request

	err := c.BindJSON(&form)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, ok := c.Get("user_id")
	if ok == false {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "no user id"})
		return
	}

	like, err := s.store.Like().Find(form.PostID, userId.(int))
	if err != nil {
		err := s.store.Like().Create(&model.Like{ImageID: form.PostID, UserID: userId.(int)})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
			return
		}
	} else {
		err := s.store.Like().Delete(like)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
