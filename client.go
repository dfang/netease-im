package netease-im

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

const NETEASE_BASE_URL = "https://api.netease.im/nimserver"

type NeteaseIM struct {
	AppKey    string
	AppSecret string
	client    http.Client
}

func Init(appKey, appSecret string) *NeteaseIM {
	c := NeteaseIM{
		AppKey:    appKey,
		AppSecret: appSecret,
		client:    http.Client{},
	}
	return &c
}

func (im *NeteaseIM) DoNeteaseImRequest(url string, postData io.Reader) (string, error) {

	req, err := http.NewRequest("POST", url, postData)
	if err != nil {
		panic(err)
	}

	im.AddHeaderToRequest(req)

	resp, err := im.client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(respBody), nil
}

// AddHeaderToRequest 添加header到请求
func (im *NeteaseIM) AddHeaderToRequest(req *http.Request) {
	curTime := strconv.FormatInt(time.Now().Unix(), 10)
	nonce := strconv.FormatInt(rand.Int63(), 10)
	hash := GetCheckSum(im.AppSecret, nonce, curTime)

	req.Header.Add("AppKey", im.AppKey)
	req.Header.Add("Nonce", nonce)
	req.Header.Add("CurTime", curTime)
	req.Header.Add("CheckSum", hash)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
}

// GetCheckSum 计算checkSum
func GetCheckSum(appSecret string, nonce string, curTime string) string {
	s := appSecret + nonce + curTime
	h := sha1.New()
	h.Write([]byte(s))
	sum := h.Sum(nil)
	// hash := fmt.Sprintf("%x", sum)
	hash := hex.EncodeToString(sum)
	return hash
}
