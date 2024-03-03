package image_trans

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/bezhai/multi-bot-task/biz/dal/mongodb"
	"github.com/bezhai/multi-bot-task/biz/dal/oss"
	"github.com/bezhai/multi-bot-task/biz/model/data_trans"
	"github.com/bezhai/multi-bot-task/biz/model/image_store"
	"github.com/bezhai/multi-bot-task/biz/service/proxy"
	"github.com/bezhai/multi-bot-task/biz/utils/langx/slicex"
)

func DownloadPixivImages(ctx context.Context, imageUrl string) error {
	fileName := slicex.Get(strings.Split(imageUrl, "/"), -1).OrElse("")
	illustId := strings.Split(fileName, "_")[0]

	count, err := mongodb.ImgCollection.CountDocuments(ctx,
		bson.M{
			"pixiv_addr": fileName,
			"tos_file_name": bson.M{
				"$ne": "",
			},
		})
	if err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	reader, _, err := proxy.Proxy(ctx, &data_trans.ProxyRequest{
		URL:     imageUrl,
		Referer: fmt.Sprintf("https://www.pixiv.net/artworks/%s", illustId),
	})

	tosFileName := fmt.Sprintf("pixiv_img_v2/%s/%s", time.Now().Format("20060102"), fileName)

	err = oss.PutObject(ctx, tosFileName, bytes.NewBuffer(reader))
	if err != nil {
		return err
	}

	err = mongodb.ImgCollection.InsertOne(ctx, &image_store.PixivImageMetaInfo{
		PixivAddr:   fileName,
		TosFileName: tosFileName,
	})
	if err != nil {
		return err
	}

	return nil
}
