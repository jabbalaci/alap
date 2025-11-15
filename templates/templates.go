package templates

import _ "embed"

//go:embed main.sh.txt
var Bash string

//go:embed main.c.txt
var C string

//go:embed Program.cs.txt
var CSharp string

//go:embed main.d.txt
var D string

//go:embed dub.json.txt
var DubJson string

//go:embed app.py.txt
var Flask string

//go:embed main.go.txt
var Go string

//go:embed Main.java.txt
var Java string

//go:embed main.kt.txt
var Kotlin string

//go:embed lorem.txt
var LoremIpsum string

//go:embed main.lua.txt
var Lua string

//go:embed mypy.ini.txt
var Mypy string

//go:embed main.nim.txt
var Nim string

//go:embed main.pas.txt
var Pascal string

//go:embed mongo.py.txt
var Pymongo string

//go:embed main.py.txt
var Python string

//go:embed main.rs.txt
var Rust string

//go:embed main.swift.txt
var Swift string

// --- Makefiles ---

//go:embed makefile.general.txt
var MakefileGeneral string

//go:embed makefile.c.txt
var MakefileC string

//go:embed makefile.d.txt
var MakefileD string

//go:embed makefile.python.txt
var MakefilePython string
