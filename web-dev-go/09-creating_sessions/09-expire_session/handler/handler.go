package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/udemy-go/web-dev-go/09-creating_sessions/09-expire_session/database"
	"golang.org/x/crypto/bcrypt"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	// database.DbSessionCleaned = time.Now()
	go func() {
		for {
			time.Sleep(time.Second * 30)
			database.CleanSession()
			fmt.Println(database.DbSession)
		}
	}()

}

func Index(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}

	user := GetUser(w, c)

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

		database.DbUser[u] = database.User{
			Username:  u,
			Password:  string(hp),
			FirstName: fn,
			LastName:  ln,
			Age:       i,
			Role:      role,
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
			MaxAge:   30,
			HttpOnly: true,
		}

		database.DbSession[cookie.Value] = database.SessionData{
			Username:     u,
			LastActivity: time.Now(),
		}

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

	c = database.DeleteSession(c)

	http.SetCookie(w, c)
	http.Redirect(w, req, "/", http.StatusSeeOther)
}

func DefaultUser(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}

	user := GetUser(w, c)

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

	user := GetUser(w, c)

	if user.Role != "admin" {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	tpl.ExecuteTemplate(w, "admin.html", user)
}

func GetUser(w http.ResponseWriter, c *http.Cookie) database.User {
	userSession := database.DbSession[c.Value]
	user := database.DbUser[userSession.Username]
	userSession.LastActivity = time.Now()
	c.MaxAge = database.SessionLength
	http.SetCookie(w, c)
	fmt.Println(userSession, c.MaxAge)
	return user
}
