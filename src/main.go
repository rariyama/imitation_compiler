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
	fmt.Println(".intel_syntax noprefix\n")
	fmt.Println(".globl main\n")
	fmt.Println("main:\n")
	fmt.Printf("  mov rax, %v\n", os.Args[1])
	fmt.Println("  ret\n")
	return
}
