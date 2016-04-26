package models

import (
	"taskManagerWeb/config"
	"net/http"
	"io/ioutil"
	"taskManagerWeb/errorHandler"
	"net/url"
	"mime/multipart"
	"fmt"
)

func Get(configObject config.ContextObject) []byte {
	request, err := http.NewRequest("GET", "http://"+configObject.ServerAddress + "/getAllTasks", nil)
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return  body
}

func Add(configObject config.ContextObject,task string, priority string)error {
	_, err := http.PostForm("http://"+configObject.ServerAddress+"/addTask", url.Values{"task": {task}, "priority": {priority}})
	if err != nil {
		errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
		return err
	}
	return nil
}

func Delete(configObject config.ContextObject,URI string) error{
	request, err := http.NewRequest("DELETE", "http://"+configObject.ServerAddress + URI, nil)
	client := &http.Client{}
	_, err = client.Do(request)
	return err
}

func UpdatePriority(configObject config.ContextObject,taskId string, priority string)error {
	_, err := http.PostForm("http://"+configObject.ServerAddress+"/updatePriority", url.Values{"taskId": {taskId}, "priority": {priority}})
	if err != nil {
		errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
		return err
	}
	return nil
}

func UpdateDescription(configObject config.ContextObject,taskId string, data string)error {
	_, err := http.PostForm("http://"+configObject.ServerAddress+"/updateTaskDescription", url.Values{"taskId": {taskId}, "data": {data}})
	if err != nil {
		errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
		return err
	}
	return nil
}

func UploadCsv(configObject config.ContextObject, multiform *multipart.Form)error{
	//file := multiform.File
	request, err := http.NewRequest("POST", "http://"+configObject.ServerAddress + "/uploadCsv",nil)
	fmt.Println("file contents:",multiform.File["uploadFile"])
	request.MultipartForm = multiform
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	return  err
}