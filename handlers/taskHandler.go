package handlers

import (
	"taskManagerWeb/config"
	"net/http"
	"taskManagerWeb/models"
	"strings"
	"taskManagerWeb/errorHandler"
	"io/ioutil"
)

func GetTasks(configObject config.ContextObject) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		data,err := models.Get(configObject)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
		}
		res.Write(data)
	}
}

func AddTask(configObject config.ContextObject) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		taskDescription := strings.Join(req.Form["task"], "")
		priority := strings.Join(req.Form["priority"], "")
		err := models.Add(configObject,taskDescription,priority)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
		}
		res.WriteHeader(http.StatusAccepted)
	}
}

func DeleteTask(configObject config.ContextObject) http.HandlerFunc{
	return func(res http.ResponseWriter, req *http.Request) {
		err := models.Delete(configObject,req.URL.Path)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
		}
		res.WriteHeader(http.StatusAccepted)
	}
}

func UpdateTask(configObject config.ContextObject) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		taskId := strings.Join(req.Form["taskId"], "")
		data := strings.Join(req.Form["data"], "")
		priority := strings.Join(req.Form["priority"], "")
		err := models.Update(configObject,taskId,data,priority)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
		}
		res.WriteHeader(http.StatusAccepted)
	}
}

func UploadTaskFromCsv(configObject config.ContextObject) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		file,_,err := req.FormFile("uploadFile")
		if err != nil {
			errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
		}
		defer file.Close()
		data,err := ioutil.ReadAll(file)
		err = models.AddTaskByCsv(configObject,string(data))
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
		}
		res.WriteHeader(http.StatusAccepted)
	}
}