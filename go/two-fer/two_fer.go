// Package twofer contains utils for the classic "one for ..." game.
package twofer

// ShareWith takes a name as input and returns a "one for ..." string.
func ShareWith(name string) string {
	if len(name) == 0 {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
