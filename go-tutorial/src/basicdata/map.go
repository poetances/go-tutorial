package basicdata

import "fmt"

/*
map使用前必须初始化，不然只声明，默认是nil。
初始化方式有两种：
m1 := make(map[string]int)

m2 := map[string]int{}

// 只声明，不初始化会报错
var m3 map[string]int

// 伪代码：map 的底层结构（实际实现更复杂）
type hmap struct {
    count     int      // 当前元素个数
    flags     uint8    // 状态标志（如是否正在写入）
    B         uint8    // 桶数量的对数（桶数为 2^B）
    hash0     uint32   // 哈希种子
    buckets   unsafe.Pointer // 指向桶数组的指针
    oldbuckets unsafe.Pointer // 扩容时保存旧桶的指针
    nevacuate uintptr  // 扩容时迁移的进度
    extra     *mapextra // 可选字段，用于优化小对象存储
}

map 底层是通过 hash 表实现的，key 是无序的。但是 map 的行为是一种类引用类型，其表现跟 slice 类似。


6. 总结
✅ map 是 Go 内置的数据结构，用于存储键值对，查找快。
✅ 常见初始化方式：make()、字面量 {}，但不能使用 new()。
✅ 增删改查：
	•	m[key] = value 添加/修改
	•	delete(m, key) 删除
	•	val, exists := m[key] 判断 key 是否存在
✅ 遍历：for key, value := range map，但无序，如需排序可先取出 key 排序。
✅ map 是引用类型，赋值时不会拷贝数据，需手动复制。
✅ key 必须是可比较类型，不能是 slice、map、function。
*/
func MapTutorial() {
	// 直接使用make
	m := make(map[string]int)
	m["name"] = 12
	m["age"] = 3
	fmt.Println(m)

	// 使用字面量
	m2 := map[string]int{
		"name": 12,
		"age": 3,
	}
	fmt.Println(m2)

	val, exists := m2["apple"]
	if exists {
		fmt.Println("apple:", val)
	} else {
		fmt.Println("apple 不存在")
	}

	// 删除元素
	delete(m2, "apple")

	for key, value := range m {
		fmt.Println(key, value)
	}
	for key := range m {
		fmt.Println("key:", key)
	}

	//（1）map 是引用类型
	m3 := m2
	m3["age"] = 10 // m3的赋值会影响m2
	
	//（2）map 的 key 只能是可比较类型
	// ✅ 可以作为 key：string、int、bool、float、struct（可比较）
	key, ok := m["age"]
	if ok {
		fmt.Println("age:", key)
	}

	key1 := m3["age"] // 如果 key 不存在，返回的是该类型的零值
	fmt.Println("age:", key1)

}