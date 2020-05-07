package request

import "testing"

func TestGet(t *testing.T) {
	//Get("http://fanyi.youdao.com/translate_o?smartresult=dict&smartresult=rule")
	Translate("word")
}