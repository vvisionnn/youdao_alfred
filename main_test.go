package main

import (
	"fmt"
	"testing"
)

func TestTranslate(t *testing.T) {
	fmt.Println(Translate("不知道这个是不是"))
	fmt.Println(Translate2("不知道这个是不是"))
}
