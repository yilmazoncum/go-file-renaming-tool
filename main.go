package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

var re = regexp.MustCompile("^(.+?) ([0-9]{4}) [(]([0-9]+) of ([0-9]+)[)][.](.+?)$")
var replaceString = "$2 - $1 - $3 of $4.$5"

type file struct {
	name string
	path string
}

func main() {
	var dry bool
	flag.BoolVar(&dry, "dry", true, "whether or not this should be a real or dry run")
	flag.Parse()

	dir := "sample"
	var toRename []string

	//add file to []toRename recursively
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if _, err := match(info.Name()); err == nil {
			toRename = append(toRename, path)
		}

		return nil
	})

	for _, oldPath := range toRename {

		dir := filepath.Dir(oldPath)
		filename := filepath.Base(oldPath)
		newFilename, _ := match(filename)
		newPath := filepath.Join(dir, newFilename)
		fmt.Printf("mv %s => %s\n", oldPath, newPath)

		if !dry {
			err := os.Rename(oldPath, newPath)
			if err != nil {
				fmt.Println("Error renaming:", oldPath, newPath, err.Error())
			}
		}
	}

}

func match(filename string) (string, error) {
	if !re.MatchString(filename) {
		return "", fmt.Errorf("%s didn't match our pattern", filename)
	}
	return re.ReplaceAllString(filename, replaceString), nil
}
