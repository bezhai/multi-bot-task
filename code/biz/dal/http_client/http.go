package http_client

import (
	"context"
	"crypto/tls"
	"time"

	"github.com/bezhai/go_utils/httpx"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/network/standard"
)

var c *client.Client

func Init() {
	var err error
	clientCfg := &tls.Config{
		InsecureSkipVerify: true,
	}
	c, err = client.NewClient(
		client.WithTLSConfig(clientCfg),
		client.WithDialer(standard.NewDialer()),
		client.WithDialTimeout(5*time.Second),
		client.WithClientReadTimeout(120*time.Second),
	)

	if err != nil {
		panic(err)
	}
}

func SendRequest[T any](ctx context.Context, url string, extraFuncList ...httpx.RequestExtraFunc) (respBody T, err error) {
	return httpx.SendRequest[T](ctx, c, url, extraFuncList...)
}

func SendRequestWithHeaderResp[T any](ctx context.Context, url string, extraFuncList ...httpx.RequestExtraFunc) (respBody T, respHeader []map[string]string, err error) {
	return httpx.SendRequestWithHeaderResp[T](ctx, c, url, extraFuncList...)
}
