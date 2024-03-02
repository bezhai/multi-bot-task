struct MultiTag {
    1: required string name(go.tag = "bson:\"name\"");
    2: optional string translation (go.tag = "bson:\"translation\"");
    3: optional bool visible (go.tag = "bson:\"visible\"");
}

struct PixivImageMetaInfo {
    1: required string pixiv_addr (go.tag = "bson:\"pixiv_addr,omitempty\"");
    2: required bool visible (go.tag = "bson:\"visible,omitempty\"");
    3: required string author (go.tag = "bson:\"author,omitempty\"");
    4: optional list<MultiTag> multi_tags (go.tag = "bson:\"multi_tags,omitempty\"");
    5: required i64 create_time (go.tag = "bson:\"create_time,omitempty\"");
    6: required i64 update_time (go.tag = "bson:\"update_time,omitempty\"");
    7: required string tos_file_name (go.tag = "bson:\"tos_file_name,omitempty\"");
    8: optional string author_id (go.tag = "bson:\"author_id,omitempty\"");
    9: required bool del_flag (go.tag = "bson:\"del_flag,omitempty\"");
    10: required i32 illust_id (go.tag = "bson:\"illust_id,omitempty\"");
    11: required string title (go.tag = "bson:\"title,omitempty\"");
    12: optional string image_key (go.tag = "bson:\"image_key,omitempty\"");
    13: optional i32 width (go.tag = "bson:\"width,omitempty\"");
    14: optional i32 height (go.tag = "bson:\"height,omitempty\"");
}

// 0: 未删除 1: 可见 2: 已删除 3: 全部 4: 不可见
enum StatusMode {
    StatusNotDelete = 0,
    StatusVisible = 1,
    StatusDelete = 2,
    StatusAll = 3,
    StatusNoVisible = 4,
}

struct ListPixivImageMetaInfoRequest {
    1: optional string author; // 模糊查询
    2: optional string author_id; // 精确查询
    3: optional list<string> tags; // 按每个tag模糊查询的交集
    4: optional StatusMode status = StatusMode.StatusVisible;
    5: optional i32 page = 1;
    6: optional i32 page_size = 10;
    7: optional i32 illust_id;
}

struct PixivImageMetaInfoWithUrl {
    1: required PixivImageMetaInfo pixiv_image_meta_info;
    2: optional string show_url;
    3: optional string download_url;
}

struct ListPixivImageMetaInfoResponseData {
    1: required list<PixivImageMetaInfoWithUrl> pixiv_image_meta_infos;
    2: required i32 total;
}

struct ListPixivImageMetaInfoResponse {
    1: required i32 code;
    2: required string msg;
    3: optional ListPixivImageMetaInfoResponseData data;
}

struct UpdatePixivImageStatusRequest {
    1: required list<string> pixiv_addr_list(api.vd="len($) > 0; msg:'pixiv_addr_list is empty'");
    2: required StatusMode status(api.vd="in($, 1, 2, 4); msg:'valid status is visible, no_visible or deleted'"); // 1: 可见 2: 删除 4: 不可见
}

struct UpdatePixivImageStatusResponse {
    1: required i32 code;
    2: required string msg;
}