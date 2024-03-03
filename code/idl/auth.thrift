struct RegisterRequest {
    1: required string username (api.vd="(len($) > 0 && len($) < 128); msg:'Illegal format'");
    2: required string password (api.vd="(len($) > 0 && len($) < 128); msg:'Illegal format'");
}

struct RegisterResponse {
    1: required i32 code;
    2: required string message;
}