package main

import "fmt"

func gugudan() func(int) {
	return func(a int) {
		for i := 1; i < 10; i++ {
			fmt.Printf("%d x %d = %d\n", a, i, a*i)
		}
	}

}
func main() {
	var input int
	fgugu := gugudan()
	for {
		fmt.Println("구구단 출력 (1~9, 0:exit): ")
		fmt.Scanf("%d", &input)
		if input == 0 {
			fmt.Println("프로그램을 종료합니다 ... ")
			break
		}
		if 1 <= input && input <= 9 {
			fgugu(input)
		} else {
			fmt.Println("1-9사이의 숫자를 입력")
		}
	}

}
