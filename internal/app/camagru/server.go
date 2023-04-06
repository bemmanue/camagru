package camagru

import (
	"github.com/bemmanue/camagru/internal/mail"
	"github.com/bemmanue/camagru/internal/model"
	"github.com/bemmanue/camagru/internal/store"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const (
	sessionName = "camagru"
	imagesPath  = "data/"
)

type server struct {
	router       *gin.Engine
	store        store.Store
	sessionStore sessions.Store
	mail         mail.Mail
}

// newServer ...
func newServer(store store.Store, sessionStore sessions.Store, mail mail.Mail) *server {
	s := &server{
		router:       gin.Default(),
		store:        store,
		sessionStore: sessionStore,
		mail:         mail,
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
	s.router.GET("/confirm", s.getConfirm)
	s.router.GET("/verify", s.getVerify)

	s.router.POST("/sign_in", s.postSignIn)
	s.router.POST("/sign_up", s.postSignUp)

	authorized := s.router.Group("")
	authorized.Use(AuthenticateUser())
	{
		authorized.GET("/feed", s.getFeed)
		authorized.GET("/new", s.getNew)
		authorized.GET("/profile", s.getProfile)
		authorized.GET("/settings", s.getSettings)
		authorized.GET("/logout", s.getLogout)

		authorized.POST("/new", s.postNew)
		authorized.POST("/comment", s.postComment)
		authorized.POST("/like", s.postLike)
		authorized.POST("/settings", s.postSettings)

		authorized.DELETE("/profile/:post_id", s.deletePost)
	}

	s.router.NoRoute(s.noRoute)
}

// AuthenticateUser ...
func AuthenticateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionID := session.Get("user_id")
		if sessionID == nil {
			c.HTML(http.StatusUnauthorized, "status.html", gin.H{
				"Status":       http.StatusUnauthorized,
				"ReasonPhrase": "Unauthorized",
			})
			c.Abort()
			return
		}

		c.Set("user_id", sessionID.(int))
		c.Next()
	}
}

// getIndex ...
func (s *server) getIndex(c *gin.Context) {
	session := sessions.Default(c)
	sessionID := session.Get("user_id")

	if sessionID == nil {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Link1": "/sign_in", "LinkName1": "Sign in",
			"Link2": "/sign_up", "LinkName2": "Sign up",
		})
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Link1": "/profile", "LinkName1": "Profile",
			"Link2": "/logout", "LinkName2": "Log out",
		})
	}
}

// getSignIn ...
func (s *server) getSignIn(c *gin.Context) {
	session := sessions.Default(c)

	session.Clear()

	err := session.Save()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "status.html", gin.H{
			"Status":       http.StatusInternalServerError,
			"ReasonPhrase": "Internal Server Error",
		})
		return
	}

	c.File("./web/templates/sign_in.html")
}

// getSignUp ...
func (s *server) getSignUp(c *gin.Context) {
	session := sessions.Default(c)

	session.Clear()

	err := session.Save()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "status.html", gin.H{
			"Status":       http.StatusInternalServerError,
			"ReasonPhrase": "Internal Server Error",
		})
		return
	}

	c.File("./web/templates/sign_up.html")
}

// getLogout ...
func (s *server) getLogout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()

	err := session.Save()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "status.html", gin.H{
			"Status":       http.StatusInternalServerError,
			"ReasonPhrase": "Internal Server Error",
		})
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"Link1": "/sign_in", "LinkName1": "Sign in",
		"Link2": "/sign_up", "LinkName2": "Sign up",
	})
}

// getConfirm ...
func (s *server) getConfirm(c *gin.Context) {
	address := c.DefaultQuery("email", "your address")

	c.HTML(http.StatusOK, "confirm.html", gin.H{
		"Address": address,
	})
}

// getVerify ...
func (s *server) getVerify(c *gin.Context) {
	email := c.Query("email")
	codeStr := c.Query("code")

	code, err := strconv.Atoi(codeStr)
	if err != nil {
		c.HTML(http.StatusUnprocessableEntity, "status.html", gin.H{
			"Status":       http.StatusUnprocessableEntity,
			"ReasonPhrase": "Unprocessable Entity",
		})
		return
	}

	v, err := s.store.Verify().FindByEmail(strings.ToLower(email))
	if err != nil {
		c.HTML(http.StatusUnprocessableEntity, "status.html", gin.H{
			"Status":       http.StatusUnprocessableEntity,
			"ReasonPhrase": "Unprocessable Entity",
		})
		return
	}

	if code != v.Code {
		c.HTML(http.StatusUnprocessableEntity, "status.html", gin.H{
			"Status":       http.StatusUnprocessableEntity,
			"ReasonPhrase": "Unprocessable Entity",
		})
		return
	}

	// update email status
	if err := s.store.User().VerifyEmail(v.Email); err != nil {
		c.HTML(http.StatusInternalServerError, "status.html", gin.H{
			"Status":       http.StatusInternalServerError,
			"ReasonPhrase": "Internal Server Error",
		})
		return
	}

	c.File("./web/templates/email_verified.html")
}

