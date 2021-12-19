package main

import (
	"github.com/universe-30/USafeGo"
	"log"
	"time"
)

func divide(a, b int) int {
	return a / b
}

type SomePointer struct {
	Name string
}

func (p *SomePointer) Run() {
	log.Println("some run", p.Name)
}

func main() {
	USafeGo.Go(func(args ...interface{}) {
		log.Println("go start")
		log.Println(divide(2, 1))
		log.Println(divide(2, 0))
	}, func(err interface{}) {
		log.Println("panic catch")
		log.Println(err)
	})

	USafeGo.Go(func(args ...interface{}) {
		var p *SomePointer
		//nil
		p.Run()
	}, func(err interface{}) {
		log.Println("panic catch")
		log.Println(err)
	})

	time.Sleep(5 * time.Second)
}
