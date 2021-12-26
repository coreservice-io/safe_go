# USafeGo

use this package to replace golang build-in goroutine to prevent the program crash from unexpected panic

### how to use
```go
import "github.com/universe-30/USafeGo"
```

### example
```go
import (
	"log"
	"time"

	"github.com/universe-30/USafeGo"
)

func divide(a, b int) int {
	return a / b
}

func main() {
	//use USafeGo package
	USafeGo.Go(
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
```

