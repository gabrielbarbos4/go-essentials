package internal

import (
	"fmt"
	"unicode/utf8"
)

func DataTypes() {
	var int1 int16 = -32767
	var int2 int32 = -1
	var int3 int64 = -2

	var uint1 uint16 = 32767
	var uint2 uint32 = 32767
	var uint3 uint64 = 32767

	var floatNum1 float32 = 12345678.9
	var floatNum2 float64 = 12345678.9

	fmt.Println("float 32 usage")
	fmt.Printf("%f\n\n", floatNum1)

	fmt.Println("float 64 usage")
	fmt.Printf("%f\n\n", floatNum2)

	fmt.Println(int1)
	fmt.Println(int2)
	fmt.Println(int3)
	fmt.Println(uint1)
	fmt.Println(uint2)
	fmt.Println(uint3)

	var string1 string = "testing String"

	fmt.Println()

	//this code does not print length of the string stead of it prints bytes
	fmt.Println("Print String bytes with len built in 'len' function")
	fmt.Println(len(string1))

	fmt.Println()

	fmt.Println("Print string content")
	fmt.Println(utf8.RuneCountInString(string1))

	fmt.Println()

	//omitting var keyword and datatype and initializing multiple variables
	boolean1, boolean2 := false, true
	fmt.Println("Print string content")
	fmt.Println(boolean1, boolean2)
}
