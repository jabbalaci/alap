package special

import (
	"alap/jio"
	"alap/templates"
	"fmt"
)

func CreateNuOn() {
	var writeOk bool

	if jio.IsDir(".venv") {
		fmt.Println("# .venv/ is present")
		content := "source .venv/bin/activate.nu"
		writeOk = jio.WriteSourceToFile(content, "on")
		if writeOk {
			fmt.Println("# `on` was created")
		}
		// ---
		var fname string
		fname = ".venv/bin/activate.nu"
		if !jio.IsFile(fname) {
			content := templates.ACTIVATE_NU
			writeOk = jio.WriteSourceToFile(content, fname)
			if writeOk {
				fmt.Printf("# `%s` was created\n", fname)
			}
		} else {
			fmt.Printf("# `%s` already exists\n", fname)
		}
		// ---
		fname = ".venv/bin/deactivate.nu"
		if !jio.IsFile(fname) {
			content := templates.DEACTIVATE_NU
			writeOk = jio.WriteSourceToFile(content, fname)
			if writeOk {
				fmt.Printf("# `%s` was created\n", fname)
			}
		} else {
			fmt.Printf("# `%s` already exists\n", fname)
		}
	}

	if jio.IsFile("pyproject.toml") {
		fmt.Println("# pyproject.toml is present, poetry project detected")
		cmd := "poetry env info -p"
		output := jio.ExecCmd(cmd)
		content := fmt.Sprintf("source %s/bin/activate.nu", output)
		writeOk = jio.WriteSourceToFile(content, "on")
		if writeOk {
			fmt.Println("# `on` was created")
		}
	}
}
