package main

import (
	"fmt"

	"github.com/brave0duck/duckstring"
	"github.com/fatih/color"
)

func main() {
	fmt.Printf("%v %v\n", color.CyanString("Hello"), color.GreenString("World!"))
	s1, s2 := "This is ", "DuckString"
	fmt.Println(duckstring.AddStr(s1, s2))
}
