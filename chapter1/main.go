package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	defer func(now time.Time) {
		elapseTime := time.Since(now).Milliseconds()
		fmt.Printf("cost time is %d ms\n", elapseTime)
	}(time.Now())
}

// Curl sends http get request to specified url
func Curl() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "curl: %v\n", err)
			os.Exit(1)
		}
		defer resp.Body.Close()
		fmt.Printf("%s\n", resp.Status)
		// _, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "curl: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}

func dup2() {
	files := os.Args[1:]
	if len(files) == 0 {
		counts := make(map[string]int)
		temp := make(map[string]int)
		fmt.Println("Enter something (type 'exit' to quit): ")
		countLines(os.Stdin, counts, temp)
		for line, count := range counts {
			if count > 1 {
				fmt.Printf("%d\t%s\n", count, line)
			}
		}
	} else {
		counts := make(map[string]int)
		for _, file := range files {
			temp := make(map[string]int)
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "open file err: %v\n", err)
				continue
			}
			countLines(f, counts, temp)
			for line, count := range temp {
				if counts[line] > 1 {
					fmt.Printf("%d\t%s", counts[line], line)
				}
				if count > 0 && counts[line] > 1 {
					fmt.Printf("\t%s\n", f.Name())
				} else {
					fmt.Printf("\n")
				}
			}
			f.Close()
		}
	}
}

func countLines(f *os.File, counts map[string]int, temp map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() && input.Text() != "exit" {
		counts[input.Text()]++
		temp[input.Text()]++
	}
}

func echo1() {
	var s string
	var sep string
	for _, v := range os.Args[1:] {
		s += sep + v
		sep = " "
	}
	fmt.Println(s)
}

func echo2() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
