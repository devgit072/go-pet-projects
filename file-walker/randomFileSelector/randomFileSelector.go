package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/golang/glog"
	"github.com/karrick/godirwalk"
)

// env GOOS=linux GOARCH=amd64 go build -v github.com/constabulary/gb/cmd/gb
const (
	fileName        = "randomSelectedFiles.txt"
	randomFileCount = 500
	randomRange     = 30
	valueToCheck    = 1
)

func main() {
	list, err := traverseFileSystem()
	if err != nil {
		panic(err.Error())
	}
	glog.Infof("Total random file count: %d", len(list))
	for _, l := range list {
		glog.Infoln(l)
	}
	err = writeFilesListIntoFile(list)
	if err != nil {
		panic(err.Error())
	}
}

func traverseFileSystem() ([]string, error) {
	glog.Infof("Calculating list of random 500 files")
	count := 0
	t1 := time.Now()
	randomFileList := make([]string, 0)
	rootPath := os.Args[1]
	fmt.Println("Getting for filePath: ", rootPath)
	rand.Seed(time.Now().UnixNano())
	err := godirwalk.Walk(rootPath, &godirwalk.Options{
		Unsorted: true,
		Callback: func(osPathname string, de *godirwalk.Dirent) error {
			count++
			r := rand.Intn(randomRange)
			if r == valueToCheck {
				randomFileList = append(randomFileList, osPathname)
			}
			return nil
		},
		ErrorCallback: func(osPathname string, err error) godirwalk.ErrorAction {
			return godirwalk.SkipNode
		},
	})
	if err != nil {
		return nil, err
	}
	t2 := time.Since(t1)
	glog.Infof("Total files count in the system: %d", count)
	glog.Infof("Time Taken: %f", t2.Seconds())
	return randomFileList, nil
}

func writeFilesListIntoFile(files []string) error {
	err := truncateFileIfExists(fileName)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		return err
	}
	sep := "\n"
	for _, line := range files {
		l := fmt.Sprintf("%s%s", line, sep)
		glog.Infof("Ye: %s", l)
		if _, err = file.WriteString(l); err != nil {
			if true {
				return nil
			}
			return err
		}
	}
	return nil
}

func truncateFileIfExists(fileName string) error {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		glog.Infof("%s doenst exist, so simply returning and no need to truncate")
		return nil
	}
	f, err := os.OpenFile(fileName, os.O_RDWR, 0666)
	defer f.Close()
	if err != nil {
		return err
	}
	f.Truncate(0)
	f.Seek(0, 0)
	return nil
}
