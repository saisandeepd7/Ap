package main

import (
    "log"
    "net/http"

    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

// HandleWebSocket handles WebSocket requests
func (s *Scheduler) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
        return
    }
    defer conn.Close()

    s.clients[conn] = true

    for {
        select {
        case job := <-s.broadcast:
            for client := range s.clients {
                if err := client.WriteJSON(job); err != nil {
                    log.Println(err)
                    client.Close()
                    delete(s.clients, client)
                }
            }
        }
    }
}
