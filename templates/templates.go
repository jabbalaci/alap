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

const Python = `
#!/usr/bin/env python3

def main():
    print("Py3")

##############################################################################

if __name__ == "__main__":
    main()
`

const Rust = `
fn main() {
    println!("rust");
}
`
