package leetcode

// fib
// 509. Fibonacci Number
// https://leetcode.com/problems/fibonacci-number/
func fib(n int) int {
	if n < 2 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

// climbStairs
// 70. Climbing Stairs
// https://leetcode.com/problems/climbing-stairs/
func climbStairs(n int) int {
	if n < 4 {
		return n
	}

	i1, i2 := 1, 1

	for i := 2; i <= n; i++ {
		i1, i2 = i2, i1+i2
	}
	return i2
}
