package main

import (
    "log"
    "net/http"
    "time"

    "github.com/gorilla/mux"
    "github.com/gorilla/websocket"
)

func main() {
    router := mux.NewRouter()
    scheduler := NewScheduler()

    // REST API endpoints
    router.HandleFunc("/jobs", scheduler.CreateJob).Methods("POST")
    router.HandleFunc("/jobs", scheduler.GetJobs).Methods("GET")

    // WebSocket endpoint
    router.HandleFunc("/ws", scheduler.HandleWebSocket)

    srv := &http.Server{
        Handler: router,
        Addr:    "127.0.0.1:8080",
        //  timeouts for servers 
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
        IdleTimeout:  60 * time.Second,
    }

    log.Fatal(srv.ListenAndServe())
}
