package main

import (
	"io"
	"os"
	"sync"
)

const fileToCopy = "large.txt"

var wg sync.WaitGroup

func copy1() {
	src, _ := os.ReadFile(fileToCopy)

	dst, _ := os.Create("copy1.txt")
	defer dst.Close()

	dst.Write(src)

	wg.Done()
}

func copy2() {
	src, _ := os.ReadFile(fileToCopy)

	os.WriteFile("copy2.txt", src, 0644)

	wg.Done()
}

func copy3() {
	src, _ := os.Open(fileToCopy)
	defer src.Close()

	dst, _ := os.Create("copy3.txt")
	defer dst.Close()

	io.Copy(dst, src)

	wg.Done()
}

func copy4() {
	src, _ := os.Open(fileToCopy)
	defer src.Close()
	fileInfo, _ := src.Stat()
	fileSize := fileInfo.Size()

	dst, _ := os.Create("copy4.txt")
	defer dst.Close()

	var bufSize int64 = 1024
	if fileSize < bufSize {
		bufSize = fileSize
	}

	var buf = make([]byte, bufSize)
	io.CopyBuffer(dst, src, buf)

	wg.Done()
}

func main() {
	wg.Add(4)

	go copy1()
	go copy2()
	go copy3()
	go copy4()

	wg.Wait()
}
