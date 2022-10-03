# alap

A little CLI tool for scaffolding files. Originally,
it was made to create basic source codes.

The word "alap" means "base" or "basic" in Hungarian.

# Motivation

When you start a new project, usually you always start
with the same basic source code. This tool will create
that basic source code for you.

The tool is written in Go, thus you have a single,
self-contained executable that you just need to put
somewhere in the PATH.

# Usage

        $ alap -h
        Usage: alap <template_id>

        Available templates:

        * c             - C source code
        * go            - Go source code
        * java          - Java source code
        * nuon          - create `on` for activating a virt. env. from Nushell
        * py            - Python 3 source code
        * rust          - Rust source code

# Example

        $ ls -al
        drwxr-xr-x   2 jabba jabba   40 okt    3 13.14 .
        drwxrwxrwt 186 root  root  4920 okt    3 13.11 ..

        $ alap go
        # `main.go` was created

        $ ls -al
        drwxr-xr-x   2 jabba jabba   60 okt    3 13.14 .
        drwxrwxrwt 186 root  root  4920 okt    3 13.11 ..
        -rw-r--r--   1 jabba jabba   68 okt    3 13.14 main.go

        $ cat main.go
        package main

        import (
                "fmt"
        )

        func main() {
                fmt.Println("go")
        }

# Changelog

* **v0.1.1** - Support for Python's virtual environments in Nushell. Now
you can activate / deactivate a virt. env. easily in Nushell.
* **v0.1.0** - initial version
