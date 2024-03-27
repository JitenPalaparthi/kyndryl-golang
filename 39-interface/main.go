package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Fprintln(os.Stdout, "Hello World!")
	// os.Stdout has implemented io.Writer interface
	/// Write(p []byte) (n int, err error)

	fw := FileWriter{FileName: "output.log"}
	fmt.Fprintln(&fw, "Hello World")

}

type FileWriter struct {
	FileName string
}

func (fw *FileWriter) Write(p []byte) (n int, err error) {
	f, err := os.OpenFile(fw.FileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	if n, err := f.Write(p); err != nil {
		return 0, err
	} else {
		return n, nil
	}
}
