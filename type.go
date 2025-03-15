package main

import "fmt"

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body string
	toAddress string
}

type sms struct {
	isSubscribed bool
	body string
	toPhoneNumber string
}

func (e email) cost() float64 {
	return float64(len(e.body))
}

func (s sms) cost() float64 {
	return float64(len(s.body))
}

func getExpenseReport(e expense) (address string, cost float64) {
	// em, ok := e.(email)
	// if ok {
	// 	return em.toAddress, em.cost()
	// }

	// s, ok := e.(sms)
	// if ok {
	// 	return s.toPhoneNumber, s.cost()
	// }

	// return

	// The code above can be re-written using the type switching syntax
	switch v := e.(type) {
	case email:
		return v.toAddress, v.cost()
	case sms:
		return v.toPhoneNumber, v.cost()
	default: 
		return
	}
}

// Type switch
func printNumericalValue(num interface{}) {
	switch v := num.(type) {
	case int:
		fmt.Printf("%T\n", v) // The code here prints the type of a variable
	case string:
		fmt.Printf("%T\n", v)
	default:
		fmt.Printf("%T\n", v)
	} 
}

func test(e expense) {
	i, c := getExpenseReport(e)

	fmt.Printf("%s costs %.2f\n", i, c)
}

func main() {
	test(email{
		body: "Test email snet to a random non-existing user.",
		toAddress: "ahmadkx@gmail.com",
	})
	printNumericalValue(1)
	printNumericalValue("1")
	printNumericalValue(struct{}{})
}