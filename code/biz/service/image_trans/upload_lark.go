package image_trans

import (
	"context"
	"errors"
	"io"

	"github.com/bytedance/gopkg/util/logger"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
	"github.com/spf13/cast"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/bezhai/multi-bot-task/biz/dal/lark_client"
	"github.com/bezhai/multi-bot-task/biz/dal/mongodb"
	"github.com/bezhai/multi-bot-task/biz/dal/oss"
	"github.com/bezhai/multi-bot-task/biz/utils/imagex"
)

func UploadImageToLark(ctx context.Context, pixivAddr string) error {
	imageInfo, err := mongodb.ImgCollection.FindOne(ctx, bson.M{
		"pixiv_addr": pixivAddr,
	})
	if err != nil {
		return err
	}

	reader, err := oss.GetObject(ctx, imageInfo.TosFileName)
	if err != nil {
		return err
	}
	imageSize, err := oss.GetObjectDetailMeta(ctx, imageInfo.TosFileName, oss.ContentLength)

	newReader, width, height, err := imagex.ResizeImage(reader, cast.ToInt(imageSize))
	if err != nil {
		logger.CtxWarnf(ctx, "resize fail err:%v", err)
		return err
	}

	imageKey, err := uploadImage(ctx, newReader)
	if err != nil {
		return err
	}

	err = mongodb.ImgCollection.UpdateOne(ctx, bson.M{
		"pixiv_addr": pixivAddr,
	}, bson.M{
		"image_key": imageKey,
		"width":     width,
		"height":    height,
	})

	return nil
}

func uploadImage(ctx context.Context, file io.Reader) (string, error) {

	uploadReq := larkim.NewCreateImageReqBuilder().Body(
		larkim.NewCreateImageReqBodyBuilder().
			ImageType("message").
			Image(file).
			Build(),
	).Build()

	uploadResp, uploadErr := lark_client.OfficialBot.Im.Image.Create(ctx, uploadReq)

	if uploadErr != nil {
		return "", uploadErr
	}

	if !uploadResp.Success() {
		return "", errors.New(uploadResp.Msg)
	}

	return *uploadResp.Data.ImageKey, nil

}
