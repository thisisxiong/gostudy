package main

import (
	"fmt"
	"simple/src/utils"
	"strings"
)

func main(){
	//str()
	//arr()
	utils.Count(2,3)


}

func arr(){
	/***数组常用操作***/
	//赋值和传参会复制整个数组
	var arr0 [3]int = [3]int{1,2,3}
	var arr1 = [2]string{"php","go"}
	var arr2 = [...]int{4,5,6,7,8} //通过初始化的值确定长度
	var arr3 = [4]string{1:"hello",3:"world"}
	var arr4 = [2][3]int{{11,12,13},{21,22,23}}
	arr5 := []int{1,2,3,8:88}
	fmt.Println(arr0,arr1,arr2,arr3,arr4,arr5)
	//数组长度
	fmt.Println(len(arr3),cap(arr3)) //4 4
	fmt.Println(len(arr5),cap(arr5)) //9 9
	//数组循环
	for key,val := range arr1{
		//0 php
		//1 go
		fmt.Println(key,val)
	}
	//切片
	fmt.Println(arr0[1:]) // [2 3]
	fmt.Println(arr1[:]) // [php go]
	fmt.Println(arr2[:2]) // [4 5]
	var slice1 = make([]int,2)
	fmt.Println(slice1) // [0 0]
	var slice2 = make([]int,2,4)
	fmt.Println(slice2) // [0 0]
	fmt.Println(len(slice2),cap(slice2)) //2 4
	p := &arr5[2]
	*p += 100
	fmt.Println(arr5) //[1 2 103 0 0 0 0 0 88]
	//切片追加
	arr6 := []int{81,82}
	// ...被打散传入相当于 c := append(arr5,81,82)
	c := append(arr5,arr6...)
	fmt.Println(c)
	data := [...]int{0,1,2,3,4,5,10:0}
	s := data[:2:3]
	fmt.Println(s,len(s),cap(s)) //[0 1] 2 3
	d :=append(s,11,22,33,44)
	fmt.Println(d,len(d),cap(d)) //[0 1 11 22 33 44] 6 6



}

func str(){
	/***字符串常用操作***/
	var title string = "hello go"
	content := "好好学习"
	//字符串长度
	fmt.Println(len(title)) //8
	fmt.Println(len(content)) //12
	//字符串拼接
	fmt.Println(title + ":" + content) //hello go:好好学习
	fmt.Println(fmt.Sprintf("%s:%s",title, content)) //hello go:好好学习
	//字符串分割
	fmt.Println(strings.Split(title," ")) //[hello go]
	//是否包含
	fmt.Println(strings.Contains(content,"学")) //true
	//前缀/后缀判断
	fmt.Println(strings.HasPrefix(title,"ha")) //false
	fmt.Println(strings.HasSuffix(title,"go")) //true
	//子串出现的位置
	fmt.Println(strings.Index(title,"o")) //4
	fmt.Println(strings.LastIndex(title,"o")) //7
	fmt.Println(strings.Index(content,"学")) //6
}




