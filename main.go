package main

import (
	"alap/jio"
	"alap/jmisc"
	"fmt"
	"log"
	"os"
	"reflect"
	"sort"
	"strings"
)

const VERSION = "0.1.0"

// language: source template
var lang2source = map[string]string{
	"c":    CTemplate,
	"go":   GoTemplate,
	"java": JavaTemplate,
	"py":   PythonTemplate,
	"rust": RustTemplate,
}

// language: file to create
var lang2file = map[string]string{
	"c":    "main.c",
	"go":   "main.go",
	"java": "Main.java",
	"py":   "main.py",
	"rust": "main.rs",
}

// Check if lang2source and lang2file contain the same keys.
// Error if there's a difference.
func checkIfMapsAreInSync() {
	var a, b []string
	for key := range lang2source {
		a = append(a, key)
	}
	for key := range lang2file {
		b = append(b, key)
	}
	sort.Strings(a)
	sort.Strings(b)
	jmisc.Assert(reflect.DeepEqual(a, b), "the two maps are not in sync")
}

// help about the usage of the program
func printHelp() {
	fmt.Println("Usage: alap <language>")
	fmt.Println()
	fmt.Println("Available language templates:")
	fmt.Println()
	var keys []string
	for key := range lang2file {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		fmt.Println("*", key)
	}
}

// Create file and save source code in it.
// Print an error message if the file exists.
func writeSourceToFile(source, fname string) {
	if jio.Exists(fname) {
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
}

// entry point
func main() {
	checkIfMapsAreInSync() // may panic

	args := os.Args[1:]
	if len(args) == 0 {
		printHelp()
		os.Exit(0)
	}
	// else
	key := args[0]
	if key == "-h" || key == "--help" {
		printHelp()
		os.Exit(0)
	}
	if key == "-v" || key == "--version" {
		fmt.Printf("alap v%s\n", VERSION)
		os.Exit(0)
	}
	fname, exists := lang2file[key]
	if !exists {
		fmt.Printf("Error: the language '%s' was not found\n", key)
		os.Exit(1)
	}
	// else
	source := lang2source[key]
	source = strings.TrimSpace(source)
	writeSourceToFile(source, fname)
	fmt.Printf("# %s was created\n", fname)
}
