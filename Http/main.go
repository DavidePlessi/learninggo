package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func (logWriter) Write(bs []byte) (int, error) {
	text := string(bs)
	fmt.Println(text)
	return len(bs), nil
}

func main() {
	res, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	lw := logWriter{}
	_, _ = io.Copy(lw, res.Body)
}
