package main

import (
	"fmt"
	"os"
	"time"

	"github.com/karrick/godirwalk"
)

func main() {
	count := 0
	t1 := time.Now()
	rootPath := os.Args[1]
	fmt.Println("Getting for filePath: ", rootPath)
	godirwalk.Walk(rootPath, &godirwalk.Options{
		Unsorted: true,
		Callback: func(osPathname string, de *godirwalk.Dirent) error {
			count++
			//fmt.Println(osPathname)
			return nil
		},
		ErrorCallback: func(osPathname string, err error) godirwalk.ErrorAction {
			return godirwalk.SkipNode
		},
	})
	t2 := time.Since(t1)
	fmt.Println(count)
	fmt.Println("Time Taken:", t2.Seconds())
}
