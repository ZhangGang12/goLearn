package main

import "fmt"

func main() {
	f(3)
}
func f(x int) {
	defer func() {
		if p := recover(); p != nil { // recover() 捕获异常
			fmt.Printf("internal error: %v\n", p)
		}
	}()
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}
