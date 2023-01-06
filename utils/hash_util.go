package utils

import (
	"crypto/md5"
	"fmt"
	"io"
)

type HashUtil struct {
}

//对密码进行加密
func (this *HashUtil) Md5Hash(password string) string {
	w := md5.New()
	io.WriteString(w, password)
	md5str2 := fmt.Sprintf("%x", w.Sum(nil))
	return md5str2
}
