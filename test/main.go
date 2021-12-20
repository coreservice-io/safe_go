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
	//goroutine
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

	//infinite loop
	count := 0
	USafeGo.GoInfiniteLoop(func() bool {
		count++
		log.Println(count)
		if count%3 == 0 {
			divide(3, 0)
		}
		if count > 6 {
			return false
		}
		return true
	}, func(err interface{}) {
		log.Println(err)
	}, 3, 5)

	time.Sleep(600 * time.Second)
}
