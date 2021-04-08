package main

import "fmt"

type Cat struct{}
type Dog struct {
	name string
}
type Car struct {
	brand string
}

func (c Cat) Say() string { return "喵喵喵" }
func (d Dog) Say() string { return "汪汪汪" }

type Sayer interface {
	say()
}

type Mover interface {
	move()
}

func (c Cat) say() {
	fmt.Println("喵喵喵~~")
}
func (d Dog) say() {
	fmt.Println("汪汪汪~~")
}

func (d Dog) move() {
	fmt.Printf("%s会动--\n", d.name)
}
func (c Car) move() {
	fmt.Printf("%s速度80迈\n", c.brand)
}

func main() {
	c := Cat{}
	d := Dog{}

	//fmt.Println("猫：",c.Say())
	//fmt.Println("狗：",d.Say())
	var x Sayer
	x = c
	x.say()
	x = d
	d.say()

	var m Mover = Dog{
		name: "旺财",
	}
	var move Mover = Car{
		brand: "奔驰",
	}
	m.move()
	move.move()

	var s interface{}
	s = "baidu.com"
	v, ok := s.(string)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("类型断言失败")
	}
	justifyType(s)

}

//使用switch断言多次判断
func justifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string value is %v", v)
	case int:
		fmt.Printf("x is a int value is %v", v)
	case bool:
		fmt.Printf("x is a bool value is %v", v)
	default:
		fmt.Println("unsupport type")
	}
}
