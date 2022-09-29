package database

type User struct {
	Username, Password, FirstName, LastName, Role string
	Age                                           int
}

var DbUser = map[string]User{}
var DbSession = map[string]string{}
