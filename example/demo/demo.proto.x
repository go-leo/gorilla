syntax = "proto3";
package leo.gors.example.demo.v1;
option go_package = "github.com/go-leo/gorilla/example/demo/v1;demo";

import "google/protobuf/empty.proto";
import "google/api/annotations.proto";

service Demo {

  // CreateUser 创建用户
  // `POST /v1/user { "name": "Leo" }` | `CreateUserRequest(name: "Leo")`
  rpc CreateUser (CreateUserRequest) returns (CreateUserResponse) {
    option (google.api.http) = {
      post : "/v1/user"
      body : "*"
    };
  }

  // DeleteUser 删除用户
  // `DELETE /v1/user/10000 | `DeleteUserRequest(id: 10000)`
  rpc DeleteUser (DeleteUserRequest) returns (google.protobuf.Empty) {
    option (google.api.http) = {
      delete : "/v1/user/{id}"
    };
  }

  // ModifyUser 修改用户
  // `PUT /v1/user/10000 { "name": "Leo" }` | `ModifyUserRequest(id: 10000, name: "Leo")`
  rpc ModifyUser (ModifyUserRequest) returns (google.protobuf.Empty) {
    option (google.api.http) = {
      put : "/v1/user/{id}"
      body : "*"
    };
  }

  // GetUser 获取用户
  // `GET /v1/user/10000` | `GetUserRequest(id: 10000)`
  rpc GetUser (GetUserRequest) returns (GetUserResponse) {
    option (google.api.http) = {
      get : "/v1/user/{id}"
    };
  }

  // ListUser 获取用户列表
  // `GET /v1/users?page_num=1&page_size=10` | `ListUserRequest(page_num: 1, page_size: 10)`
  rpc ListUser (ListUserRequest) returns (ListUserResponse) {
    option (google.api.http) = {
      get : "/v1/users"
    };
  }
}

message User {
  int64 id = 1;
  string name = 2;
}

message CreateUserRequest {
  string name = 1;
}

message CreateUserResponse {
  User item = 1;
}

message DeleteUserRequest {
  int64 id = 1;
}

message ModifyUserRequest {
  int64 id = 1;
  string name = 2;
}

message GetUserRequest {
  int64 id = 1;
}

message GetUserResponse {
  User item = 1;
}

message ListUserRequest {
  int64 page_num = 1;
  int64 page_size = 2;
}

message ListUserResponse {
  repeated User list = 1;
}