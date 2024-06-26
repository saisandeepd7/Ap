package main

import "time"

// Job represents a job with a name and duration
type Job struct {
    Name     string        `json:"name"`
    Duration time.Duration `json:"duration"`
    Status   string        `json:"status"`
}
