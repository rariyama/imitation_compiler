package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <sys/stat.h>
#include <errno.h>
*/
import "C"
import (
	"fmt"
	"os"
)

func goStrtol(s string, base int) (int, error) {
	p := C.CString(s)
	n, err := C.strtol(p, nil, C.int(base))
	return int(n), err
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("the number of arguments is not correct")
		return
	}
	var input string
	input = os.Args[1]

	fmt.Printf(".intel_syntax noprefix\n")
	fmt.Printf(".globl main\n")
	fmt.Printf("main:\n")
	//get first number before symbol.
	initialNum, err := goStrtol(input, 10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("  mov rax, %v\n", initialNum)
	// get numbers after first symbol.
	for i, char := range input {
		if char == '+' {
			var num, err = goStrtol(string(input)[i+1:], 10)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("  add rax, %v\n", num)
		} else if char == '-' {
			var num, err = goStrtol(string(input)[i+1:], 10)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("  sub rax, %v\n", num)
		}
	}
	fmt.Printf("  ret\n")
	return
}
