package tree

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

const (
	separater = "/"
)

func Tree(path string) error {
	fmt.Println("path: ", path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if file.IsDir() {
			Tree(filepath.Join(path, file.Name()))
			continue
		}
		fmt.Println(file.Name())
	}
	return nil
}
