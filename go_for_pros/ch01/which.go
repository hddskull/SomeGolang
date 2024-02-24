package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func which() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Give argument!")
		return
	}

	file := arguments[1]

	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)

	for _, dir := range pathSplit {
		fullPath := filepath.Join(dir, file)
		//path exists?
		fileInfo, err := os.Stat(fullPath)
		if err == nil {
			mode := fileInfo.Mode()
			//is regular file
			if mode.IsRegular() {
				//is it executable
				if mode&0111 != 0 {
					fmt.Println(fullPath)
					return
				}
			}
		}
	}
}
