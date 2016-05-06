package handlers

import (
	"taskManagerWeb/config"
	"net/http"
	"taskManagerWeb/models"
	"strings"
	"taskManagerWeb/errorHandler"
	"io/ioutil"
	"github.com/codegangsta/negroni"
	"taskManagerWeb/tokenValidator"
)

func Auth(context config.Context) negroni.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
		cookie,err := req.Cookie("taskManagerToken")
		if err != nil {
			http.Redirect(res,req,"auth",http.StatusTemporaryRedirect)
			return
		}
		if !(tokenValidator.IsValidToken(cookie.Value,req)){
			http.Redirect(res,req,"auth",http.StatusTemporaryRedirect)
			return
		}
		next( res,req)
		return
	}
}

func GetTasks(context config.Context) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		cookie,err := req.Cookie("taskManagerId")
		data,err := models.Get(context,cookie.Value)
		if err != nil {
			http.Error(res, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		res.Write(data)
	}
}

func AddTask(context config.Context) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		cookie,err := req.Cookie("taskManagerId")
		req.ParseForm()
		taskDescription := strings.Join(req.Form["task"], "")
		priority := strings.Join(req.Form["priority"], "")
		err = models.Add(context,taskDescription,priority,cookie.Value)
		if err != nil {
			http.Error(res, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		res.WriteHeader(http.StatusAccepted)
	}
}

func DeleteTask(context config.Context) http.HandlerFunc{
	return func(res http.ResponseWriter, req *http.Request) {
		cookie,err := req.Cookie("taskManagerId")
		err = models.Delete(context,req.URL.Path,cookie.Value)
		if err != nil {
			http.Error(res, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		res.WriteHeader(http.StatusAccepted)
	}
}

func UpdateTask(context config.Context) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		taskId := req.URL.Path
		data := strings.Join(req.Form["data"], "")
		priority := strings.Join(req.Form["priority"], "")
		cookie,err := req.Cookie("taskManagerId")
		err = models.Update(context,taskId,data,priority,cookie.Value)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			http.Error(res, "Internal Server Error", http.StatusInternalServerError)
			return
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
		cookie,err := req.Cookie("taskManagerId")
		err = models.AddTaskByCsv(context,string(data),cookie.Value)
		if err != nil {
			res.Write([]byte(err.Error()))
			http.Error(res, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		res.WriteHeader(http.StatusAccepted)

	}
}

func DownloadCsv(context config.Context) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		cookie,err := req.Cookie("taskManagerId")
		data,err := models.GetCsv(context,cookie.Value)
		if err != nil {
			http.Error(res, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		res.Header().Set("Content-Disposition","attachment; filename=task.csv")
		res.Write(data)
	}
}