package respx

import (
	"github.com/cloudwego/hertz/pkg/app"
)

func SuccessWith(c *app.RequestContext, data interface{}) {
	resp := map[string]interface{}{
		"code": 0,
		"msg":  "success",
		"data": data,
	}

	c.JSON(200, resp)
}

func FailWith(c *app.RequestContext, code int, msg string, data interface{}) {
	resp := map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": data,
	}

	c.JSON(200, resp)
}

func Success(c *app.RequestContext) {
	resp := map[string]interface{}{
		"code": 0,
		"msg":  "success",
	}

	c.JSON(200, resp)
}

func Fail(c *app.RequestContext, code int, msg string) {
	resp := map[string]interface{}{
		"code": code,
		"msg":  msg,
	}

	c.JSON(200, resp)
}
