package main

import (
	"flag"

	"youdao_alfred/alfred"
	"youdao_alfred/youdao"
)

func main() {

	key := flag.String("k", "Hello World", "the sentence you want to translate")
	flag.Parse()

	result := alfred.NewResult()

	dst := youdao.Translate(*key)
	result.Append(alfred.NewItem(dst, "test subtitle", dst))

	result.End()
}
