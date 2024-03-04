package translation_store

import (
	"context"
	"github.com/bezhai/multi-bot-task/biz/model/translation"
	"github.com/bezhai/multi-bot-task/biz/utils/langx/ptr"
	"go.mongodb.org/mongo-driver/bson"
)

func ListTranslation(ctx context.Context, req *translation.ListTranslationRequest) (
	*translation.ListTranslationResponseData, error) {

	if req.PageSize <= 0 {
		req.PageSize = 10
	}
	if req.Page <= 0 {
		req.Page = 1
	}

	filter := bson.D{}

	if ptr.Value(req.SearchKey) != "" {
		filter = append(filter, bson.E{
			Key: "$or",
			Value: bson.D{
				{
					Key: "origin",
					Value: bson.M{
						"$regex":   *req.SearchKey,
						"$options": "i",
					},
				},
				{
					Key: "extra_info",
					Value: bson.E{
						Key: "$elemMatch",
						Value: bson.E{
							Key: "zh",
							Value: bson.M{
								"$regex":   *req.SearchKey,
								"$options": "i",
							},
						},
					},
				},
				{
					Key: "extra_info",
					Value: bson.E{
						Key: "$elemMatch",
						Value: bson.E{
							Key: "en",
							Value: bson.M{
								"$regex":   *req.SearchKey,
								"$options": "i",
							},
						},
					},
				},
			},
		})
	}

	if req.OnlyUntranslated {
		filter = append(filter, bson.E{
			Key:   "has_translate",
			Value: "false",
		})
	}

	return nil, nil

	//count, err := mongodb.TranslateMap.CountDocuments(ctx, filter)
	//
	//if err != nil {
	//	return nil, err
	//}
	//
	//originRes, err := mongodb.TranslateMap.Find(ctx, filter,
	//	options.Find().SetSkip(int64((req.Page-1)*req.PageSize)).SetLimit(int64(req.PageSize)))
	//
	//if err != nil {
	//	return nil, err
	//}
	//
	//wg := new(errgroup.Group)

}
