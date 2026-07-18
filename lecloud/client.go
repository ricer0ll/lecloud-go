package lecloud

import (
	"fmt"

	"resty.dev/v3"
)

type Client struct {
	RestyClient *resty.Client
	Credentials *Credentials
	APIKey      string
	JWT         string
}

func NewClient(credentials *Credentials, apiKey string) *Client {
	restyClient := resty.New()
	restyClient.SetHeader("api-key", apiKey)

	client := &Client{
		RestyClient: restyClient,
		Credentials: credentials,
	}

	restyClient.SetAuthToken(client.getJWT())
	return client
}

func (c Client) getJWT() string {
	req := GetJWTRequest{
		Email:    c.Credentials.Username,
		Password: c.Credentials.Password,
	}
	resp := GetJWTResponse{}

	result, err := c.RestyClient.R().
		SetBody(req).
		SetResult(&resp).
		Post("https://api.riceroll.fyi/v1/jwt")

	if err != nil {
		panic(err)
	}
	if result.IsStatusFailure() {
		panic(fmt.Errorf("Failed to validate credentials: %s", result.String()))
	}

	return resp.Token
}

func (c Client) GetSecret(secretID string) string {
	resp := GetSecretResponse{}

	result, err := c.RestyClient.R().
		SetPathParams(map[string]string{
			"account_id": c.Credentials.AccountID,
			"id":         secretID,
		}).
		SetResult(&resp).
		Get("https://api.riceroll.fyi/v1/accounts/{account_id}/secrets/{id}/decrypt")
	if err != nil {
		panic(err)
	}
	if result.IsStatusFailure() {
		panic(fmt.Errorf("Failed to get secret: %s", result.String()))
	}

	return resp.SecretValue
}
