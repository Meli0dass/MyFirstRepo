package main

import "fmt"

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var num uint64 = 7120531
	fmt.Printf("%b\n", num)
	fmt.Println(PopCount1(num))
	fmt.Println(PopCount2(num))
}

func PopCount1(x uint64) int {
	ans := 0
	for x != 0 {
		if x&1 == 1 {
			ans++
		}
		x >>= 1
	}
	return ans
}

func PopCount2(x uint64) int {
	ans := 0
	for ; x != 0; x = x & (x - 1) {
		ans++
	}
	return ans
}
