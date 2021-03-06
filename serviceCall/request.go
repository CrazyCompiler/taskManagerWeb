package serviceCall

import (
	"github.com/afex/hystrix-go/hystrix"
	"io/ioutil"
	"taskManagerWeb/errorHandler"
	"net/http"
	"taskManagerWeb/config"
	"github.com/golang/protobuf/proto"
	"github.com/CrazyCompiler/taskManagerContract"
	"errors"
	"crypto/tls"
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

func Make(context config.Context,request *http.Request)([]byte,error){
	receiveData := make(chan []byte)
	errorToReturn := make(chan error)
	configureHystrix()
	hystrix.Go("task", func() error {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}

		resp, err := client.Do(request)
		if err != nil {
			errorHandler.ErrorHandler(context.ErrorLogFile,err)
			return err
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			errorHandler.ErrorHandler(context.ErrorLogFile,err)
			return  err
		}
		dataProvided := &contract.Response{}
		err = proto.Unmarshal(body,dataProvided)
		if err != nil {
			errorHandler.ErrorHandler(context.ErrorLogFile,err)
		}
		receiveData <-dataProvided.Data
		if dataProvided.Err == nil{
			errorToReturn <- nil
			return nil
		}
		if dataProvided.Err != nil{
			errorToReturn <- errors.New(*dataProvided.Err.Description)
		}
		return  nil
	}, func(err error) error {
		receiveData <- nil
		errorToReturn <- err
		return nil
	})
	return <-receiveData,<-errorToReturn
}
