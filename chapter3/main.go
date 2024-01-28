package main

import (
	"bytes"
	"fmt"
	"strings"
)

type WeekDay int

const (
	Sunday WeekDay = iota
	Monday
	Tuesday
)

type Flags uint

const (
	FlagUp Flags = 1 << iota
	FlagBroadcast
	FlagLoopBack
	FlagPointToPoint
	FlagMulticast
)

const (
	KB = 1000
	MB = 1000 * KB
	GB = 1000 * MB
	TB = 1000 * GB
	PB = 1000 * TB
	EB = 1000 * PB
	ZB = 1000 * EB
	YB = 1000 * ZB
)

func main() {
	var f float64 = 212
	a := 5.0 / 9.0 * (f - 32)
	b := 5.0 * (f - 32) / 9.0
	fmt.Printf("%T %[1]v\n", f)
	fmt.Printf("%T %[1]v\n", a)
	fmt.Printf("%T %[1]v\n", b)
}

func comma(s string) string {
	startIdx := 0
	endIdx := len(s) - 1
	if strings.HasPrefix(s, "+") || strings.HasPrefix(s, "-") {
		startIdx = 1
	}
	if strings.LastIndex(s, ".") != -1 {
		endIdx = strings.LastIndex(s, ".") - 1
	}
	buf := bytes.Buffer{}
	for i := endIdx; i >= startIdx; i-- {
		buf.WriteByte(s[i])
		if i > startIdx && (endIdx+1-i)%3 == 0 {
			buf.WriteString(",")
		}
	}
	tmp := buf.String()
	ans := bytes.Buffer{}
	for i := len(tmp) - 1; i >= 0; i-- {
		ans.WriteByte(tmp[i])
	}
	var prefix, suffix string
	if startIdx != 0 {
		prefix = s[:startIdx]
	}
	if endIdx != len(s)-1 {
		suffix = s[endIdx+1:]
	}
	return prefix + ans.String() + suffix
}

func isShuffle(a, b string) bool {
	diff := [1000]int{}
	for i := range a {
		diff[a[i]]++
	}
	for i := range b {
		diff[b[i]]--
	}
	for _, v := range diff {
		if v != 0 {
			return false
		}
	}
	return true
}
