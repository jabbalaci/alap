// Package jio facilitates I/O operations.
package jio

import "os"

// Returns true if the entry (file, directory, symbolic link) exists.
// Returns false if the entry doesn't exist.
func Exists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	// else, doesn't exist
	return false
}
