package handlers

import (
	"taskManagerWeb/config"
	"net/http"
	"taskManagerWeb/models"
	"strings"
	"taskManagerWeb/errorHandler"
)

func GetTasks(configObject config.ContextObject) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		rows := models.Get(configObject)
		res.Write(rows)
	}
}

func AddTask(configObject config.ContextObject) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		taskDescription := strings.Join(req.Form["task"], "")
		priority := strings.Join(req.Form["priority"], "")
		err := models.Add(configObject,taskDescription,priority)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
		}
		res.WriteHeader(http.StatusAccepted)
	}
}

func DeleteTask(configObject config.ContextObject) http.HandlerFunc{
	return func(res http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		err := models.Delete(configObject,req.URL.Path)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
		}
		res.WriteHeader(http.StatusAccepted)
	}
}

func UpdateTaskPriority(configObject config.ContextObject) http.HandlerFunc{
	return func(res http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		taskId := strings.Join(req.Form["taskId"], "")
		priority := strings.Join(req.Form["priority"], "")
		err := models.UpdatePriority(configObject,taskId,priority)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
		}
		res.WriteHeader(http.StatusAccepted)
	}
}


func UpdateTaskDescription(configObject config.ContextObject) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		taskId := strings.Join(req.Form["taskId"], "")
		priority := strings.Join(req.Form["data"], "")
		err := models.UpdateDescription(configObject,taskId,priority)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
		}
		res.WriteHeader(http.StatusAccepted)
	}
}

func UploadTaskFromCsv(configObject config.ContextObject) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		err := req.ParseMultipartForm(32 << 20)
		if err != nil {
			errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
		}
		m := req.MultipartForm
		err = models.UploadCsv(configObject,m)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
		}
		res.WriteHeader(http.StatusAccepted)
	}
}