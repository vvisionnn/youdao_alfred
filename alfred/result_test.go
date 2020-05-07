package alfred

import (
	"fmt"
	"testing"
)

func TestNewResult(t *testing.T) {
	r := NewResult()

	r.Append(&ResultItem{
		Title:        "test",
		Subtitle:     "test subtitle",
		Arg:          "ok",
		QuickLookUrl: "https://www.baidu.com",
		Mods:         nil,
	})

	fmt.Println((*r.Items)[0])
}
