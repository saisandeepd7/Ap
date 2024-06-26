# Job Scheduler

## Overview

This project is a simplified job scheduler with a React frontend that visualizes job statuses and allows users to submit new jobs. The scheduler prioritizes jobs using the Shortest Job First (SJF) algorithm. Real-time UI updates are achieved using WebSockets.

## How to Run

### Backend (Go)

1. **Navigate to the backend directory:**
   ```bash
   cd backend
   go mod tidy
   go run main.go

### Frontend (React)

2. **Navigate to the backend directory:**
   ```bash
   cd frontend
   npm install
   npm start

## Design Choices and Approaches

### Backend

Language: Go was chosen for its performance and simplicity in handling concurrency, which is beneficial for implementing the job scheduler.

Shortest Job First (SJF) Algorithm: The scheduler prioritizes jobs based on their duration, running the shortest jobs first to minimize the average waiting time.

WebSockets: Used for real-time communication to keep the frontend updated with job statuses without the need for polling.

Data Structures:

Scheduler: Manages the job queue, WebSocket clients, and broadcasting updates.
Job: Represents a job with a name, duration, and status.

### Frontend

React: Chosen for its component-based architecture, making it easy to manage and update the UI.

WebSocket Integration: Ensures the UI receives real-time updates from the backend, providing an interactive user experience.

Job Status Indicators: Clear visual indicators (color-coded) for different job statuses (pending, running, completed).

Job Submission Form: A simple form to submit new jobs to the backend

## Conclusion

This project demonstrates a simple yet effective job scheduler with a Go backend and React frontend, showcasing real-time updates and prioritization using the SJF algorithm. The clear separation of concerns and use of modern web technologies ensures maintainability and scalability.
