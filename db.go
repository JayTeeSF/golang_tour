package main
/*
 * mysql -uroot -p -e 'create database go_db; use go_db; create table users (id int NOT NULL AUTO_INCREMENT, email varchar(255), PRIMARY KEY (id)); insert into users(email) values("jaytee@jayteesf.com");' mysql
 * go get github.com/go-sql-driver/mysql
 */

import (
  "fmt"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

func main() {
  db, err := sql.Open("mysql",
  "root:@tcp(127.0.0.1:3306)/go_db")
  if err != nil {
    panic(err)
  }
  defer db.Close()

  err = db.Ping()
  if err != nil {
    panic(err)
  }
  fmt.Println("opened & pinged the db")

  var (
    id int
    email string
  )
  // rows, err := db.Query("select id, email from users where id = ?", 1)
  rows, err := db.Query("select id, email from users")
  if err != nil {
    panic(err)
  }
  defer rows.Close()
  //fmt.Println("querying...")
  //fmt.Println("row: id: ",id, " email: ", email)
  for rows.Next() {
    err := rows.Scan(&id, &email)
    if err != nil {
      panic(err)
    }
    fmt.Println("row: id: ",id, " email: ", email)
  }
  //fmt.Println("eoR")
  err = rows.Err()
  if err != nil {
      panic(err)
  }
  fmt.Println("done")
}
