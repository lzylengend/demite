package wx_api

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	//1c34b3cbf36f0fcc8bd6c670503c0201
	appId     = "wx3b5ad2c147bf98b1"
	appSecret = "1c34b3cbf36f0fcc8bd6c670503c0201"
)

type accessRespose struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
}

func GetAccessTokenFromWx() (string, error) {
	rsp, err := clientDo("GET", "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid="+
		appId+
		"&secret="+appSecret, []byte{})
	if err != nil {
		return "", err
	}

	res := &accessRespose{}
	err = json.Unmarshal(rsp, res)
	if err != nil {
		return "", err
	}

	if res.ErrCode != 0 {
		return "", errors.New(strconv.Itoa(res.ErrCode) + res.ErrMsg)
	}

	return res.AccessToken, nil
}

func clientDo(reqMethod string, url string, body []byte) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest(reqMethod, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	bodyRsp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bodyRsp, err
}
