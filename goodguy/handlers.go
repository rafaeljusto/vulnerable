package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "/favicon.ico" {
		handleFavicon(w)
		return
	}

	if _, err := validateSession(r); err == nil {
		w.Header().Set("Location", "/update")
		w.WriteHeader(http.StatusSeeOther)
		return
	}

	username := r.FormValue("username")
	username = strings.TrimSpace(username)
	username = strings.ToLower(username)

	password := r.FormValue("password")
	password = strings.TrimSpace(password)

	data := loginData{
		basicData: basicData{
			SuccessMessage: r.FormValue("successMessage"),
			ErrorMessage:   r.FormValue("errorMessage"),
		},
		Username: username,
	}

	if r.Method == "GET" {
		buildLoginOutput(w, data)
		return
	}

	if username == "" || password == "" {
		data.ErrorMessage = "Missing username or password for login"
		buildLoginOutput(w, data)
		return
	}

	usersLock.Lock()
	defer usersLock.Unlock()

	if user, ok := users[username]; ok && user.password == password {
		w.Header().Set("Location", "/update")
		http.SetCookie(w, &http.Cookie{
			Name:  "session",
			Value: generateSessionHash(user, r.RemoteAddr) + user.Username,
		})
		w.WriteHeader(http.StatusSeeOther)

	} else {
		data.ErrorMessage = "Wrong username or password"
		buildLoginOutput(w, data)
	}
}

func handleLogout(w http.ResponseWriter, r *http.Request) {
	if _, err := validateSession(r); err != nil {
		w.Header().Set("Location", "/?errorMessage="+url.QueryEscape("You're not logged"))
		w.WriteHeader(http.StatusSeeOther)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "session",
		Expires: time.Now().Add(-time.Hour),
	})

	w.Header().Set("Location", "/?successMessage="+url.QueryEscape("Logged out successfully"))
	w.WriteHeader(http.StatusSeeOther)
}

func buildLoginOutput(w http.ResponseWriter, data loginData) {
	var err error

	data.Menu = menuLogin

	t := template.New("login")
	t, err = t.Parse(loginTmpl)

	if err != nil {
		log.Printf("error parsing login template. details: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = t.Execute(w, data); err != nil {
		log.Printf("error executing login template. details: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "/favicon.ico" {
		handleFavicon(w)
		return
	}

	name := r.FormValue("name")
	name = strings.TrimSpace(name)

	username := r.FormValue("username")
	username = strings.TrimSpace(username)
	username = strings.ToLower(username)

	password := r.FormValue("password")
	password = strings.TrimSpace(password)

	data := registerData{
		user: user{
			Name:         name,
			Username:     username,
			password:     password,
			randomSource: rand.Int(),
		},
	}

	if user, err := validateSession(r); err == nil {
		data.LoggedUsername = user.Username
	}

	if r.Method == "GET" {
		buildRegisterOutput(w, data)
		return
	}

	if username == "" || password == "" || name == "" {
		data.ErrorMessage = "Not all fields were filled"
		buildRegisterOutput(w, data)
		return
	}

	usersLock.Lock()
	defer usersLock.Unlock()

	if _, ok := users[username]; ok {
		data.ErrorMessage = "Username already in use"
		buildRegisterOutput(w, data)
		return
	}

	users[username] = data.user

	w.Header().Set("Location", "/?successMessage="+url.QueryEscape("Your account was created successfully"))
	w.WriteHeader(http.StatusSeeOther)
}

func buildRegisterOutput(w http.ResponseWriter, data registerData) {
	var err error

	data.Menu = menuRegister

	t := template.New("register")
	t, err = t.Parse(registerTmpl)

	if err != nil {
		log.Printf("error parsing register template. details: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = t.Execute(w, data); err != nil {
		log.Printf("error executing register template. details: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func handleUpdate(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "/favicon.ico" {
		handleFavicon(w)
		return
	}

	user, err := validateSession(r)
	if err != nil {
		w.Header().Set("Location", "/?errorMessage="+url.QueryEscape(err.Error()))
		w.WriteHeader(http.StatusSeeOther)
		return
	}

	data := registerData{
		basicData: basicData{
			LoggedUsername: user.Username,
		},
		user: user,
	}

	name := r.FormValue("name")
	name = strings.TrimSpace(name)

	password := r.FormValue("password")
	password = strings.TrimSpace(password)

	if password == "" && name == "" {
		buildUpdateOutput(w, data)
		return
	}

	data.Name = name
	if password != "" {
		data.password = password
	}

	usersLock.Lock()
	defer usersLock.Unlock()

	users[user.Username] = data.user
	data.SuccessMessage = "Updated successfully"
	buildUpdateOutput(w, data)
}

func buildUpdateOutput(w http.ResponseWriter, data registerData) {
	var err error

	data.Menu = menuUpdate

	t := template.New("update")
	t, err = t.Parse(updateTmpl)

	if err != nil {
		log.Printf("error parsing update template. details: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = t.Execute(w, data); err != nil {
		log.Printf("error executing update template. details: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func handleSearch(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "/favicon.ico" {
		handleFavicon(w)
		return
	}

	username := r.FormValue("username")
	username = strings.TrimSpace(username)
	username = strings.ToLower(username)

	data := searchData{
		NumberOfUsers: len(users),
		Username:      username,
	}

	if user, err := validateSession(r); err == nil {
		data.LoggedUsername = user.Username
	}

	if r.Method == "GET" || data.Username == "" {
		buildSearchOutput(w, data)
		return
	}

	usersLock.Lock()
	defer usersLock.Unlock()

	if user, ok := users[data.Username]; ok {
		data.Result = user
	} else {
		data.ErrorMessage = "Username not found"
	}

	buildSearchOutput(w, data)
}

func buildSearchOutput(w http.ResponseWriter, data searchData) {
	var err error

	data.Menu = menuSearch

	t := template.New("search").Funcs(template.FuncMap{
		"toHTML": func(input string) template.HTML {
			return template.HTML(input)
		},
	})
	t, err = t.Parse(searchTmpl)

	if err != nil {
		log.Printf("error parsing search template. details: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = t.Execute(w, data); err != nil {
		log.Printf("error executing search template. details: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func handleFavicon(w http.ResponseWriter) {
	favicon, err := base64.StdEncoding.DecodeString(faviconData)

	if err != nil {
		log.Printf("error decoding favicon. details: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "image/jpeg")
	if _, err = w.Write(favicon); err != nil {
		log.Printf("error writing body. details: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func validateSession(r *http.Request) (user, error) {
	session, err := r.Cookie("session")
	if err == http.ErrNoCookie {
		return user{}, fmt.Errorf("Invalid session")
	}

	if len(session.Value) < 64 {
		return user{}, fmt.Errorf("Invalid session size")
	}

	hash := session.Value[:64]
	username := session.Value[64:]

	usersLock.Lock()
	user, ok := users[username]
	usersLock.Unlock()

	if !ok {
		return user, fmt.Errorf("Invalid session username")
	}

	if hash != generateSessionHash(user, r.RemoteAddr) {
		return user, fmt.Errorf("Invalid session auth")
	}

	return user, nil
}

func generateSessionHash(u user, remoteAddr string) string {
	ip, _, _ := net.SplitHostPort(remoteAddr)

	hash := sha256.New()
	hash.Write([]byte(fmt.Sprintf("%s%d%s", u.Username, u.randomSource, ip)))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
