package handler

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/udemy-go/web-dev-go/09-creating_sessions/04-sign_up/database"
	"golang.org/x/crypto/bcrypt"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func Index(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		tpl.ExecuteTemplate(w, "index.html", nil)
		return
	}

	userId := database.DbSession[c.Value]
	user := database.DbUser[userId]

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

		// Check username taken
		_, ok := database.DbUser[u]
		if ok {
			tpl.ExecuteTemplate(w, "register.html", "username already taken")
			return
		}

		p := req.FormValue("password")
		fn := req.FormValue("firstname")
		ln := req.FormValue("lastname")
		age := req.FormValue("age")

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

		database.DbUser[u] = database.User{
			Username:  u,
			Password:  string(hp),
			FirstName: fn,
			LastName:  ln,
			Age:       i,
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

		user, ok := database.DbUser[u]

		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(p))

		if !ok || u != user.Username || err != nil {
			tpl.ExecuteTemplate(w, "login.html", "wrong")
			return
		}

		sId := uuid.New()

		cookie := &http.Cookie{
			Name:     "session",
			Value:    sId.String(),
			MaxAge:   60 * 10,
			HttpOnly: true,
		}

		database.DbSession[cookie.Value] = u

		http.SetCookie(w, cookie)
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "login.html", nil)
}

func Logout(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	c.MaxAge = -1

	delete(database.DbSession, c.Value)

	c.Value = ""

	http.SetCookie(w, c)
	http.Redirect(w, req, "/", http.StatusSeeOther)
}
