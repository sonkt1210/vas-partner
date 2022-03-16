package services

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/sonkt1210/tiki-vas/constant"
	"github.com/sonkt1210/tiki-vas/helpers"
	"github.com/sonkt1210/tiki-vas/logger"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type REST interface {
	Get(url string, headers ...map[string]string) ([]byte, error)
	Post(url string, body []byte, headers ...map[string]string) ([]byte, error)
	Put(url string, body []byte, headers ...map[string]string) ([]byte, error)
	Delete(url string, headers ...map[string]string) ([]byte, error)
}

type rEST struct {
	logger logger.Logger
}

func NewREST(logger logger.Logger) *rEST {
	return &rEST{
		logger: logger,
	}
}

// Get method
func (rest *rEST) Get(url string, headers ...map[string]string) ([]byte, error) {
	res, reqErr := rest.api("GET", url, nil, headers...)
	if reqErr != nil {
		rest.logger.Errorf(
			"RESTService: Fail to make GET request, url = %s, header = %v, error = %s",
			url, headers, reqErr.Error())
	}
	return res, reqErr
}

// Post make POST request
func (rest *rEST) Post(url string, body []byte, headers ...map[string]string) ([]byte, error) {
	res, reqErr := rest.api("POST", url, bytes.NewReader(body), headers...)

	if reqErr != nil {
		rest.logger.Errorf(
			"RESTService : Fail to make POST request, url = %s, body = %s, header = %v, error = %s",
			url, string(body), headers, reqErr.Error())
	}
	return res, reqErr
}

// Put make PUT request
func (rest *rEST) Put(url string, body []byte, headers ...map[string]string) ([]byte, error) {
	res, reqErr := rest.api("PUT", url, bytes.NewReader(body), headers...)
	if reqErr != nil {
		rest.logger.Errorf(
			"RESTService : Fail to make PUT request, url = %s, body = %s, header = %v, error = %s",
			url, string(body), headers, reqErr.Error())
	}
	return res, reqErr
}

// Delete make DELETE request
func (rest *rEST) Delete(url string, headers ...map[string]string) ([]byte, error) {
	res, reqErr := rest.api("DELETE", url, nil, headers...)
	if reqErr != nil {
		rest.logger.Errorf(
			"RESTService: Fail to make DELETE request, url = %s, header = %v, error = %s",
			url, headers, reqErr.Error())
	}
	return res, reqErr
}

// api make a http request
func (rest *rEST) api(method string, url string, body io.Reader, headers ...map[string]string) ([]byte, error) {
	client := http.Client{}
	req, clientErr := http.NewRequest(method, url, body)
	if clientErr != nil {
		return nil, clientErr
	}

	if len(headers) >= 1 {
		header := headers[0]
		for k, v := range header {
			req.Header.Add(k, v)
		}
		if val, ok := header[constant.HeaderRequestTimeout]; ok {
			if timeout := helpers.ParseInt32(val); timeout > 0 {
				client.Timeout = time.Duration(timeout) * time.Second
			}
		}
	}

	resp, reqErr := client.Do(req)
	if reqErr != nil {
		return nil, clientErr
	}
	if resp.StatusCode >= 400 {
		return nil, errors.New(fmt.Sprintf("Request error | status: %d", resp.StatusCode))
	}
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	rest.logger.Infof("resp %s", resp)
	return ioutil.ReadAll(resp.Body)
}
