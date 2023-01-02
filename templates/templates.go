package templates

import _ "embed"

//go:embed main.c.txt
var C string

//go:embed main.go.txt
var Go string

//go:embed Main.java.txt
var Java string

//go:embed Program.cs.txt
var CSharp string

//go:embed main.py.txt
var Python string

//go:embed app.py.txt
var Flask string

//go:embed main.rs.txt
var Rust string
