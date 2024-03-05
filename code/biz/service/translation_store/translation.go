package translation_store

import (
	"context"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/sync/errgroup"

	"github.com/bezhai/multi-bot-task/biz/dal/mongodb"
	"github.com/bezhai/multi-bot-task/biz/model/translation"
	"github.com/bezhai/multi-bot-task/biz/utils/langx/ptr"
)

func UpdateTranslate(ctx context.Context, word string, translation string) error {
	err := mongodb.TranslateMap.UpdateMany(ctx, bson.D{
		{
			"origin", word,
		},
	}, bson.M{"translation": translation, "has_translate": true}, options.Update().SetUpsert(true))
	if err != nil {
		return err
	}

	err = mongodb.ImgCollection.UpdateMany(ctx, bson.D{
		{
			"multi_tags.name", word,
		},
	}, bson.D{
		{
			"multi_tags.$.translation", translation,
		},
	})
	if err != nil {
		return err
	}

	return nil
}

func DeleteTranslate(ctx context.Context, word string) error {
	hlog.CtxInfof(ctx, "delete translate: %s", word)
	return mongodb.TranslateMap.DeleteMany(ctx, bson.M{
		"origin": word,
	})
}

func ListTranslation(ctx context.Context, req *translation.ListTranslationRequest) (
	*translation.ListTranslationResponseData, error) {

	if req.PageSize <= 0 {
		req.PageSize = 10
	}
	if req.Page <= 0 {
		req.Page = 1
	}

	var filter []bson.M

	if ptr.Value(req.SearchKey) != "" {
		filter = append(filter, bson.M{
			"$or": []bson.M{
				{
					"origin": bson.M{
						"$regex":   *req.SearchKey,
						"$options": "i",
					},
				},
				{
					"extra_info.zh": bson.M{
						"$regex":   *req.SearchKey,
						"$options": "i",
					},
				},
				{
					"extra_info.en": bson.M{
						"$regex":   *req.SearchKey,
						"$options": "i",
					},
				},
			},
		})
	}

	if req.OnlyUntranslated {
		filter = append(filter, bson.M{
			"has_translate": false,
		})
	}

	count, err := mongodb.TranslateMap.CountDocuments(ctx, bson.M{
		"$and": filter,
	})

	if err != nil {
		return nil, err
	}

	originRes, err := mongodb.TranslateMap.Find(ctx, bson.M{
		"$and": filter,
	}, options.Find().SetSkip(int64((req.Page-1)*req.PageSize)).SetLimit(int64(req.PageSize)))

	if err != nil {
		return nil, err
	}

	res := make([]*translation.TranslateWithNum, len(originRes))

	var g errgroup.Group
	ch := make(chan struct{}, 10)
	for i := range originRes {
		ch <- struct{}{}
		index := i
		g.Go(func() error {
			defer func() {
				<-ch
			}()
			useCount, countErr := mongodb.ImgCollection.CountDocuments(ctx, bson.D{
				{
					"multi_tags.name", originRes[index].Origin,
				},
			})
			if countErr != nil {
				return countErr
			}
			res[index] = &translation.TranslateWithNum{
				Word: originRes[index],
				Num:  useCount,
			}
			return nil
		})
	}

	if err = g.Wait(); err != nil {
		return nil, err
	}

	return &translation.ListTranslationResponseData{
		Total: ptr.Ptr(count),
		Data:  res,
	}, nil
}
