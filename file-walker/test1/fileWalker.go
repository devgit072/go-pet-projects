package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

var count int

func VisitFile(fp string, fi os.FileInfo, err error) error {
	if err != nil {
		fmt.Println(err) // can't walk here,
		return nil       // but continue walking elsewhere
	}
	count++
	//fmt.Println(fp)
	//if fi.IsDir() {
	//	return nil // not a file.  ignore.
	//}
	//matched, err := filepath.Match("*.go", fi.Name())
	//if err != nil {
	//	fmt.Println(err) // malformed pattern
	//	return err       // this is fatal.
	//}
	//if matched {
	//	fmt.Println(fp)
	//}
	return nil
}

func main() {
	t1 := time.Now()
	rootPath := os.Args[1]
	fmt.Println("Getting for filePath: ", rootPath)
	filepath.Walk(rootPath, VisitFile)
	t2 := time.Since(t1)
	fmt.Println(count)
	fmt.Println("Time Taken:", t2.Seconds())
}
