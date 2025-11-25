package basicdata

import (
	"fmt"
	"strconv"
)

func TypeConversionTutorial() { 
	intToIntExample()
	intToFloatExample()
	stringToIntExample()
}

func intToIntExample() {
	var a int8 = 20
	var b int16 = int16(a) // int8 转 int16

	fmt.Printf("a = %d, b = %d\n", a, b)
}

func intToFloatExample() {
	var a int8 = 20
	var b float32 = float32(a)
	fmt.Printf("a = %d, b = %f\n", a, b)
}

func stringToIntExample() {
	var str string = "123"
	num, err := strconv.Atoi(str)
	fmt.Printf("str = %s, num = %d, err = %v\n", str, num, err)

	var str2 string = "123.45"
	num3, err2 := strconv.ParseFloat(str2, 64)
	fmt.Printf("str2 = %s, num3 = %f, err2 = %v\n", str2, num3, err2)
}

