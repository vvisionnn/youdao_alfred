package request

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Request struct {
	Header   map[string]string
	FormData map[string]string
	URL      string
}

func (r *Request) DO(method string) (string, error) {
	data := url.Values{}
	for k, v := range r.FormData {
		data.Set(k, v)
	}

	req, err := http.NewRequest(method, r.URL, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return "", err
	}

	for k, v := range r.Header {
		req.Header.Set(k, v)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println(err)
		}
	}()

	buf, _ := ioutil.ReadAll(resp.Body)

	return string(buf), nil
}

func (r *Request) Get() (string, error) {
	return r.DO("GET")
}

func (r *Request) defaultPOST() (string, error) {
	return r.DO("POST")
}

func (r *Request) POST(url string) (string, error) {
	r.URL = url
	return r.defaultPOST()
}
