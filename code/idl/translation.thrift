struct TranslateWordExtra {
    1: optional string zh (go.tag = "bson:\"zh\"");
    2: optional string en (go.tag = "bson:\"en\"");
}

struct TranslateWord {
    1: required string origin (go.tag = "bson:\"origin\"");
    2: required string translation (go.tag = "bson:\"translation\"");
    3: required bool has_translate (go.tag = "bson:\"has_translate\"");
    4: required TranslateWordExtra extra_info (go.tag = "bson:\"extra_info\"");
}

struct TranslateWithNum {
    1: required TranslateWord word (go.tag = "bson:\"word\"");
    2: required i64 num (go.tag = "bson:\"num\"");
}

struct UpdateTranslationRequest {
    1: required string origin (go.tag = "bson:\"origin\"");
    2: required string translation (go.tag = "bson:\"translation\"");
}

struct UpdateTranslationResponse {
    1: required i32 code;
    2: required string msg;
}

struct ListTranslationRequest {
    1: optional string search_key;
    2: optional i32 page = 1;
    3: optional i32 page_size = 10;
    4: optional bool only_untranslated = true;
}

struct ListTranslationResponseData {
    1: optional list<TranslateWithNum> data;
    2: optional i64 total;
}

struct ListTranslationResponse {
    1: required i32 code;
    2: required string msg;
    3: optional ListTranslationResponseData data;
}

struct DeleteTranslationRequest {
    1: required string origin;
}

struct DeleteTranslationResponse {
    1: required i32 code;
    2: required string msg;
}
