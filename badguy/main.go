package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var port int
var attackServer string

func init() {
	flag.IntVar(&port, "port", 80, "server listen port")
	flag.StringVar(&attackServer, "attack-server", "", "ip address and port of the good guy")
}

func main() {
	flag.Parse()

	http.HandleFunc("/", handle)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	log.Fatalf("error starting http server. details: %s", err)
}
