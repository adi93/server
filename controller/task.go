package controller

import (
	_ "io/ioutil" // gonna use this later
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"server/api"
	"server/service"
)

// TaskController is for task crud operations
type TaskController struct {
	TaskService service.ITaskService
}

// CreateTask ...
func (pc TaskController) CreateTask(w http.ResponseWriter, r *http.Request) {
	var createTaskRequest api.CreateTaskRequest
	err := NewValidationDecoder(r).DecodeAndValidate(&createTaskRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("createTaskRequest:[%v]", createTaskRequest)

	resp := pc.TaskService.CreateTask(r.Context(), createTaskRequest)
	log.Printf("createTaskResponse:[%v]", resp)
	handleResponse(resp, w)
}

// GetAllTasks ...
func (pc TaskController) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	resp := pc.TaskService.GetAllTasks(r.Context())
	log.Printf("GetAllTasksResponse: [%v]", resp)
	handleResponse(resp, w)
}

// GetTask ...
func (pc TaskController) GetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name, ok := vars["name"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp := pc.TaskService.GetTask(r.Context(), name)
	log.Printf("GetTaskResponse:[%v]", resp)
	handleResponse(resp, w)
}

// UpdateTask ...
func (pc TaskController) UpdateTask(w http.ResponseWriter, r *http.Request) {
	var updateTaskRequest api.UpdateTaskRequest
	err := NewValidationDecoder(r).DecodeAndValidate(&updateTaskRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("updateTaskRequest:[%v]", updateTaskRequest)

	resp := pc.TaskService.UpdateTask(r.Context(), updateTaskRequest)
	log.Printf("updateTaskResponse:[%v]", resp)
	handleResponse(resp, w)

}

// DeleteTask ...
func (pc TaskController) DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name, ok := vars["name"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	resp := pc.TaskService.DeleteTask(r.Context(), name)
	log.Printf("DeleteTaskResponse:[%v]", resp)
	handleResponse(resp, w)
}
