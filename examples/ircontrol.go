package main

import (
	"fmt"
	"log"
	"time"

	"github.com/barnex/ev3dev-lang-go/ev3"
)

func main() {
	fmt.Println("v5")
	m, err := ev3.OpenMotor("")
	if err != nil {
		log.Fatal(err)
	}

	m.SetDutyCycleSP(50)
	m.SetCommand("run-forever")
	time.Sleep(1 * time.Second)
	m.SetCommand("stop")
}
