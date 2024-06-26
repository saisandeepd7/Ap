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