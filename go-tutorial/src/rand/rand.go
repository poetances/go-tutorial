package rand

import (
	crand "crypto/rand"
	"fmt"
	"math/rand"
)

func DemonstrateMathRand() {

	// go 语言中 math/rand 包用于生成伪随机数。它提供了多种方法来生成不同类型的随机数，如整数、浮点数等。
	// 伪随机数，一般使用两种算法：
	// 线性同余生成器（Linear Congruential Generator, LCG）
	// 梅森旋转算法（Mersenne Twister）
	// 如果使用相同的种子，则生成的随机数序列是确定的。

	r1 := rand.New(rand.NewSource(42))
	r2 := rand.New(rand.NewSource(42))

	fmt.Println("Generator 1:")
	for i := 0; i < 5; i++ {
		fmt.Print(r1.Intn(100), " ")
	}

	fmt.Println("\nGenerator 2:")
	for i := 0; i < 5; i++ {
		fmt.Print(r2.Intn(100), " ")
	}
}

func DemonostrateCryptoRand() {
	// crypto/rand 包用于生成加密安全的随机数。它使用操作系统提供的随机数生成器，适用于需要高安全性的场景，如密钥生成、令牌生成等。
	// 生成加密安全的随机字节
	b := make([]byte, 10)
	_, err := crand.Read(b)
	if err != nil {
		fmt.Println("Error reading random bytes:", err)
		return
	}
	fmt.Println("Random bytes:", b)
}
