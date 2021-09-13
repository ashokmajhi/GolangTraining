/*
Building on the previous hands-on exercise, create a program that uses “else if” and “else”.
*/
package main

import "fmt"

func main() {
	x := "Ashok Majhi"

	if x == "Moneypenny" {
		fmt.Println(x)
	} else if x == "Ashok Majhi" {
		fmt.Println("IBM Employee : ", x)
	} else {
		fmt.Println("neither")
	}
}
