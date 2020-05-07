package youdao

import (
	"crypto/md5"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
	"youdao_alfred/request"
)

const (
	cookie = "OUTFOX_SEARCH_USER_ID=1721894360@59.111.179.141; _ntes_nnid=" +
		"28c86721ce2c2392d0b4bc1c066195c2,1562810189644; OUTFOX_SEARCH_USER_I" +
		"D_NCOO=1717529083.05212; P_INFO=qducst_xmt@163.com|1572920094|0|othe" +
		"r|00&99|shd&1572744638&mail163#shd&null#10#0#0|&0|mail163|qducst_xmt" +
		"@163.com; JSESSIONID=aaadMecWfzOYVgeMhs8dx; ___rl__test__cookies=158" +
		"4780646068"
	youdaoUrl = "http://fanyi.youdao.com/translate_o?smartresult=dict&smartresult=rule"
)

func GetTime() string {
	return strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
}

func MD5(str string) string {
	// don't add strings.Lower
	myHash := md5.Sum([]byte(str))
	return fmt.Sprintf("%x", myHash)
}

func GetSalt() string {
	return GetTime() + strconv.Itoa(rand.Intn(10))
}

func GetSign(key, salt string) string {
	str := "fanyideskweb" + key + salt + "Nw(nmmbP%A-r6U3EUn]Aj"
	return MD5(str)
}

func Translate(key string) string {
	ts := GetTime()
	salt := GetSalt()
	sign := GetSign(key, salt)

	req := request.Request{}
	req.Header = map[string]string{
		// key step for post
		"Content-Type": "application/x-www-form-urlencoded; charset=UTF-8",
		"Cookie":       cookie,
		"Referer":      "http://fanyi.youdao.com/",
		"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36",
	}
	req.FormData = map[string]string{
		"i":           key,
		"from":        "AUTO",
		"to":          "AUTO",
		"smartresult": "dict",
		"client":      "fanyideskweb",
		"action":      "FY_BY_REALTlME",
		"salt":        salt,
		"sign":        sign,
		"ts":          ts,
		"bv":          "70244e0061db49a9ee62d341c5fed82a",
		"doctype":     "json",
		"keyfrom":     "fanyi.web",
		"version":     "2.1",
	}

	resp, err := req.POST(youdaoUrl)
	if err != nil {
		log.Println(err)
	}

	return resp
}
