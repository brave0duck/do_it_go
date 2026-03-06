package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Println("input : ")
	fmt.Scanln(&input)
	fmt.Println("you input :  ", input)

	// reader := bufio.NewReader(os.Stdin)

	// fmt.Println("input any words : ")
	// txt, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Error : %v", err)
	// }
	// txt = strings.Trim(txt, "")
	// fmt.Println("You typing word is : ", txt)

}
