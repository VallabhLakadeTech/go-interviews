package main

import (
	"fmt"
	"sync"
)

func main() {

	fileList := []string{"1", "2", "3", "4"}
	processedFileList := []string{}

	ch := make(chan string)
	var wg sync.WaitGroup

	go func() {
		for processedFile := range ch {
			processedFileList = append(processedFileList, processedFile)
		}
		//1min

	}()

	for _, file := range fileList {
		wg.Add(1)
		go processFiles(file, &wg, ch)
	}
	// wg.Wait()
	// close(ch)

	fmt.Println(processedFileList)
}

func processFiles(fileName string, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()
	//Some processing here
	fileName = fileName + ".json"
	ch <- fileName
}
