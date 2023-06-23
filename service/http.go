package service

import (
	"github.com/tranvannghia021/gocore/config"
	"io"
	"io/ioutil"
	"net/http"
)

func HandleErrorReg(err error) config.ResReq {
	return config.ResReq{
		Status: false,
		Data:   nil,
		Error:  err,
	}
}
func GetRequest(url string, headers map[string]string) config.ResReq {
	req, err := http.NewRequest("GET", url, nil)
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	client := &http.Client{}
	res, err := client.Do(req)
	HandleErrorReg(err)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	HandleErrorReg(err)
	return config.ResReq{
		Status: true,
		Data:   body,
		Error:  nil,
	}
}

func PostRequest(url string, headers map[string]string, body io.Reader) config.ResReq {
	req, err := http.NewRequest("POST", url, body)
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	client := &http.Client{}
	res, err := client.Do(req)
	HandleErrorReg(err)
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	HandleErrorReg(err)
	return config.ResReq{
		Status: true,
		Data:   data,
		Error:  nil,
	}
}

func PostFormDataRequest(url string, headers map[string]string, body io.Reader) config.ResReq {
	req, err := http.NewRequest("POST", url, body)
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	res, err := client.Do(req)
	HandleErrorReg(err)
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	HandleErrorReg(err)
	return config.ResReq{
		Status: true,
		Data:   data,
		Error:  nil,
	}
}

func PutRequest(url string, headers map[string]string, body io.Reader) config.ResReq {
	req, err := http.NewRequest("PUT", url, body)
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	client := &http.Client{}
	res, err := client.Do(req)
	HandleErrorReg(err)
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	HandleErrorReg(err)
	return config.ResReq{
		Status: true,
		Data:   data,
		Error:  nil,
	}
}

func DeleteRequest(url string, headers map[string]string, body io.Reader) config.ResReq {
	req, err := http.NewRequest("DELETE", url, body)
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	client := &http.Client{}
	res, err := client.Do(req)
	HandleErrorReg(err)
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	HandleErrorReg(err)
	return config.ResReq{
		Status: true,
		Data:   data,
		Error:  nil,
	}
}
