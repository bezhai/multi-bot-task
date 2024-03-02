package image_trans

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/bezhai/multi-bot-task/biz/clients/mongo_client"
	"github.com/bezhai/multi-bot-task/biz/clients/oss_client"
	"github.com/bezhai/multi-bot-task/biz/model/data_trans"
	"github.com/bezhai/multi-bot-task/biz/model/image_store"
	"github.com/bezhai/multi-bot-task/biz/service/proxy"
	"github.com/bezhai/multi-bot-task/biz/utils/langx/slicex"
)

func DownloadPixivImages(ctx context.Context, imageUrl string) error {

	fileName := slicex.Get(strings.Split(imageUrl, "/"), -1).OrElse("")
	illustId := strings.Split(fileName, "_")[0]

	reader, _, err := proxy.Proxy(ctx, &data_trans.ProxyRequest{
		URL:     imageUrl,
		Referer: fmt.Sprintf("https://www.pixiv.net/artworks/%s", illustId),
	})

	tosFileName := fmt.Sprintf("pixiv_img_v2/%s/%s", time.Now().Format("20060102"), fileName)

	err = oss_client.PutObject(ctx, tosFileName, bytes.NewBuffer(reader))
	if err != nil {
		return err
	}

	err = mongo_client.ImgCollection.InsertOne(ctx, &image_store.PixivImageMetaInfo{
		PixivAddr:   fileName,
		TosFileName: tosFileName,
	})
	if err != nil {
		return err
	}

	return nil
}
