package main

import (
	"fmt"
)

func main() {
	intVar := 123
	floatVar := 123.456
	strVar := "Hello, World!"
	boolVar := true
	pointerVar := &intVar

	// %v
	// value 기본 출력 - 자료형에 맞게 알아서 출력
	fmt.Printf("Default format: %v\n", intVar)         // 123
	fmt.Printf("Default format: %v\n", floatVar)       // 123.456
	fmt.Printf("Default format: %v\n", strVar)         // Hello, World!
	fmt.Printf("Default format: %v\n", boolVar)        // true
	fmt.Printf("Default format: %v\n", pointerVar)     // 0xc0000a6010

	// %T
	// value의 타입 출력
	fmt.Printf("Type: %T\n", intVar)                   // int
	fmt.Printf("Type: %T\n", floatVar)                 // float64
	fmt.Printf("Type: %T\n", strVar)                   // string
	fmt.Printf("Type: %T\n", boolVar)                  // bool
	fmt.Printf("Type: %T\n", pointerVar)               // *int

	// %x
	// 정수 -> 16진수
	// 문자열 -> 각 바이트를 16 진수로 인코딩
	fmt.Printf("Hexadecimal: %x\n", intVar)            // 7b
	fmt.Printf("Hexadecimal: %x\n", strVar)            // 48656c6c6f2c20576f726c6421

	// %d -> 정수 (십진수)
	fmt.Printf("Integer: %d\n", intVar)                // 123

	// %f, %e, %E
	// f -> 부동소수점 (6)
	// e/E -> 지수 표기법
 	fmt.Printf("Float: %f\n", floatVar)                		// 123.456000
	fmt.Printf("Scientific (lowercase): %e\n", floatVar) 	// 1.234560e+02
	fmt.Printf("Scientific (uppercase): %E\n", floatVar) 	// 1.234560E+02

	// %s -> 문자열
	fmt.Printf("String: %s\n", strVar)                 // Hello, World!

	// %p -> 포인터 주소
	// 항상 0x로 시작하는 16진수 형태
	fmt.Printf("Pointer address: %p\n", pointerVar)    // 0xc0000a6010
}