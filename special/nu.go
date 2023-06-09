package special

import (
	"fmt"
	"os"

	"github.com/jabbalaci/alap/lib"
	"github.com/jabbalaci/alap/templates"
)

func CreateNuOn() {
	var writeOk bool
	var fname string
	isPoetryProject := false

	if lib.IsDir(".venv") {
		fmt.Println("# .venv/ is present")
		fname = ".venv/bin/activate.nu"
	} else if lib.IsFile("pyproject.toml") {
		isPoetryProject = true
		fmt.Println("# pyproject.toml is present, poetry project detected")
		cmd := "poetry env info -p"
		folder, err := lib.ExecCmd(cmd)
		if err != nil {
			fmt.Println("Error: poetry's virt. env. folder was not found")
			return
		}
		fname = fmt.Sprintf("%s/bin/activate.nu", folder)
	} else {
		fmt.Println("Error: virt. env. was not found")
		return
	}

	if fname != "" {
		content := templates.ACTIVATE_NU
		if isPoetryProject {
			err := os.Remove(fname)
			if err != nil {
				fmt.Printf("Error: `%s` couldn't be deleted\n", fname)
				return
			}
		}
		writeOk = lib.WriteSourceToFile(content, fname)
		if writeOk {
			fmt.Printf("# `%s` was created\n", fname)
		} else {
			fmt.Printf("# `%s` already exists\n", fname)
		}
	}

	if isPoetryProject {
		// create on.nu
		cmd := fmt.Sprintf("overlay use %s", fname) // fname: full path to activate.nu
		onNuFile := "on.nu"
		writeOk = lib.WriteSourceToFile(cmd, onNuFile)
		if writeOk {
			fmt.Printf("# `%s` was created\n", onNuFile)
		} else {
			fmt.Printf("# `%s` already exists\n", onNuFile)
		}
	}
}
