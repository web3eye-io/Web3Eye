package utils

import (
	"fmt"
	"io"
	"reflect"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RawCodec struct{}

func (cb RawCodec) Marshal(v interface{}) ([]byte, error) {
	return v.([]byte), nil
}

func (cb RawCodec) Unmarshal(data []byte, v interface{}) error {
	ba, ok := v.(*[]byte)
	if !ok {
		return fmt.Errorf("cannot transfer %v type to *[]byte", reflect.TypeOf(v))
	}
	*ba = append(*ba, data...)
	return nil
}

func (cb RawCodec) Name() string { return "dtm_raw" }

// check the connection of grpc stream err
func CheckStreamErrCode(err error) bool {
	if err == io.EOF ||
		status.Code(err) == codes.Unavailable ||
		status.Code(err) == codes.Canceled ||
		status.Code(err) == codes.Unimplemented {
		return true
	}
	return false
}
