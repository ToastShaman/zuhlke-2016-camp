package main

import (
	"log"
	"net/http"
	"time"

	"./images"
	"./ws"

	"github.com/gorilla/mux"
)

func main() {
	hub := ws.NewHub()
	go hub.Run()

	r := mux.NewRouter()
	r.HandleFunc("/images", images.List).Methods("GET")
	r.HandleFunc("/images", images.Upsert).Methods("POST")
	r.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("./static"))))
	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws.ServeWs(hub, w, r)
	})

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:9001",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Listening on " + srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
