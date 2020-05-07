package main

import (
	"flag"
	"strings"

	"youdao_alfred/alfred"
	"youdao_alfred/youdao"
)

func main() {

	key := flag.String("k", "Hello World", "the sentence you want to translate")
	flag.Parse()

	//fmt.Println("SRC -> ", *key, " LEN -> ", len(*key))
	//fmt.Println("OTHERS -> ", flag.Args())
	if len(flag.Args()) > 0 {
		*key += " " + strings.Join(flag.Args(), " ")
	}
	//fmt.Println("SRC -> ", *key, " LEN -> ", len(*key))

	result := alfred.NewResult()
	dst := youdao.Translate(*key)
	result.Append(alfred.NewItem(dst, *key, dst))

	//fmt.Println("DST -> ", dst)
	//fmt.Println("result -> ", result)

	result.End()
}
