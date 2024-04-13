/*
 * Copyright (c) 2024 by hanbindsg.
 */
package openapi

const (
	API_BASE_URL = "https://chat-go.jwzhd.com/open-apis/v1"
)

type OpenApi struct {
	Token string `json:"token"`
}

func NewOpenApi(token string) *OpenApi {
	// https://chat-go.jwzhd.com/open-apis/v1/bot/send?token=6820b38681984313bdf12af71ab0a13d
	return &OpenApi{
		Token: token,
	}
}
