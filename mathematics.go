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

// combine
// 77. Combinations
// https://leetcode.com/problems/combinations/
func combine(n int, k int) [][]int {
	var result [][]int
	subCombine(&result, []int{}, k, 1, n)
	return result
}
func subCombine(result *[][]int, part []int, k int, from int, to int) {
	for from <= to {
		if k == 1 {
			newElem := make([]int, len(part))
			copy(newElem, part)
			*result = append(*result, append(newElem, from))
		} else {
			subCombine(result, append(part, from), k-1, from+1, to)
		}
		from++
	}
}

// minCostClimbingStairs
// 746. Min Cost Climbing Stairs
// https://leetcode.com/problems/min-cost-climbing-stairs/
func minCostClimbingStairs(cost []int) int {
	length := len(cost)

	maxCost := make([]int, length)
	maxCost[0] = cost[0]
	maxCost[1] = cost[1]

	for i := 2; i < length; i++ {
		maxCost[i] = cost[i] + min(maxCost[i-1], maxCost[i-2])
	}
	return min(maxCost[length-1], maxCost[length-2])
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
