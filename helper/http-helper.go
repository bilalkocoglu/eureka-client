package helper

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

type OperationBody interface {
}

func MakeGetCall(urlToGet string, headers map[string]string, params map[string]string, onlyErrorLog bool) (error, *http.Response) {
	if !onlyErrorLog {
		log.Info().Msg(fmt.Sprintf("In MakeGetCall to %s", urlToGet))
	}
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}

	paramStr := ""
	if params != nil {
		for key, value := range params {
			paramStr += key + "=" + value + ","
		}
		paramRunes := []rune(paramStr)
		paramStr = string(paramRunes[0 : len(paramRunes)-1])
	}
	if paramStr != "" {
		urlToGet += "?" + paramStr
	}
	req, err := http.NewRequest(http.MethodGet, urlToGet, nil)
	if err != nil {
		log.Print("Error while creating the http request  " + err.Error())
		return errors.New("Error while creating http request  " + err.Error()), nil
	}

	req.Header.Add("Content-Type", "application/json")
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Print("Error while making post call " + err.Error())
		return errors.New("Error while making post call " + err.Error()), nil
	}
	if !onlyErrorLog {
		log.Print("Successful GET call with HTTP Status : " + resp.Status)
	}

	return nil, resp
}

func MakePostCall(urlToPost string, body OperationBody, headers map[string]string, onlyErrorLog bool) (error, *http.Response) {
	if !onlyErrorLog {
		log.Info().Msg(fmt.Sprintf("In MakePostCall to %s with body %s", urlToPost, body))
	}
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	err := encoder.Encode(body)
	if err != nil {
		log.Print("error while encoding " + err.Error())
	}

	if !onlyErrorLog {
		log.Printf("Request body  %+v", buffer.String())
	}

	client := &http.Client{Transport: tr}
	req, err := http.NewRequest(http.MethodPost, urlToPost, &buffer)
	if err != nil {
		log.Print("Error while creating the http request  " + err.Error())
		return errors.New("Error while creating http request  " + err.Error()), nil
	}

	req.Header.Add("Content-Type", "application/json")
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Print("Error while making post call " + err.Error())
		return errors.New("Error while making post call " + err.Error()), nil
	}
	if !onlyErrorLog {
		log.Print("Successful POST call with HTTP Status : " + resp.Status)
	}

	return nil, resp
}

func MakePutCall(urlToPut string, body OperationBody, headers map[string]string, onlyErrorLog bool) (error, *http.Response) {
	if !onlyErrorLog {
		log.Info().Msg(fmt.Sprintf("In MakePutCall to %s with body %s", urlToPut, body))
	}
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	var buffer bytes.Buffer
	if body != nil {
		encoder := json.NewEncoder(&buffer)
		//encoder.SetIndent(" ", "\t")
		err := encoder.Encode(body)
		if err != nil {
			log.Print("error while encoding " + err.Error())
		}
		if !onlyErrorLog {
			log.Printf("Prepared Request body  %+v", buffer.String())
		}
	}

	client := &http.Client{Transport: tr}
	req, err := http.NewRequest(http.MethodPut, urlToPut, &buffer)
	if err != nil {
		log.Print("Error while creating the http request  " + err.Error())
		return errors.New("Error while creating http request " + err.Error()), nil
	}

	req.Header.Add("Content-Type", "application/json")
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Print("Error while making PUT call " + err.Error())
		return errors.New("Error while making PUT call " + err.Error()), nil
	}

	if !onlyErrorLog {
		log.Print("Successful PUT call with HTTP Status : " + resp.Status)
	}

	return nil, resp
}
