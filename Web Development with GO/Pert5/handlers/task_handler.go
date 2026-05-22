package handlers

import (
	"encoding/json"
	"net/http"
	"pert5/models"
	"pert5/utils"
	"strconv"
	"strings"
	"time"
)

var tasks []models.Task
var nextID = 1

// Validasi status task
var validStatuses = map[string]bool{
	"pending":     true,
	"in-progress": true,
	"done":        true,
}

// func init() {
// 	tasks = []models.Task{}
// 	nextID = 1
// }

// GET /api/tasks
func GetTasks(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, http.StatusOK, "success", "List of tasks", tasks)
}

// POST /api/tasks
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)

	// Validasi status
	if !validStatuses[task.Status] {
		utils.WriteJSON(w, http.StatusBadRequest, "error", "Invalid status", nil)
		return
	}

	task.ID = nextID
	task.CreatedAt = time.Now()
	nextID++

	tasks = append(tasks, task)

	utils.WriteJSON(w, http.StatusCreated, "success", "Task created", task)
}

// PUT /api/tasks/{id}
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/tasks/")
	id, _ := strconv.Atoi(idStr)

	for i, task := range tasks {
		if task.ID == id {
			var updatedTask models.Task
			json.NewDecoder(r.Body).Decode(&updatedTask)

			// Validasi status
			if !validStatuses[updatedTask.Status] {
				utils.WriteJSON(w, http.StatusBadRequest, "error", "Invalid status", nil)
				return
			}

			tasks[i].Title = updatedTask.Title
			tasks[i].Description = updatedTask.Description
			tasks[i].Status = updatedTask.Status

			utils.WriteJSON(w, http.StatusOK, "success", "Task updated", tasks[i])
			return
		}
	}

	utils.WriteJSON(w, http.StatusNotFound, "error", "Task not found", nil)
}

// PATCH /api/tasks/{id}
func PatchTask(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/tasks/")
	id, _ := strconv.Atoi(idStr)

	for i, task := range tasks {
		if task.ID == id {
			var patchData map[string]interface{}
			json.NewDecoder(r.Body).Decode(&patchData)

			if title, ok := patchData["title"].(string); ok {
				tasks[i].Title = title
			}
			if description, ok := patchData["description"].(string); ok {
				tasks[i].Description = description
			}
			if status, ok := patchData["status"].(string); ok {
				if !validStatuses[status] {
					utils.WriteJSON(w, http.StatusBadRequest, "error", "Invalid status", nil)
					return
				}
				tasks[i].Status = status
			}

			utils.WriteJSON(w, http.StatusOK, "success", "Task patched", tasks[i])
			return
		}
	}

	utils.WriteJSON(w, http.StatusNotFound, "error", "Task not found", nil)
}

// DELETE /api/tasks/{id}
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/tasks/")
	id, _ := strconv.Atoi(idStr)

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			utils.WriteJSON(w, http.StatusOK, "success", "Task deleted", nil)
			return
		}
	}

	utils.WriteJSON(w, http.StatusNotFound, "error", "Task not found", nil)
}

func GetTaskByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/tasks/")
	id, _ := strconv.Atoi(idStr)

	for _, task := range tasks {
		if task.ID == id {
			utils.WriteJSON(w, http.StatusOK, "success", "Task found", task)
			return
		}
	}

	utils.WriteJSON(w, http.StatusNotFound, "error", "Task not found", nil)
}
