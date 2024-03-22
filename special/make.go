package special

import (
	"fmt"
	"os"
	"strings"

	"github.com/jabbalaci/alap/lib"
	path "github.com/jabbalaci/alap/lib"
	"github.com/jabbalaci/alap/templates"
)

type MakeInfo struct {
	_id         string
	description string
	sourceCode  string
}

const makefile = "Makefile"

// template IDs and their associated structs
var langMap = map[string]MakeInfo{
	"0": {_id: "general", description: " - general", sourceCode: templates.MakefileGeneral},
	"1": {_id: "c", description: " - C", sourceCode: templates.MakefileC},
	"2": {_id: "py", description: " - Python", sourceCode: templates.MakefilePython},
}

func printList() {
	fmt.Println("Available Makefile templates:")
	fmt.Println()

	for k, entry := range langMap {
		fmt.Printf("* %2v%v\n", k, entry.description)
	}
	fmt.Println("*  q - quit")
}

// Makefile uses leading tab(s)
func replaceLeadingSpacesWithTabs(source string) string {
	return strings.Replace(source, "    ", "\t", -1)
}

func CreateMakefile(to_stdout bool) {
	if path.Exists(makefile) {
		fmt.Printf("Error: the file '%s' already exists\n", makefile)
		os.Exit(1)
	}
	printList()
	choice := lib.Input("> ")
	if choice == "q" {
		fmt.Println("bye")
		os.Exit(0)
	}
	if _, ok := langMap[choice]; !ok {
		fmt.Println("Error: invalid choice")
		os.Exit(1)
	}
	entry := langMap[choice]
	source := replaceLeadingSpacesWithTabs(entry.sourceCode)
	if to_stdout {
		fmt.Println(source)
	} else { // if writing to file
		writeOk := lib.WriteSourceToFile(source, makefile)
		if writeOk {
			fmt.Printf("# `%s` was created\n", makefile)
		}
	}
}
