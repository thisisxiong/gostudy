package main

import (
	"fmt"
)

func main() {
	//s1 := test(func() int {return 100})
	//fmt.Println(s1) //100
	//s2 := format(func(s string,x,y int) string{
	//	return fmt.Sprintf(s,x,y)
	//},"%d,%d",10,20)
	//fmt.Println(s2) //10.20
	//
	//var a,b int = 1,2
	//swap(&a,&b)
	//fmt.Println(a,b) //2 1
	//
	//sum :=myfunc(1,2,3,4)
	//fmt.Println(sum) //10
	//s := []int{5,6,7,8}
	//res := myfunc(s...)
	//fmt.Println(res) //26
	//add := add(1,5)
	//fmt.Println(add) //6
	//fmt.Println(myfunc(mulit())) //4
	//
	////匿名函数
	//getSqrt := func(a float64) float64{
	//	return math.Sqrt(a)
	//}
	//fmt.Println(getSqrt(9)) //3
	//
	//fmt.Println(factorial(5)) //120
	//
	////defer延迟调用 先进后出
	//var whatever [5]struct{}
	//for i := range whatever {
	//	defer func() {fmt.Println(i)}()
	//}
	//
	//test2()
	//exception()

}

func test(fn func() int) int {
	return fn()
}

//定义函数类型
type FormatFunc func(s string, x, y int) string

func format(fn FormatFunc, s string, x, y int) string {
	return fn(s, x, y)

}

//函数参数 默认是值传递 map。slice、chan、指针、interface 默认是引用方式传递
func swap(x, y *int) {
	var tmp int
	tmp = *x
	*x = *y
	*y = tmp
}

//0个或多个参数
func myfunc(args ...int) int {
	sum := 0
	for _, v := range args {
		sum += v
	}
	return sum
}

//没有参数的 return 语句返回各个返回变量的当前值
func add(a, b int) (sum int) {
	sum = a + b
	return
}

//多返回值可直接作为其他函数调用实参
func mulit() (int, int) {
	return 1, 3
}

//递归阶乘
func factorial(i int) int {
	if i <= 1 {
		return 1
	}
	return i * factorial(i-1)
}

//延迟调用参数在注册时求值或者复制 可用指针或者闭包延迟读取
func test2() {
	x, y := 10, 20
	defer func(i int) {
		fmt.Println("defer:", i, y) //10 120
	}(x)
	x += 10
	y += 100
	fmt.Println("x=", x, " y=", y) //x= 20  y= 120
}

//异常处理
func exception() {
	defer func() {
		//捕获异常
		if err := recover(); err != nil {
			fmt.Println(err.(string))
		}
	}()

	panic("exception") //抛出异常
}
