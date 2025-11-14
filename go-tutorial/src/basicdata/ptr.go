package basicdata

import "fmt"

/*
var var_name *var-type
var ip *int
var fp *float32

当一个指针被定义后没有分配到任何变量时，它的值为 nil。
nil 指针也称为空指针。
nil在概念上和其它语言的null、None、nil、NULL一样，都指代零值或空值。
*/
func PtrTutorial() {

	var a int = 10
    // 指针类型，默认是nil，一般用new分配地址
    var p *int

    p = &a  // 取变量 a 的地址，并赋值给 p
    fmt.Println("变量 a 的值:", a) 
    fmt.Println("变量 a 的地址:", &a)
    fmt.Println("指针 p 存储的地址:", p)
    fmt.Println("指针 p 指向的值:", *p) // 通过指针访问 a 的值

    *p = 20  // 通过指针修改变量 a 的值
    fmt.Println("修改后 a 的值:", a)
}