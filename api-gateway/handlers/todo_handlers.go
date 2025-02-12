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
	userID := r.Context().Value("userID").(string)
	if userID == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

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
	userID := r.Context().Value("userID").(string) // Get user ID from AuthMiddleware
	if userID == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todos := &proto.CreateTodoRequest{
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		Completed:   boolPtr(false),
		UserId:      userID,
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
	userID := r.Context().Value("userID").(string)
	if userID == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

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
		UserId:      userID,
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
	userID, ok := r.Context().Value("userID").(string)
	if !ok || userID == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	todoID := mux.Vars(r)["id"]
	if todoID == "" {
		http.Error(w, "Todo ID is required!", http.StatusBadRequest)
		return
	}

	resp, err := todoClient.DeleteTodo(context.Background(), todoID, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userID").(string)
	if !ok || userID == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	resp, err := todoClient.GetAllTodos(context.Background(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// fmt.Println("getAllTodos Resp: ", resp)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func MarkAsDone(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userID").(string)
	if !ok || userID == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	todoID := mux.Vars(r)["id"]
	if todoID == "" {
		http.Error(w, "Todo ID is required!", http.StatusBadRequest)
		return
	}

	resp, err := todoClient.MarkAsDone(context.Background(), todoID, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func boolPtr(b bool) *bool {
	return &b
}

/*

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	stream, err := todoClient.client.GetAllTodos(ctx, &proto.GetAllTodosRequest{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	encoder := json.NewEncoder(w)

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_ = encoder.Encode(resp.Todo) // Stream each JSON object
		w.(http.Flusher).Flush()       // Force sending data to the client
	}
}

*/
