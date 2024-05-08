package main

import "fmt"

func plus(a, b int) int {
	return a + b
}

func multiReturn() (int, int) {
	return 2, 3
}

//variadic functions
func variadicSum(nums ...int) {
	fmt.Print(nums, " ")
	sum := 0

	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)
}

//closures and anonymous functions
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

//recursion
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2=", res)

	res1, res2 := multiReturn()
	fmt.Println(res1, res2)

	nums := []int{1, 2, 3, 4}
	variadicSum(nums...)

	nextInt := intSeq()
	fmt.Print(nextInt(), " ")
	fmt.Print(nextInt(), " ")
	fmt.Println(nextInt(), " ")

	nextInts := intSeq()
	fmt.Println(nextInts())

	fmt.Println(nextInt())

	fmt.Println(factorial(4))

	fmt.Println(fibonacci(6))

}
