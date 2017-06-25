package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var port int

func init() {
	flag.IntVar(&port, "port", 80, "server listen port")
}

func main() {
	flag.Parse()

	http.HandleFunc("/", handleLogin)
	http.HandleFunc("/logout", handleLogout)
	http.HandleFunc("/register", handleRegister)
	http.HandleFunc("/update", handleUpdate)
	http.HandleFunc("/search", handleSearch)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	log.Fatalf("error starting http server. details: %s", err)
}
