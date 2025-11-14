package process

import "fmt"

/*
✅ 条件判断
	•	if-else：支持变量初始化，作用域受限
	•	switch：可用于多种匹配，默认不 fallthrough
✅ 循环
	•	for：唯一的循环结构
	•	支持 range 遍历数组、切片、map
✅ 跳转
	•	break：终止循环
	•	continue：跳过本次循环
	•	goto：不推荐
✅ defer
	•	延迟执行，常用于资源清理
	•	多个 defer 按 LIFO 顺序执行
*/
func GotoTutorial() {
	fmt.Println("开始")
	goto End
	fmt.Println("这行不会执行")
	End:
	fmt.Println("结束")
}