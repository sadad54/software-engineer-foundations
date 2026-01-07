package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main(){
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <directory_path>")
		os.Exit(1)
	}
	
	root:= os.Args[1]
	fileCount := 0
	dirCount := 0

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error)error{
		if err !=nil{
			return err
		}

		if info.IsDir(){
			dirCount++
		} else {
			fileCount++
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("Directory Scanned:", root)
	fmt.Println("Total files:", fileCount)
	fmt.Println("Total directories:", dirCount)
}
go func() {
	fmt.Println("go routine executed concurrently")
}()