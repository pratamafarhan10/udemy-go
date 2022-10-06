package controllers

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/udemy-go/web-dev-go/15-go_and_mongodb/crud/models"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (uc UserController) GetUserBySessionId(c *http.Cookie) (models.User, error) {
	user := models.User{}
	user.Session.SessionId = c.Value

	err := models.GetUserBySessionId(user, &user)
	return user, err
}

func (uc UserController) UpdateLastActivity(sessionId string) error {
	err := models.UpdateUserSessionBySessionId(sessionId, time.Now().Format(time.Layout))

	return err
}

func (uc UserController) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Check cookie
	c, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Get user
	user, err := uc.GetUserBySessionId(c)
	if err != nil {
		log.Println("INDEX ERROR", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Update last activity
	err = uc.UpdateLastActivity(c.Value)
	if err != nil {
		log.Println("FAILED UPDATE LAST ACTIVITY", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	tpl.ExecuteTemplate(w, "index.html", user)
}

func (uc UserController) UpdateView(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Check cookie
	c, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Get user
	user, err := uc.GetUserBySessionId(c)
	if err != nil {
		log.Println("INDEX ERROR", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Update last activity
	err = uc.UpdateLastActivity(c.Value)
	if err != nil {
		log.Println("FAILED UPDATE LAST ACTIVITY", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	tpl.ExecuteTemplate(w, "updateUser.html", user)
}

func (uc UserController) Update(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Check cookie
	c, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Get user
	user, err := uc.GetUserBySessionId(c)
	if err != nil {
		log.Println("INDEX ERROR", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Update last activity
	err = uc.UpdateLastActivity(c.Value)
	if err != nil {
		log.Println("FAILED UPDATE LAST ACTIVITY", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Check if username already taken
	username := r.FormValue("username")
	if user.Username != username {
		container := models.User{
			Username: username,
		}

		err = models.GetUserByUsername(container, &container)
		if err != mongo.ErrNoDocuments {
			log.Println("USERNAME ALREADY TAKEN", err)
			data := struct{ UsernameTaken string }{UsernameTaken: "Username already taken"}
			tpl.ExecuteTemplate(w, "updateUser.html", data)
			return
		}
	}

	// Chcek password
	password := r.FormValue("password")
	oldPassword := r.FormValue("oldpassword")
	passwordConfirm := r.FormValue("passwordconfirmation")
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword))
	if passwordConfirm != password {
		log.Println("PASSWORD NOT EQUAL", err)
		data := struct{ Wrong string }{Wrong: "password dont match"}
		tpl.ExecuteTemplate(w, "updateUser.html", data)
		return
	} else if err != nil {
		log.Println("ERROR COMPARING PASSWORD IN UPDATE USER CONTROLLER", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	firstname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")
	age := r.FormValue("age")
	role := r.FormValue("role")

	f, h, err := r.FormFile("profilepicture")
	if err != nil {
		log.Println("PROFILE PICTURE", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer f.Close()

	// Hash password
	hp, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("ERROR GENERATING PASSWORD IN UPDATE USER CONTROLLER", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Save file
	fileType := strings.Split(h.Header["Content-Type"][0], "/")
	filename := `/assets/` + username + `-` + time.Now().Format("020106030405") + fileType[1]
	nf, err := os.Create(`.` + filename)
	if err != nil {
		log.Println("ERROR CREATING A NEW FILE", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer nf.Close()
	io.Copy(nf, f)

	// Insert data do Database
	user.Username = username
	user.Password = string(hp)
	user.FirstName = firstname
	user.LastName = lastname
	user.Age = age
	user.Role = role
	user.ProfilePicture = filename

	err = models.UpdateUser(user)
	if err != nil {
		log.Println("ERROR UPDATING USER", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
