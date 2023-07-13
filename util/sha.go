package util

import (
	"crypto/md5"
	"encoding/hex"
	"hash"
)

type Sha1Stream struct {
	_sha1 hash.Hash
}

func MD5(data []byte) string {
	_md5 := md5.New()
	_md5.Write(data)
	return hex.EncodeToString(_md5.Sum([]byte("")))
}
