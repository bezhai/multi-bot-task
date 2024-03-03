package lark_client

import (
	lark "github.com/larksuite/oapi-sdk-go/v3"

	"github.com/bezhai/multi-bot-task/utils/env_utils"
)

var OfficialBot *lark.Client

func Init() {
	OfficialBot = lark.NewClient(
		env_utils.Value("LARK_APP_ID"),
		env_utils.Value("LARK_APP_SECRET"),
		lark.WithEnableTokenCache(true))
}
