package main

import "sync"

var users map[string]user
var usersLock sync.Mutex

func init() {
	users = make(map[string]user)
}

const (
	menuLogin    menu = "login"
	menuUpdate   menu = "update"
	menuRegister menu = "register"
	menuSearch   menu = "search"
)

type menu string

type basicData struct {
	Menu           menu
	LoggedUsername string
	ErrorMessage   string
	SuccessMessage string
}

type loginData struct {
	basicData
	Username string
}

type registerData struct {
	basicData
	user
}

type user struct {
	Name         string
	Username     string
	password     string
	randomSource int
}
