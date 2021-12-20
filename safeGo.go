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

func GoInfiniteLoop(function func(), onPanic func(err interface{}), interval int, redoDelaySec int) {
	runToken := make(chan struct{})
	go func() {
		for {
			<-runToken
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
					function()
					time.Sleep(time.Duration(interval) * time.Second)
				}
			}()
		}
	}()
	runToken <- struct{}{}
}
