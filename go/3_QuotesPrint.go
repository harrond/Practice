// 3_QuotesPrint
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var name string
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("What is the quote? ")
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str) //뒤에 개행문자붙어서 제거
	fmt.Print("Who said it? ")
	fmt.Scanf("%s", &name)

	fmt.Printf("%s says, \"%s\"\n", name, str)

}
