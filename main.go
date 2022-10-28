package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type file struct {
	name string
	path string
}

func main() {
	dir := "sample"
	var toRename []file

	//add file to []toRename recursively
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if _, err := match(info.Name()); err == nil {
			toRename = append(toRename, file{
				name: info.Name(),
				path: path,
			})
		}
		return nil
	})

	for _, f := range toRename {
		fmt.Printf("%q\n", f)
	}

	for _, originalFile := range toRename {

		var newFile file
		var err error

		newFile.name, err = match(originalFile.name)

		if err != nil {
			fmt.Println("Error matching:", originalFile.path, err.Error())
		}

		newFile.path = filepath.Join(dir, newFile.name)
		fmt.Printf("mv %s => %s\n", originalFile.path, newFile.path)

		err = os.Rename(originalFile.path, newFile.path)
		if err != nil {
			fmt.Println("Error renaming:", originalFile.path, err.Error())
		}
	}

}

func match(fileName string) (string, error) {

	pieces := strings.Split(fileName, ".")
	ext := pieces[len(pieces)-1]
	tmp := strings.Join(pieces[0:len(pieces)-1], ".")
	pieces = strings.Split(tmp, "_")
	name := strings.Join(pieces[0:len(pieces)-1], "_")
	number, err := strconv.Atoi(pieces[len(pieces)-1])

	if err != nil {
		return "", fmt.Errorf("%s did not match our pattern", fileName)
	}

	//Birthday - 1.txt
	return fmt.Sprintf("%s - %d.%s", strings.Title(name), number, ext), nil
}
