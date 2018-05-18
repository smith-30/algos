package column2

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

var (
	_dict     = map[string][]string{}
	dict      = []string{}
	dummyDict = []string{"from", "time", "item", "form", "off", "test"}

	notFoundErr = errors.New("word key is not found.")
)

func init() {
	// 文字列をsortした結果をキーに、アナグラムを持つ
	_dict["opst"] = []string{"stop", "post", "tops"}
}

func _getDict() map[string][]string {
	return _dict
}

func setDict() {
	//読み込みファイル準備
	src, err := os.Open("./dict.csv")
	if err != nil {
		panic(err)
	}
	defer src.Close()
	reader := csv.NewReader(transform.NewReader(src, japanese.ShiftJIS.NewDecoder()))
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		for _, v := range record {
			s := strings.Split(v, "\r")
			for _, w := range s {
				dict = append(dict, w)
			}
		}
	}
	return
}

func Anagram() {
	// 容量割り当てときたいが目安がわからないので保留
	list := make(map[string][]string)

	for _, word := range dict {
		key := sortStr(word)
		list[key] = append(list[key], word)
	}

	for _, words := range list {
		for _, w := range words {
			fmt.Print(w, " ")
		}
		fmt.Println()
	}
}

func GetAnagram(key string) ([]string, error) {
	key = sortStr(key)

	if r, ok := _dict[key]; ok {
		return r, nil
	}

	return nil, notFoundErr
}

func sortStr(k string) string {
	s := strings.Split(k, "")
	sort.Strings(s)

	return strings.Join(s, "")
}
