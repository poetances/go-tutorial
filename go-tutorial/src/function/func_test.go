package function

import (
	"fmt"
	"testing"
)

func TestFunctionTutorial(t *testing.T) {
}

// TestBasicFunctions 测试基本函数
func TestBasicFunctions(t *testing.T) {
	// 测试 add 函数
	result := add(10, 20)
	if result != 30 {
		t.Errorf("add(10, 20) = %d; 期望 30", result)
	}
	
	// 测试 addThree 函数
	result2 := addThree(1, 2, 3)
	if result2 != 6 {
		t.Errorf("addThree(1, 2, 3) = %d; 期望 6", result2)
	}
}

// TestMultipleReturnValues 测试多返回值函数
func TestMultipleReturnValues(t *testing.T) {
	// 测试 divide 函数
	quotient, remainder := divide(10, 3)
	if quotient != 3 || remainder != 1 {
		t.Errorf("divide(10, 3) = %d, %d; 期望 3, 1", quotient, remainder)
	}
}

// TestNamedReturnValues 测试命名返回值
func TestNamedReturnValues(t *testing.T) {
	sum, diff := calc(20, 10)
	if sum != 30 || diff != 10 {
		t.Errorf("calc(20, 10) = %d, %d; 期望 30, 10", sum, diff)
	}
	
	area, perimeter := rectangle(5.0, 3.0)
	if area != 15.0 || perimeter != 16.0 {
		t.Errorf("rectangle(5.0, 3.0) = %.1f, %.1f; 期望 15.0, 16.0", area, perimeter)
	}
}

// TestVariadicFunctions 测试可变参数函数
func TestVariadicFunctions(t *testing.T) {
	// 测试 sum 函数
	result := sum(1, 2, 3, 4, 5)
	if result != 15 {
		t.Errorf("sum(1, 2, 3, 4, 5) = %d; 期望 15", result)
	}
	
	// 测试切片展开
	numbers := []int{10, 20, 30}
	result2 := sum(numbers...)
	if result2 != 60 {
		t.Errorf("sum([]int{10, 20, 30}...) = %d; 期望 60", result2)
	}
	
	// 测试混合参数
	joined := joinWithSeparator(", ", "Go", "Python", "Java")
	expected := "Go, Python, Java"
	if joined != expected {
		t.Errorf("joinWithSeparator = %s; 期望 %s", joined, expected)
	}
}

// TestFirstClassFunctions 测试一等公民函数
func TestFirstClassFunctions(t *testing.T) {
	// 测试函数作为参数
	addFunc := func(a, b int) int { return a + b }
	result := calculate(10, 5, addFunc)
	if result != 15 {
		t.Errorf("calculate(10, 5, add) = %d; 期望 15", result)
	}
	
	// 测试函数作为返回值
	incrementer := createIncrementer(5)
	result2 := incrementer(10)
	if result2 != 15 {
		t.Errorf("incrementer(10) = %d; 期望 15", result2)
	}
}

// TestClosures 测试闭包
func TestClosures(t *testing.T) {
	counter := createCounter()
	
	// 测试闭包状态保持
	if counter() != 1 {
		t.Errorf("第一次调用 counter() = %d; 期望 1", counter())
	}
	if counter() != 2 {
		t.Errorf("第二次调用 counter() = %d; 期望 2", counter())
	}
	if counter() != 3 {
		t.Errorf("第三次调用 counter() = %d; 期望 3", counter())
	}
}

// TestDefer 测试 defer
func TestDefer(t *testing.T) {
	// 测试 defer 修改返回值
	result := deferWithReturn()
	fmt.Println(result)
}

// TestErrorHandling 测试错误处理
func TestErrorHandling(t *testing.T) {
	// 测试正常情况
	result, err := safeDivide(10, 2)
	if err != nil {
		t.Errorf("safeDivide(10, 2) 不应该出错: %v", err)
	}
	if result != 5 {
		t.Errorf("safeDivide(10, 2) = %d; 期望 5", result)
	}
	
	// 测试错误情况
	_, err2 := safeDivide(10, 0)
	if err2 == nil {
		t.Error("safeDivide(10, 0) 应该出错但没有")
	}
}

// BenchmarkFunctionPerformance 性能测试
func BenchmarkFunctionPerformance(b *testing.B) {
	// 测试函数调用性能
	b.Run("direct_calculation", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = i * i // 直接计算
		}
	})
	
	b.Run("function_call", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = square(i) // 函数调用
		}
	})
	
	b.Run("closure_call", func(b *testing.B) {
		sqr := func(x int) int { return x * x }
		for i := 0; i < b.N; i++ {
			_ = sqr(i) // 闭包调用
		}
	})
}

// square 辅助函数
func square(x int) int {
	return x * x
}

// TestMaximum 测试原有 Maximum 函数
func TestMaximum(t *testing.T) {
	result := Maximum(10, 20)
	if result != 20 {
		t.Errorf("Maximum(10, 20) = %d; 期望 20", result)
	}
	
	result2 := Maximum(30, 15)
	if result2 != 30 {
		t.Errorf("Maximum(30, 15) = %d; 期望 30", result2)
	}
}

// TestCalc 测试原有 Calc 函数
func TestCalc(t *testing.T) {
	sum, sub := Calc(10, 5)
	if sum != 15 || sub != 5 {
		t.Errorf("Calc(10, 5) = %d, %d; 期望 15, 5", sum, sub)
	}
}