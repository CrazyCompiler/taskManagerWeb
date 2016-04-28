package models

import (
	"taskManagerWeb/config"
	"taskManagerWeb/errorHandler"
	"strconv"
	"github.com/taskManagerContract"
	"taskManagerWeb/serviceCall"
)

func Get(configObject config.ContextObject) ([]byte,error) {
	method := "GET"
	url := "http://"+configObject.ServerAddress + "/tasks"
	data,err:= serviceCall.Make(configObject,method,url,nil)
	return  data,err
}

func Add(configObject config.ContextObject,task string, priority string)error {
	dataToSend := &contract.Task{}
	dataToSend.Task = &task
	dataToSend.Priority = &priority

	method := "POST"
	url := "http://"+configObject.ServerAddress+"/task"
	_,err := serviceCall.Make(configObject,method,url,dataToSend)
	return err
}

func Delete(configObject config.ContextObject,URI string) error{
	method := "DELETE"
	url := "http://"+configObject.ServerAddress +URI
	_,err := serviceCall.Make(configObject,method,url,nil)
	return err
}

func Update(configObject config.ContextObject,taskId string, taskDescription string, taskPriority string)error {
	id,err := strconv.Atoi(taskId)
	convertedTaskId := int32(id)
	if err != nil {
		errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
	}
	data := &contract.Task{}
	data.TaskId = &convertedTaskId
	data.Task = &taskDescription
	data.Priority = &taskPriority

	method := "POST"
	url := "http://"+configObject.ServerAddress+"/update"
	_,err = serviceCall.Make(configObject,method,url,data)
	return err
}

func AddTaskByCsv(configObject config.ContextObject, csvFileData string)error{
	data := &contract.Task{}
	data.Task = &csvFileData
	method := "POST"
	url := "http://"+configObject.ServerAddress + "/tasks/csv"
	_,err := serviceCall.Make(configObject,method,url,data)
	return err
}