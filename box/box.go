package box

import "time"

func Fib(a int) int {
	time.Sleep(time.Millisecond * 10)
	if a < 2 {
		return a
	}
	return Fib(a-1) + Fib(a-2)
}
