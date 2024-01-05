/*
Write a program that generates a tree-like structure representing the hierarchy of files and directories starting from a specified root directory. Use os and os.FileInfo to gather information about files and directories.
*/
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	path := `C:\Users\Admin\Documents\GoLang\Go_Programs\GoProgram\BasicTopics`
	createFileTree(path, " ")

}
func createFileTree(path string, indent string) {
	file, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, files := range file {
		fmt.Println(indent + files.Name())
		if files.IsDir() {
			createFileTree(filepath.Join(path, files.Name()), indent+"  ")
		}
	}
}
