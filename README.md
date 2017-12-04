
#install
* git get github.com/aiyeyun/logger.git

#Instructions

```golang
package main

import "github.com/aiyeyun/logger"

func main() {
    logger.Info("this is info message")
	logger.Infof("this is info message %s", "hello world")

	logger.Error("this is error message")
	logger.Errorf("this is error message %s", "hello world")

	logger.Warning("this is warning message")
	logger.Warningf("this is warning message %s", "hello world")

	logger.Fatal("this is fatal message")
	logger.Fatalf("this is fatal message %s", "hello world")
}

```

#output shell

```shell
[INFO] 2017-12-04 17:47:42 main.go:6 ▶ this is info message
[INFO] 2017-12-04 17:47:42 main.go:7 ▶ this is info message hello world
[ERROR] 2017-12-04 17:47:42 main.go:9 ▶ this is error message
[ERROR] 2017-12-04 17:47:42 main.go:10 ▶ this is error message hello world
[WARNING] 2017-12-04 17:47:42 main.go:12 ▶ this is warning message
[WARNING] 2017-12-04 17:47:42 main.go:13 ▶ this is warning message hello world
[FATAL] 2017-12-04 17:47:42 main.go:15 ▶ this is fatal message
```