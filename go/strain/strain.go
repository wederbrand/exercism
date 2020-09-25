// Package strain has been documented, did you think I would forget?
package strain

// Ints all
type Ints []int

// Lists types
type Lists [][]int

// Strings are documented
type Strings []string

// Keep keeps all the elements that matches the func
func (mySelf Ints) Keep(theFunc func(int) bool) Ints {
	var result Ints

	for _, value := range mySelf {
		if theFunc(value) {
			result = append(result, value)
		}
	}

	return result
}

// Discard discards all the elements that matches the func
func (mySelf Ints) Discard(theFunc func(int) bool) Ints {
	var result Ints

	for _, value := range mySelf {
		if !theFunc(value) {
			result = append(result, value)
		}
	}

	return result
}

// Keep keeps all the elements that matches the func
func (mySelf Lists) Keep(theFunc func([]int) bool) Lists {
	var result Lists

	for _, value := range mySelf {
		if theFunc(value) {
			result = append(result, value)
		}
	}

	return result
}

// Keep keeps all the elements that matches the func
func (mySelf Strings) Keep(theFunc func(string) bool) Strings {
	var result Strings

	for _, value := range mySelf {
		if theFunc(value) {
			result = append(result, value)
		}
	}

	return result

}
