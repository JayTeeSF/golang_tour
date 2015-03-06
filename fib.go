package main

import "fmt"

/*
// need to pass-in the condition
func IntPromptWithDefaultInt(message string, default_response int) int {
	response := IntPrompt(fmt.Sprintf(message, default_response))
	if response < 0 { // the condition
    response = default_response
	}
}
*/

func IntPrompt(message string) int {
	response int
	fmt.Print(message, " ")
  fmt.Scanln(&response)
	fmt.Println()
  return response
}

func Fib(count int) int {
	if count < 2 {
		return 1
	} else {
	  return Fib(count - 1) + Fib(count - 2)
	}
}

func main() {
	default_count = 1
	//message := "Fib [%d]:"
	//count := IntPromptWithDefaultInt(message, default_count)
	message := fmt.Sprintf("Fib [%d]:", default_count)
	count := IntPrompt(message)
	if count < 0 {
    count = default_count
	}
  fmt.Println(" => ", Fib(count))
}
