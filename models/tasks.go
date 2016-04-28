package models

import (
	"taskManagerWeb/config"
	"net/http"
	"io/ioutil"
	"taskManagerWeb/errorHandler"
	"github.com/golang/protobuf/proto"
	"bytes"
	"strconv"
	"github.com/taskManagerContract"
	"fmt"
)

func Get(configObject config.ContextObject) []byte {
	request, err := http.NewRequest("GET", "http://"+configObject.ServerAddress + "/tasks", nil)
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
		return nil
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
	}
	return  body
}

func Add(configObject config.ContextObject,task string, priority string)error {
	data := &contract.Task{}
	data.Task = &task
	data.Priority = &priority
	dataToBeSend,err :=  proto.Marshal(data)
	if err != nil {
		errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
	}
	request, err := http.NewRequest("POST", "http://"+configObject.ServerAddress+"/task",bytes.NewBuffer(dataToBeSend))
	if err != nil {
		errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
	}
	client := &http.Client{}
	resp, err := client.Do(request)
	fmt.Println(resp)
	return  err
}

func Delete(configObject config.ContextObject,URI string) error{
	fmt.Println("http://"+configObject.ServerAddress + "/task"+URI)
	request, err := http.NewRequest("DELETE", "http://"+configObject.ServerAddress +URI, nil)
	client := &http.Client{}
	_, err = client.Do(request)
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

	dataToBeSend,err :=  proto.Marshal(data)
	if err != nil {
		errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
	}

	request, err := http.NewRequest("POST", "http://"+configObject.ServerAddress+"/update",bytes.NewBuffer(dataToBeSend))
	client := &http.Client{}
	_, err = client.Do(request)
	return  err
}

func AddTaskByCsv(configObject config.ContextObject, csvFileData string)error{
	data := &contract.UploadCsvData{}
	data.CsvData = &csvFileData
	dataToBeSend,err :=  proto.Marshal(data)
	if err != nil {
		errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
	}
	request, err := http.NewRequest("POST", "http://"+configObject.ServerAddress + "/tasks/csv",bytes.NewBuffer(dataToBeSend))
	client := &http.Client{}
	_, err = client.Do(request)
	return  err
}