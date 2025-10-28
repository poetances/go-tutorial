package variable

import "testing"

/*
Go测试规范：
测试文件名：以_test.go结尾
测试函数名：以Test/Benchmark/Example/Fuzz开头，函数签名固定为func (t *testing.T)
测试参数：单元测试接收 *testing.T，性能测试接收 *testing.B，模糊测试接收 *testing.F
测试文件位置：与被测试代码放在同一包下
测试报名称：与被测试代码包名相同
*/

/// 单元测试，用于验证代码逻辑是否正确
func TestVariable(t *testing.T) {
	VarTutorial()
	VarTutorial2()
}


/// 性能测试，用于测试代码性能，通常用于对比不同实现的性能差异
func BenchmarkVariable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		VarTutorial()
	}
}

func ExampleVarTutorial() {
	VarTutorial()
	// Output:
	// 0 12 true (5+5i) 255 [1 2 0]
}

/// 模糊测试，用于发现代码中的边界情况和潜在错误
func FuzzVariable(f *testing.F) {
	f.Add(1)
	f.Add(2)
	f.Fuzz(func(t *testing.T, a int) {
		if a < 0 {
			t.Errorf("a should be non-negative")
		}
	})
}