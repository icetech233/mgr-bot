/*
 * Copyright (c) 2024 by hanbindsg.
 */
package main

import (
	"fmt"

	"github.com/icetech233/gin"
	"github.com/icetech233/mgr-bot/subscription"
)

func main() {
	port := 8805
	s := subscription.NewSubscription(port)
	s.OnMessageNormal = onMessageNormal
	s.Router = CustomRouter(s)
	s.Start()
}

func CustomRouter(s *subscription.Subscription) *gin.Engine {
	router := gin.Default()
	router.POST("/custom", func(c *gin.Context) {
		var sr subscription.SubScriptionResp
		if err := c.BindJSON(&sr); err != nil {
			return
		}
		s.Parse(sr)
	})

	return router
}

// 接收普通消息
func onMessageNormal(event subscription.MessageEvent) {
	fmt.Println(event)
}
