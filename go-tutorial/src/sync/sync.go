package sync

import (
	"fmt"
	"sync"
	"time"
)

// 1.sync 包概述
// sync 包提供了基本的同步原语，如互斥锁、条件变量、等待组等，用于在并发编程中协调多个 Goroutine 的执行。

// 1.1 为什么需要 sync包
// Go 语言的 Goroutine 是并发执行的，当多个 Goroutine 访问共享资源时，可能会导致数据竞争（Data Race）。 sync 包提供了同步机制来避免这些问题。
func DemonstrateDataRace() {
	var counter int
	for i := 0; i < 1000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(time.Second)
	fmt.Printf("最终计算值：%d (期望值: 1000)\n", counter)
}
func DemonstrateMutexSolution() {
	var counter int
	var mutex sync.Mutex

	for i := 0; i < 1000; i++ {
		go func() {
			mutex.Lock()
			defer mutex.Unlock()
			counter++
		}()
	}

	time.Sleep(time.Second)
	fmt.Printf("最终计算值：%d (期望值: 1000)\n", counter)
}

// 2.Mutex 互斥锁
// 2.1 基本用法
// Mutex是互斥锁，提供两个基本方法 Lock() 和 Unlock()
func DemonstrateBasicMutex() {
	var mutex sync.Mutex
	var shareData int

	// write
	mutex.Lock()
	shareData = 100
	mutex.Unlock()

	// Read
	mutex.Lock()
	value := shareData
	mutex.Unlock()
	fmt.Printf("最终计算值：%d (期望值: 100)\n", value)
}

// 2.2 使用 defer 确保解锁
func DemonstrateMutexWithDefer() {
	var mutext sync.Mutex
	var shareData []string

	func() {
		mutext.Lock()
		defer mutext.Unlock()
		shareData = append(shareData, "hello")
		fmt.Printf("共享数据: %v\n", shareData)
	}()

	// 即使发生 panic，也会解锁
	func() {
		mutext.Lock()
		defer mutext.Unlock()

		shareData = append(shareData, "item3")
		// 模拟可能发生 panic 的操作
		if len(shareData) > 2 {
			fmt.Println("模拟 panic")
			// panic("something went wrong") // 即使 panic，defer 也会执行
		}
	}()
}

// 2.3 实际应用：计数器
type SafeCounter struct {
	mutext sync.Mutex
	value  int
}

// 注意书写格式
func (c *SafeCounter) Increment() {
	c.mutext.Lock()
	defer c.mutext.Unlock()
	c.value++
}
func (c *SafeCounter) Value() int {
	c.mutext.Lock()
	defer c.mutext.Unlock()
	return c.value
}
func (c *SafeCounter) Reset() {
	c.mutext.Lock()
	defer c.mutext.Unlock()
	c.value = 0
}
func DemonstrateSafeCounter() {
	counter := &SafeCounter{}
	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				counter.Increment()
			}
		}()
	}

	time.Sleep(time.Second)
	fmt.Printf("最终计数值：%d (期望值: 10000)\n", counter.Value())
}

// 3. RWMutex(读写互斥锁)
// 3.1 读写锁的基本使用
// RWMutex 是读写互斥锁，它允许多个读操作同时进行，但写操作需要互斥。
func DemonstrateRWMutex() {
	var rwMutext sync.RWMutex
	var shareData map[string]int = make(map[string]int)

	// write
	rwMutext.Lock()
	shareData["key1"] = 100
	rwMutext.Unlock()

	// read
	rwMutext.RLock()
	value := shareData["key1"]
	rwMutext.RUnlock()
	fmt.Printf("最终计算值：%d (期望值: 100)\n", value)
}

type Cache struct {
	rwMutext sync.RWMutex
	data     map[string]interface{}
	stats    CacheStats
}
type CacheStats struct {
	Hits   int
	Misses int
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]interface{}),
	}
}
func (c *Cache) Get(key string) (interface{}, bool) {
	c.rwMutext.RLock()
	defer c.rwMutext.RUnlock()
	value, ok := c.data[key]
	if ok {
		c.rwMutext.Lock()
		c.stats.Hits++
		c.rwMutext.Unlock()
		return value, true
	}

	c.rwMutext.Lock()
	c.stats.Misses++
	c.rwMutext.Unlock()
	return nil, false
}
func (c *Cache) Set(key string, value interface{}) {
	c.rwMutext.Lock()
	defer c.rwMutext.Unlock()
	c.data[key] = value
}
func (c *Cache) Stats() CacheStats {
	c.rwMutext.RLock()
	defer c.rwMutext.RUnlock()
	return c.stats
}

func DemonstrateCache() {
	cache := NewCache()
	cache.Set("user:1", "Alice")
	cache.Set("user:2", "Bob")

	// 并发读取
	for i := 0; i < 10; i++ {
		go func(id int) {
			if value, exists := cache.Get(fmt.Sprintf("user:%d", id%2+1)); exists {
				fmt.Printf("Goroutine %d 获取到值: %v\n", id, value)
			} else {
				fmt.Printf("Goroutine %d 缓存未命中\n", id)
			}
		}(i)
	}

	time.Sleep(1 * time.Second)
	stats := cache.Stats()
	fmt.Printf("缓存统计: 命中 %d, 未命中 %d\n", stats.Hits, stats.Misses)
}

