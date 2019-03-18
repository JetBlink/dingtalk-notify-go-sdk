# DingTalk Notify
钉钉机器人通知。dingtalk robot notification sdk.

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

## Usage

### 获取实例

  ```
import (
	dingtalk_robot "github.com/JetBlink/dingtalk-notify-go-sdk"
)

func main() {
	robot := dingtalk_robot.NewRobot("your dingtalk robot token")
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

robot := dingtalk_robot.NewRobot(os.Getenv("ROBOT_TOKEN"))
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

文本消息和Markdown消息都支持**@指定手机号**和**@所有人**，参数位置见具体方法。

## 官方文档

* [自定义机器人](https://open-doc.dingtalk.com/docs/doc.htm?spm=a219a.7629140.0.0.karFPe&treeId=257&articleId=105735&docType=1)
* [消息类型及数据格式](https://open-doc.dingtalk.com/docs/doc.htm?treeId=172&articleId=104972&docType=1)
> 每个机器人每分钟最多发送**20条**。

## License

[MIT](https://opensource.org/licenses/MIT)
