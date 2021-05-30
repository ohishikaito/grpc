package response

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func RecordNotFound(err error) error {
	// どっちでも使えるから、場合によって使い分ける？
	return status.Errorf(codes.NotFound, err.Error())
	// return status.New(codes.NotFound, err.Error()).Err()
}
