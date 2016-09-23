package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"./home"
	"./images"
	"./ws"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	httpAddr = flag.String("http", ":8080", "Listen address")
	static   = flag.String("static", "./static/", "Static Folder")
)

func main() {
	flag.Parse()

	hub := ws.NewHub()
	go hub.Run()

	r := mux.NewRouter()
	r.HandleFunc("/home", home.Get).Methods("GET")
	r.HandleFunc("/images", images.List).Methods("GET")
	r.HandleFunc("/images", images.Upsert).Methods("POST")
	r.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir(*static))))
	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws.ServeWs(hub, w, r)
	})

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	srv := &http.Server{
		Handler:      loggedRouter,
		Addr:         *httpAddr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Listening on " + srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
