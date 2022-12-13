package utils

import (
	"bytes"
	"encoding/binary"
	"strconv"
)

func MinInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func Str2uint64(str string) (uint64, error) {
	baseNum := 16
	targetNum := 64
	i, err := strconv.ParseUint(str, baseNum, targetNum)
	return i, err
}

func Uint642Bytes(n uint64) ([]byte, error) {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.BigEndian, n)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func Bytes2Uint64(b []byte) (uint64, error) {
	buf := bytes.NewBuffer(b)
	var n uint64
	err := binary.Read(buf, binary.BigEndian, &n)
	return n, err
}
