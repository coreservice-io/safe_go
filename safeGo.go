package USafeGo

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
