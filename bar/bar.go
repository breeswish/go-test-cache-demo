package bar

import "time"

func Fib(a int) int {
	time.Sleep(time.Millisecond * 99)
	if a < 2 {
		return a
	}
	return Fib(a-1) + Fib(a-2)
}
