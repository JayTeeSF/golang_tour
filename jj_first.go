package main

import (
  "encoding/csv"
  "fmt"
  "io"
  "os"
)


func main() {
  var default_file_name, file_name string
  default_file_name = "../jj_first/questions_and_answers.csv"
  fmt.Printf("CSV File Path [e.g. '%s']: ", default_file_name)
  fmt.Scanln(&file_name)
  fmt.Println()

  _, err := os.Stat(file_name)
  if err != nil {
    file_name = default_file_name
  }

  f, err := os.Open(file_name)
  if err != nil {
    panic(err)
  }

  reader := csv.NewReader(f)

  var count int = 0
  for {
    record, err := reader.Read()
    if count > 0 { // skip header
      if err == io.EOF {
        break
      } else if err != nil {
        panic(err)
      }
      var question string = string(record[0])
      var answer string = string(record[1])
      var guess string

      fmt.Print(question, " ")
      fmt.Scanln(&guess)
      fmt.Println()
      if guess == answer {
        fmt.Println("winner winner, chicken dinner")
      } else {
        fmt.Println("sorry, chuck it was: ", answer)
      }
    }
    count += 1
  }
}
