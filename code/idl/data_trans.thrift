struct UploadTosFileToLarkRequest {
    1: required string pixiv_addr;
}

struct UploadTosFileToLarkResponse {
    1: required i32 code;
    2: required string msg;
}

struct ProxyRequest {
    1: required string url;
    2: required string referer;
    3: optional map<string, string> headers;
}

struct ProxyResponse {
}

struct DownloadPixivImageRequest {
    1: required string pixiv_url;
}

struct DownloadPixivImageResponse {
    1: required i32 code;
    2: required string msg;
}