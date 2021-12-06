// Provide utitlity functions
package utils

// Check if given integer an odd number.
// This func utilises bit operator '&'.
// Because all odd numbers in binary end
// with 1, while even numbers end with 0,
// hence an '&1' operation will result in
// true for odd number and false for even number.
func IsOddNumber(n int) bool {
	return n&1 == 1
}
