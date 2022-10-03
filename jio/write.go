package jio

import (
	"fmt"
	"log"
	"os"
)

// Create file and save source code in it.
// Print an error message if the file exists.
func WriteSourceToFile(source, fname string) bool {
	writeOk := false
	if Exists(fname) {
		fmt.Printf("Error: the file '%s' already exists\n", fname)
		os.Exit(1)
	}
	// else
	f, err := os.Create(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	f.WriteString(source)
	f.WriteString("\n")
	writeOk = true

	return writeOk
}
