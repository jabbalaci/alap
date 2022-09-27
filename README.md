# alap

A little CLI tool for scaffolding basic source codes.

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
    Usage: alap <language>

    Available language templates:

    * c
    * go
    * java
    * py
    * rust

# Example

    $ alap go
    # main.go was created

    $ cat main.go
    package main

    import (
            "fmt"
    )

    func main() {
            fmt.Println("go")
    }
