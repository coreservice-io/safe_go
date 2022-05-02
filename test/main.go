package main

import (
	"log"
	"time"

	"github.com/coreservice-io/safe_go"
)

func divide(a, b int) int {
	return a / b
}

func main() {
	safe_go.Go(
		// process
		func(args ...interface{}) {
			log.Println("go start")
			log.Println(divide(2, 1))
			log.Println(divide(2, 0))
		},
		// onPanic callback
		func(err interface{}) {
			log.Println("panic catch")
			log.Println(err)
		})

	time.Sleep(600 * time.Second)
}
