package main

import (
	"context"

	"gorm.io/gen"

	"github.com/bezhai/multi-bot-task/biz/dal/mysql"
	"github.com/bezhai/multi-bot-task/config"
)

func main() {
	// init db
	config.InitEnv()
	mysql.Init()

	// GrowthPromotion
	func() {
		g := gen.NewGenerator(gen.Config{
			OutPath:      "biz/mysql/dao",
			ModelPkgPath: "biz/mysql/model",

			FieldNullable:     true, // 对于数据库中nullable的数据，在生成代码中自动对应为指针类型
			FieldWithIndexTag: true, // 从数据库同步的表结构代码包含gorm的index tag
			FieldWithTypeTag:  true, // 同步的表结构代码包含gorm的type tag(数据库中对应数据类型)
		})

		// 复用工程原本使用的SQL连接配置db (*gorm.DB)，也可以根据需求在此处之间建立连接
		g.UseDB(mysql.GetDB(context.Background()))

		basicModels := []string{
			"roles",
			"permissions",
			"module_access",
			"users",
			"user_permissions",
			"user_module_access",
		}

		for _, model := range basicModels {
			g.ApplyBasic(g.GenerateModel(model))
		}

		// 执行并生成代码
		g.Execute()
	}()

}
