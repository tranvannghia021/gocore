package service

import (
	"github.com/tranvannghia021/gocore/vars"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type IHttpRequest interface {
	HandleErrorReq(body []byte) vars.ResReq
	GetRequest() vars.ResReq
	PostRequest() vars.ResReq
	PostFormDataRequest() vars.ResReq
	PutRequest() vars.ResReq
	DeleteRequest() vars.ResReq
	SetAuth(token string) *SHttpRequest
}
type SHttpRequest struct {
	Url             string            `json:"url,omitempty"`
	Headers         map[string]string `json:"headers,omitempty"`
	Body            io.Reader         `json:"body,omitempty"`
	FormData        url.Values        `json:"form_data,omitempty"`
	Err             error             `json:"err,omitempty"`
	HeadersResponse map[string][]string
}

func NewHttpRequest() *SHttpRequest {
	return &SHttpRequest{
		HeadersResponse: make(map[string][]string),
	}

}

func (s *SHttpRequest) SetAuth(token string) *SHttpRequest {
	if s.Headers == nil {
		s.Headers = make(map[string]string)
	}
	s.Headers["Authorization"] = "Bearer " + token
	return s
}
func (s *SHttpRequest) HandleErrorReq(body []byte) vars.ResReq {
	if s.Err != nil {

		return vars.ResReq{
			Status: false,
			Data:   nil,
			Error:  s.Err,
		}
	}

	res := vars.ResReq{
		Status:     true,
		Data:       body,
		Error:      nil,
		HeadersRes: make(map[string][]string),
	}
	res.HeadersRes = s.HeadersResponse
	return res
}
func (s *SHttpRequest) GetRequest() vars.ResReq {
	req, err := http.NewRequest("GET", s.Url, nil)
	for k, v := range s.Headers {
		req.Header.Add(k, v)
	}
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	s.Err = err
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	s.Err = err

	for key, value := range res.Header {
		s.HeadersResponse[key] = value
	}
	return s.HandleErrorReq(body)
}

func (s *SHttpRequest) PostRequest() vars.ResReq {
	req, err := http.NewRequest("POST", s.Url, s.Body)
	for k, v := range s.Headers {
		req.Header.Add(k, v)
	}
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	s.Err = err
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	s.Err = err
	return s.HandleErrorReq(data)
}

func (s *SHttpRequest) PostFormDataRequest() vars.ResReq {
	req, err := http.NewRequest("POST", s.Url, strings.NewReader(s.FormData.Encode()))
	for k, v := range s.Headers {
		req.Header.Add(k, v)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	res, err := client.Do(req)
	s.Err = err
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	s.Err = err
	return s.HandleErrorReq(data)
}

func (s *SHttpRequest) PutRequest() vars.ResReq {
	req, err := http.NewRequest("PUT", s.Url, s.Body)
	for k, v := range s.Headers {
		req.Header.Add(k, v)
	}
	client := &http.Client{}
	res, err := client.Do(req)
	s.Err = err
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	s.Err = err
	return s.HandleErrorReq(data)
}

func (s *SHttpRequest) DeleteRequest() vars.ResReq {
	req, err := http.NewRequest("DELETE", s.Url, s.Body)
	for k, v := range s.Headers {
		req.Header.Add(k, v)
	}
	client := &http.Client{}
	res, err := client.Do(req)
	s.Err = err
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	s.Err = err
	return s.HandleErrorReq(data)
}
