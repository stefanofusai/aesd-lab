package main

import "fmt"

var counter_f_rec int
var counter_f_riter int

func f_rec(n int) uint64 {
	counter_f_rec++

	if n == 1 || n == 2 {
		return 1
	}

	return f_rec(n-1) + f_rec(n-2)
}

func f_iter1(n int) uint64 {
	var f uint64
	var f1, f2 uint64 = 1, 1

	if n == 1 || n == 2 {
		return 1
	}

	for n >= 3 {
		n--
		f = f1 + f2
		f1 = f2
		f2 = f
	}

	return f
}

func f_iter2(n int) uint64 {
	var f uint64
	var f1, f2 uint64 = 1, 1

	if n == 1 || n == 2 {
		return 1
	}

	for i := 3; i <= n; i++ {
		f = f1 + f2
		f1 = f2
		f2 = f
	}

	return f
}

func f_riter(a uint64, b uint64, n int) uint64 {
	counter_f_riter++

	if n == 2 {
		return a
	}

	if n == 1 {
		return b
	}

	return f_riter(a+b, a, n-1)
}

func main() {
	fmt.Println(f_rec(30))
	fmt.Printf("Number of calls: %d\n", counter_f_rec)
	fmt.Println(f_iter1(100))
	fmt.Println(f_iter2(100))
	fmt.Println(f_riter(1, 1, 100))
	fmt.Printf("Number of calls: %d\n", counter_f_riter)
}
