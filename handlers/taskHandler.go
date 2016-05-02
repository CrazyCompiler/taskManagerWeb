package handlers

import (
	"taskManagerWeb/config"
	"net/http"
	"taskManagerWeb/models"
	"strings"
	"taskManagerWeb/errorHandler"
	"io/ioutil"
)

func GetTasks(context config.Context) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		data,err := models.Get(context)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
		}
		res.Write(data)
	}
}

func AddTask(context config.Context) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		taskDescription := strings.Join(req.Form["task"], "")
		priority := strings.Join(req.Form["priority"], "")
		err := models.Add(context,taskDescription,priority)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
		}
		res.WriteHeader(http.StatusAccepted)
	}
}

func DeleteTask(context config.Context) http.HandlerFunc{
	return func(res http.ResponseWriter, req *http.Request) {
		err := models.Delete(context,req.URL.Path)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
		}
		res.WriteHeader(http.StatusAccepted)
	}
}

func UpdateTask(context config.Context) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		taskId := strings.Join(req.Form["taskId"], "")
		data := strings.Join(req.Form["data"], "")
		priority := strings.Join(req.Form["priority"], "")
		err := models.Update(context,taskId,data,priority)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
		}
		res.WriteHeader(http.StatusAccepted)
	}
}

func UploadTaskFromCsv(context config.Context) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		file,_,err := req.FormFile("uploadFile")
		if err != nil {
			errorHandler.ErrorHandler(context.ErrorLogFile,err)
		}
		defer file.Close()
		data,err := ioutil.ReadAll(file)
		err = models.AddTaskByCsv(context,string(data))
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
		}
		res.WriteHeader(http.StatusAccepted)
	}
}

func DownloadCsv(context config.Context) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		data,err := models.GetCsv(context)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
		}
		res.Header().Set("Content-Disposition","attachment; filename=task.csv")
		res.Write(data)
	}
}