package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func getMD5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))

	fmt.Println(fmt.Sprintf("%x", h.Sum(nil)))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

func main()  {
	fmt.Println(getMD5Encode(""))
	fmt.Println(getMD5Encode("aaa1"))
	fmt.Println(getMD5Encode("aaa"))
}
