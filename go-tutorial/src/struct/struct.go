package structpackge

import (
	"encoding/json"
	"fmt"
)

/*
struct是值类型，所有有零值
*/
type Person struct {
	Name string
	Age  int
}

//  1.1 结构体的零值
func  StuctZeroValue() {
	var p Person
	fmt.Printf("name: %s, age: %d\n", p.Name, p.Age)
}

// 1.2 结构体的字面量
func StructLiterals() {
	// 方式 1：按顺序初始化，必须全部初始化，并且顺序要对。一般不推荐
	p1 := Person{"Alice", 20}
	fmt.Printf("name: %s, age: %d\n", p1.Name, p1.Age)
	
	// 方式 2：按字段名初始化，可以不全部初始化
	p2 := Person{Name: "Bob"}
	fmt.Printf("name: %s, age: %d\n", p2.Name, p2.Age)
}

// 2.1 结构体指针
func StructPointer() {
	 // 使用 new() 创建结构体指针
	p1 := new(Person)
	// p1: 0x1400000c090, name: , age: 0。new 创建指针，是会分配内存的，同时有零值
	fmt.Printf("p1: %p, name: %s, age: %d\n", p1, p1.Name, p1.Age) 

	var p11 *Person
	// p11: 0x0， 下面打印会崩溃，因为空指针没法访问其属性。var 创建指针，是不会分配内存的，没有零值
	fmt.Printf("p11: %p, name: %s, age: %d\n", p11, p11.Name, p11.Age) 

	p1.Name = "Charlie"
	p1.Age = 25
	fmt.Printf("p1 name: %s, age: %d\n", p1.Name, p1.Age)

	// 使用 & 创建结构体指针，等价于 new()
	p2 := &Person{}
	p2.Name = "David"
	p2.Age = 30
	fmt.Printf("p2 name: %s, age: %d\n", p2.Name, p2.Age)
}

// 2.3 结构体指针的访问
func StructPointerAccess() {
	p := Person{Name: "Eve", Age: 28}
	ptr := &p

	// 通过指针访问结构体字段，语法糖：可以直接使用点号访问
	ptr.Name = "Alice" // 等价于 (*ptr).Name = "Alice"
	ptr.Age = 30 // 等价于 (*ptr).Age = 30
	
	// 通过指针修改结构体字段，原始结构体也会被修改
	fmt.Printf("p name: %s, age: %d\n", p.Name, p.Age)
}

// 3.1 结构体作为函数参数
func StructAsFunctionParameter() {
	p := Person{Name: "Frank", Age: 35}
	fmt.Println("传值调用前:", p)
	modifyStructByValue(p)
	fmt.Println("传值调用后:", p)

	fmt.Println("传指针调用前:", p)
	modifyStructByPointer(&p)
	fmt.Println("传指针调用后:", p)
}
// 传值调用，修改的是副本
func modifyStructByValue(p Person) {
	p.Name = "Grace"
	p.Age = 40
}
// 传指针调用，修改的是原始结构体
func modifyStructByPointer(p *Person) {
	p.Name = "Hank"
	p.Age = 45
}

// 3.2 结构体作为函数返回值
func StructAsFunctionReturnValue() {
	p1 := createPerson("Alice", 20)
	p2 := createPersonPtr("Bob", 25)	

	fmt.Printf("p1 name: %s, age: %d\n", p1.Name, p1.Age)
	fmt.Printf("p2 name: %s, age: %d\n", p2.Name, p2.Age)
}
func createPerson(name string, age int) Person {
    return Person{Name: name, Age: age} // 不安全：返回局部变量，Go 会将其分配到栈上
}
func createPersonPtr(name string, age int) *Person {
    p := Person{Name: name, Age: age}
    return &p  // 安全：返回局部变量的指针，Go 会将其分配到堆上
}

