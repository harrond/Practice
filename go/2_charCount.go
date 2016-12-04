// 2_charCount
package main

import (
	"fmt"
	"strconv"
) //한글 3바이트
func main() {
	var str string
	var strlen int
	for {
		fmt.Print("What is the input string? ")
		fmt.Scanf("%s", &str)
		strlen = len(str)
		if strlen != 0 {
			break
		} else {
			fmt.Println("insert string plz")
		}
	}
	str = str + " has " + strconv.Itoa(strlen) + " characters."
	fmt.Println(str)

}
