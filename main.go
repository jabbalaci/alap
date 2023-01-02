package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/jabbalaci/alap/lib"
	"github.com/jabbalaci/alap/special"
	"github.com/jabbalaci/alap/templates"
)

const VERSION = "0.1.5"

type LangInfo struct {
	fname       string
	sourceCode  string
	description string
	executable  bool
}

// template IDs and their associated structs
var langMap = map[string]LangInfo{
	"c":     {fname: "main.c", sourceCode: templates.C, description: "\t\t- C source code"},
	"go":    {fname: "main.go", sourceCode: templates.Go, description: "\t\t- Go source code"},
	"java":  {fname: "Main.java", sourceCode: templates.Java, description: "\t\t- Java source code"},
	"cs":    {fname: "Program.cs", sourceCode: templates.CSharp, description: "\t\t- C# source code"},
	"py":    {fname: "main.py", sourceCode: templates.Python, description: "\t\t- Python 3 source code", executable: true},
	"flask": {fname: "app.py", sourceCode: templates.Flask, description: "\t\t- Flask source code", executable: true},
	"rust":  {fname: "main.rs", sourceCode: templates.Rust, description: "\t\t- Rust source code"},
	"nuon":  {fname: "on", sourceCode: "", description: "\t\t- create `on` for activating a virt. env. from Nushell"},
}

// help about the usage of the program
func printHelp() {
	fmt.Printf("alap v%v, https://github.com/jabbalaci/alap\n", VERSION)
	fmt.Println()
	fmt.Println("Usage: alap <template_id>")
	fmt.Println()
	fmt.Println("Available templates:")
	fmt.Println()
	var keys []string
	for key := range langMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		entry := langMap[key]
		fmt.Printf("* %v%v [%v]\n", key, entry.description, entry.fname)
	}
}

func handleSpecialCase(fname string) {
	if fname == "nuon" {
		special.CreateNuOn()
	}
}

// entry point
func main() {
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
	entry, exists := langMap[key]
	if !exists {
		fmt.Printf("Error: the language '%s' was not found\n", key)
		os.Exit(1)
	}
	// else
	source := strings.TrimSpace(entry.sourceCode)
	writeOk := false
	if len(source) > 0 {
		writeOk = lib.WriteSourceToFile(source, entry.fname)
		if writeOk {
			fmt.Printf("# `%s` was created\n", entry.fname)
		}
	} else {
		handleSpecialCase(key)
	}

	if entry.executable {
		lib.MakeExecutable(entry.fname)
	}
}
