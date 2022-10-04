package models

type User struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Age       int    `json:"Age"`
	Id        string `json:"Id"`
}
