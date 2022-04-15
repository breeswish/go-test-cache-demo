package foo

import "time"

func Fib(a int) int {
	time.Sleep(time.Millisecond * 100)
	if a < 2 {
		return a
	}
	return Fib(a-1) + Fib(a-2)
}
