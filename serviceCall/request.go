package serviceCall

import (
	"github.com/afex/hystrix-go/hystrix"
	"io/ioutil"
	"taskManagerWeb/errorHandler"
	"net/http"
	"taskManagerWeb/config"
	"github.com/golang/protobuf/proto"
	"bytes"
	"github.com/taskManagerContract"
)

const TIMEOUT int = 1000
const MAXCONCURRENTREQUESTS int = 100
const ERRORPERCENTTHRESHOLD int = 25

func configureHystrix(){
	hystrix.ConfigureCommand("task", hystrix.CommandConfig{
		Timeout: TIMEOUT,
		MaxConcurrentRequests: MAXCONCURRENTREQUESTS,
		ErrorPercentThreshold: ERRORPERCENTTHRESHOLD,
	})
}


func Make(context config.Context,method string, url string, data *contract.Task)([]byte,error){
	receiveData := make(chan []byte)
	errorToReturn := make(chan error)
	configureHystrix()
	hystrix.Go("task", func() error {
		if data == nil{
			data = &contract.Task{}
		}

		dataToBeSend,err :=  proto.Marshal(data)
		if err != nil {
			errorHandler.ErrorHandler(context.ErrorLogFile,err)
		}
		request, err := http.NewRequest(method,url, bytes.NewBuffer(dataToBeSend))
		client := &http.Client{}
		resp, err := client.Do(request)
		if err != nil {
			errorHandler.ErrorHandler(context.ErrorLogFile,err)
			return nil
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			errorHandler.ErrorHandler(context.ErrorLogFile,err)
		}
		receiveData <- body
		errorToReturn <- nil
		return  nil
	}, func(err error) error {
		receiveData <- nil
		errorToReturn <- err
		return nil
	})
	return <-receiveData,<-errorToReturn
}
