package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"hash"
	"strconv"
	"strings"
	"time"
)

func main() {
	err := Authorize("b4cfeeadca80f6f5", "b6e9e018302b96bc089a4c23046d16", "DELETE", "live", "4.jpg", "b4cfeeadca80f6f5:bja6AWDC97Q1wck2DfQmsL4El8Y=:1534762900")
	fmt.Println(err)
}

func Authorize(keyID, secret string, method, bucket, file, token string) (err error) {
	// token keyid:sign:time
	var (
		expire int64
		delta  int64
		now    int64
		kid    string
		ss     = strings.Split(token, ":")
	)
	if len(ss) != 3 {
		return errors.New("1")
	}
	kid = ss[0]
	if kid != keyID {
		fmt.Println(kid)
		fmt.Println(keyID)
		return errors.New("2")
	}
	if expire, err = strconv.ParseInt(ss[2], 10, 64); err != nil {
		return errors.New("3")
	}
	now = time.Now().Unix()
	fmt.Println("expire:", expire)
	// > ±15 min is forbidden
	if expire > now {
		delta = expire - now
	} else {
		delta = now - expire
	}
	if delta > 900 {
		return errors.New("4")
	}
	err = sign(ss[1], method, bucket, file, secret, expire)
	return
}

func sign(src, method, bucket, file, keySecret string, expire int64) (err error) {
	var (
		content string
		mac     hash.Hash
	)
	content = fmt.Sprintf("%s\n%s\n%s\n%d\n", method, bucket, file, expire)
	mac = hmac.New(sha1.New, []byte(keySecret))
	mac.Write([]byte(content))
	if ets := base64.StdEncoding.EncodeToString(mac.Sum(nil)); ets != src {
		fmt.Println(ets)
		fmt.Println(src)
		return errors.New("5")
	}
	return
}
