package main

import (
	"os"
	"fmt"
	"taskManagerWeb/config"
	"taskManagerWeb/errorHandler"
	"taskManagerWeb/routers"
	"net/http"
	"flag"
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

	var serverAddressFlag = flag.String("sa","127.0.0.1:8080","listening to the service")
	var portFlag = flag.String("p","8888","To which port it will listen")

	flag.Parse()

	configObject.ServerAddress = *serverAddressFlag
	port := *portFlag

	if err != nil {
		errorHandler.ErrorHandler(configObject.ErrorLogFile, err)
	}

	routers.HandleRequests(configObject)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("their was error ", err)
	}

}