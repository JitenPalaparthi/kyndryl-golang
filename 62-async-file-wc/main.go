package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

func main() {

	f, err := os.OpenFile("data1.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		os.Exit(1) // non zero exit is a fatal. Error from unix systems perspective
	}
	defer f.Close()
	wcMap := make(map[int]int)
	buf := make([]byte, 1)
	lineNum := 1
	var line string
	for {
		n, err := f.Read(buf)
		if string(buf) == "\n" {
			wc := WordCount(line)
			//fmt.Println(wc)
			if wc != 0 {
				wcMap[lineNum] = wc
			}
			line = ""
			lineNum++
		}
		line = line + string(buf)

		if err != nil || n == 0 {
			break
		}
	}

	fmt.Println(wcMap)

	runtime.Goexit()
}

func WordCount(line string) int {
	fmt.Println(line)
	strs := strings.Fields(line)
	return len(strs)
}
