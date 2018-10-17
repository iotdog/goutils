package goutils

import (
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
func JSONGet(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err.Error())
	}

	return body
}
