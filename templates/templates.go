package templates

const C = `
#include <stdio.h>

int main()
{
    printf("hello\n");

    return 0;
}
`

const Go = `
package main

import (
	"fmt"
)

func main() {
	fmt.Println("go")
}
`

const Java = `
public class Main
{
    public static void main(String[] args)
    {
        System.out.println("hello");
    }
}
`

const CSharp = `
using System;
using static System.Console;
using System.Linq;
using System.Collections;
using System.Collections.Generic;

namespace Example
{
    public class Program
    {
        public static void Main(string[] args)
        {
            WriteLine("hello");
        }
    }
}
`

const Python = `
#!/usr/bin/env python3

def main():
    print("Py3")

##############################################################################

if __name__ == "__main__":
    main()
`

const Flask = `
#!/usr/bin/env python3

from flask import Flask, render_template

app = Flask(__name__)


@app.route("/")
def index():
    context = {
        "text": "hello"
    }
    return render_template("index.html", **context)


if __name__ == "__main__":
    app.run(debug=True)  # listen on localhost ONLY
#    app.run(debug=True, host='0.0.0.0')    # listen on all public IPs
`

const Rust = `
fn main() {
    println!("rust");
}
`
