package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		input := bufio.NewScanner(os.Stdin)
		fmt.Println("请输入单位转换的类型和具体数值(类型可选temperature weight length)，以exit结束:")
		var arr []string
		for input.Scan() && input.Text() != "exit" {
			arr = append(arr, input.Text())
		}
		switch arr[0] {
		case "temperature":
			TemperatureConvert(arr[1:])
		case "weight":
			WeightConvert(arr[1:])
		case "length":
			LengthConvert(arr[1:])
		default:
			log.Fatal("不支持该类型的单位转换")
			os.Exit(1)
		}
	} else {
		if len(os.Args[2:]) == 0 {
			log.Fatal("未输入待转换数值")
			os.Exit(2)
		}
		switch os.Args[1] {
		case "temperature":
			TemperatureConvert(os.Args[2:])
		case "weight":
			WeightConvert(os.Args[2:])
		case "length":
			LengthConvert(os.Args[2:])
		default:
			log.Fatal("不支持该类型的单位转换")
			os.Exit(1)
		}
	}
}

func TemperatureConvert(temp []string) {
	CToF(temp)
	FToC(temp)
}

func CToF(cels []string) {
	for _, v := range cels {
		num, err := strconv.ParseFloat(v, 64)
		if err != nil {
			log.Fatal("输入值不是数值类型")
			os.Exit(3)
		}
		fmt.Printf("%s 摄氏度 = %.2f华氏度\n", v, num*9/5+32)
	}
}

func FToC(fahs []string) {
	for _, v := range fahs {
		num, err := strconv.ParseFloat(v, 64)
		if err != nil {
			log.Fatal("输入值不是数值类型")
			os.Exit(3)
		}
		fmt.Printf("%s 华氏度 = %.2f 摄氏度\n", v, (num-32)*5/9)
	}
}

func WeightConvert(weis []string) {

}

func LengthConvert(lens []string) {

}
