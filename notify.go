package dingtalk_robot

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func NewRobot(token string) *Robot {
	return &Robot{
		token: token,
	}
}

const baseUrl = "https://oapi.dingtalk.com/robot/send?access_token=%s"

type Robot struct {
	token string
}

func (robot *Robot) SendMessage(msg interface{}) error {
	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		return errors.New(fmt.Sprintf("msg json failed, msg: %v, err: %v", msg, err.Error()))
	}

	body := bytes.NewBuffer([]byte(jsonMsg))
	url := fmt.Sprintf(baseUrl, robot.token)
	res, err := http.Post(url, "application/json;charset=utf-8", body)
	if err != nil {
		return errors.New(fmt.Sprintf("send dingTalk message failed, error: %v", err.Error()))
	}

	if res.StatusCode != 200 {
		return errors.New(fmt.Sprintf("send dingTalk message failed, status code: %v", res.StatusCode))
	}

	defer res.Body.Close()
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errors.New(fmt.Sprintf("send dingTalk message failed, error: %v", err.Error()))
	}

	type response struct {
		ErrCode int `json:"errcode"`
	}

	var ret response

	if err := json.Unmarshal([]byte(result), &ret); err != nil {
		return errors.New(fmt.Sprintf("send dingTalk message failed, result: %s, error: %v", result, err.Error()))
	}

	if ret.ErrCode != 0 {
		return errors.New(fmt.Sprintf("send dingTalk message failed, result: %s", result))
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
