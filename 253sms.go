package goutils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// C253SMSClient 创蓝短信接口客户端
type C253SMSClient struct {
	Username string
	Password string
}

// NewC253SMSClient 创建短信客户端
func NewC253SMSClient(username, pwd string) *C253SMSClient {
	return &C253SMSClient{
		username,
		pwd,
	}
}

type c253APIResp struct {
	Code     string `json:"code"`
	MsgID    string `json:"msgId"`
	Time     string `json:"time"`
	ErrorMsg string `json:"errorMsg"`
}

// Send 发送短信，其中msg类似：【253云通讯】您好，您的验证码是999999
func (c253 *C253SMSClient) Send(phone, msg string) bool {
	params := make(map[string]interface{})

	params["account"] = c253.Username    //创蓝API账号
	params["password"] = c253.Password   //创蓝API密码
	params["phone"] = phone              // 手机号码
	params["msg"] = url.QueryEscape(msg) // 短信内容

	bytesData, err := json.Marshal(params)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	reader := bytes.NewReader(bytesData)
	url := "http://smssh1.253.com/msg/send/json" //短信发送URL
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	apiResp := c253APIResp{}
	err = json.Unmarshal(respBytes, &apiResp)
	if err != nil {
		fmt.Println(err.Error())
	}
	if "0" != apiResp.Code {
		fmt.Println(apiResp.ErrorMsg)
		return false
	}
	return true
}
