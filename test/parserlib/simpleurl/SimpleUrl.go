package simpleurl

import (
	"io/ioutil"
	"net/http"
)

type simpleUrl struct {
	Proxy string
}

func (s *simpleUrl) Get(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func NewSimpleUrl() *simpleUrl {
	return new(simpleUrl)
}