// 4.WaitGroup 等待组
// 4.1 基本用法
// WaitGroup 用于等待一组 Goroutinue完成
func DemonstrateBasicWaitGroup() {
	var wg sync.WaitGroup

	// 启动多个 Goroutine
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d 开始工作\n", id)
			time.Sleep(time.Duration(id) * 100 * time.Millisecond)
			fmt.Printf("Goroutine %d 完成工作\n", id)
		}(i)
	}

	// 等待所有 goroutine 完成
	wg.Wait()
	fmt.Println("所有 goroutine 已完成")
}

// 5.Once（一次性操作)
// 5.1 基本用法
// Once 确保某个函数只执行一次，即使多个 Goroutine 同时 diaoyong
func DemonstrateOnce() {
	var once sync.Once
	var initialized bool

	initialize := func() {
		fmt.Println("执行初始化操作")
		initialized = true
		time.Sleep(1 * time.Second) // 模拟耗时操作
	}

	// 启动多个 Goroutine 同时尝试初始化
	for i := 0; i < 10; i++ {
		go func(id int) {
			fmt.Printf("Goroutine %d 尝试初始化\n", id)
			once.Do(initialize)
			fmt.Printf("Goroutine %d 初始化完成\n", id)
		}(i)
	}

	time.Sleep(2 * time.Second)
	fmt.Printf("初始化状态: %t\n", initialized)
}
// 5.2 实际应用：单例模式
type Singleton struct {
	data string
}
var (
	// 这是懒汉式单例模式，也是最推荐的单例模式
	instances *Singleton
	once sync.Once
)
func GetInstance() *Singleton {
	once.Do(func(){
		instances = &Singleton{
			data: "单例数据",
		}
		fmt.Println("创建单例示例")
	})
	return instances
}
func DemonstrateSingleton() {
	for i := 0; i < 5; i++ {
		go func(id  int) {
            singleton := GetInstance()
            fmt.Printf("Goroutine %d 获取单例: %p\n", id, singleton)
		}(i)
	}
}
// 6.Cond(条件变量)
// 6.1 基本用法
// Cond 用于等待特定条件的发生，它提供了一组用于等待和通知的原子操作。
func DemonstrateCond() {
	var mutex sync.Mutex
	cond := sync.NewCond(&mutex)
	var ready bool 

	// 等待 goroutine
	go func() {
		cond.L.Lock()
		defer cond.L.Unlock()

		fmt.Println("等待条件满足...")
		for !ready {
			cond.Wait() // 等待条件满足
		}
		fmt.Println("条件满足，开始执行...")
	}()

	// 通知 gorutinue
	time.Sleep(time.Second)
	cond.L.Lock()
	ready = true
	cond.Signal() 
	cond.L.Unlock()
	
	// 等待一段时间确保所有 goroutine 完成
	time.Sleep(2 * time.Second)

}

// 7.Pool 对象池
// 7.1 基本用法
// Pool 用于缓存和重用对象，减少 GC 压力
func DemonstratePool() {
	pool := sync.Pool{
		New: func() interface{} {
			return make([]byte, 1024)
		},
	}

	// 获取对象
	buf1 := pool.Get().([]byte)
	fmt.Printf(" 获取到缓冲区: %p\n", buf1)

	// 使用对象
	copy(buf1, "hello, word")
	fmt.Printf(" 使用缓冲区: %s\n", buf1)

	// 归还对象
	pool.Put(buf1)
	fmt.Printf(" 归还缓冲区: %p\n", buf1)

	 // 再次获取（可能会重用之前归还的对象）
    buf2 := pool.Get().([]byte)
    fmt.Printf("再次获取缓冲区: %p\n", buf2)
    fmt.Printf("缓冲区内容: %s\n", string(buf2))
}

// 8.Map 并发安全的 map
// 8.1 基本用法
// sync.Map 是并发安全的映射，适用于读多写少的场景
func DemonstrateMap() {
	var m sync.Map

	// 存储值
	m.Store("key1", "value1")
	m.Store("key2", 42)
	m.Store("key3", []int{1, 2, 3})

	// 读取值
	if value, ok := m.Load("key1"); ok {
		fmt.Printf("key1 的值 %v\n", value)
	}

	// 读取或存储（不存在）
	value, loaded := m.LoadOrStore("key2", 100)
    fmt.Printf("key2 的值: %v, 是否已存在: %t\n", value, loaded)

	// 删除键
	m.Delete("key3")

	// 遍历
	m.Range(func(key, value interface{}) bool {
        fmt.Printf("键: %v, 值: %v\n", key, value)
		return true
	})
}

// 9. 实际开发中需要注意几个细节：
// 9.1 避免死锁
// 9.2 减少颗粒度（加锁的内容越小越好）
// 9.3 使用 defer 释放锁
