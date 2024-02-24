package mongo_client

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/bezhai/multi-bot-task/biz/model/image_store"
	"github.com/bezhai/multi-bot-task/biz/model/translation"
	"github.com/bezhai/multi-bot-task/utils/env_utils"
)

var ImgCollection *MongoCollection[image_store.PixivImageMetaInfo]
var TranslateMap *MongoCollection[translation.TranslateWord]

var db *mongo.Client

func InitMongoDb() {
	var err error
	url := fmt.Sprintf(`mongodb://%s:%s@%s/chiwei?connectTimeoutMS=2000&authSource=admin`,
		env_utils.Value("MONGO_INITDB_ROOT_USERNAME"),
		env_utils.Value("MONGO_INITDB_ROOT_PASSWORD"),
		env_utils.Value("MONGO_SERVER_IP"),
	)
	db, err = mongo.Connect(context.Background(),
		options.Client().ApplyURI(url))
	if err != nil {
		panic(err)
	}
	ImgCollection = GenCollection[image_store.PixivImageMetaInfo]("img_map")
	TranslateMap = GenCollection[translation.TranslateWord]("trans_map")
}

func GenCollection[T any](name string) *MongoCollection[T] {
	return NewCollection[T](db.Database("chiwei").Collection(name))
}
