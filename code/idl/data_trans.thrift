struct UploadTosFileToLarkRequest {
    1: required string pixiv_addr;
}

struct UploadTosFileToLarkResponseData {
    1: required string image_key;
    2: required i32 width;
    3: required i32 height;
}

struct UploadTosFileToLarkResponse {
    1: required i32 code;
    2: required string msg;
    3: optional UploadTosFileToLarkResponseData data;
}

struct ProxyRequest {
    1: required string url;
    2: required string referer;
}

struct ProxyResponse {
}