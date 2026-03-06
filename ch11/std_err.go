package main

import (
	"fmt"
	"os"
)

func main() {
	if _, err := fmt.Println("This is std err message"); err != nil {
		fmt.Fprintf(os.Stderr, "Error occurred : %v", err)
	}
	if _, err := os.Open("nonexistendtfile.txt"); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
}
