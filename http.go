package goutils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/leesper/holmes"
)

// Jsonify 返回JSON应答
func Jsonify(w http.ResponseWriter, message interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(message); err != nil {
		holmes.Errorln(err)
	}
}

// JSONGet 发送Get请求，返回数据为JSON格式
func JSONGet(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

// JSONPost 发送Post请求， 返回数据为JSON格式
func JSONPost(url string, jsonData []byte) ([]byte, error) {
	resp, err := http.Post(url, "application/json", bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
