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

func parImparvalue(number int) (message string) {
	switch result := number % 2; result {
	case 0:
		return "par number"
	default:
		return "impar number"
	}
}

func main() {
	//FUnction learing and returns
	salaryDay, _ := myWalletMonth(1000000, 28, 500000)
	fmt.Printf(" my day salary %v \n", salaryDay)
	salaryDaySecond, securitySocial := myWalletMonth(1000000, 28, 500000)
	fmt.Printf("%v my day salary and \n  my final sallary %v \n", salaryDaySecond, securitySocial)
	counter := 0
	//For testing
	for counter < 15 {
		printFunctionTest(counter)
		listNumberFunction(counter-1, counter+1, "Final name ")
		counter++
	}
	//switch syntax
	fmt.Println(parImparvalue(2))
}
