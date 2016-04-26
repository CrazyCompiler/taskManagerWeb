package main

import (
	"os"
	"fmt"
	"taskManagerWeb/config"
	"taskManagerWeb/errorHandler"
	"taskManagerWeb/routers"
	"net/http"
)

func main() {
	configObject := config.ContextObject{}
	errorLogFilePath := "errorLog"
	errorFile, err := os.OpenFile(errorLogFilePath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer errorFile.Close()

	configObject.ErrorLogFile = errorFile
	configObject.ServerAddress = os.Args[2]

	if err != nil {
		errorHandler.ErrorHandler(configObject.ErrorLogFile, err)
	}

	routers.HandleRequests(configObject)
	err = http.ListenAndServe(":"+os.Args[1], nil)
	if err != nil {
		fmt.Println("their was error ", err)
	}

}