// postSignIn ...
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

	u, err := s.store.User().FindByUsernameVerified(strings.ToLower(form.Username))
	if err != nil || !u.ComparePassword(form.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "wrong username or password"})
		return
	}

	session := sessions.Default(c)
	session.Set("user_id", u.ID)
	if err = session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ok": "you signed in"})
}

// postSignUp ...
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
		Username:      strings.ToLower(form.Username),
		Email:         strings.ToLower(form.Email),
		Password:      form.Password,
		EmailVerified: false,
	}

	if err := u.Validate(); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	// check username
	exists, err := s.store.User().UsernameExists(u.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else if exists == true {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Username is taken by another account."})
		return
	}

	// check email
	exists, err = s.store.User().EmailExists(u.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else if exists == true {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Email is taken by another account."})
		return
	}

	// create user
	if err := s.store.User().Create(u); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// generate verification code and store it to database
	v := model.VerifyCode{
		Email:  u.Email,
		Code:   rand.Intn(1000000),
		UserID: u.ID,
	}

	if err := s.store.Verify().Create(&v); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// send verification letter
	if err := s.mail.Verify(v.Email, strconv.Itoa(v.Code)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"ok": "email verification needed"})
}

// postNew ...
func (s *server) postNew(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uu, err := uuid.NewUUID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	name := uu.String()
	extension := filepath.Ext(file.Filename)
	path := imagesPath + name + extension

	if err := c.SaveUploadedFile(file, path); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userId, ok := c.Get("user_id")
	if ok == false {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	p := &model.Post{
		ImageID:      i.ID,
		AuthorID:     userId.(int),
		CreationTime: time.Now(),
	}

	if err := s.store.Post().Create(p); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Header("Location", path)
	c.JSON(http.StatusCreated, gin.H{"ok": "created"})
}

// getFeed ...
func (s *server) getFeed(c *gin.Context) {
	userId, ok := c.Get("user_id")
	if ok == false {
		c.HTML(http.StatusInternalServerError, "status.html", gin.H{
			"Status":       http.StatusInternalServerError,
			"ReasonPhrase": "Internal Server Error",
		})
		return
	}

	page := c.DefaultQuery("page", "1")

	pageNum, err := strconv.Atoi(page)
	if err != nil || pageNum < 1 {
		c.HTML(http.StatusNotFound, "status.html", gin.H{
			"Status":       http.StatusNotFound,
			"ReasonPhrase": "Not Found",
		})
		return
	}

	maxPageCount, err := s.store.Post().GetPageCount()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "status.html", gin.H{
			"Status":       http.StatusInternalServerError,
			"ReasonPhrase": "Internal Server Error",
		})
		return
	}

	if pageNum > maxPageCount {
		c.HTML(http.StatusNotFound, "status.html", gin.H{
			"Status":       http.StatusNotFound,
			"ReasonPhrase": "Not Found",
		})
		return
	}

	posts, err := s.store.Post().GetPage(pageNum, userId.(int))
	if err != nil {
		c.HTML(http.StatusInternalServerError, "status.html", gin.H{
			"Status":       http.StatusInternalServerError,
			"ReasonPhrase": "Internal Server Error",
		})
		return
	}

	// Calculate neighbour pages
	PreviousPage := pageNum - 1
	NextPage := pageNum + 1
	if NextPage > maxPageCount {
		NextPage = 0
	}

	c.HTML(http.StatusOK, "feed.html", gin.H{
		"Posts":        posts,
		"PreviousPage": PreviousPage,
		"NextPage":     NextPage,
	})
}

// getNew ...
func (s *server) getNew(c *gin.Context) {
	c.File("./web/templates/new.html")
}

// getProfile ...
func (s *server) getProfile(c *gin.Context) {
	userId, ok := c.Get("user_id")
	if ok == false {
		c.HTML(http.StatusInternalServerError, "status.html", gin.H{
			"Status":       http.StatusInternalServerError,
			"ReasonPhrase": "Internal Server Error",
		})
		return
	}

	page := c.DefaultQuery("page", "1")
	pageNum, err := strconv.Atoi(page)
	if err != nil || pageNum < 1 {
		c.HTML(http.StatusNotFound, "status.html", gin.H{
			"Status":       http.StatusNotFound,
			"ReasonPhrase": "Not Found",
		})
		return
	}

	maxPageCount, err := s.store.Post().GetUserPageCount(userId.(int))
	if err != nil {
		c.HTML(http.StatusInternalServerError, "status.html", gin.H{
			"Status":       http.StatusInternalServerError,
			"ReasonPhrase": "Internal Server Error",
		})
		return
	}

	if pageNum > maxPageCount {
		c.HTML(http.StatusNotFound, "status.html", gin.H{
			"Status":       http.StatusNotFound,
			"ReasonPhrase": "Not Found",
		})
		return
	}

	posts, err := s.store.Post().GetUserPage(pageNum, userId.(int))
	if err != nil {
		c.HTML(http.StatusNotFound, "status.html", gin.H{
			"Status":       http.StatusNotFound,
			"ReasonPhrase": "Not Found",
		})
		return
	}

	PreviousPage := pageNum - 1
	NextPage := pageNum + 1
	if NextPage > maxPageCount {
		NextPage = 0
	}

	c.HTML(http.StatusOK, "profile.html", gin.H{
		"Posts":        posts,
		"PreviousPage": PreviousPage,
		"NextPage":     NextPage,
	})
}

