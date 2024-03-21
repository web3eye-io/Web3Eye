package utils

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"

	"github.com/mr-tron/base58"
)

func Sha256Hash(i any) (string, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, i)
	if err != nil {
		return "", err
	}
	h := sha256.New()
	_, err = h.Write(buf.Bytes())
	if err != nil {
		return "", err
	}
	return base58.FastBase58Encoding(h.Sum([]byte{})), nil
}
