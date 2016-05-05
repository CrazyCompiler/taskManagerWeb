package models

import (
	"taskManagerWeb/config"
	"taskManagerWeb/errorHandler"
	"taskManagerWeb/serviceCall"
	"github.com/golang/protobuf/proto"
	"bytes"
	"net/http"
	"github.com/CrazyCompiler/taskManagerContract"
)

func createNewRequest(method string, url string, data *contract.Task)(*http.Request,error)  {
	if data == nil{
		data = &contract.Task{}
	}

	dataToBeSend,err :=  proto.Marshal(data)
	request, err := http.NewRequest(method,url, bytes.NewBuffer(dataToBeSend))
	return request,err;
}

func Get(context config.Context,userId string) ([]byte,error) {
	method := "GET"
	url := context.ServerAddress + "/tasks/"+userId
	requestToService,err := createNewRequest(method,url,nil)
	if err != nil {
		errorHandler.ErrorHandler(context.ErrorLogFile,err)
		return nil,err
	}
	data,err:= serviceCall.Make(context,requestToService)
	return  data,err
}

func Add(context config.Context,task string, priority string)error {
	data := &contract.Task{}
	data.Task = &task
	data.Priority = &priority
	method := "POST"
	url := context.ServerAddress+"/tasks"
	requestToService,err := createNewRequest(method,url,data)
	if err != nil {
		errorHandler.ErrorHandler(context.ErrorLogFile,err)
		return err
	}
	_,err = serviceCall.Make(context,requestToService)
	return err
}

func Delete(context config.Context,URI string) error{
	method := "DELETE"
	url := context.ServerAddress +URI
	requestToService,err := createNewRequest(method,url,nil)
	if err != nil {
		errorHandler.ErrorHandler(context.ErrorLogFile,err)
		return err
	}
	_,err = serviceCall.Make(context,requestToService)
	return err
}

func Update(context config.Context,taskId string, taskDescription string, taskPriority string)error {
	data := &contract.Task{}
	data.Task = &taskDescription
	data.Priority = &taskPriority

	method := "PATCH"
	url := context.ServerAddress+taskId
	requestToService,err := createNewRequest(method,url,data)
	if err != nil {
		errorHandler.ErrorHandler(context.ErrorLogFile,err)
		return err
	}
	_,err = serviceCall.Make(context,requestToService)
	return err
}

func AddTaskByCsv(context config.Context, csvFileData string)error{
	data := &contract.Task{}
	data.Task = &csvFileData
	method := "POST"
	url := context.ServerAddress + "/tasks/csv"
	requestToService,err := createNewRequest(method,url,data)
	if err != nil {
		errorHandler.ErrorHandler(context.ErrorLogFile,err)
		return err
	}
	_,err = serviceCall.Make(context,requestToService)
	return err
}

func GetCsv(context config.Context)([]byte,error){
	method := "GET"
	url := context.ServerAddress + "/tasks/csv"
	requestToService,err := createNewRequest(method,url,nil)
	if err != nil {
		errorHandler.ErrorHandler(context.ErrorLogFile,err)
		return []byte(""),err
	}
	data,err:= serviceCall.Make(context,requestToService)
	return  data,err
}