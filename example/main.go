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
