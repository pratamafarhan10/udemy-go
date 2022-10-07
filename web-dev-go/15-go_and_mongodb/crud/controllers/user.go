package controllers

import (
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
	auth "github.com/udemy-go/web-dev-go/15-go_and_mongodb/crud/middleware"
	"github.com/udemy-go/web-dev-go/15-go_and_mongodb/crud/models"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (uc UserController) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	user, _, err := auth.Authenticate(r)

	if err == http.ErrNoCookie {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	} else if err != nil {
		if err.Error() == "session timeout" {
			http.Redirect(w, r, "/logout", http.StatusSeeOther)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	tpl.ExecuteTemplate(w, "index.html", user)
}

func (uc UserController) UpdateView(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	user, _, err := auth.Authenticate(r)

	if err == http.ErrNoCookie {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	} else if err != nil {
		if err.Error() == "session timeout" {
			http.Redirect(w, r, "/logout", http.StatusSeeOther)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	tpl.ExecuteTemplate(w, "updateUser.html", user)
}

func (uc UserController) Update(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	user, _, err := auth.Authenticate(r)

	if err == http.ErrNoCookie {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	} else if err != nil {
		if err.Error() == "session timeout" {
			http.Redirect(w, r, "/logout", http.StatusSeeOther)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	userReq := models.User{
		Username:  r.FormValue("username"),
		Password:  r.FormValue("password"),
		FirstName: r.FormValue("firstname"),
		LastName:  r.FormValue("lastname"),
		Age:       r.FormValue("age"),
		Role:      r.FormValue("role"),
	}

	// Check if username already taken
	if user.Username != userReq.Username {
		uat, err := userReq.UsernameAlreadyTaken()
		if uat {
			data := struct{ UsernameTaken string }{UsernameTaken: "Username already taken"}
			tpl.ExecuteTemplate(w, "updateUser.html", data)
			return
		} else if err != nil && err != mongo.ErrNoDocuments {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}

	// Chcek password
	oldPassword := r.FormValue("oldpassword")
	passwordConfirm := r.FormValue("passwordconfirmation")
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword))
	if passwordConfirm != userReq.Password {
		data := struct{ Wrong string }{Wrong: "password dont match"}
		tpl.ExecuteTemplate(w, "updateUser.html", data)
		return
	} else if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Save file to server
	f, h, err := r.FormFile("profilepicture")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer f.Close()

	fileType := strings.Split(h.Header["Content-Type"][0], "/")
	filename := `/assets/` + userReq.Username + `-` + time.Now().Format("020106030405") + fileType[1]
	nf, err := os.Create(`.` + filename)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer nf.Close()
	io.Copy(nf, f)

	// Hash password
	hp, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Insert data do Database
	userReq.Id = user.Id
	userReq.Password = string(hp)
	userReq.ProfilePicture = filename

	err = userReq.UpdateUser()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (uc UserController) AdminPage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	user, _, err := auth.Authenticate(r)

	if err == http.ErrNoCookie {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	} else if err != nil {
		if err.Error() == "session timeout" {
			http.Redirect(w, r, "/logout", http.StatusSeeOther)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	tpl.ExecuteTemplate(w, "admin.html", user)
}

func (uc UserController) UserPage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	user, _, err := auth.Authenticate(r)

	if err == http.ErrNoCookie {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	} else if err != nil {
		if err.Error() == "session timeout" {
			http.Redirect(w, r, "/logout", http.StatusSeeOther)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	tpl.ExecuteTemplate(w, "defaultUser.html", user)
}
