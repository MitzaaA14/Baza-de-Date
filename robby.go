package main

import (
	"log"
	"net/http"
	"robby/db"
	"robby/pages"
	"time"
)

func main() {
	if err := db.InitDb(); err != nil {
		log.Fatalf("DB error: %s", err)
	}

	if err := pages.InitTemplates(); err != nil {
		log.Fatalf("Templates error: %s", err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/cautare", pages.HandlerCautarefunc)
	mux.HandleFunc("/turneu", pages.HandlerTurneufunc)
	mux.HandleFunc("/", pages.HandlerIndexfunc)

	log.Printf("Server started on http://localhost:8080")

	s := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
