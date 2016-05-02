package models

import (
	"taskManagerWeb/config"
	"taskManagerWeb/errorHandler"
	"strconv"
	"taskManagerWeb/serviceCall"
	"github.com/taskManagerContract"
)

func Get(context config.Context) ([]byte,error) {
	method := "GET"
	url := context.ServerAddress + "/tasks"
	data,err:= serviceCall.Make(context,method,url,nil)
	return  data,err
}

func Add(context config.Context,task string, priority string)error {
	dataToSend := &contract.Task{}
	dataToSend.Task = &task
	dataToSend.Priority = &priority
	method := "POST"
	url := context.ServerAddress+"/task"
	_,err := serviceCall.Make(context,method,url,dataToSend)
	return err
}

func Delete(context config.Context,URI string) error{
	method := "DELETE"
	url := context.ServerAddress +URI
	_,err := serviceCall.Make(context,method,url,nil)
	return err
}

func Update(context config.Context,taskId string, taskDescription string, taskPriority string)error {
	id,err := strconv.Atoi(taskId)
	convertedTaskId := int32(id)
	if err != nil {
		errorHandler.ErrorHandler(context.ErrorLogFile,err)
	}
	data := &contract.Task{}
	data.TaskId = &convertedTaskId
	data.Task = &taskDescription
	data.Priority = &taskPriority

	method := "POST"
	url := context.ServerAddress+"/update"
	_,err = serviceCall.Make(context,method,url,data)
	return err
}

func AddTaskByCsv(context config.Context, csvFileData string)error{
	data := &contract.Task{}
	data.Task = &csvFileData
	method := "POST"
	url := context.ServerAddress + "/tasks/csv"
	_,err := serviceCall.Make(context,method,url,data)
	return err
}

func GetCsv(context config.Context)([]byte,error){
	method := "GET"
	url := context.ServerAddress + "/tasks/download/csv"
	data,err:= serviceCall.Make(context,method,url,nil)
	return  data,err
}