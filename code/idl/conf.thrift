struct GetStringValueRequest {
    1: required string key;
}

struct GetStringValueResponseData {
    1: required string value;
}

struct GetStringValueResponse {
    1: required i32 code;
    2: required string msg;
    3: optional GetStringValueResponseData data;
}
