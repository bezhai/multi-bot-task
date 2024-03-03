package image_db

import (
	"context"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/sync/errgroup"

	"github.com/bezhai/multi-bot-task/biz/dal/mongodb"
	"github.com/bezhai/multi-bot-task/biz/dal/oss"
	"github.com/bezhai/multi-bot-task/biz/model/image_store"
	"github.com/bezhai/multi-bot-task/biz/utils/langx/ptr"
)

func ListImages(ctx context.Context, req *image_store.ListPixivImageMetaInfoRequest) ([]*image_store.PixivImageMetaInfoWithUrl, int, error) {
	images, total, err := listImages(ctx, req)
	if err != nil {
		return nil, 0, err
	}

	wg := new(errgroup.Group)
	limitChan := make(chan struct{}, 10)
	result := make([]*image_store.PixivImageMetaInfoWithUrl, len(images))

	for i := range images {
		index := i
		wg.Go(func() error {
			limitChan <- struct{}{}
			defer func() {
				<-limitChan
			}()

			result[index] = &image_store.PixivImageMetaInfoWithUrl{
				PixivImageMetaInfo: images[index],
			}

			if images[index].TosFileName != "" {
				downloadUrl, downloadErr := oss.GenUrl(images[index].TosFileName, true)
				if downloadErr != nil {
					return downloadErr
				}
				result[index].DownloadURL = ptr.Ptr(downloadUrl)
				url, downloadErr := oss.GenUrl(images[index].TosFileName, false)
				if downloadErr != nil {
					return downloadErr
				}
				result[index].ShowURL = ptr.Ptr(url)
			}
			return nil
		})
	}

	err = wg.Wait()
	if err != nil {
		return nil, 0, err
	}

	return result, total, nil
}

func listImages(ctx context.Context, req *image_store.ListPixivImageMetaInfoRequest) ([]*image_store.PixivImageMetaInfo, int, error) {
	var filters []bson.M

	if ptr.Value(req.Author) != "" {
		filters = append(filters, bson.M{
			"author": bson.M{
				"$regex":   *req.Author,
				"$options": "i",
			},
		})
	}

	if ptr.Value(req.IllustID) != 0 {
		filters = append(filters, bson.M{
			"illust_id": *req.IllustID,
		})
	}

	for _, tag := range req.Tags {
		m := bson.M{
			"translation": bson.M{
				"$regex":   tag,
				"$options": "i",
			},
		}
		filters = append(filters,
			bson.M{
				"multi_tags": bson.M{
					"$elemMatch": m,
				},
			})
	}

	if ptr.Value(req.AuthorID) != "" {
		filters = append(filters, bson.M{
			"author_id": *req.AuthorID,
		})
	}

	switch req.Status {
	case image_store.StatusMode_StatusVisible:
		filters = append(filters, bson.M{
			"visible":  true,
			"del_flag": false,
		})
	case image_store.StatusMode_StatusNoVisible:
		filters = append(filters, bson.M{
			"visible":  false,
			"del_flag": false,
		})
	case image_store.StatusMode_StatusDelete:
		filters = append(filters, bson.M{
			"del_flag": true,
		})
	case image_store.StatusMode_StatusAll:
	case image_store.StatusMode_StatusNotDelete:
		fallthrough
	default:
		filters = append(filters, bson.M{
			"del_flag": false,
		})
	}

	count, err := mongodb.ImgCollection.CountDocuments(ctx, bson.M{
		"$and": filters,
	})
	if err != nil {
		return nil, 0, err
	}

	if req.Page <= 0 {
		req.Page = 1
	}

	if req.PageSize <= 0 {
		req.PageSize = 20
	}

	opts := options.
		Find().
		SetSkip(int64((req.Page - 1) * req.PageSize)).SetLimit(int64(req.PageSize)).
		SetSort(bson.M{"create_time": -1})

	res, err := mongodb.ImgCollection.Find(
		ctx,
		bson.M{
			"$and": filters,
		},
		opts,
	)

	if err != nil {
		hlog.CtxWarnf(ctx, "get images fail, err:%v", err)
		return nil, 0, err
	}

	return res, int(count), nil
}

func UpdateImageStatus(ctx context.Context, req *image_store.UpdatePixivImageStatusRequest) error {

	if len(req.PixivAddrList) <= 0 {
		hlog.CtxWarnf(ctx, "empty pixiv addr list")
		return nil
	}

	updateParam := bson.M{}
	switch req.Status {
	case image_store.StatusMode_StatusNoVisible:
		updateParam = bson.M{
			"visible":  false,
			"del_flag": false,
		}
	case image_store.StatusMode_StatusVisible:
		updateParam = bson.M{
			"visible":  true,
			"del_flag": false,
		}
	case image_store.StatusMode_StatusDelete:
		updateParam = bson.M{
			"del_flag": true,
		}
	default:
		hlog.CtxWarnf(ctx, "unknown status: %d", req.Status)
		return nil
	}

	err := mongodb.ImgCollection.UpdateMany(ctx, bson.M{
		"pixiv_addr": bson.M{
			"$in": req.PixivAddrList,
		},
	}, updateParam)

	if err != nil {
		hlog.CtxWarnf(ctx, "update status fail, err:%v", err)
	}
	return err
}
