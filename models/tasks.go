package models

import (
	"taskManagerWeb/config"
	"net/http"
	"io/ioutil"
	"taskManagerWeb/errorHandler"
	"fmt"
	"github.com/golang/protobuf/proto"
	"bytes"
	"github.com/CrazyCompiler/taskManagerContract"
	"strconv"
)

func Get(configObject config.ContextObject) []byte {
	request, err := http.NewRequest("GET", "http://"+configObject.ServerAddress + "/getAllTasks", nil)
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
	data := &contract.AddTask{}
	data.Task = &task
	data.Priority = &priority
	dataToBeSend,err :=  proto.Marshal(data)
	if err != nil {
		errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
	}
	request, err := http.NewRequest("POST", "http://"+configObject.ServerAddress+"/addTask",bytes.NewBuffer(dataToBeSend))
	if err != nil {
		errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
	}
	client := &http.Client{}
	_, err = client.Do(request)

	return  err
}

func Delete(configObject config.ContextObject,URI string) error{
	request, err := http.NewRequest("DELETE", "http://"+configObject.ServerAddress + URI, nil)
	client := &http.Client{}
	_, err = client.Do(request)
	return err
}

func UpdatePriority(configObject config.ContextObject,taskId string, priority string)error {
	id,err := strconv.Atoi(taskId)
	convertedTaskId := int32(id)
	if err != nil {
		errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
	}
	data := &contract.UpdatePriority{}
	data.TaskId = &convertedTaskId
	data.Priority = &priority

	dataToBeSend,err :=  proto.Marshal(data)
	if err != nil {
		fmt.Println("error :",err.Error())
	}

	request, err := http.NewRequest("POST", "http://"+configObject.ServerAddress+"/updatePriority",bytes.NewBuffer(dataToBeSend))
	client := &http.Client{}
	_, err = client.Do(request)
	return  err
}

func UpdateDescription(configObject config.ContextObject,taskId string, taskDescription string)error {
	id,err := strconv.Atoi(taskId)
	convertedTaskId := int32(id)
	if err != nil {
		errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
	}
	data := &contract.UpdateTaskDescription{}
	data.TaskId = &convertedTaskId
	data.Data = &taskDescription

	dataToBeSend,err :=  proto.Marshal(data)
	if err != nil {
		fmt.Println("error :",err.Error())
	}

	request, err := http.NewRequest("POST", "http://"+configObject.ServerAddress+"/updateTaskDescription",bytes.NewBuffer(dataToBeSend))
	client := &http.Client{}
	_, err = client.Do(request)
	return  err
}

func AddTaskByCsv(configObject config.ContextObject, csvFileData string)error{

	data := &contract.UploadCsvData{}
	data.CsvData = &csvFileData
	dataToBeSend,err :=  proto.Marshal(data)
	if err != nil {
		fmt.Println("error :",err.Error())
	}

	request, err := http.NewRequest("POST", "http://"+configObject.ServerAddress + "/uploadCsv",bytes.NewBuffer(dataToBeSend))
	client := &http.Client{}
	_, err = client.Do(request)
	return  err





	//file := multiform.File
	//request, err := http.NewRequest("POST", "http://"+configObject.ServerAddress + "/uploadCsv",nil)
	//fmt.Println("file contents:",multiform.File["uploadFile"])
	//request.MultipartForm = multiform
	//client := &http.Client{}
	//resp, err := client.Do(request)
	//if err != nil {
	//	errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
	//}
	//defer resp.Body.Close()
	//
	//_, err = ioutil.ReadAll(resp.Body)
	//return  err
}