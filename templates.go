package main

const CTemplate = `
#include <stdio.h>

int main()
{
    printf("hello\n");

    return 0;
}
`

const GoTemplate = `
package main

import (
	"fmt"
)

func main() {
	fmt.Println("go")
}
`

const JavaTemplate = `
public class Main
{
    public static void main(String[] args)
    {
        System.out.println("hello");
    }
}
`

const PythonTemplate = `
#!/usr/bin/env python3

def main():
    print("Py3")

##############################################################################

if __name__ == "__main__":
    main()
`

const RustTemplate = `
fn main() {
    println!("rust");
}
`
