// 1_Hello
package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Println("What is your name? ")
	fmt.Scanln(&name)
	name = "Hello, " + name + ", nice to meet you!"
	fmt.Println(name)
}
