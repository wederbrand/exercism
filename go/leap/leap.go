// Package leap has a package comment that summarizes what it's about.
package leap

// IsLeapYear has a comment documenting it.
func IsLeapYear(x int) bool {
	if x%400 == 0 {
		return true
	}
	if x%100 == 0 {
		return false
	}
	if x%4 == 0 {
		return true
	}
	return false
}
