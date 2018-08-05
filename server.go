package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type server struct {
	router *mux.Router
}

func newServer() *server {
	s := server{
		mux.NewRouter(),
	}

	s.router.HandleFunc("/", authenticateRequest).Methods("POST")

	http.Handle("/", s.router)

	return &s
}

func (s *server) start() {
	srv := &http.Server{
		Handler:      s.router,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Starting http server.")
	log.Fatal(srv.ListenAndServe())
}

func authenticateRequest(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("X-Access-Token")

	pubkey, err := loadPublicKey("test/cert1.pem")
	if err != nil {
		fmt.Println("Error loading pubkey", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	parsedToken, err := validateToken(pubkey, token)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, parsedToken.UserId)
}
