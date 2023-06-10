package special

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/jabbalaci/alap/lib"
	"github.com/jabbalaci/alap/templates"
)

func CreateVenvInCurrentDir() {
	fmt.Println("# .venv/ is present")
	fileActivateNu := ".venv/bin/activate.nu"

	content := templates.VENV_ACTIVATE_NU
	here, _ := os.Getwd()
	content = strings.Replace(content, "__VIRTUAL_ENV__", filepath.Join(here, ".venv"), -1)
	content = strings.Replace(content, "__BIN_NAME__", "bin", -1)
	content = strings.Replace(content, "(if ('__VIRTUAL_PROMPT__' | is-empty)", "(if ('' | is-empty)", -1)
	content = strings.Replace(content, "($virtual_env | path basename)", "($virtual_env | path dirname | path basename)", -1)

	writeOk := lib.WriteSourceToFile(content, fileActivateNu)
	if writeOk {
		fmt.Printf("# `%s` was created\n", fileActivateNu)
	} else {
		fmt.Printf("# `%s` already exists\n", fileActivateNu)
	}
}

func CreateVenvWithPoetry() {
	fmt.Println("# pyproject.toml is present, poetry project detected")
	cmd := "poetry env info -p"
	poetryVenvFolder, err := lib.ExecCmd(cmd)
	if err != nil {
		fmt.Println("Error: poetry's virt. env. folder was not found")
		return
	}
	fileActivateNu := fmt.Sprintf("%s/bin/activate.nu", poetryVenvFolder)

	content := templates.VENV_ACTIVATE_NU
	content = strings.Replace(content, "__BIN_NAME__", "bin", -1)
	content = strings.Replace(content, "__VIRTUAL_ENV__", poetryVenvFolder, -1)
	content = strings.Replace(content, "(if ('__VIRTUAL_PROMPT__' | is-empty)", "(if ('' | is-empty)", -1)

	err = os.Remove(fileActivateNu)
	if err != nil {
		fmt.Printf("Error: `%s` couldn't be deleted\n", fileActivateNu)
		return
	}

	writeOk := lib.WriteSourceToFile(content, fileActivateNu)
	if writeOk {
		fmt.Printf("# `%s` was created\n", fileActivateNu)
	} else {
		fmt.Printf("# `%s` already exists\n", fileActivateNu)
	}

	// create on
	content = templates.VENV_ON
	content = strings.Replace(content, "{{ path_of_activate.nu }}", fileActivateNu, -1)
	onFile := "on"
	writeOk = lib.WriteSourceToFile(content, onFile)
	if writeOk {
		fmt.Printf("# `%s` was created\n", onFile)
		lib.MakeExecutable(onFile)
	} else {
		fmt.Printf("# `%s` already exists\n", onFile)
	}
}

func CreateNuOn() {
	if lib.IsDir(".venv") {
		CreateVenvInCurrentDir()
	} else if lib.IsFile("pyproject.toml") {
		CreateVenvWithPoetry()
	} else {
		fmt.Println("Error: virt. env. was not found")
		return
	}
}
