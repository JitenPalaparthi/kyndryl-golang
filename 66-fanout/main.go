package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	wc, chars := make(chan int), make(chan int)
	//line := readFileLines("data2.txt")
	line := readFileLineByLine("data2.txt")
	done := GetLineStats(line, wc, chars)
out:
	for {
		select {
		case c := <-wc:
			fmt.Println("Word Count", c)
		case crs := <-chars:
			fmt.Println("Chars Count", crs)
		case <-done:
			break out
		}
	}
}

func readFileLines(logfile string) chan string {
	line := make(chan string)
	f, err := os.OpenFile(logfile, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return nil
	}
	//defer f.Close()
	go func(*os.File) {
		defer f.Close()
		rd := bufio.NewReader(f)
		for {
			l, err := rd.ReadString('\n')

			//fmt.Println("", l, err)
			if err != nil {
				//close(line)
				if err == io.EOF {
					close(line)
					break
				}
				return
			}
			line <- l
			//_ = l
			//_ = line // GET the line string
		}
	}(f)
	return line
}

func GetLineStats(line chan string, wc, chars chan int) <-chan struct{} {
	done := make(chan struct{})
	go func() {
		for l := range line {
			select {
			case wc <- len(strings.Fields(l)):
			case chars <- len(l):
			}
		}

		close(done)
	}()
	return done
}

func readFileLineByLine(fn string) chan string {
	chLine := make(chan string)
	//	chLineNo := make(chan int)
	f, err := os.OpenFile("data.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	go func(*os.File) {
		defer f.Close()
		buf := make([]byte, 1)
		lineNum := 1
		var line string
		for {
			n, err := f.Read(buf)
			if string(buf) == "\n" {
				chLine <- line
				//chLineNo <- lineNum
				//}
				line = ""
				lineNum++
			} else {
				line = line + string(buf)
			}
			if err != nil || n == 0 {
				chLine <- line
				//	chLineNo <- lineNum
				close(chLine)
				//	close(chLineNo)
				break
			}
		}
	}(f)
	return chLine
}
