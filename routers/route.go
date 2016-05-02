package routers

import (
	"taskManagerWeb/config"
	"github.com/gorilla/mux"
	"taskManagerWeb/handlers"
	"net/http"
)

func HandleRequests(context config.Context) {
	r := mux.NewRouter()
	r.HandleFunc("/tasks/download/csv",handlers.DownloadCsv(context)).Methods("GET")
	r.HandleFunc("/tasks/{id:[0-9]+}",handlers.UpdateTask(context)).Methods("PATCH")
	r.HandleFunc("/tasks/csv",handlers.UploadTaskFromCsv(context)).Methods("POST")
	r.HandleFunc("/task/{id:[0-9]+}", handlers.DeleteTask(context)).Methods("DELETE")
	r.HandleFunc("/tasks", handlers.GetTasks(context)).Methods("GET")
	r.HandleFunc("/tasks", handlers.AddTask(context)).Methods("POST")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./")))
	http.Handle("/", r)
}