// getSettings ...
func (s *server) getSettings(c *gin.Context) {
	// get user id
	userID, ok := c.Get("user_id")
	if ok == false {
		c.HTML(http.StatusInternalServerError, "status.html", gin.H{
			"Status":       http.StatusNotFound,
			"ReasonPhrase": "Not Found",
		})
		return
	}

	u, err := s.store.User().Find(userID.(int))
	if err != nil {
		c.HTML(http.StatusInternalServerError, "status.html", gin.H{
			"Status":       http.StatusNotFound,
			"ReasonPhrase": "Not Found",
		})
	}

	c.HTML(http.StatusOK, "settings.html", gin.H{
		"LikeNotify":    u.LikeNotify,
		"CommentNotify": u.CommentNotify,
	})
}

// postComment ...
func (s *server) postComment(c *gin.Context) {
	type request struct {
		PostID  int    `json:"post_id"`
		Comment string `json:"comment"`
	}

	var form request

	err := c.BindJSON(&form)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// get comment author id
	userId, ok := c.Get("user_id")
	if ok == false {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "no user id"})
		return
	}

	// get commented post
	p, err := s.store.Post().Find(form.PostID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := s.store.Comment().Create(&model.Comment{
		AuthorID:     userId.(int),
		PostID:       form.PostID,
		CommentText:  form.Comment,
		CreationTime: time.Now(),
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
		return
	}

	// find comment author
	commentAuthor, err := s.store.User().Find(p.AuthorID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
		return
	}

	// find post author
	u, err := s.store.User().Find(p.AuthorID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
		return
	}

	// send notification letter
	if u.CommentNotify == true {
		if err := s.mail.CommentNotify(u.Email, commentAuthor.Username); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// postLike ...
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

	// get like author id
	userId, ok := c.Get("user_id")
	if ok == false {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "no user id"})
		return
	}

	// get like author
	p, err := s.store.Post().Find(form.PostID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	like, err := s.store.Like().Find(form.PostID, userId.(int))
	if err != nil {
		if err := s.store.Like().Create(&model.Like{
			PostID: form.PostID,
			UserID: userId.(int),
		}); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
			return
		}
	} else {
		if err := s.store.Like().Delete(like); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
			return
		}
	}

	// find like author
	likeAuthor, err := s.store.User().Find(p.AuthorID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
		return
	}

	// find post author
	u, err := s.store.User().Find(p.AuthorID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
		return
	}

	// send notification letter
	if u.LikeNotify == true {
		if err := s.mail.LikeNotify(u.Email, likeAuthor.Username); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (s *server) noRoute(c *gin.Context) {
	c.HTML(http.StatusNotFound, "status.html", gin.H{
		"Status":       http.StatusNotFound,
		"ReasonPhrase": "Not Found",
	})
}

func (s *server) postSettings(c *gin.Context) {
	likeNotify := c.Query("like_notify")
	commentNotify := c.Query("comment_notify")

	// get user id
	userId, ok := c.Get("user_id")
	if ok == false {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "no user id"})
		return
	}

	value, err := strconv.ParseBool(likeNotify)
	if err != nil {
		err := s.store.User().UpdateLikeNotify(userId.(int), value)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}
	}

	value, err = strconv.ParseBool(commentNotify)
	if err != nil {
		err := s.store.User().UpdateCommentNotify(userId.(int), value)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"ok": "ok"})
}

func (s *server) deletePost(c *gin.Context) {
	post_id := c.Param("post_id")

	postID, err := strconv.Atoi(post_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong post identifier"})
		return
	}

	// find post data
	p, err := s.store.Post().Find(postID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong post identifier"})
		return
	}

	// fund user image data
	i, err := s.store.Image().Find(p.ImageID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	// remove user image
	if err := os.RemoveAll(i.Path); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	// remove records from post database
	if err := s.store.Post().Delete(postID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	// remove comments to related post
	if err := s.store.Comment().DeleteByPostID(postID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	// remove likes to related post
	if err := s.store.Like().DeleteByPostID(postID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"ok": "no content"})
}
