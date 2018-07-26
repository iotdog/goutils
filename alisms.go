package goutils

import (
	"fmt"

	"github.com/GiterLab/aliyun-sms-go-sdk/dysms"
	"github.com/tobyzxj/uuid"
)

// AliSMSClient 阿里云短信客户端
type AliSMSClient struct {
	keyID     string
	keySecret string
}

// NewAliSMSClient 创建阿里云短信客户端
func NewAliSMSClient(keyID, keySecret string) *AliSMSClient {
	alisms := new(AliSMSClient)
	alisms.keyID = keyID
	alisms.keySecret = keySecret
	return alisms
}

// Send 发送短信
func (client *AliSMSClient) Send(phone, smsParams, templateID, signName string) bool {
	dysms.SetACLClient(client.keyID, client.keySecret)
	respSendSms, err := dysms.SendSms(uuid.New(), phone, signName, templateID, smsParams).DoActionWithException()
	if err != nil {
		fmt.Println("send sms failed", err, respSendSms.Error())
		return false
	}
	fmt.Println("send sms succeed", respSendSms.GetRequestID())
	return true
}
