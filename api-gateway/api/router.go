package api

import (
	"api-gateway/handlers"
	"api-gateway/middleware"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/register", handlers.Register).Methods("POST")
	router.HandleFunc("/login", handlers.Login)

	protected := router.PathPrefix("/").Subrouter()

	protected.Use(middleware.AuthMiddleware)

	protected.HandleFunc("/logout", handlers.Logout)

	protected.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
	// router.HandleFunc("/create", handlers.CreateUser).Methods("POST")
	protected.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	protected.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")

	protected.HandleFunc("/todos/{id}", handlers.GetTodo).Methods("GET")
	protected.HandleFunc("/todo", handlers.CreateTodo).Methods("POST")
	protected.HandleFunc("/todos/{id}", handlers.UpdateTodo).Methods("PUT")
	protected.HandleFunc("/todos/{id}", handlers.DeleteTodo).Methods("DELETE")
	protected.HandleFunc("/mark/{id}", handlers.MarkAsDone).Methods("PUT")

	protected.HandleFunc("/all", handlers.GetAllTodos)

	router.HandleFunc("/welcome", welcome)

	return router
}

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome")
}
