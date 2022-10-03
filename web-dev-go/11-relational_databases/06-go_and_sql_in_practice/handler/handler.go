package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/udemy-go/web-dev-go/11-relational_databases/06-go_and_sql_in_practice/model"
	"golang.org/x/crypto/bcrypt"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func Index(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}

	ok := StillAuthenticated(c.Value)

	if !ok {
		http.Redirect(w, req, "/logout", http.StatusSeeOther)
		return
	}

	user, err := GetUser(w, c)
	if err != nil {
		log.Println(err)
	}

	tpl.ExecuteTemplate(w, "index.html", user)
}

func Register(w http.ResponseWriter, req *http.Request) {
	_, err := req.Cookie("session")

	// Check if the user already logged in
	if err == nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	if req.Method == http.MethodPost {
		u := req.FormValue("username")
		p := req.FormValue("password")
		fn := req.FormValue("firstname")
		ln := req.FormValue("lastname")
		age := req.FormValue("age")
		role := req.FormValue("role")

		i, err := strconv.Atoi(age)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		hp, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		d := model.User{
			Username:  u,
			Password:  string(hp),
			FirstName: fn,
			LastName:  ln,
			Role:      role,
			Age:       i,
		}

		err = model.CreateUser(d)

		if err != nil {
			log.Println(err)
			return
		}

		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "register.html", nil)
}

func Login(w http.ResponseWriter, req *http.Request) {
	_, err := req.Cookie("session")

	// Check if the user already logged in
	if err == nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	if req.Method == http.MethodPost {
		u := req.FormValue("username")
		p := req.FormValue("password")

		user, err := model.GetUserByUsername(u)

		if err != nil || (model.User{}) == user {
			fmt.Println("error disini")
			tpl.ExecuteTemplate(w, "login.html", "wrong")
			return
		}

		fmt.Println(user)

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(p))

		if u != user.Username || err != nil {
			fmt.Println("error disana")
			tpl.ExecuteTemplate(w, "login.html", "wrong")
			return
		}

		sId := uuid.New()

		cookie := &http.Cookie{
			Name:     "session",
			Value:    sId.String(),
			HttpOnly: true,
		}

		sd := model.SessionData{
			Id:           sId.String(),
			LastActivity: time.Now().Format(time.Layout),
			User_id:      user.Id,
		}

		err = model.CreateSession(sd)

		if err != nil {
			log.Println(err)
			return
		}

		http.SetCookie(w, cookie)
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "login.html", nil)
}

func UpdateUser(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}

	ok := StillAuthenticated(c.Value)

	if !ok {
		http.Redirect(w, req, "/logout", http.StatusSeeOther)
		return
	}

	user, err := GetUser(w, c)
	if err != nil {
		log.Println(err)
	}

	if req.Method == http.MethodPost {
		u := req.FormValue("username")
		p := req.FormValue("password")
		fn := req.FormValue("firstname")
		ln := req.FormValue("lastname")
		age := req.FormValue("age")
		role := req.FormValue("role")

		i, err := strconv.Atoi(age)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		hp, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		d := model.User{
			Id:        user.Id,
			Username:  u,
			Password:  string(hp),
			FirstName: fn,
			LastName:  ln,
			Role:      role,
			Age:       i,
		}

		err = model.UpdateUser(d)

		if err != nil {
			log.Println(err)
			return
		}

		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "updateUser.html", user)
}

func Logout(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	c = RevokeSession(c)

	http.SetCookie(w, c)
	http.Redirect(w, req, "/", http.StatusSeeOther)
}

func DefaultUser(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}

	ok := StillAuthenticated(c.Value)
	if !ok {
		http.Redirect(w, req, "/logout", http.StatusSeeOther)
		return
	}

	user, err := GetUser(w, c)
	if err != nil {
		log.Println(err)
	}

	if user.Role != "user" {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	tpl.ExecuteTemplate(w, "defaultUser.html", user)
}

func Admin(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}

	ok := StillAuthenticated(c.Value)
	if !ok {
		http.Redirect(w, req, "/logout", http.StatusSeeOther)
		return
	}

	user, err := GetUser(w, c)
	if err != nil {
		log.Println(err)
	}

	if user.Role != "admin" {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	tpl.ExecuteTemplate(w, "admin.html", user)
}

func GetUser(w http.ResponseWriter, c *http.Cookie) (model.User, error) {
	userSession, err := model.GetSessionById(c.Value)
	if err != nil {
		return model.User{}, err
	}
	user, err := model.GetUserById(userSession.User_id)
	if err != nil {
		return model.User{}, err
	}

	// Update last LastActivity
	userSession.LastActivity = time.Now().Format(time.Layout)
	err = model.UpdateSession(userSession)

	if err != nil {
		fmt.Println(err, "ERRORRRR")
		return model.User{}, err
	}

	return user, nil
}
