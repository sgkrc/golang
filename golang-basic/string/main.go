package main

import (
	"fmt"
)

func main() {
	var str string = "Hello, Cosmos"

	fmt.Println(str)
	fmt.Println(len(str))
    fmt.Println(str[0]) // 72 (ASCII value of 'H')

	for index, runeValue := range str {
		fmt.Printf("%d: %c\n", index, runeValue)
	}

	str1 := "Hello, "
	str2 := "World!"
	combined := str1 + str2
	fmt.Println(combined)

	substr := str[7:13]
	fmt.Println(substr)

	// 문자열을 byte 배열로 변환하기
	byteSlice := []byte(str)
	byteSlice[0] = 'h'
	newStr := string(byteSlice)
	fmt.Println(newStr)
}