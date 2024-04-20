package main

import (
	"io"
	"os"
)

const fileToCopy = "large.txt"

func copy1() {
	src, _ := os.ReadFile(fileToCopy)

	dst, _ := os.Create("copy1.txt")
	defer dst.Close()

	dst.Write(src)
}

func copy2() {
	src, _ := os.ReadFile(fileToCopy)

	os.WriteFile("copy2.txt", src, 0644)
}

func copy3() {
	src, _ := os.Open(fileToCopy)
	defer src.Close()

	dst, _ := os.Create("copy3.txt")
	defer dst.Close()

	io.Copy(dst, src)
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
}

func main() {
	go copy1()
	go copy2()
	go copy3()
	copy4()
}
