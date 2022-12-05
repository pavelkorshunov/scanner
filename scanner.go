package main

import (
	"fmt"
	"os"
)

var scannerFiles = []string{
	"./test.txt",
}

func main() {
	args := os.Args[1:]
	fmt.Println(args)

	for _, val := range scannerFiles {
		content, err := os.ReadFile(val)
		if err != nil {
			fmt.Println("error read file")
			os.Exit(1)
		}

		fmt.Printf("%s", content)
	}
}
