package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	//for i := 0; i < 10; i++ {
	//	wg.Add(1)
	//	go hello(i)
	//}
	//wg.Wait()

	//gomain()

	//goruntime()

	//goexit()
	runtime.GOMAXPROCS(4)
	go a()
	go b()
	time.Sleep(time.Second)
}

func a() {
	for i := 0; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 0; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

func goexit() {
	go func() {
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			runtime.Goexit()
			defer fmt.Println("C.defer")
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()
	for {

	}
}

func goruntime() {
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
		}
	}("world")
	for i := 0; i < 2; i++ {
		runtime.Gosched()
		fmt.Println("hello ")
	}
}

func hello(i int) {
	defer wg.Done()
	fmt.Println("Hello Goroutine", i)
}

func gomain() {
	go func() {
		i := 0
		for {
			i++
			fmt.Printf("new goroutine: i= %d \n", i)
			time.Sleep(time.Second)
		}
	}()

	i := 0
	for {
		i++
		fmt.Printf("mian goroutine: i=%d \n", i)
		time.Sleep(time.Second)
		if i == 2 {
			break
		}
	}
}
