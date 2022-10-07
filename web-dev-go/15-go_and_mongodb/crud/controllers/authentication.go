package controllers

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	auth "github.com/udemy-go/web-dev-go/15-go_and_mongodb/crud/middleware"
	"github.com/udemy-go/web-dev-go/15-go_and_mongodb/crud/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct{}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (ac AuthController) LoginFormValidation(username, password string) map[string][]string {
	v := map[string][]string{}

	if username == "" {
		v["username"] = append(v["username"], "username cannot be empty")
	}

	if password == "" {
		v["password"] = append(v["password"], "password cannot be empty")
	}

	return v
}

func (ac AuthController) RegisterFormValidation(username, password, firstname, lastname, age, role string) map[string][]string {
	v := map[string][]string{}

	if username == "" {
		v["username"] = append(v["username"], "username cannot be empty")
	}

	if password == "" {
		v["password"] = append(v["password"], "password cannot be empty")
	}

	if firstname == "" {
		v["firstname"] = append(v["firstname"], "firstname cannot be empty")
	}

	if lastname == "" {
		v["lastname"] = append(v["lastname"], "lastname cannot be empty")
	}

	if age == "" {
		v["age"] = append(v["age"], "age cannot be empty")
	}

	if role == "" {
		v["role"] = append(v["role"], "role cannot be empty")
	}

	return v
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

	valid := ac.LoginFormValidation(username, password)
	if len(valid) != 0 {
		data := struct{ Wrong string }{Wrong: "form not valid"}
		tpl.ExecuteTemplate(w, "login.html", data)
		return
	}

	user := models.User{
		Username: username,
	}

	err = user.GetUser(bson.M{"username": user.Username}, &user)
	if err == mongo.ErrNoDocuments {
		data := struct{ Wrong string }{Wrong: "wrong username or password"}
		tpl.ExecuteTemplate(w, "login.html", data)
		return
	} else if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil || username != user.Username {
		data := struct{ Wrong string }{Wrong: "wrong username or password"}
		tpl.ExecuteTemplate(w, "login.html", data)
		return
	}

	sId := uuid.New()

	user.Session.SessionId = sId.String()
	user.Session.LastActivity = time.Now().Format(time.Layout)

	err = models.UpdateUserSession(user)
	if err != nil {
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

	valid := ac.RegisterFormValidation(username, password, firstname, lastname, age, role)
	if len(valid) != 0 {
		data := struct{ Wrong string }{Wrong: "form not valid"}
		tpl.ExecuteTemplate(w, "register.html", data)
		return
	}

	ps, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	user := models.User{
		Username: username,
	}

	// Check if username already taken
	uat, err := user.UsernameAlreadyTaken()
	if uat {
		data := struct{ UsernameTaken string }{UsernameTaken: "Username already taken"}
		tpl.ExecuteTemplate(w, "updateUser.html", data)
		return
	} else if err != nil && err != mongo.ErrNoDocuments {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
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

	err = user.InsertUser()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (ac AuthController) Logout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	user, c, err := auth.Authenticate(r)

	if err == http.ErrNoCookie {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	} else if err != nil && err.Error() != "session timeout" {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	c.MaxAge = -1

	user.Session.SessionId = ""
	user.Session.LastActivity = ""

	go func() {
		err = models.UpdateUserSession(user)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}()

	c.Value = ""

	http.SetCookie(w, c)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
