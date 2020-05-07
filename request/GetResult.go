package request

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

/*
	ts 是时间
	sign 固定加密 input + 其他
	salt 是时间加随机
*/

const (
	appVersion = "5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 " +
		"(KHTML, like Gecko) Chrome/81.0.4044.129 Safari/537.36"
	cookie = "DICT_UGC=be3af0da19b5c5e6aa4e17bd8d90b28a|; OUTFOX_SEARCH_USER_ID" +
		"=-1102944674@117.151.182.145; JSESSIONID=abcc6jQCXf64J8uXVcPhx; OUTFOX_SE" +
		"ARCH_USER_ID_NCOO=1582404644.2664843; SESSION_FROM_COOKIE=unknown; ___rl__te" +
		"st__cookies=1588771748533"
)

func GetTime() string {
	return string(time.Now().UnixNano() / 1e6)
}

func MD5( str string ) string {
	myHash := md5.Sum([]byte(strings.ToLower(str)))
	fmt.Printf("%x\n", myHash)
	return fmt.Sprintf("%x", myHash)
}

func GetSalt() string {
	return GetTime() + strconv.Itoa(rand.Intn(10))
}

func GetSign( key, salt string ) string {
	str := "fanyideskweb" + key + salt + "Nw(nmmbP%A-r6U3EUn]Aj"
	return MD5(str)
}

func getBV() string {
	return MD5(appVersion)
}


func Translate( input string ) ( output string ) {
	youdaoUrl := "http://fanyi.youdao.com/translate?smartresult=dict&smartresult=rule"
	req, err := http.NewRequest("POST", youdaoUrl, nil)
	if err != nil {
		log.Println(err)
	}

	salt := GetSalt()

	// form data
	data := url.Values{}
	data.Add("i", input)
	data.Add("from", "AUTO")
	data.Add("to", "AUTO")
	data.Add("smartresult", "dict")
	data.Add("client", "fanyideskweb")
	data.Add("action", "FY_BY_REALTlME")
	data.Add("salt", GetSalt())
	data.Add("sign", GetSign(input, salt))
	data.Add("ts", GetTime())
	data.Add("bv", "b396e111b686137a6ec711ea651ad37c")
	data.Add("doctype", "json")
	data.Add("keyfrom", "fanyi.web")
	data.Add("version", "2.1")

	// build header
	req.Header.Add("Accept","application/json, text/javascript, */*; q=0.01")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Add("Content-Length", string(len(data)))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("Cookie", cookie)
	req.Header.Add("Host", "fanyi.youdao.com")
	req.Header.Add("Origin", "http://fanyi.youdao.com")
	req.Header.Add("Referer", "http://fanyi.youdao.com/")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.129 Safari/537.36")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")

	fmt.Println(data)
	// add form data
	req.Form = data

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		//log.Println(resp)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//log.Println(err)
	}

	fmt.Println(string(body))
	return ""
}
