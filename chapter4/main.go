package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

type tree struct {
	val         int
	left, right *tree
}

type Point struct {
	X, Y int
}

func main() {
	//point := Point{1, 2}
	//fmt.Printf("%#v\n", point)
	//data, _ := json.Marshal(&point)
	//fmt.Printf("%s\n", data)
	//data, _ = json.MarshalIndent(&point, "", "    ")
	//fmt.Printf("%s\n", data)
	//
	//var positions struct{ X int }
	//if err := json.Unmarshal(data, &positions); err != nil {
	//	log.Fatalln("JSON unmarshalling failed")
	//}
	//fmt.Printf("%#v\n", positions)
	//since := time.Now().Unix()
	//time.Sleep(2 * time.Second)
	//now := time.Now().Unix()
	//fmt.Println(now - since)

	allocateIpToUpstream([]string{"1", "2", "3"}, []string{"a", "b", "c", "d", "e"}, 20)

}

func reverse(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func rotate(nums []int, n int) {
	length := len(nums)
	if length == 0 {
		return
	}
	n %= length
	if n == 0 {
		return
	}
	start := 0
	cnt := 0
	for cnt < length {
		i := start
		pre := nums[i]
		j := (i + length - n) % length
		for j != start {
			tmp := nums[j]
			nums[j] = pre
			pre = tmp
			i = j
			j = (i + length - n) % length
			cnt++
		}
		nums[j] = pre
		cnt++
		start++
	}
}

func removeDuplicateElement(arr []string) []string {
	if len(arr) <= 1 {
		return arr
	}
	i := 1
	j := 1
	for i < len(arr) {
		if arr[i] != arr[i-1] {
			arr[j] = arr[i]
			j++
		}
		i++
	}
	return arr[:j]
}

func removeAdjacentSpace(arr []byte) string {
	if len(arr) == 1 {
		return string(arr)
	}
	i := 1
	j := 1
	for i < len(arr) {
		if !(unicode.IsSpace(rune(arr[i])) && unicode.IsSpace(rune(arr[i-1]))) {
			arr[j] = arr[i]
			j++
		}
		i++
	}
	return string(arr[:j])
}

func wordFreq() {
	count := make(map[string]int)
	fmt.Println("Please input something: ")
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		text := input.Text()
		if text == "exit" {
			break
		}
		count[text]++
	}
	for key, val := range count {
		fmt.Printf("the freq of %s is %d\n", key, val)
	}
}

func allocateIpToUpstream(ips []string, upstreams []string, n int) {
	var ptr, count int
	for i := 0; i < n; i++ {
		if count == len(upstreams) {
			count = 0
			ptr = (ptr + 1) % len(ips)
		}
		fmt.Printf("ip_%s -> server_%s\n", ips[ptr], upstreams[count])
		count++
	}
}
