package main

import (
  "database/sql"
  "fmt"
  _ "github.com/lib/pq"
)

const (
  host = "localhost"
  port = 5432
  user = "postgres"
  password = "postgres"
  dbname = "test"
)

func handleError(err error) {
  if err != nil {
    panic(err)
  }
}

func main() {
  psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
  db, err := sql.Open("postgres", psqlconn)
  handleError(err)

  defer db.Close()

  err = db.Ping()
  handleError(err)

  fmt.Println("Connected to postgres db")

  insertStatement := `insert into "accounts"("first_name", "last_name", "email") values($1, $2, $3)`
  _, err = db.Exec(insertStatement, "Devon", "Morris", "devonmorris1992@gmail.com")
  handleError(err)

  fmt.Println("Inserted record")
}
