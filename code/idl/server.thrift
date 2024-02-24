namespace go chiwei_bot

include "conf.thrift"
include "data_trans.thrift"
include "image_store.thrift"
include "translation.thrift"

service ChiweiBotService {
    conf.GetStringValueResponse GetStringValue(1: conf.GetStringValueRequest request) (api.get="/api/need-auth/conf/get-string-value")
    conf.SetStringValueResponse SetStringValue(1: conf.SetStringValueRequest request) (api.post="/api/need-auth/conf/set-string-value")
    conf.GetMemberValueResponse GetMemberValue(1: conf.GetMemberValueRequest request) (api.get="/api/need-auth/conf/get-member-value")
    conf.SetMemberValueResponse SetMemberValue(1: conf.SetMemberValueRequest request) (api.post="/api/need-auth/conf/set-member-value")

    image_store.ListPixivImageMetaInfoResponse ListPixivImageMetaInfo(1: image_store.ListPixivImageMetaInfoRequest request) (api.get="/api/need-auth/image-store/list-info")
    image_store.UpdatePixivImageStatusResponse UpdatePixivImageStatus(1: image_store.UpdatePixivImageStatusRequest request) (api.post="/api/need-auth/image-store/update-status")
    image_store.AddDownloadTaskResponse AddDownloadTask(1: image_store.AddDownloadTaskRequest request) (api.post="/api/need-auth/image-store/add-task")

    translation.DeleteTranslationResponse DeleteTranslation(1: translation.DeleteTranslationRequest request) (api.post="/api/need-auth/translation/delete")
    translation.ListTranslationResponse ListTranslation(1: translation.ListTranslationRequest request) (api.get="/api/need-auth/translation/list")
    translation.UpdateTranslationResponse UpdateTranslation(1: translation.UpdateTranslationRequest request) (api.post="/api/need-auth/translation/update")

    data_trans.UploadTosFileToLarkResponse UploadTosFileToLark(1: data_trans.UploadTosFileToLarkRequest request) (api.post="/api/need-sk/data-trans/upload-to-lark")
}