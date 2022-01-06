# Go - Ctrlc

[![Go Report Card](https://goreportcard.com/badge/github.com/danielgatis/go-ctrlc?style=flat-square)](https://goreportcard.com/report/github.com/danielgatis/go-ctrlc)
[![License MIT](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/danielgatis/go-ctrlc/master/LICENSE)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/danielgatis/go-ctrlc)

Gracefully quit when you press ctrl-c.

## Install

```bash
go get -u github.com/danielgatis/go-ctrlc
```

And then import the package in your code:

```go
import "github.com/danielgatis/go-ctrlc"
```

### Example

An example described below is one of the use cases.

```go
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/danielgatis/go-ctrlc"
)

func main() {
	tick := time.NewTicker(time.Second)
	exit := make(chan struct{})

	ctrlc.Watch(func() {
		exit <- struct{}{}
	})

	for {
		select {
		case <-exit:
			fmt.Println("bye!")
			os.Exit(0)
		case <-tick.C:
			fmt.Println(time.Now())
		}
	}
}
```

```
❯ go run main.go
2022-01-06 12:08:46.015813 -0300 -03 m=+1.000276951
2022-01-06 12:08:47.016224 -0300 -03 m=+2.000688652
^Cbye!
```

### License

Copyright (c) 2021-present [Daniel Gatis](https://github.com/danielgatis)

Licensed under [MIT License](./LICENSE)

### Buy me a coffee

Liked some of my work? Buy me a coffee (or more likely a beer)

<a href="https://www.buymeacoffee.com/danielgatis" target="_blank"><img src="https://bmc-cdn.nyc3.digitaloceanspaces.com/BMC-button-images/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: auto !important;width: auto !important;"></a>
