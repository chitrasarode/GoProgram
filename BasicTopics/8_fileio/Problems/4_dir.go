/*
Write a program that generates a tree-like structure representing the hierarchy of files and directories starting from a specified root directory. Use os and os.FileInfo to gather information about files and directories.
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("Error reading directories")
		return
	}

	for _, file := range files {
		if file.IsDir() {
			fmt.Println(file.Name())
			subDir, err := os.ReadDir(file.Name())
			if err != nil {
				fmt.Println("Error reading directories : ", err)
				return
			}
			for _, file := range subDir {
				fmt.Println("  ", file.Name())
			}

		} else {
			fmt.Println(file.Name())
		}
	}
}
