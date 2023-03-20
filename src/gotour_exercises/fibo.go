package main

import "fmt"

func fibonacci() func(i int) int {
	fibo := 0

	return func(x int) int {
		if x <= 1 {
			fibo = x
		} else {
			fibo = fibonacci()(x-1) + fibonacci()(x-2)
		}
		return fibo
	}

}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
