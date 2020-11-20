package qingyun

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"net/url"
	"strings"
)

const ()

// Signature 签名机制
// https://docs.qingcloud.com/product/api/common/signature.html
func Signature(reqMethod string, platform string, param url.Values, secret string) string {

	// fmt.Println("param.Encode=", param.Encode())
	urls := param.Encode()
	source := reqMethod + "\n" + platform + "\n" + urls

	// fmt.Println(source)
	// return ShaHmac1(source, secret)
	return ShaHmac256(source, secret)
}

// ShaHmac1 计算 HMAC-SHA1 校验码
func ShaHmac1(source, secret string) string {
	key := []byte(secret)
	shaHmac := hmac.New(sha1.New, key)
	_, _ = shaHmac.Write([]byte(source))
	signedBytes := shaHmac.Sum(nil)
	signedString := base64.StdEncoding.EncodeToString(signedBytes)
	return signedString
}

// ShaHmac1 计算 HMAC-SHA1 校验码
func ShaHmac256(source, secret string) string {
	key := []byte(secret)
	shaHmac := hmac.New(sha256.New, key)
	_, _ = shaHmac.Write([]byte(source))
	signedBytes := shaHmac.Sum(nil)
	signedString := base64.StdEncoding.EncodeToString(signedBytes)
	signedString = strings.TrimSpace(signedString)
	signedString = strings.Replace(signedString, " ", "+", -1)
	signedString = url.PathEscape(signedString)
	/*
		https://docs.qingcloud.com/product/api/common/signature.html
		编码时空格要转换成 “%20” , 而不是 “+”
		resolv: https://www.jianshu.com/p/2ba7dda583b5
	*/
	return signedString
}
