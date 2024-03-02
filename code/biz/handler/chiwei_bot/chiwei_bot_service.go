// Code generated by hertz generator.

package chiwei_bot

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	"github.com/bezhai/multi-bot-task/biz/model/conf"
	"github.com/bezhai/multi-bot-task/biz/model/data_trans"
	"github.com/bezhai/multi-bot-task/biz/model/image_store"
	"github.com/bezhai/multi-bot-task/biz/model/translation"
	"github.com/bezhai/multi-bot-task/biz/service/conf_value"
	"github.com/bezhai/multi-bot-task/biz/service/image_db"
	"github.com/bezhai/multi-bot-task/biz/service/image_trans"
	"github.com/bezhai/multi-bot-task/biz/service/proxy"
	"github.com/bezhai/multi-bot-task/biz/utils/respx"
)

// GetStringValue .
// @router /api/conf/string_value [GET]
func GetStringValue(ctx context.Context, c *app.RequestContext) {
	var err error
	var req conf.GetStringValueRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		respx.FailWithError(c, consts.StatusBadRequest, err)
		return
	}

	value, err := conf_value.GetConfStringValue(ctx, req.Key)
	if err != nil {
		respx.Fail(c, 500, err.Error())
		return
	}

	respx.SuccessWith(c, &conf.GetStringValueResponseData{Value: value})
}

// UploadTosFileToLark .
// @router /api/need-sk/data-trans/upload-tos-file-to-lark [POST]
func UploadTosFileToLark(ctx context.Context, c *app.RequestContext) {
	var err error
	var req data_trans.UploadTosFileToLarkRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		respx.FailWithError(c, consts.StatusBadRequest, err)
		return
	}

	err = image_trans.UploadImageToLark(ctx, req.PixivAddr)
	if err != nil {
		respx.Fail(c, 500, err.Error())
		return
	}

	respx.Success(c)
}

// ListPixivImageMetaInfo .
// @router /api/need-auth/image-store/pixiv-image-meta-info [GET]
func ListPixivImageMetaInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req image_store.ListPixivImageMetaInfoRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		respx.FailWithError(c, consts.StatusBadRequest, err)
		return
	}

	infos, total, err := image_db.ListImages(ctx, &req)
	if err != nil {
		respx.FailWithError(c, consts.StatusInternalServerError, err)
		return
	}

	respx.SuccessWith(c, &image_store.ListPixivImageMetaInfoResponseData{
		PixivImageMetaInfos: infos,
		Total:               int32(total),
	})

}

// UpdatePixivImageStatus .
// @router /api/need-auth/image-store/update-pixiv-image-status [POST]
func UpdatePixivImageStatus(ctx context.Context, c *app.RequestContext) {
	var err error
	var req image_store.UpdatePixivImageStatusRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		respx.FailWithError(c, consts.StatusBadRequest, err)
		return
	}

	err = image_db.UpdateImageStatus(ctx, &req)
	if err != nil {
		respx.FailWithError(c, consts.StatusInternalServerError, err)
		return
	}

	respx.Success(c)
}

// DeleteTranslation .
// @router /api/need-auth/translation/delete [POST]
func DeleteTranslation(ctx context.Context, c *app.RequestContext) {
	var err error
	var req translation.DeleteTranslationRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		respx.FailWithError(c, consts.StatusBadRequest, err)
		return
	}

	resp := new(translation.DeleteTranslationResponse)

	c.JSON(consts.StatusOK, resp)
}

// ListTranslation .
// @router /api/need-auth/translation/list [GET]
func ListTranslation(ctx context.Context, c *app.RequestContext) {
	var err error
	var req translation.ListTranslationRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		respx.FailWithError(c, consts.StatusBadRequest, err)
		return
	}

	resp := new(translation.ListTranslationResponse)

	c.JSON(consts.StatusOK, resp)
}

// UpdateTranslation .
// @router /api/need-auth/translation/update [POST]
func UpdateTranslation(ctx context.Context, c *app.RequestContext) {
	var err error
	var req translation.UpdateTranslationRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		respx.FailWithError(c, consts.StatusBadRequest, err)
		return
	}

	resp := new(translation.UpdateTranslationResponse)

	c.JSON(consts.StatusOK, resp)
}

// SetStringValue .
// @router /api/need-auth/conf/set-string-value [POST]
func SetStringValue(ctx context.Context, c *app.RequestContext) {
	var err error
	var req conf.SetStringValueRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		respx.FailWithError(c, consts.StatusBadRequest, err)
		return
	}

	_, err = conf_value.SetConfStringValue(ctx, req.Key, req.Value)
	if err != nil {
		respx.Fail(c, 500, err.Error())
		return
	}

	respx.Success(c)
}

// GetMemberValue .
// @router /api/need-auth/conf/get-member-value [GET]
func GetMemberValue(ctx context.Context, c *app.RequestContext) {
	var err error
	var req conf.GetMemberValueRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		respx.FailWithError(c, consts.StatusBadRequest, err)
		return
	}

	value, err := conf_value.GetConfMemberValue(ctx, req.Key)
	if err != nil {
		respx.Fail(c, 500, err.Error())
		return
	}

	respx.SuccessWith(c, &conf.GetMemberValueResponseData{Value: value})
}

// SetMemberValue .
// @router /api/need-auth/conf/set-member-value [POST]
func SetMemberValue(ctx context.Context, c *app.RequestContext) {
	var err error
	var req conf.SetMemberValueRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		respx.FailWithError(c, consts.StatusBadRequest, err)
		return
	}

	err = conf_value.SetConfMemberValue(ctx, req.Key, req.Value)
	if err != nil {
		respx.Fail(c, 500, err.Error())
		return
	}

	respx.Success(c)
}

// Proxy .
// @router /api/need-auth/data-trans/proxy [GET]
func Proxy(ctx context.Context, c *app.RequestContext) {
	var err error
	var req data_trans.ProxyRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		respx.FailWithError(c, consts.StatusBadRequest, err)
		return
	}

	respBody, respHeader, err := proxy.Proxy(ctx, &req)
	if err != nil {
		respx.FailWithError(c, consts.StatusInternalServerError, err)
		return
	}

	for _, header := range respHeader {
		for k, v := range header {
			c.Header(k, v)
		}
	}

	c.Data(200, "application/json", respBody)
}

// DownloadPixivImage .
// @router /api/need-sk/data-trans/download-pixiv-image [POST]
func DownloadPixivImage(ctx context.Context, c *app.RequestContext) {
	var err error
	var req data_trans.DownloadPixivImageRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		respx.FailWithError(c, consts.StatusBadRequest, err)
		return
	}

	err = image_trans.DownloadPixivImages(ctx, req.PixivURL)
	if err != nil {
		respx.Fail(c, 500, err.Error())
		return
	}

	respx.Success(c)
}
