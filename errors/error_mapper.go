package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var errorCodeToStatus = map[string]error{
	"23505": status.Errorf(codes.AlreadyExists, "Failed to create menu item"),
}

