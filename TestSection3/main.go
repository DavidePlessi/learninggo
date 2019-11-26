package main

import (
	"fmt"
	"io"
	"os"
	"path"
)

type logWriter struct{}

func (logWriter) Write(bs []byte) (int, error) {
	text := string(bs)
	fmt.Println(text)
	return len(bs), nil
}

func getFullPath(filename string) string {
	currentDirectory, err := os.Getwd()
	check(err)
	return path.Join(currentDirectory, "TestSection3", filename)
}

func check(e error) {
	if e != nil {
		fmt.Println("Error: ", e)
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Argument missing")
		os.Exit(1)
	}
	filePath := getFullPath(os.Args[1])
	file, err := os.Open(filePath)
	check(err)

	lw := logWriter{}
	_, err = io.Copy(lw, file)
	check(err)
}
