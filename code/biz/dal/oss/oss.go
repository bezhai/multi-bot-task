package oss

import (
	"context"
	"fmt"
	"github.com/spf13/cast"
	"io"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/hertz/pkg/common/hlog"

	"github.com/bezhai/multi-bot-task/biz/utils/langx"
	"github.com/bezhai/multi-bot-task/biz/utils/langx/slicex"
	"github.com/bezhai/multi-bot-task/utils/env_utils"
)

var internalBucket *oss.Bucket
var cnameBucket *oss.Bucket

func Init() {
	internalBucket = newOssClient(false)
	cnameBucket = newOssClient(true)
}

func newOssClient(useCname bool) *oss.Bucket {
	var err error
	provider, err := oss.NewEnvironmentVariableCredentialsProvider()
	if err != nil {
		panic(err)
	}

	ossClient, err := oss.New(env_utils.Value(langx.IfElse(useCname, "END_POINT", "INTERNAL_END_POINT")),
		"",
		"",
		oss.SetCredentialsProvider(&provider),
		oss.UseCname(useCname))
	if err != nil {
		panic(err)
	}

	bucket, err := ossClient.Bucket(env_utils.Value("OSS_BUCKET"))
	if err != nil {
		panic(err)
	}

	return bucket
}

func PutObject(ctx context.Context, fileName string, reader io.Reader) error {
	err := internalBucket.PutObject(fileName, reader)
	if err != nil {
		logger.CtxWarnf(ctx, "put object %s fail, err:%v", fileName, err)
	} else {
		logger.CtxInfof(ctx, "put object %s success", fileName)
	}
	return err
}

func GetObject(ctx context.Context, fileName string) (io.ReadCloser, error) {
	reader, err := internalBucket.GetObject(fileName)
	if err != nil {
		logger.CtxWarnf(ctx, "put object %s fail, err:%v", fileName, err)
	} else {
		logger.CtxInfof(ctx, "put object %s success", fileName)
	}
	return reader, err
}

func GetObjectDetailMeta(fileName string, key string) (string, error) {
	meta, err := internalBucket.GetObjectDetailedMeta(fileName)
	if err != nil {
		return "", err
	}
	return meta.Get(key), nil
}

func GenUrl(fileName string, isDownload bool) (string, error) {
	var options []oss.Option
	pixivAddrName := slicex.Get(strings.Split(fileName, "/"), -1).OrElse("")
	contentLength, err := GetObjectDetailMeta(fileName, ContentLength)
	if isDownload {
		options = append(options, oss.ResponseContentDisposition(fmt.Sprintf(`attachment; filename="%s"`, pixivAddrName)))
	} else {
		if cast.ToInt64(contentLength) < 20*1024*1024 {
			options = append(options, oss.Process("style/sort_image"))
		}
	}
	url, err := cnameBucket.SignURL(fileName, oss.HTTPGet, 7200,
		options...)
	if err != nil {
		hlog.CtxWarnf(context.Background(), "sign url %s fail, err:%v", fileName, err)
	} else {
		hlog.CtxInfof(context.Background(), "sign url %s success", fileName)
	}
	return url, err
}
