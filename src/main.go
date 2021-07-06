package main

import "fmt"

func printFunctionTest(number int) {
	fmt.Println(number)
}

func listNumberFunction(beforeNumber, afterNumber int, message string) {
	fmt.Println()
}

func myWalletMonth(salary, dayMonth, socialSecurity int) (int, int) {
	securitySocial := salary - socialSecurity
	salaryDay := securitySocial / dayMonth
	return salaryDay, securitySocial
}

func main() {
	salaryDay, _ := myWalletMonth(1000000, 28, 500000)
	fmt.Printf(" my day salary %v \n", salaryDay)
	salaryDaySecond, securitySocial := myWalletMonth(1000000, 28, 500000)
	fmt.Printf("%v my day salary and \n  my final sallary %v \n", salaryDaySecond, securitySocial)
	counter := 0
	for counter < 15 {
		printFunctionTest(counter)
		listNumberFunction(counter-1, counter+1, "Final name ")
		counter++
	}
}
