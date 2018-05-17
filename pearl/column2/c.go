package column2

import (
	"errors"
	"sort"
	"strings"
)

var (
	dict = map[string][]string{}

	notFoundErr = errors.New("word key is not found.")
)

func init() {
	// 文字列をsortした結果をキーに、アナグラムを持つ
	dict["opst"] = []string{"stop", "post", "tops"}
}

func getDict() map[string][]string {
	return dict
}

func GetAnagram(key string) ([]string, error) {
	key = sortStr(key)

	if r, ok := dict[key]; ok {
		return r, nil
	}

	return nil, notFoundErr
}

func sortStr(k string) string {
	s := strings.Split(k, "")
	sort.Strings(s)

	return strings.Join(s, "")
}
