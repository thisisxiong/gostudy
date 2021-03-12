package main

import "fmt"

func main(){
	//ptr()
	//maps()
	//structs()
	process()
}

func ptr(){
	// 指针
	a := 10
	b := &a //取变量a的地址 将指针保存在b中
	fmt.Printf("a:%d ptr:%p\n",a,&a) //a:10 ptr:0xc00000a0a8
	fmt.Printf("b:%p type:%T\n",b,b) //b:0xc00000a0a8 type:*int
	//指针取值
	fmt.Println(*b) //10
	//空指针
	var p *string
	fmt.Println(p) //nil
	fmt.Printf("p的值是%v\n",p) //p的值是<nil>
}

func maps(){
	//map类型的变量默认初始值为nil make第二个参数为cap容量
	map1 := make(map[string]string)
	map1["name"] = "张三"
	map1["email"] = "zhangsa@imooc.com"
	fmt.Println(map1) //map[email:zhangsa@imooc.com name:张三]
	fmt.Println(map1["email"]) //zhangsa@imooc.com
	fmt.Printf("type:%T\n",map1) //type:map[string]string
	//判断map中键是否存在
	v,ok := map1["name"]
	fmt.Println(v,ok) //张三 true
	val,exist := map1["address"]
	fmt.Println(val,exist) // false
	//使用delete(map,key)函数删除键值对
	delete(map1,"name")
	fmt.Println(map1) //map[email:zhangsa@imooc.com]
}

func structs(){
	//结构体定义
	type person struct{
		name string
		city string
		age int8
	}
	//基本实例化
	var p person
	p.name = "小明"
	p.city = "北京"
	p.age = 18
	fmt.Println(p) //{小明 北京 18}
	//指针类型结构体
	var p2 = new(person)
	p2.name = "测试"
	p2.city = "海淀"
	p2.age = 11
	fmt.Println(p2) //&{测试 海淀 11}
	//取结构体的地址实例化
	p3 := &person{}
	p3.name = "test"
	p3.city = "昌平"
	p3.age = 20
	fmt.Println(p3) //&{test 昌平 20}
	//使用键值对初始化
	p4 := person{
		name:"imooc",
		city:"beijing",
		age:10,
	}
	fmt.Println(p4) //{imooc beijing 10}
	//对结构体指针键值对初始化
	p5 := &person{
		city:"中国",
	}
	fmt.Println(p5) //&{ 中国 0}
	//匿名结构体
	var user struct{
		name string
		age int
	}
	user.name = "小王"
	user.age = 12
	fmt.Println(user) //{小王 12}

}

func process(){
	//初始化语句 布尔表达式
	if x :=3; x>5 {
		fmt.Println("大于5")
	}else if x<5 {
		fmt.Println("小于5")
	}else{
		fmt.Println("等于5")
	}

	a := 5
	b := 1
	if a ==5 {
		if b==1 {
			fmt.Println("a的值为5，b的值为1")
		}
	}

	var score int = 90
	var grade string = "B"
	switch score {
		case 90: grade = "A"
		case 80: grade = "B"
		case 50,60,70: grade = "C"
		default: grade = "D"
	}
	fmt.Println(grade)
	switch { //省略条件表达式可以当if else使用
		case score > 80:
			fmt.Println("优秀\n")
		case score == 80:
			fmt.Println("良好\n")
		case score < 80:
			fmt.Println("一般")
	}

	switch{
		case score >60:
			fmt.Println("及格了。。。")
			fallthrough //默认每个case后带break 使用fallthrough 强制执行后面的case代码
		case score > 80:
			fmt.Println("良好了。。。")
			fallthrough
		case score>100:
			fmt.Println("一鸣惊人了。。。")
			fallthrough
		default:
			fmt.Println("评价完了。。")
	}
	//常见的for循环支持初始化语句
	for i:=1; i<4;i++ {
		fmt.Println(i)
	}
	n :=5
	for n > 1{ //替代while(n >1){}
		fmt.Println(n)
		n--
	}
	//for { //替代while(true){}
	//	fmt.Println(n)
	//}

	s := "abc4"
	//忽略value 支持string/array/slice/map
	for i := range s{
		fmt.Println(i)
	}
	for _,v := range s{
		fmt.Println(v) //97 98 99 52 对应是ascii码
	}

	m := map[string]int{
		"a" :1,
		"b" :2,
	}
	for key,val := range m{ //range会复制对象
		fmt.Println(key,val)
	}




}

