// Package diffsquares calculates the square or sum, sum of squares and diff between the two.
package diffsquares

// SquareOfSum calculates the square of the sum of natural number up to and including n.
func SquareOfSum(n int) int {
	sum := n * (n + 1) / 2

	return sum * sum
}

// SumOfSquares calculates the sum of the squares of natural number up to and including n.
func SumOfSquares(n int) int {
	sum := n * (n + 1) * (2*n + 1) / 6

	return sum
}

// Difference calculates the diff between square of sum and sum of squares.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
