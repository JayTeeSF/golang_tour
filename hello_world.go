/*
 * go build hello_world.go
 * go run hello_world.go
 */

package main

import (
  "fmt"
  "time"
)

func main() {
  fmt.Println("hey, it's now", time.Now())
}
