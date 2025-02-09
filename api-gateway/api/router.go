package api

import (
	"api-gateway/handlers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/register", handlers.Register).Methods("POST")
	router.HandleFunc("/login", handlers.Login)

	router.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
	// router.HandleFunc("/create", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")

	router.HandleFunc("/todos/{id}", handlers.GetTodo).Methods("GET")
	router.HandleFunc("/todo", handlers.CreateTodo).Methods("POST")
	router.HandleFunc("/todos/{id}", handlers.UpdateTodo).Methods("PUT")
	router.HandleFunc("/todos/{id}", handlers.DeleteTodo).Methods("DELETE")
	router.HandleFunc("/mark/{id}", handlers.MarkAsDone).Methods("PUT")

	router.HandleFunc("/all", handlers.GetAllTodos)

	router.HandleFunc("/welcome", welcome)

	// router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
	// 	fmt.Println(route.GetPathTemplate())
	// 	return nil
	// })

	return router
}

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome")
}
