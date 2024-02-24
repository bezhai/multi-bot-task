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

struct SetStringValueRequest {
    1: required string key;
    2: required string value;
}

struct SetStringValueResponse {
    1: required i32 code;
    2: required string msg;
}

struct GetMemberValueRequest {
    1: required string key;
}

struct GetMemberValueResponseData {
    1: required list<string> value;
}

struct GetMemberValueResponse {
    1: required i32 code;
    2: required string msg;
    3: optional GetMemberValueResponseData data;
}

struct SetMemberValueRequest {
    1: required string key;
    2: required list<string> value;
}

struct SetMemberValueResponse {
    1: required i32 code;
    2: required string msg;
}
