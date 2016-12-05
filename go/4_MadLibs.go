// 4_MadLibs
package main

import (
	"fmt"
	"strings"
)

func main() {
	var none, verb, adj, adverb string
	fmt.Print("Enter a none: ")
	fmt.Scan(&none)
	fmt.Print("Enter a verb: ")
	fmt.Scan(&verb)
	fmt.Print("Enter an adjective: ")
	fmt.Scan(&adj)
	fmt.Print("Enter an adverb: ")
	fmt.Scan(&adverb)

	str := strings.Replace("Do you 1 2 3 4? That's hiarious!", "1", none, 1)
	str = strings.Replace(str, "2", verb, 1)
	str = strings.Replace(str, "3", adj, 1)
	str = strings.Replace(str, "4", adverb, 1)

	fmt.Println(str)
}
