package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("the number of arguments is not correct")
		return
	}
	fmt.Printf(".intel_syntax noprefix\n")
	fmt.Printf(".globl main\n")
	fmt.Printf("main:\n")
	fmt.Printf("  mov rax, %v\n", os.Args[1])
	fmt.Printf("  ret\n")
	return
}
