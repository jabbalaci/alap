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

```
$ alap -h
alap v0.2.0, https://github.com/jabbalaci/alap

Usage: alap <template_id> [option]

Available options:

-h, --help        show this help
-v, --version     version info
--stdout          don't create source file, print result to stdout

Available templates:

* c             - C source code [main.c]
* cs            - C# source code [Program.cs]
* flask         - Flask source code [app.py]
* go            - Go source code [main.go]
* java          - Java source code [Main.java]
* lua           - Lua source code [main.lua]
* mypy          - mypy ini file [mypy.ini]
* py            - Python 3 source code [main.py]
* rust          - Rust source code [main.rs]
* sh            - Bash source code [main.sh]
```

# Examples

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

        $ alap rust --stdout
        fn main() {
            println!("rust");
        }


# Installation

Download the source code and issue the following command:

        $ go build

Or, you can install it directly:

        $ go install github.com/jabbalaci/alap@HEAD

In this latter case, the binary file `alap` will be in the folder `$GOPATH/bin` .
