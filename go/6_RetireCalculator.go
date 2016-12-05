// 6_RetireCalculator
package main

import (
	"fmt"
	"time"
)

func main() {
	var age, retireAge int
	for {
		fmt.Print("What is your current age? ")
		fmt.Scan(&age)
		fmt.Print("At what age would you like to retire? ")
		fmt.Scan(&retireAge)
		if (retireAge - age) < 0 {
			fmt.Println("이미 정년 입니다.")
		} else {
			break
		}
	}
	now := time.Now()

	fmt.Printf("You have %d years left until you can retire.\n", retireAge-age)
	fmt.Printf("It's %d, so you can retire in %d.\n", now.Year(), now.Year()+retireAge-age)

}
