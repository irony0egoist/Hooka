package main

import (
  "fmt"
  "log"
  "github.com/irony0egoist/Hooka/pkg/hooka"
)

func main(){
  check, err := hooka.AutoCheck()
  if err != nil {
    log.Fatal(err)
  }

  if (check == true) {
    fmt.Println("Probably a sandbox")
    return
  }

  fmt.Println("Not a sandbox")
}

