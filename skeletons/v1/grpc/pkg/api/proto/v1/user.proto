syntax = "proto3";

package proto;
option go_package = "./";

// User implements a user rpc service.
service User {
  rpc Create(CreateRequest) returns (CreateResponse) {}
}

message CreateRequest {
  string request = 1;
}

message CreateResponse {
  string response = 1;
}