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
	res, err := http.Get(url)

	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err.Error())
	}

	holmes.Debugln(len(body))

	return body

	// json.Unmarshal(body, &data)
	// holmes.Debugf("Results: %v\n", data)
	// return data
}
