package mongo_client

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/bezhai/multi-bot-task/biz/model/pixiv_model"
	"github.com/bezhai/multi-bot-task/utils/env_utils"
)

var ImgCollection *MongoCollection[pixiv_model.PixivImageInfo]
var TranslateMap *MongoCollection[pixiv_model.TranslateWord]
var DownloadTaskMap *MongoCollection[pixiv_model.DownloadTask]

var db *mongo.Client

func InitMongoDb() {
	var err error
	url := fmt.Sprintf(`mongodb://%s:%s@%s/chiwei?connectTimeoutMS=2000&authSource=admin`,
		env_utils.Value("MONGO_INITDB_ROOT_PASSWORD"),
		env_utils.Value("MONGO_INITDB_ROOT_PASSWORD"),
		env_utils.Value("MONGO_SERVER_IP"),
	)
	db, err = mongo.Connect(context.Background(),
		options.Client().ApplyURI(url))
	if err != nil {
		panic(err)
	}
	ImgCollection = GenCollection[pixiv_model.PixivImageInfo]("img_map")
	TranslateMap = GenCollection[pixiv_model.TranslateWord]("trans_map")
	DownloadTaskMap = GenCollection[pixiv_model.DownloadTask]("download_task")
}

func GenCollection[T any](name string) *MongoCollection[T] {
	return NewCollection[T](db.Database("chiwei").Collection(name))
}
