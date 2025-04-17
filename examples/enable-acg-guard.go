package main

import (
	"fmt"
	"github.com/irony0egoist/Hooka/pkg/hooka"
	"log"
	"time"
)

func main() {
	err := hooka.EnableACG()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("ACG Guard enabled!")
	time.Sleep(1000 * time.Second)

  // Do some other stuff
}
