package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleLogin)
	http.HandleFunc("/logout", handleLogout)
	http.HandleFunc("/register", handleRegister)
	http.HandleFunc("/update", handleUpdate)

	// http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
	// })

	log.Fatalf("error starting http server. details: %s", http.ListenAndServe(":8080", nil))
}
