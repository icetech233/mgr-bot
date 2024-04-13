/*
 * Copyright (c) 2024 by hanbindsg.
 */
package main

import (
	"fmt"

	"github.com/icetech233/mgr-bot/openapi"
	"github.com/icetech233/mgr-bot/subscription"
)

func main() {
	port := 8805
	subscription := subscription.NewSubscription(port)
	subscription.OnMessageNormal = onMessageNormal
	subscription.OnMessageInstruction = onMessageInstruction
	subscription.OnGroupJoin = onGroupJoin
	subscription.OnGroupLeave = onGroupLeave
	subscription.OnBotFollowed = onBotFollowed
	subscription.Start()
}

// 接收普通消息
func onMessageNormal(event subscription.MessageEvent) {
	fmt.Println("message.receive.normal:\n", event)
}

// 接收指令消息
func onMessageInstruction(event subscription.MessageEvent) {
	fmt.Println("message.receive.instruction:\n", event)
}

// 接收群新成员加入事件
func onGroupLeave(event subscription.GroupLeaveEvent) {
	fmt.Println(event)
}

func onGroupJoin(event subscription.GroupJoinEvent) {
	fmt.Println(event)
}

// 接收机器人被关注事件
func onBotFollowed(event subscription.BotFollowedEvent) {
	fmt.Println(event)
}

/* 示例方法
 * 发送文本消息
 * token来自于云湖官网控制台
 */
func SendTextMessage(recvId string, recvType string, text string) (openapi.BasicResponse, error) {
	openApi := openapi.NewOpenApi("token")
	textMessage := openapi.TextMessage{
		RecvId:   recvId,
		RecvType: recvType,
		Text:     text,
	}
	return openApi.SendTextMessage(textMessage)
}

/* 示例方法
 * 发送文本消息
 * token来自于云湖官网控制台
 * buttons消息结构如下：(参考代码文件：openapi/openapi_test.go)
 * 1、[]openapi.Button{}
 * 2、[][]openapi.Button{}
 */
func SendTextMessage1(recvId string, recvType string, text string, buttons interface{}) (openapi.BasicResponse, error) {
	openApi := openapi.NewOpenApi("token")
	textMessage := openapi.TextMessage{
		RecvId:   recvId,
		RecvType: recvType,
		Text:     text,
		Buttons:  buttons,
	}
	return openApi.SendTextMessage(textMessage)
}

/* 示例方法
 * 编辑文本消息
 * msgId字段是来着发送消息时返回的消息ID
 */
func EditTextMessage(msgId string, recvId string, recvType string, text string) (openapi.BasicResponse, error) {
	openApi := openapi.NewOpenApi("token")
	newTextMessage := openapi.EditTextMessage{
		MsgId:    msgId,
		RecvId:   recvId,
		RecvType: recvType,
		Text:     text,
		Buttons:  nil,
	}
	return openApi.EditTextMessage(newTextMessage)
}
