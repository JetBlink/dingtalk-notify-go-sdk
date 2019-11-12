# DingTalk Notify
钉钉机器人通知(支持加签）。dingtalk robot notification sdk.

## Overview

* [Installation](#Installation)
* [Usage](#Usage)
  * [获取实例](#获取实例)
  * [发送消息](#发送消息)
    * [发送原始消息](#发送原始消息)
    * [发送文本消息](#发送文本消息)
    * [发送Markdown消息](#发送markdown消息)
    * [发送链接消息](#发送链接消息)
  * [Tips](#Tips)
* [官方文档](#官方文档)
* [License](#license)

## Installation

```
go get github.com/JetBlink/dingtalk-notify-go-sdk
```

## Run Test

* 不带签名机器人
```
ROBOT_TOKEN=your_robot_token go test
```

* 带签名的机器人
```
ROBOT_TOKEN=your_robot_token ROBOT_SECRET=your_robot_secret go test
```

## Usage

### 获取实例

  ```
import (
	dingtalk_robot "github.com/JetBlink/dingtalk-notify-go-sdk"
)

func main() {
	robot := dingtalk_robot.NewRobot("your robot token", "your robot secret")
}
  ```

### 发送消息

#### 发送原始消息

```
msg := map[string]interface{}{
	"msgtype": "text",
	"text": map[string]string{
		"content": "这是一条golang钉钉消息测试.",
	},
	"at": map[string]interface{}{
		"atMobiles": []string{},
		"isAtAll":   false,
	},
}

robot := dingtalk_robot.NewRobot(os.Getenv("ROBOT_TOKEN"), os.Getenv("ROBOT_SECRET"))
if err := robot.SendMessage(msg); err != nil {
	t.Error(err)
}
```

#### 发送文本消息

```
robot.SendTextMessage("普通文本消息", []string{}, false)
```

#### 发送Markdown消息

```
robot.SendMarkdownMessage(
	"Markdown Test Title",
	"### Markdown 测试消息\n* 谷歌: [Google](https://www.google.com/)\n* 一张图片\n ![](https://avatars0.githubusercontent.com/u/40748346)",
	[]string{},
	false,
)
```

#### 发送链接消息

```
robot.SendLinkMessage(
	"Link Test Title",
	"这是一条链接测试消息",
	"https://github.com/JetBlink",
	"https://avatars0.githubusercontent.com/u/40748346",
)
```

### Tips

文本消息和Markdown消息都支持**@指定手机号**和**@所有人**，参数位置见具体方法

## 官方文档

* [自定义机器人](https://ding-doc.dingtalk.com/doc#/serverapi2/qf2nxq)
> 每个机器人每分钟最多发送**20条**。

## License

[MIT](https://opensource.org/licenses/MIT)
