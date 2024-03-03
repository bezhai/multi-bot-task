package mysql

import (
	"context"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/bezhai/multi-bot-task/biz/mysql/dao"
	"github.com/bezhai/multi-bot-task/utils/env_utils"
)

var db *gorm.DB
var GenDao *dao.Query

func Init() {
	var err error
	dsn := fmt.Sprintf("root:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		env_utils.Value("MYSQL_ROOT_PASSWORD"),
		env_utils.Value("MYSQL_SERVER_IP"),
		env_utils.Value("MYSQL_DBNAME"))
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	GenDao = dao.Use(db)
}

func GetDB(ctx context.Context) *gorm.DB {
	return db.WithContext(ctx)
}
