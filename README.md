Go Logging Library
==================

Use:
----

```
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
	Info, Warning, Error, err := logs.Open(filepath.Join(dir, "output.log"))
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