package util

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/nu7hatch/gouuid"
)

func makeToken() string {
	u4, _ := uuid.NewV4()
	h := md5.New()
	h.Write([]byte(u4.String()))
	return hex.EncodeToString(h.Sum(nil))

}
