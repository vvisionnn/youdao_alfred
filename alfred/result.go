package alfred

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type ModItem struct {
	Valid     bool              `json:"valid,omitempty"`
	Arg       string            `json:"arg,omitempty"`
	Subtitle  string            `json:"subtitle,omitempty"`
	Variables map[string]string `json:"variables,omitempty"`
}

type ResultItem struct {
	//Valid        bool                   `json:"valid,omitempty"`
	//Uid          string                 `json:"uid,omitempty"`
	//Type         string                 `json:"type,omitempty"`
	//Autocomplete string                 `json:"autocomplete,omitempty"`
	Title        string              `json:"title,omitempty"`
	Subtitle     string              `json:"subtitle,omitempty"`
	Arg          string              `json:"arg,omitempty"`
	QuickLookUrl string              `json:"quicklookurl,omitempty"`
	Mods         map[string]*ModItem `json:"mods,omitempty"`
}

type Result struct {
	Items *[]*ResultItem `json:"items"`
}

func NewResult() *Result {
	items := make([]*ResultItem, 0)
	return &Result{Items: &items}
}

func NewItem(title, subtitle, arg string) *ResultItem {
	return &ResultItem{
		Title:    title,
		Arg:      arg,
		Subtitle: subtitle,
	}
}

func (r *Result) Append(items ...*ResultItem) {
	*r.Items = append(*r.Items, items...)
}

func (r *Result) Count() int {
	return len(*r.Items)
}

func (r *Result) End() {
	b := new(bytes.Buffer)

	if err := json.NewEncoder(b).Encode(r); err != nil {
		log.Println(err)
	}
	fmt.Print(b.String())

	os.Exit(0)
}
