package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/mozillazg/request"
)

func main() {
	fmt.Println(Translate("Не уверен, что это он"))
	fmt.Println(Translate2("Не уверен, что это он"))
}

const (
	cookie = "OUTFOX_SEARCH_USER_ID=1721894360@59.111.179.141; _ntes_nnid=28c86721ce2c2392d0b4bc1c066195c2,1562810189644; OUTFOX_SEARCH_USER_ID_NCOO=1717529083.05212; P_INFO=qducst_xmt@163.com|1572920094|0|other|00&99|shd&1572744638&mail163#shd&null#10#0#0|&0|mail163|qducst_xmt@163.com; JSESSIONID=aaadMecWfzOYVgeMhs8dx; ___rl__test__cookies=1584780646068"
	youdaoUrl = "http://fanyi.youdao.com/translate_o?smartresult=dict&smartresult=rule"
)

func GetTime() string {
	return strconv.FormatInt(time.Now().UnixNano() / 1e6, 10)
}

func MD5( str string ) string {
	// don't add strings.Lower
	myHash := md5.Sum([]byte(str))
	return fmt.Sprintf("%x", myHash)
}

func GetSalt() string {
	return GetTime() + strconv.Itoa(rand.Intn(10))
}

func GetSign( key, salt string ) string {
	str := "fanyideskweb" + key + salt + "Nw(nmmbP%A-r6U3EUn]Aj"
	return MD5(str)
}

//func getBV() string {
//	return MD5(appVersion)
//}

func Translate( key string ) string {
	ts := GetTime()
	salt := GetSalt()
	sign := GetSign(key, salt)
	//fmt.Println(ts)
	//fmt.Println(salt)
	//fmt.Println(sign)

	// form data
	data := url.Values{}
	data.Set("i", key)
	data.Set("from", "AUTO")
	data.Set("to", "AUTO")
	data.Set("smartresult", "dict")
	data.Set("client", "fanyideskweb")
	data.Set("action", "FY_BY_REALTlME")
	data.Set("salt", salt)
	data.Set("sign", sign)
	data.Set("ts", ts)
	data.Set("bv", "70244e0061db49a9ee62d341c5fed82a")
	data.Set("doctype", "json")
	data.Set("keyfrom", "fanyi.web")
	data.Set("version", "2.1")

	// 不能直接将 data 加入 req.Form 应该按照下面操作
	body := bytes.NewBufferString(data.Encode())
	req, err := http.NewRequest(http.MethodPost, youdaoUrl, body)
	if err != nil {
		log.Println(err)
	}

	// build header
	//req.Header.Add("Accept","application/json, text/javascript, */*; q=0.01")
	//req.Header.Add("Accept-Encoding", "gzip, deflate")
	//req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	//req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Cookie", cookie)
	//req.Header.Add("Host", "fanyi.youdao.com")
	//req.Header.Add("Origin", "http://fanyi.youdao.com")
	req.Header.Set("Referer", "http://fanyi.youdao.com/")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36")
	//req.Header.Add("X-Requested-With", "XMLHttpRequest")

	resp, e := http.DefaultClient.Do(req)
	if e != nil {
		log.Println(e)
	}
	defer resp.Body.Close()
	//fmt.Println(resp)
	//fmt.Println(resp.StatusCode)

	buf, _ := ioutil.ReadAll(resp.Body)

	return string(buf)
}

func Translate2( key string ) string {
	//fmt.Println(key)
	c := new(http.Client)
	req := request.NewRequest(c)
	ts := GetTime()
	salt := GetSalt()
	sign := GetSign(key, salt)
	//fmt.Println(ts, salt, sign)


	req.Data = map[string]string{
		"i": key,
		"from": "AUTO",
		"to": "AUTO",
		"smartresult": "dict",
		"client": "fanyideskweb",
		"action": "FY_BY_REALTlME",
		"salt": salt,
		"sign": sign,
		"ts": ts,
		"bv": "70244e0061db49a9ee62d341c5fed82a",
		"doctype": "json",
		"keyfrom": "fanyi.web",
		"version": "2.1",
	}

	//fmt.Println(strconv.Itoa(len(req.Data)))
	req.Headers = map[string]string{
		//"Accept":"application/json, text/javascript, */*; q=0.01",
		//"Accept-Encoding": "gzip, deflate",
		//"Accept-Language": "zh-CN,zh;q=0.9",
		//"Content-Length": strconv.Itoa(len(data.Encode())),
		//"Content-Type": "application/x-www-form-urlencoded; charset=UTF-8",
		"Cookie": cookie,
		//"Host": "fanyi.youdao.com",
		//"Origin": "http://fanyi.youdao.com",
		"Referer": "http://fanyi.youdao.com/",
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36",
		//"X-Requested-With": "XMLHttpRequest",
	}

	resp, err := req.Post(youdaoUrl)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	j, _ := resp.Json()
	jj, _ := j.MarshalJSON()
	return string(jj)
}
