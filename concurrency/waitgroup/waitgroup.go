package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Printf("CPUs\t\t%d", runtime.NumCPU)
	fmt.Printf("Go routines\t%d", runtime.NumGoroutine)

	wg.Add(1)
	go foo()
	bar()

	fmt.Println("CPUs\t\t", runtime.NumCPU)
	fmt.Println("Go routines\t", runtime.NumGoroutine)
	wg.Done()

}

func foo() {
	for i := 0; i <= 5; i++ {
		fmt.Println("foo: ", i)
		wg.Done()
	}
}

func bar() {
	for i := 0; i <= 5; i++ {
		fmt.Println("bar: ", i)
	}
}
