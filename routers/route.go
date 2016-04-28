package routers

import (
	"taskManagerWeb/config"
	"github.com/gorilla/mux"
	"taskManagerWeb/handlers"
	"net/http"
)

func HandleRequests(configObject config.ContextObject) {
	r := mux.NewRouter()
	r.HandleFunc("/update",handlers.UpdateTask(configObject)).Methods("POST")
	r.HandleFunc("/tasks/csv",handlers.UploadTaskFromCsv(configObject)).Methods("POST")
	r.HandleFunc("/task/{id:[0-9]+}", handlers.DeleteTask(configObject)).Methods("DELETE")
	r.HandleFunc("/tasks", handlers.GetTasks(configObject)).Methods("GET")
	r.HandleFunc("/task", handlers.AddTask(configObject)).Methods("POST")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./")))
	http.Handle("/", r)
}
