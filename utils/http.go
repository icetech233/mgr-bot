/*
 * Copyright (c) 2024 by hanbindsg.
 */
package utils

import (
	"github.com/go-resty/resty/v2"
)

// 发送POST请求
// 重试3次
func HttpPost(url string, data []byte) (resp *resty.Response, err error) {
	client := resty.New()
	resp, err = client.SetRetryCount(3).R().
		SetHeader("Content-Type", "application/json").
		SetBody(data).
		Post(url)
	return resp, err
}
