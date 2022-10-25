package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	/* fileName := "birthday_001.txt"

	newName, err := match(fileName, 4)
	if err != nil {
		fmt.Println("no match")
		os.Exit(1)
	}
	fmt.Println(newName) */
	dir := "./sample"
	count := 0
	var toRename []string

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if file.IsDir() {
		} else {
			_, err := match(file.Name(), 0)
			if err == nil {
				count++
				toRename = append(toRename, file.Name())
			}
		}
	}

	for _, originalFileName := range toRename {
		originalPath := filepath.Join(dir, originalFileName)
		newFileName, _ := match(originalFileName, count)
		if err != nil {
			panic(err)
		}
		newPath := filepath.Join(dir, newFileName)
		fmt.Printf("mv %s => %s\n", originalPath, newPath)
		err := os.Rename(originalPath, newPath)
		if err != nil {
			panic(err)
		}

	}

}

func match(fileName string, total int) (string, error) {

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
	return fmt.Sprintf("%s - %d of %d.%s", strings.Title(name), number, total, ext), nil
}
