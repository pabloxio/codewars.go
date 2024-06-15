// https://www.codewars.com/kata/5541f58a944b85ce6d00006a/train/go
package kata

// fibs stores the known fibonacci numbers
var fibs = map[uint64]uint64{0: 0, 1: 1}

func fibonacci(n uint64) uint64 {
	f, ok := fibs[n]
	if ok {
		return f
	}

	fibs[n-1] = fibonacci(n - 1)
	fibs[n-2] = fibonacci(n - 2)

	return fibs[n-1] + fibs[n-2]
}

func ProductFib(prod uint64) [3]uint64 {
	var i, fib, fibPrev, fibNext uint64

	for ; ; i++ {
		fib = fibonacci(i)
		fibNext = fibonacci(i + 1)
		if fib*fibNext == prod {
			return [3]uint64{fib, fibNext, uint64(1)}
		}

		if fib*fibNext > prod {
			fibPrev = fib
			fib = fibNext
			break
		}
	}

	return [3]uint64{fibPrev, fib, uint64(0)}
}
