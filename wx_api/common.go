package wx_api

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"demite/util"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"strconv"
)

const (
	//wx3b5ad2c147bf98b1
	//1c34b3cbf36f0fcc8bd6c670503c0201
	appId     = "wx61f3efbb27ee263d"
	appSecret = "ac4e20c220ba6bc53618287bc8ff778a"
)

type accessRespose struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
}

type sessionRespose struct {
	Openid     string `json:"openid"`      //用户唯一标识
	SessionKey string `json:"session_key"` //会话密钥
	Unionid    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

func GetAccessTokenFromWx() (string, error) {
	rsp, err := util.ClientDo("GET", "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid="+
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

func CodeToSession(code string) (*sessionRespose, error) {
	rsp, err := util.ClientDo("GET", "https://api.weixin.qq.com/sns/jscode2session?appid="+appId+
		"&secret="+appSecret+
		"&js_code="+code+"&grant_type=authorization_code", []byte{})
	if err != nil {
		return nil, err
	}

	res := &sessionRespose{}
	err = json.Unmarshal(rsp, res)
	if err != nil {
		return nil, err
	}

	if res.ErrCode != 0 {
		return nil, errors.New(strconv.Itoa(res.ErrCode) + res.ErrMsg)
	}

	return res, nil
}

func Decrypt(encryptedData, iv, sessionId string) error {
	ivBytes, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return err
	}

	sessionIdBytes, err := base64.StdEncoding.DecodeString(sessionId)
	if err != nil {
		return err
	}

	res, err := aesDecrypt(encryptedData, sessionIdBytes, ivBytes)
	if err != nil {
		return err
	}
	fmt.Println(string(res))
	return nil
}

func aesEncrypt(encodeStr string, key []byte, iv []byte) (string, error) {
	encodeBytes := []byte(encodeStr)
	//根据key 生成密文
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	encodeBytes = pKCS5Padding(encodeBytes, blockSize)

	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(encodeBytes))
	blockMode.CryptBlocks(crypted, encodeBytes)

	return base64.StdEncoding.EncodeToString(crypted), nil
}

func pKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	//填充
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(ciphertext, padtext...)
}

func aesDecrypt(decodeStr string, key []byte, iv []byte) ([]byte, error) {
	//先解密base64
	decodeBytes, err := base64.StdEncoding.DecodeString(decodeStr)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(decodeBytes))

	blockMode.CryptBlocks(origData, decodeBytes)
	origData = pKCS5UnPadding(origData)
	return origData, nil
}

func pKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
