// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package twofer should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package twofer

import (
	"fmt"
)

// ShareWith should have a comment documenting it.
// 	// Write some code here to pass the test suite.
// 	// Then remove all the stock comments.
// 	// They're here to help you get started but they only clutter a finished solution.
// 	// If you leave them in, reviewers may protest!

// 	// if name is given:
// 	// 	print one for "name" one for me
// 	// else:
// 	// 	print One for you one for me

// 	// return ""
// }

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
	// comment
	if len(name) == 0 {
		name = "you"
	}
	// return "One for " + name + ", one for me."
	return fmt.Sprintf("One for %s, one for me.", name)
}
