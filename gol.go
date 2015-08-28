package main

import (
	"fmt"
	// "io/ioutil"
	// "flag"
	"os"
	"path/filepath"
	"strings"
)

type matcher func(candidate string) bool

func main() {
	root := "/home/murat"
	filewalker(root, getMatcher("blah"))
}

func getMatcher(substring string) matcher {
	return func(candidate string) bool {
		return strings.Contains(candidate, substring)
	}
}

/** http://stackoverflow.com/questions/6608873/file-system-scanning-in-golang */

func filewalker(root string, matcher matcher) {
	visit := func(path string, f os.FileInfo, err error) error {
		if matcher(path) {
			fmt.Println(path)
		}
		return nil
	}

	// flag.Parse()
	err := filepath.Walk(root, visit)
	fmt.Printf("filepath.Walk() returned %v\n", err)
}

// https://github.com/monochromegane/the_platinum_searcher/blob/master/find.go
// func readdir() {
// 	fileInfos, _ := ioutil.ReadDir("/home/murat")
// 	for _, fd := range fileInfos {
// 		fileInfo := newFileInfo("/home/murat", l, info.follow)
// 		fmt.Println("Name: ", fileInfo.Name())
// 	}
// }
