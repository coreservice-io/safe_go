package USafeGo

import "time"

func Go(function func(args ...interface{}), onPanic func(err interface{}), args ...interface{}) {
	go func() {
		//catch panic
		defer func() {
			if err := recover(); err != nil {
				if onPanic != nil {
					onPanic(err)
				}
			}
		}()

		//func
		function(args...)
	}()
}

func GoInfiniteLoop(function func() bool, onPanic func(err interface{}), interval int, redoDelaySec int) {
	runToken := make(chan struct{})
	stopSignal := make(chan struct{})
	go func() {
		for {
			select {
			case <-runToken:
				go func() {
					defer func() {
						if err := recover(); err != nil {
							if onPanic != nil {
								onPanic(err)
							}
							time.Sleep(time.Duration(redoDelaySec) * time.Second)
							runToken <- struct{}{}
						}
					}()
					for {
						isGoOn := function()
						if !isGoOn {
							stopSignal <- struct{}{}
							return
						}
						time.Sleep(time.Duration(interval) * time.Second)
					}
				}()
			case <-stopSignal:
				return
			}
		}
	}()
	runToken <- struct{}{}
}
