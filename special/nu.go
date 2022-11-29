package special

import (
	"fmt"

	"github.com/jabbalaci/alap/lib"
	"github.com/jabbalaci/alap/templates"
)

func CreateNuOn() {
	var writeOk bool

	if lib.IsDir(".venv") {
		fmt.Println("# .venv/ is present")
		content := "source .venv/bin/activate.nu"
		writeOk = lib.WriteSourceToFile(content, "on")
		if writeOk {
			fmt.Println("# `on` was created")
		}
		// ---
		var fname string
		fname = ".venv/bin/activate.nu"
		if !lib.IsFile(fname) {
			content := templates.ACTIVATE_NU
			writeOk = lib.WriteSourceToFile(content, fname)
			if writeOk {
				fmt.Printf("# `%s` was created\n", fname)
			}
		} else {
			fmt.Printf("# `%s` already exists\n", fname)
		}
		// ---
		fname = ".venv/bin/deactivate.nu"
		if !lib.IsFile(fname) {
			content := templates.DEACTIVATE_NU
			writeOk = lib.WriteSourceToFile(content, fname)
			if writeOk {
				fmt.Printf("# `%s` was created\n", fname)
			}
		} else {
			fmt.Printf("# `%s` already exists\n", fname)
		}
	}

	if lib.IsFile("pyproject.toml") {
		fmt.Println("# pyproject.toml is present, poetry project detected")
		cmd := "poetry env info -p"
		output := lib.ExecCmd(cmd)
		content := fmt.Sprintf("source %s/bin/activate.nu", output)
		writeOk = lib.WriteSourceToFile(content, "on")
		if writeOk {
			fmt.Println("# `on` was created")
		}
	}
}
