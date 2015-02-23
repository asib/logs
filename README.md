Go Logging Library
==================

Use:
----

```go
package main

import (
	"github.com/asib/logs"
	"log"
	"os"
	"path/filepath"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln("Can't get working dir: ", err)
	}
    // second argument is timeout duration
    // 0 timeout means log everything, timeout > 0 means only log if time since last log is > timeout duration
	Info, Warning, Error, err := logs.NewLogger(filepath.Join(dir, "output.log"), time.Duration(0))
	if err != nil {
		log.Fatalln("Unable to open log file: ", err)
	}

	Info.Println("INFO LOG")
	Warning.Println("WARNING LOG")
	Error.Println("ERROR LOG")
}
```

In output.log:
--------------

```
INFO: 2015/02/19 12:29:17 main.go:20: INFO LOG
WARNING: 2015/02/19 12:29:17 main.go:21: WARNING LOG
ERROR: 2015/02/19 12:29:17 main.go:22: ERROR LOG
```
