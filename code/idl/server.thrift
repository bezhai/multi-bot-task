namespace go chiwei_bot

include "conf.thrift"
include "data_trans.thrift"
include "image_store.thrift"
include "translation.thrift"
include "auth.thrift"

service ChiweiBotService {
    auth.RegisterResponse Register(1: auth.RegisterRequest request) (api.post="/api/auth/register")

    conf.GetStringValueResponse GetStringValue(1: conf.GetStringValueRequest request) (api.get="/api/need-auth/conf/get-string-value")
    conf.SetStringValueResponse SetStringValue(1: conf.SetStringValueRequest request) (api.post="/api/need-auth/conf/set-string-value")
    conf.GetMemberValueResponse GetMemberValue(1: conf.GetMemberValueRequest request) (api.get="/api/need-auth/conf/get-member-value")
    conf.SetMemberValueResponse SetMemberValue(1: conf.SetMemberValueRequest request) (api.post="/api/need-auth/conf/set-member-value")

    image_store.ListPixivImageMetaInfoResponse ListPixivImageMetaInfo(1: image_store.ListPixivImageMetaInfoRequest request) (api.post="/api/need-auth/image-store/list-info")
    image_store.UpdatePixivImageStatusResponse UpdatePixivImageStatus(1: image_store.UpdatePixivImageStatusRequest request) (api.post="/api/need-auth/image-store/update-status")

    translation.DeleteTranslationResponse DeleteTranslation(1: translation.DeleteTranslationRequest request) (api.post="/api/need-auth/translation/delete")
    translation.ListTranslationResponse ListTranslation(1: translation.ListTranslationRequest request) (api.get="/api/need-auth/translation/list")
    translation.UpdateTranslationResponse UpdateTranslation(1: translation.UpdateTranslationRequest request) (api.post="/api/need-auth/translation/update")

    data_trans.ProxyResponse Proxy(1: data_trans.ProxyRequest request) (api.post="/api/need-sk/data-trans/proxy") // 代理请求
    data_trans.DownloadPixivImageResponse DownloadPixivImage(1: data_trans.DownloadPixivImageRequest request) (api.post="/api/need-sk/data-trans/download-pixiv-image") // 下载pixiv图片
    data_trans.UploadTosFileToLarkResponse UploadTosFileToLark(1: data_trans.UploadTosFileToLarkRequest request) (api.post="/api/need-sk/data-trans/upload-to-lark") // pixivId换imageKey
}