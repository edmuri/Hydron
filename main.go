package main

import (
	"fmt"
	"os"
)

func verifyArgs(args []string) {
	fmt.Print(args)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("[!] Not enough arguments")
		return
	}

	verifyArgs(os.Args)

	fmt.Println(os.Args[1:])

}
