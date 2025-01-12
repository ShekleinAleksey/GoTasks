package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ShekleinAleksey/GoTasks/database"
	"github.com/ShekleinAleksey/GoTasks/handlers"
	"github.com/ShekleinAleksey/GoTasks/telegram"
	"github.com/gorilla/mux"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	taskHandlers := handlers.NewTaskHandkers(db)

	go telegram.StartBot()

	router := mux.NewRouter()
	router.HandleFunc("/tasks", taskHandlers.CreateTask).Methods("POST")
	router.HandleFunc("/tasks", taskHandlers.GetTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", taskHandlers.GetTaskById).Methods("GET")
	router.HandleFunc("/tasks/{id}", taskHandlers.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", taskHandlers.DeleteTask).Methods("DELETE")
	fmt.Println("hello")
	http.ListenAndServe(":8080", nil)
}
