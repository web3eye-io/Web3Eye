package utils

import (
	"fmt"
	"reflect"
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
