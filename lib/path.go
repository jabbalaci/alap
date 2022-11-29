// file operations (exists, isfile, etc.)
package lib

import (
	"io/fs"
	"os"
)

// Returns true if the entry (file, directory, symbolic link, etc.) exists.
// Returns false if the entry doesn't exist.
func Exists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	// else, doesn't exist
	return false
}

// Is it a directory?
// Returns true for a directory and for a symbolic link pointing on a directory.
func IsDir(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}

	return fileInfo.IsDir()
}

// Is it a regular file?
// Returns true for a regular file and for a symbolic link pointing on a regular file.
func IsFile(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}

	return fileInfo.Mode().IsRegular()
}

// Is it a symbolic link?
func IsLink(path string) bool {
	// can handle symbolic link, but will not follow the link
	fileInfo, err := os.Lstat(path)
	if err != nil {
		return false
	}
	return fileInfo.Mode()&os.ModeSymlink == os.ModeSymlink
}

// Make a file executable for the owner. Equivalent of "chmod u+x <file>" .
// You can also call it on Windows. It won't do anything, but you won't get an error either.
func MakeExecutable(fname string) error {
	stats, err := os.Stat(fname)
	if err != nil {
		return err
	}
	mode := stats.Mode()
	executable := uint32(mode) | (1 << 6) // bit for "executable for user"
	err = os.Chmod(fname, fs.FileMode(executable))
	if err != nil {
		return err
	}
	return nil
}
