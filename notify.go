package dingtalk_robot

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func NewRobot(token, secret string) *Robot {
	return &Robot{
		token:  token,
		secret: secret,
	}
}

const (
	baseUrl = "https://oapi.dingtalk.com/robot/send?access_token=%s&timestamp=%d&sign=%s"
)

func sign(t int64, secret string) string {
	str := fmt.Sprintf("%d\n%s", t, secret)
	hash := sha256.New()
	hash.Write([]byte(str))
	data := hash.Sum(nil)

	bs64Encoder := base64.StdEncoding
	bs64 := bs64Encoder.EncodeToString(data)
	src := make([]byte, bs64Encoder.EncodedLen(len(data)))
	bs64Encoder.Encode(src, data)

	urlEncoder := base64.URLEncoding
	return urlEncoder.EncodeToString([]byte(bs64))
}

type Robot struct {
	token, secret string
}

func (robot *Robot) SendMessage(msg interface{}) error {
	body := bytes.NewBuffer(nil)
	err := json.NewEncoder(body).Encode(msg)
	if err != nil {
		return fmt.Errorf("msg json failed, msg: %v, err: %v", msg, err.Error())
	}

	t := time.Now().Unix()

	url := fmt.Sprintf(baseUrl, robot.token, t, sign(t, robot.secret))

	res, err := http.Post(url, "application/json;charset=utf-8", body)
	if err != nil {
		return fmt.Errorf("send dingTalk message failed, error: %v", err.Error())
	}

	if res.StatusCode != 200 {
		return fmt.Errorf("send dingTalk message failed, status code: %v", res.StatusCode)
	}

	defer func() { _ = res.Body.Close() }()
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("send dingTalk message failed, error: %v", err.Error())
	}

	type response struct {
		ErrCode int `json:"errcode"`
	}
	var ret response

	if err := json.Unmarshal(result, &ret); err != nil {
		return fmt.Errorf("send dingTalk message failed, result: %s, error: %v", result, err.Error())
	}

	if ret.ErrCode != 0 {
		return fmt.Errorf("send dingTalk message failed, result: %s", result)
	}

	return nil
}

func (robot *Robot) SendTextMessage(content string, atMobiles []string, isAtAll bool) error {
	msg := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]string{
			"content": content,
		},
		"at": map[string]interface{}{
			"atMobiles": atMobiles,
			"isAtAll":   isAtAll,
		},
	}

	return robot.SendMessage(msg)
}

func (robot *Robot) SendMarkdownMessage(title string, text string, atMobiles []string, isAtAll bool) error {
	msg := map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]string{
			"title": title,
			"text":  text,
		},
		"at": map[string]interface{}{
			"atMobiles": atMobiles,
			"isAtAll":   isAtAll,
		},
	}

	return robot.SendMessage(msg)
}

func (robot *Robot) SendLinkMessage(title string, text string, messageUrl string, picUrl string) error {
	msg := map[string]interface{}{
		"msgtype": "link",
		"link": map[string]string{
			"title":      title,
			"text":       text,
			"messageUrl": messageUrl,
			"picUrl":     picUrl,
		},
	}

	return robot.SendMessage(msg)
}