// 4 结构体方法
// 4.1 值接收者方法
func (p Person) GetName() string {
	return p.Name
}
func (p Person) GetAge() int {
	return p.Age
}
func StructValueReceiverMethod() {
	p := Person{Name: "Ivy", Age: 22}
	name := p.GetName()
	age := p.GetAge()
	fmt.Printf("name: %s, age: %d\n", name, age)

	ptr := &p
	name1 := ptr.GetName() // 语法糖：指针调用值接收者方法（Go 自动解引用）
	age1 := ptr.GetAge() // 等价于 (*ptr).GetAge()
	fmt.Printf("name: %s, age: %d\n", name1, age1)
}

// 5.结构体嵌套和组合
// 5.1 匿名嵌套
type Address struct {
	City string
	Street string
	ZipCode string
}
type Employee struct {
	Name string 
	Age int 
	Address // 匿名嵌套 Address 结构体
}
func DemonstrateAnonymousEmbedding() {
	emp := Employee{
		Name: "John",
		Age: 30,
		Address: Address{
			City: "New York",
			Street: "5th Avenue",
			ZipCode: "10001",
		},
	}

	// 直接访问嵌套结构体的字段（提升字段）
	fmt.Println("城市", emp.City)  // 等价于 emp.Address.City
	fmt.Println("街道", emp.Street) // 等价于 emp.Address.Street

	// 显式访问嵌套结构体的字段
	fmt.Println("邮编", emp.Address.ZipCode)
}

// 5.2 命名嵌套
type PersonWithNamedAddress struct {
	Name string 
	Age int 
	Addr Address // 命名嵌套 Address 结构体
}
func DemonstrateNamedEmbedding() {
	p := PersonWithNamedAddress{
		Name: "David",
		Age: 28, 
		Addr: Address{
			City: "Beijing",
			Street: "Changan Avenue",
			ZipCode: "100000",
		},
	}

	// 必须通过字段名访问嵌套结构体
	fmt.Println("城市", p.Addr.City)  
	fmt.Println("街道", p.Addr.Street) 
	fmt.Println("邮编", p.Addr.ZipCode)
}

// 5.3 方法重写
type Vehicle struct {
	Brand string
	Speed int
}
func (v Vehicle) Move() string {
	return fmt.Sprintf("品牌: %s, 速度: %d km/h", v.Brand, v.Speed)
}

type Car struct {
	Vehicle
	Doors int 
}
// 通过重写方法，实现方法的覆盖。override
func (c Car)Move() string {
	return fmt.Sprintf("品牌: %s, 速度: %d km/h, 门数: %d", c.Brand, c.Speed, c.Doors)
}

// 5.4 多态
type Mover interface {
	Move() string
}
type Walker interface {
	Walk() string
}
type Animal struct {
	Name string
}
func (a Animal)Move() string {
	return fmt.Sprintf("%s 在移动", a.Name)
}
func (a Animal)Walk() string {
	return fmt.Sprintf("%s 在行走", a.Name)
}

type Bird struct {
	Animal
	WingSpan int
}
func (b Bird)Move() string {
	return fmt.Sprintf("%s 在飞行，翼展 %d cm", b.Name, b.WingSpan)
}
func (b Bird)Walk() string {
	return fmt.Sprintf("%s 在行走", b.Name)
}
func DemonstratePolymorphism() {
	animals := []Mover{
		Animal{Name: "狗"},
		Bird{Animal: Animal{Name: "鸟"}, WingSpan: 100},
	}
	
	for _, animal := range animals {
		fmt.Println(animal.Move())
	}
}
// 6.结构体标签
// 6.1 基本标签使用
type User struct {
	ID int 			`json:"id"`
	Name string 	`json:"name"`
	Password string `json:"-"` // 忽略该字段
	Email string 	`json:"email,omitempty"` // 为空时忽略该字段
}
func DemonstrateStructTags() {
	user := User{
		ID: 1,
		Name: "Alice",
		Password: "secret123",
		Email: "",
	}

	jsonData, _ := json.Marshal(user)
	fmt.Println(string(jsonData))
}