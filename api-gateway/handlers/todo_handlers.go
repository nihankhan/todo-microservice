package handlers

import (
	"api-gateway/grpc"
	"api-gateway/proto"
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var todoClient = grpc.NewTodoClient()

func GetTodo(w http.ResponseWriter, r *http.Request) {
	todoID := mux.Vars(r)["id"]
	if todoID == "" {
		http.Error(w, "Todo ID is required!", http.StatusBadRequest)
		return
	}

	resp, err := todoClient.GetTodoByID(context.Background(), todoID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp.Todo)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todos := &proto.CreateTodoRequest{
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
	}

	resp, err := todoClient.CreateTodo(context.Background(), todos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	ID := mux.Vars(r)["id"]
	if ID == "" {
		http.Error(w, "Todo ID is required!", http.StatusBadRequest)
		return 
	}
	
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todos := &proto.UpdateTodoRequest{
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
	}

	resp, err := todoClient.UpdateTodo(context.Background(), todos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	todoID := mux.Vars(r)["id"]
	if todoID == "" {
		http.Error(w, "Todo ID is required!", http.StatusBadRequest)
		return
	}

	resp, err := todoClient.DeleteTodo(context.Background(), todoID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
