package USafeGo

func Go(function func(args ...interface{}), onPanic func(panic interface{}), args ...interface{}) {
	go func() {
		//catch panic
		defer func() {
			if panic_err := recover(); panic_err != nil {
				if onPanic != nil {
					onPanic(panic_err)
				}
			}
		}()

		//func
		function(args...)
	}()
}
