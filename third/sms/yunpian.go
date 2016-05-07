package sms

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	ApiKey           = "7293f1f6f2528c622069fa9f4b221bd0" //云片的apikey可从后台看到
	ApiBase          = "https://sms.yunpian.com/v2/"
	SendSingleSMSUrl = "sms/single_send.json"
)

type SMSResponse struct {
	Code   int
	Msg    string
	Count  int
	Fee    float32
	Unit   string
	Mobile string
	Sid    int64
}

func SendSingleSMS(phone, content string) error {
	v := make(url.Values)
	v.Set("apikey", ApiKey)
	v.Set("mobile", phone)
	v.Set("text", content)
	resp, err := HttpRequest("POST", ApiBase+SendSingleSMSUrl, v)
	if err != nil {
		return err
	}
	var r SMSResponse
	err = json.Unmarshal(resp, &r)
	if err != nil {
		return err
	}
	if r.Code == 0 {
		return nil
	} else {
		return errors.New(r.Msg)
	}
}

func HttpRequest(mothod string, url string, params url.Values) (data []byte, err error) {
	var resp *http.Response
	if strings.ToUpper(mothod) == "POST" {
		resp, err = http.PostForm(url, params)
	} else if strings.ToUpper(mothod) == "GET" {
		paramstr := params.Encode()
		if len(paramstr) >= 3 {
			url += "?" + paramstr
		}
		resp, err = http.Get(url)
	} else {
		return nil, errors.New("Not Support Mothod")
	}
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err = ioutil.ReadAll(resp.Body)
	return
}
