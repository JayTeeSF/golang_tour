package main

func main() {
	var c chan bool
	// ↪ go run stuck.go
	// fatal error: all goroutines are asleep - deadlock!
	c <- true
}
