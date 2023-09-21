package main

import (
	"log"
	"net/http"
)

func rootHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
}

func main() {
	port := "8000"

	mux := http.NewServeMux()

	mux.Handle("/", rootHandler())

	log.Println("server is listening on port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
