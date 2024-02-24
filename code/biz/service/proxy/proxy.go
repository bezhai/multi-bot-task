package proxy

import (
	"context"

	"github.com/bezhai/go_utils/httpx"

	"github.com/bezhai/multi-bot-task/biz/clients/http_client"
	"github.com/bezhai/multi-bot-task/biz/clients/redis_client"
)

func Proxy(ctx context.Context, url, referer string) (respBody []byte, respHeaders []map[string]string, err error) {
	headers := map[string]string{
		"user-agent":         redis_client.MustGet(ctx, "user-agent"),
		"referer":            referer,
		"cookie":             redis_client.MustGet(ctx, "cookie"),
		"sec-ch-ua":          redis_client.MustGet(ctx, "sec-ch-ua"),
		"sec-ch-ua-mobile":   "?0",
		"sec-ch-ua-platform": "macOS",
	}

	return http_client.SendRequestWithHeaderResp[[]byte](ctx, url, httpx.GetFunc(nil), httpx.AddHeaderFunc(headers))
}
