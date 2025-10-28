package process

import "fmt"


/*
for init statement; expression; post statement {
	execute statement
}

2. 循环语句
Go 只有 for 循环，没有 while 或 do-while 语法。
简化版的for循环相当于while循环
for expression {
	execute statement
}

当省略所有语句时，for 会创建一个无限循环。
for {
	execute statement
}

使用range遍历数组，切片，通道或集合时，返回索引和值。
同时也可以使用，break、continue、goto、return等控制语句。


*/ 
func ForTutorial() {

	for i := 0; i < 10; i++ {
		// println(i) 一般不建议使用println
		fmt.Println(i)
	}	

	// for省略init statement和post statement
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// for无限循环
	for {
		fmt.Println(i)
		i++
		if i > 10 {
			break
		}
	}

	// for range
	nums := []int{1, 2, 3}
	for i, v := range nums {
		fmt.Println(i, v)
	}
}
