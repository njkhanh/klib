package restful

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

type DefaultMessage struct {
	ErrorCode int         `json:"errorCode"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}

func Sent(w http.ResponseWriter, data []byte, err error, params map[string]string) {
	if err == nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")

		if len(params) > 0 {
			for key, val := range params {
				w.Header().Set(key, val)
			}
		}

		w.WriteHeader(http.StatusOK)
		_, err = w.Write(data)
	} else {
		w.WriteHeader(http.StatusExpectationFailed)
		_, _ = fmt.Fprintln(w, err)
	}
}

func Post(url string, data []byte) ([]byte, error) {
	result := *new([]byte)
	body := bytes.NewReader(data)

	req, err := http.NewRequest("POST", url, body)

	if err == nil {
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)

		if err == nil {
			defer resp.Body.Close()
			return ioutil.ReadAll(resp.Body)

		}
	}

	return result, err
}

func Get(url string, data []byte) ([]byte, error) {
	result := *new([]byte)
	body := bytes.NewReader(data)

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, body)

	if err == nil {
		var resp *http.Response
		resp, err = client.Do(req)

		if err == nil {
			defer resp.Body.Close()
			return ioutil.ReadAll(resp.Body)
		}

	}

	return result, err
}

func Put(url string, data []byte) ([]byte, error) {
	result := *new([]byte)
	body := bytes.NewReader(data)

	req, err := http.NewRequest("PUT", url, body)

	if err == nil {
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)

		if err == nil {
			defer resp.Body.Close()
			return ioutil.ReadAll(resp.Body)
		}
	}

	return result, err
}

func Delete(url string, data []byte) ([]byte, error) {
	result := *new([]byte)
	body := bytes.NewReader(data)

	req, err := http.NewRequest("DELETE", url, body)

	if err == nil {
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)

		if err == nil {
			defer resp.Body.Close()
			return ioutil.ReadAll(resp.Body)
		}
	}

	return result, err
}
