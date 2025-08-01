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

const VERSION = "0.2.8"

const SPECIAL_CASE = "--"

type LangInfo struct {
	fname       string
	sourceCode  string
	description string
	executable  bool
}

// template IDs and their associated structs
var langMap = map[string]LangInfo{
	"c":       {fname: "main.c", sourceCode: templates.C, description: "\t\t- C source code"},
	"cs":      {fname: "Program.cs", sourceCode: templates.CSharp, description: "\t\t- C# source code"},
	"d":       {fname: "main.d", sourceCode: templates.D, description: "\t\t- D source code", executable: true},
	"dub":     {fname: "dub.json", sourceCode: templates.DubJson, description: "\t\t- dub.json for D source code"},
	"flask":   {fname: "app.py", sourceCode: templates.Flask, description: "\t\t- Flask source code", executable: true},
	"go":      {fname: "main.go", sourceCode: templates.Go, description: "\t\t- Go source code"},
	"java":    {fname: "Main.java", sourceCode: templates.Java, description: "\t\t- Java source code"},
	"lorem":   {fname: "lorem.txt", sourceCode: templates.LoremIpsum, description: "\t\t- lorem ipsum text"},
	"lua":     {fname: "main.lua", sourceCode: templates.Lua, description: "\t\t- Lua source code", executable: true},
	"mypy":    {fname: "mypy.ini", sourceCode: templates.Mypy, description: "\t\t- mypy ini file"},
	"nim":     {fname: "main.nim", sourceCode: templates.Nim, description: "\t\t- Nim source code"},
	"pas":     {fname: "main.pas", sourceCode: templates.Pascal, description: "\t\t- Pascal source code"},
	"py":      {fname: "main.py", sourceCode: templates.Python, description: "\t\t- Python 3 source code", executable: true},
	"pymongo": {fname: "mongo.py", sourceCode: templates.Pymongo, description: "\t- MongoDB example (Python 3 + pymongo)", executable: true},
	"rust":    {fname: "main.rs", sourceCode: templates.Rust, description: "\t\t- Rust source code"},
	"sh":      {fname: "main.sh", sourceCode: templates.Bash, description: "\t\t- Bash source code", executable: true},
	// "swift":   {fname: "main.swift", sourceCode: templates.Swift, description: "\t\t- Swift source code"},
	//
	// "nuon": {fname: "on", sourceCode: SPECIAL_CASE, description: "\t\t- prepare a virt. env. for Nushell", executable: true},
}

func verify(d map[string]LangInfo) {
	for k, entry := range d {
		msg := fmt.Sprintf("the source code for '%s' cannot be empty", k)
		lib.Assert(len(entry.sourceCode) > 0, msg)
	}
}

func printMake() {
	fmt.Println("* ---")
	fmt.Println("* make\t\t- create [Makefile] interactively")
}

// help about the usage of the program
func printHelp() {
	fmt.Printf("alap v%v, https://github.com/jabbalaci/alap\n", VERSION)
	fmt.Println()
	fmt.Println("Usage: alap <template_id> [option]")
	fmt.Println()
	fmt.Println("Available options:")
	fmt.Println()
	fmt.Println("-h, --help        show this help")
	fmt.Println("-v, --version     version info")
	fmt.Println("--stdout          don't create source file, print result to stdout")
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
	printMake()
}

func handleSpecialCase(fname string) {
	if fname == "nuon" {
		special.CreateNuOn()
	}
}

func process(key string, to_stdout bool) {
	entry, exists := langMap[key]
	if !exists {
		fmt.Printf("Error: the language/option '%s' is unknown\n", key)
		os.Exit(1)
	}
	// else
	source := strings.TrimSpace(entry.sourceCode)

	// if writing to file:
	if !to_stdout {
		if source == SPECIAL_CASE {
			handleSpecialCase(key)
		} else {
			writeOk := lib.WriteSourceToFile(source, entry.fname)
			if writeOk {
				fmt.Printf("# `%s` was created\n", entry.fname)
			}
		}

		if entry.executable {
			lib.MakeExecutable(entry.fname)
		}
	}

	if to_stdout {
		if source == SPECIAL_CASE {
			fmt.Printf("Error: this is a special case; printing to stdout is not supported\n")
			os.Exit(1)
		} else {
			fmt.Println(source)
		}
	}
}

// entry point
func main() {
	verify(langMap)

	args := os.Args[1:]

	if lib.Contains(args, "-h") || lib.Contains(args, "--help") {
		printHelp()
		os.Exit(0)
	}
	if lib.Contains(args, "-v") || lib.Contains(args, "--version") {
		fmt.Printf("alap v%s\n", VERSION)
		os.Exit(0)
	}

	to_stdout := false
	if lib.Contains(args, "--stdout") {
		to_stdout = true
		args, _ = lib.Remove(args, "--stdout")
	}
	if len(args) == 0 {
		printHelp()
		os.Exit(0)
	}

	key := args[0]

	if key == "make" {
		special.CreateMakefile(to_stdout)
		return
	}
	// else
	process(key, to_stdout)
	if key == "d" {
		// for a .d file, we also create a simple dub.json file
		process("dub", to_stdout)
	}
}
