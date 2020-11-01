package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println(".intel_syntax noprefix\n")
	fmt.Println(".globl main\n")
	fmt.Println("main:\n")
	flag.Parse()
	fmt.Printf("  mov rax, %v\n", flag.Arg(0))
	fmt.Println("  ret\n")
}
