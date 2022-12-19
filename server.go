package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func main() {

	m := http.NewServeMux()
	m.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("hola mundo :)")); err != nil {
			log.Errorf("failed to stream response with '%v'", err)
		}
	})
	const addr = ":8080"
	server := &http.Server{
		Addr:    addr,
		Handler: m,
	}
	log.Infof("listening on '%s'", addr)
	log.Fatal(server.ListenAndServe())
}
