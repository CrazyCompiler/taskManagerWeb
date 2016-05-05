package main

import (
	"os"
	"taskManagerWeb/config"
	"taskManagerWeb/errorHandler"
	"taskManagerWeb/routers"
	"flag"
)

func main() {
	context := config.Context{}
	errorLogFilePath := "errorLog"
	errorFile, err := os.OpenFile(errorLogFilePath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer errorFile.Close()

	context.ErrorLogFile = errorFile

	var serverAddressFlag = flag.String("sa","127.0.0.1:8080","listening to the service")
	var portFlag = flag.String("p","8888","To which port it will listen")

	flag.Parse()

	context.ServerAddress = "http://"+*serverAddressFlag
	port := *portFlag

	if err != nil {
		errorHandler.ErrorHandler(context.ErrorLogFile, err)
	}

	routers.HandleRequests(context,port)

}