package main

import "fmt"

func test_colors() {
	for i := 0; i < 256; i++ {
		fmt.Printf("\033[48;2;%d;0;0m Haiiiii", i)
	}
}

func main() {
	test_colors()
	fmt.Println("")
}
