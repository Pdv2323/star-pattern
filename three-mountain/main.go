package main

import "fmt"

func main() {
	var num = 5
	for loop := 0; loop < 3; loop++ {
		for i := 0; i < num; i++ {
			for j := 0; j < i; j++ {
				fmt.Print(" ")
			}
			fmt.Println("*")
		}
		for i := num - 2; i > 0; i-- {
			for j := i; j > 0; j-- {
				fmt.Print(" ")
			}
			fmt.Println("*")
		}
	}
	fmt.Println("*")

}

// for i := 0; i < num; i++ {
// 	for j := 0; j < i; j++ {
// 		fmt.Print(" ")
// 	}
// 	fmt.Println("*")
// }
// for i := num % 2; i > 0; i-- {
// 	for j := i; j > 0; j-- {
// 		fmt.Print(" ")
// 	}
// 	fmt.Println("*")
// }
