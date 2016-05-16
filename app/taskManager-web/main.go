package main

import (
	"os"
	"taskManagerWeb/config"
	"taskManagerWeb/errorHandler"
	"taskManagerWeb/routers"
	"flag"
	"crypto/tls"
)


func createTlsConnection(context config.Context){
	cert, _ := tls.LoadX509KeyPair("./server.cert", "./server.key")
	config := tls.Config{Certificates: []tls.Certificate{cert}}
	_, err := tls.Listen("tcp", "127.0.0.1:8000", &config)
	if err != nil {
		errorHandler.ErrorHandler(context.ErrorLogFile,err)
	}
}


func main() {
	context := config.Context{}
	errorLogFilePath := "errorLog"
	errorFile, err := os.OpenFile(errorLogFilePath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer errorFile.Close()

	context.ErrorLogFile = errorFile
	createTlsConnection(context)

	var serverAddressFlag = flag.String("sa","https://localhost:8080/","listening to the service")
	var portFlag = flag.String("p","8888","To which port it will listen")

	flag.Parse()

	context.ServerAddress = *serverAddressFlag
	port := *portFlag

	if err != nil {
		errorHandler.ErrorHandler(context.ErrorLogFile, err)
	}

	routers.HandleRequests(context,port)

}