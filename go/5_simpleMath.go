// 5_simpleMath
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var firstNum, seconNum string
	var fn, sn int
	var err1, err2 error
	for {
		fmt.Print("What is the first number? ")
		fmt.Scan(&firstNum)
		fn, err1 = strconv.Atoi(firstNum)
		if (err1 != nil) || (fn < 0) {
			fmt.Println("음수가 아닌 \"수\"를 입력해주세요.")
		} else {
			break
		}
	}
	for {
		fmt.Print("What is the second number? ")
		fmt.Scan(&seconNum)
		sn, err2 = strconv.Atoi(seconNum)
		if (err2 != nil) || (sn <= 0) {
			fmt.Println("음수나 0이 아닌 \"수\"를 입력해주세요.")
		} else {
			break
		}
	}

	fmt.Printf("%s + %s = %d\n%s - %s = %d\n%s * %s = %d\n%s / %s = %d\n", firstNum, seconNum, fn+sn, firstNum, seconNum, fn-sn, firstNum, seconNum, fn*sn, firstNum, seconNum, fn/sn)

}
