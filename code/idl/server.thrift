namespace go chiwei_bot

include "conf.thrift"

service ChiweiBotService {
    conf.GetStringValueResponse GetStringValue(1: conf.GetStringValueRequest request) (api.get="/api/conf/string_value")
}