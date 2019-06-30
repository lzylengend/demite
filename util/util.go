package util

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func Md5(str string) string {
	w := md5.New()
	io.WriteString(w, str)               //将str写入到w中
	return fmt.Sprintf("%x", w.Sum(nil)) //w.Sum(nil)将w的hash转成[]byte格式
}

func ClientDo(reqMethod string, url string, body []byte, header map[string]string) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest(reqMethod, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	for k, v := range header {
		req.Header.Add(k, v)
	}

	resp, err := client.Do(req)
	defer resp.Body.Close()

	bodyRsp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bodyRsp, err
}
