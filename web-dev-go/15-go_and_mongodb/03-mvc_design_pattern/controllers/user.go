package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/udemy-go/web-dev-go/15-go_and_mongodb/03-mvc_design_pattern/models"
)

// Insert db connection to user
type UserController struct{}

// NewUserController return a pointer to UserController so it can call all the method sets
func NewUserController() *UserController {
	return &UserController{}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u1 := models.User{
		FirstName: "James",
		LastName:  "Bond",
		Age:       47,
		Id:        p.ByName("id"),
	}

	bs, err := json.Marshal(u1)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bs)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u1 := models.User{}

	err := json.NewDecoder(r.Body).Decode(&u1)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	bs, err := json.Marshal(u1)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(bs)
}
