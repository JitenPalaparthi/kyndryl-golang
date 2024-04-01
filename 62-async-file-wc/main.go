package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	done := make(chan struct{})

	f, err := os.OpenFile("data.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		os.Exit(1) // non zero exit is a fatal. Error from unix systems perspective
	}
	defer f.Close()
	wcMap := make(map[int]int)
	go func() {
		buf := make([]byte, 1)
		lineNum := 1
		var line string

		for {
			n, err := f.Read(buf)
			if string(buf) == "\n" {
				chWC := WordCount(line) // go routine internally
				v := <-chWC
				if v != 0 {
					wcMap[lineNum] = v
				}
				//}
				line = ""
				lineNum++
			} else {
				line = line + string(buf)
			}
			if err != nil || n == 0 {
				chWC := WordCount(line) // go routine internally
				v := <-chWC
				if v != 0 {
					wcMap[lineNum] = v
				}
				done <- struct{}{}
			}
		}
	}()
	<-done // Future pattern
	fmt.Println(wcMap)

	//runtime.Goexit()
	//time.Sleep(time.Second * 5)
}

func WordCount(line string) chan int {
	countCh := make(chan int)
	go func() {
		strs := strings.Fields(line)
		countCh <- len(strs)
		close(countCh)
	}()
	return countCh
}

// map[1:7 2:11 4:4 5:4 6:15 7:15 8:6]
