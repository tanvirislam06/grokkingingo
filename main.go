package main

import (
	"grokkingingo/algorithms"
)

type employee struct {
	Name   string
	Role   string
	Salary int
}

func (e *employee) SetSalary(salary int) {
	e.Salary = salary
}

func main() {
	// algorithms.RunFind3Sum()
	// algorithms.RunReverseString()
	algorithms.RunRemoveNthLastNode()

	// Basic variable:
	// var num int = 1
	// var numero = 1
	// var number int
	// number = num
	// fmt.Println("num, number, numero: ", num, number, numero)

	// // Declaring array
	// var nums [3]int = [3]int{1, 2, 3}
	// numList := [3]int{3, 4, 5}
	// fmt.Println("nums, numList: ", nums, numList)

	// // Declaring array again
	// var strings [2]string = [2]string{"Hi", "Hello"}
	// fmt.Println("strings: ", strings)

	// // Declaring 2D arrays
	// twoDArray := [1][3]int{
	// 	{1, 2, 3},
	// }
	// fmt.Println("twoDArray: ", twoDArray)

	// var twoD [1][1]int
	// twoD = [1][1]int{{1}}
	// fmt.Println("twoD: ", twoD)

	// // Pointers
	// var value int = 5
	// var pointer *int = &value
	// fmt.Println("pointer: ", pointer)

	// // Structs
	// emp := &employee{
	// 	Name: "Tanvir",
	// }

	// emp.SetSalary(100000)
	// fmt.Println("emp: ", *emp)

	// // Maps
	// scores := map[string]int{"Alice": 10, "Dan": 5}
	// fmt.Println("scores: ", scores["Alice"])

}
