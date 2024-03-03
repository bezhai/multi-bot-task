package dal

import (
	"github.com/bezhai/multi-bot-task/biz/dal/http_client"
	"github.com/bezhai/multi-bot-task/biz/dal/lark_client"
	"github.com/bezhai/multi-bot-task/biz/dal/mongodb"
	"github.com/bezhai/multi-bot-task/biz/dal/mysql"
	"github.com/bezhai/multi-bot-task/biz/dal/oss"
	"github.com/bezhai/multi-bot-task/biz/dal/redis"
)

func Init() {
	mysql.Init()
	http_client.Init()
	lark_client.Init()
	mongodb.Init()
	oss.Init()
	redis.Init()
}
