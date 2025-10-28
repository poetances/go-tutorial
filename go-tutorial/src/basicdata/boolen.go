package basicdata

import "fmt"

// boolen类型只有两个值 true 和 false，不能像C语言那样用0和非0表示
func BoolenTutorial() {

	var a bool = true
	b := false
	fmt.Println("A is", a, "B is", b)
}