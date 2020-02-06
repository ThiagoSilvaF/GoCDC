package main 

import (
    "fmt"
  
    _ "github.com/lib/pq"
    "gocdc/postgres"
    "gocdc/kafka"

)

const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "postgres"
    dbname   = "inventory"
)

var db_name = "POSTGRES"

func main() {

  fmt.Printf("going to call kafka")

  kafka.SendMessage()

  if db_name == "POSTGRES" {
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
      "password=%s sslmode=disable",
      host, port, user, password)

    postgres.InitDB(psqlInfo)
  }

}
