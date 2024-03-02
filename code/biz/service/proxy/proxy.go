package proxy

import (
	"context"

	"github.com/bezhai/go_utils/httpx"

	"github.com/bezhai/multi-bot-task/biz/clients/http_client"
	"github.com/bezhai/multi-bot-task/biz/clients/redis_client"
	"github.com/bezhai/multi-bot-task/biz/model/data_trans"
	"github.com/bezhai/multi-bot-task/biz/utils/langx/mapx"
)

func Proxy(ctx context.Context, req *data_trans.ProxyRequest) (respBody []byte, respHeaders []map[string]string, err error) {
	headers := map[string]string{
		"user-agent":         redis_client.MustGet(ctx, "user-agent"),
		"referer":            req.Referer,
		"cookie":             redis_client.MustGet(ctx, "cookie"),
		"sec-ch-ua":          redis_client.MustGet(ctx, "sec-ch-ua"),
		"sec-ch-ua-mobile":   "?0",
		"sec-ch-ua-platform": "macOS",
	}

	headers = mapx.ConcatReplace(headers, req.Headers)

	return http_client.SendRequestWithHeaderResp[[]byte](ctx, req.URL,
		httpx.GetFunc(nil), httpx.AddHeaderFunc(headers))
}
