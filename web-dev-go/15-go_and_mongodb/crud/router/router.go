package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/udemy-go/web-dev-go/15-go_and_mongodb/crud/controllers"
)

func Routers() *httprouter.Router {
	router := httprouter.New()
	uc := controllers.NewUserController()
	ac := controllers.NewAuthController()
	router.GET("/", uc.Index)
	router.GET("/login", ac.LoginView)
	router.POST("/login", ac.Login)
	router.GET("/register", ac.RegisterView)
	router.POST("/register", ac.Register)
	router.GET("/update", uc.UpdateView)
	router.POST("/update", uc.Update)
	router.Handler("GET", "/assets/*filepath", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	return router
}
