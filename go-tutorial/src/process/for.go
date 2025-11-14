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
	for i := 0; i < 5; i++ {
		fmt.Println("i =", i)
	}

	// 简化版的for循环相当于while循环
	i := 0
	for i < 5 {
		fmt.Println("i =", i)
		i++
	}

	// 无限循环 类似 while(true)
	j := 0
	for {
		fmt.Println("j =", j)
		j++
		if j > 5 {
			break
		}
	}

	// 使用range遍历数组
	arr := []int{10, 20, 30, 40, 50}
	arr = append(arr, 60)
	for index, value := range arr {
		fmt.Printf("arr[%d] = %d\n", index, value)
	}

}
