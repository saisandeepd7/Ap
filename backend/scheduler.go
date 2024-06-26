package main

import (
    "encoding/json"
    "log"
    "net/http"
    "sort"
    "sync"
    "time"

    "github.com/gorilla/websocket"
)

// Job represents a job with a name and duration
type Job struct {
    Name     string        `json:"name"`
    Duration time.Duration `json:"duration"`
    Status   string        `json:"status"`
}

// Scheduler manages job scheduling
type Scheduler struct {
    Jobs      []*Job
    mu        sync.Mutex
    clients   map[*websocket.Conn]bool
    broadcast chan *Job
}

// NewScheduler creates a new Scheduler
func NewScheduler() *Scheduler {
    return &Scheduler{
        Jobs:      make([]*Job, 0),
        clients:   make(map[*websocket.Conn]bool),
        broadcast: make(chan *Job),
    }
}

// CreateJob handles job creation
func (s *Scheduler) CreateJob(w http.ResponseWriter, r *http.Request) {
    var job Job
    if err := json.NewDecoder(r.Body).Decode(&job); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    job.Status = "pending"

    s.mu.Lock()
    s.Jobs = append(s.Jobs, &job)
    sort.Slice(s.Jobs, func(i, j int) bool {
        return s.Jobs[i].Duration < s.Jobs[j].Duration
    })
    s.mu.Unlock()

    go s.runJob(&job)
    s.broadcast <- &job

    w.WriteHeader(http.StatusCreated)
}

// GetJobs retrieves the list of jobs
func (s *Scheduler) GetJobs(w http.ResponseWriter, r *http.Request) {
    s.mu.Lock()
    defer s.mu.Unlock()

    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(s.Jobs); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

// runJob runs a job
func (s *Scheduler) runJob(job *Job) {
    job.Status = "running"
    s.broadcast <- job

    time.Sleep(job.Duration)

    job.Status = "completed"
    s.broadcast <- job
}

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
