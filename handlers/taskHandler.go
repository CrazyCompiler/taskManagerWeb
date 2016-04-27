package handlers

import (
	"taskManagerWeb/config"
	"net/http"
	"taskManagerWeb/models"
	"strings"
	"taskManagerWeb/errorHandler"
	"github.com/afex/hystrix-go/hystrix"
	"io/ioutil"
)

const TIMEOUT int = 1000
const MAXCONCURRENTREQUESTS int = 100
const ERRORPERCENTTHRESHOLD int = 25

func GetTasks(configObject config.ContextObject) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		rows := make(chan []byte, 1)
		status := make(chan int, 1)

		hystrix.ConfigureCommand("getTask", hystrix.CommandConfig{
			Timeout: TIMEOUT,
			MaxConcurrentRequests: MAXCONCURRENTREQUESTS,
			ErrorPercentThreshold: ERRORPERCENTTHRESHOLD,
		})

		hystrix.Go("getTask", func() error {
			rows <- models.Get(configObject)
			status <- http.StatusOK
			return nil
		}, func(err error) error {
			status <- http.StatusBadRequest
			return nil
		})

		if <-status == http.StatusBadGateway{
			res.WriteHeader(<-status)
		}
		res.Write(<-rows)
	}
}

func AddTask(configObject config.ContextObject) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		taskDescription := strings.Join(req.Form["task"], "")
		priority := strings.Join(req.Form["priority"], "")
		status := make(chan int, 1)

		hystrix.ConfigureCommand("addTask", hystrix.CommandConfig{
			Timeout: TIMEOUT,
			MaxConcurrentRequests: MAXCONCURRENTREQUESTS,
			ErrorPercentThreshold: ERRORPERCENTTHRESHOLD,
		})
		hystrix.Go("addTask", func() error {
			err := models.Add(configObject,taskDescription,priority)
			if err == nil {
				status <-http.StatusAccepted
			}
			return nil
		}, func(err error) error {
			status <- http.StatusBadRequest
			return nil
		})
		res.WriteHeader(<-status)

	}
}

func DeleteTask(configObject config.ContextObject) http.HandlerFunc{
	return func(res http.ResponseWriter, req *http.Request) {
		status := make(chan int, 1)
		hystrix.ConfigureCommand("deleteTask", hystrix.CommandConfig{
			Timeout: TIMEOUT,
			MaxConcurrentRequests: MAXCONCURRENTREQUESTS,
			ErrorPercentThreshold: ERRORPERCENTTHRESHOLD,
		})
		hystrix.Go("deleteTask", func() error {
			err := models.Delete(configObject,req.URL.Path)
			if err == nil {
				status <-http.StatusAccepted
			}
			return nil
		}, func(err error) error {
			status <- http.StatusBadRequest
			return nil
		})
		res.WriteHeader(<-status)
	}
}

func UpdateTaskPriority(configObject config.ContextObject) http.HandlerFunc{
	return func(res http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		taskId := strings.Join(req.Form["taskId"], "")
		priority := strings.Join(req.Form["priority"], "")

		status := make(chan int, 1)

		hystrix.ConfigureCommand("updateTaskPriority", hystrix.CommandConfig{
			Timeout: TIMEOUT,
			MaxConcurrentRequests: MAXCONCURRENTREQUESTS,
			ErrorPercentThreshold: ERRORPERCENTTHRESHOLD,
		})
		hystrix.Go("updateTaskPriority", func() error {
			err := models.UpdatePriority(configObject,taskId,priority)
			if err == nil {
				status <-http.StatusAccepted
			}
			return nil
		}, func(err error) error {
			status <- http.StatusBadRequest
			return nil
		})
		res.WriteHeader(<-status)
	}
}


func UpdateTaskDescription(configObject config.ContextObject) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		taskId := strings.Join(req.Form["taskId"], "")
		data := strings.Join(req.Form["data"], "")
		status := make(chan int, 1)
		hystrix.ConfigureCommand("updateTaskDescription", hystrix.CommandConfig{
			Timeout: TIMEOUT,
			MaxConcurrentRequests: MAXCONCURRENTREQUESTS,
			ErrorPercentThreshold: ERRORPERCENTTHRESHOLD,
		})
		hystrix.Go("updateTaskPriority", func() error {
			err := models.UpdateDescription(configObject,taskId,data)
			if err == nil {
				status <-http.StatusAccepted
			}
			return nil
		}, func(err error) error {
			status <- http.StatusBadRequest
			return nil
		})
		res.WriteHeader(<-status)
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
		status := make(chan int, 1)
		hystrix.ConfigureCommand("uploadCsv", hystrix.CommandConfig{
			Timeout: TIMEOUT,
			MaxConcurrentRequests: MAXCONCURRENTREQUESTS,
			ErrorPercentThreshold: ERRORPERCENTTHRESHOLD,
		})
		hystrix.Go("uploadCsv", func() error {
			err := models.AddTaskByCsv(configObject,string(data))
			if err == nil {
				status <-http.StatusAccepted
			}
			return nil
		}, func(err error) error {
			status <- http.StatusBadRequest
			return nil
		})
		res.WriteHeader(<-status)
	}
}