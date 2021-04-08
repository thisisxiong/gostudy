package main

import (
	"errors"
	"fmt"
)

/**方法**/
//方法定义 func(receiver type) methodName(参数列表)(返回值列表){}
type User struct {
	Name  string
	Email string
}

//匿名字段
type Manager struct {
	User
}

func (u User) Notify() {
	fmt.Printf("name：%s, email: %s \n", u.Name, u.Email)
}

func (self *User) ToString() string {
	return fmt.Sprintf("User:%p,%v", self, self)
}

func main() {
	u1 := User{"golang", "golang@golang.com"}
	u1.Notify()

	u2 := new(User)
	u2.Name = "php"
	u2.Email = "php@php.cn"
	u2.Notify()

	m := Manager{User{"php", "admin@php.com"}}
	fmt.Println(m.ToString())

	//test01()
	//getCircleArea(-3)
	//test02()
	area, err := getCircleArea2(-3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(area)
	}
}

//系统抛异常
func test01() {
	a := [5]int{1, 2, 3, 4, 5}
	a[1] = 123
	fmt.Println(a)
	index := 10
	a[index] = 10
	fmt.Println(a)
}

//自己抛异常
func getCircleArea(radius float32) (area float32) {
	if radius < 0 {
		panic("半径不能为负")
	}
	return 3.14 * radius * radius
}

//返回异常
func getCircleArea2(radius float32) (area float32, err error) {
	if radius < 0 {
		err = errors.New("半径不能为负")
		return
	}
	area = 3.14 * radius * radius
	return
}

func test02() {
	defer func() {
		if error := recover(); error != nil {
			fmt.Println(error)
		}
	}()
	getCircleArea(-4)
	fmt.Println("这里没有执行了")
}
