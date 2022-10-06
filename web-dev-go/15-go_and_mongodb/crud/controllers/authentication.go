package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/udemy-go/web-dev-go/15-go_and_mongodb/crud/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct{}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (ac AuthController) LoginView(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, err := r.Cookie("session")
	if err == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "login.html", nil)
}

func (ac AuthController) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, err := r.Cookie("session")
	if err == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	user := models.User{
		Username: username,
	}

	err = models.GetUserByUsername(user, &user)
	if err == mongo.ErrNoDocuments {
		log.Println("LOGIN USERNAME NOT FOUND", err)
		data := struct{ Wrong string }{Wrong: "wrong username or password"}
		tpl.ExecuteTemplate(w, "login.html", data)
		return
	} else if err != nil {
		log.Println("LOGIN ERROR", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil || username != user.Username {
		log.Println("LOGIN USERNAME NOT FOUND", err)
		data := struct{ Wrong string }{Wrong: "wrong username or password"}
		tpl.ExecuteTemplate(w, "login.html", data)
		return
	}

	sId := uuid.New()

	user.Session.SessionId = sId.String()
	user.Session.LastActivity = time.Now().Format(time.Layout)

	err = models.UpdateUserSession(user)
	if err != nil {
		log.Println("ERROR UPDATE USER SESSION IN LOGIN", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	cookie := &http.Cookie{
		Name:  "session",
		Value: sId.String(),
	}

	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (ac AuthController) RegisterView(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, err := r.Cookie("session")
	if err == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "register.html", nil)
}

func (ac AuthController) Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, err := r.Cookie("session")
	if err == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	firstname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")
	age := r.FormValue("age")
	role := r.FormValue("role")

	ps, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	user := models.User{
		Username: username,
	}

	err = models.GetUserByUsername(user, &user)
	if err != mongo.ErrNoDocuments {
		log.Println("USERNAME ALREADY TAKEN", err)
		data := struct{ UsernameTaken string }{UsernameTaken: "Username already taken"}
		tpl.ExecuteTemplate(w, "register.html", data)
		return
	}

	user = models.User{
		Id:        primitive.NewObjectID(),
		Username:  username,
		Password:  string(ps),
		FirstName: firstname,
		LastName:  lastname,
		Age:       age,
		Role:      role,
	}

	err = models.InsertUser(user)
	if err != nil {
		log.Println("ERROR INSERT USER IN REGISTER", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
