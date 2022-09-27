// Package jmisc contains miscellaneous stuff.
//
// If you can't decide where to put a function, place it here
// and later you move it to a proper package.
package jmisc

import (
	"fmt"
)

// Assert mimics assert from other languages.
// It panics if the expression is false.
// If msg is given, it is printed as an error message.
func Assert(expr bool, msg string) {
	if !expr {
		panic(fmt.Sprintf("Assertion error: %s", msg))
	}
}
