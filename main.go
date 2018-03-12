package main

import (
	"fmt"
	"github.com/383530895/myGo/logs"
)

func main() {
	fmt.Println("this is test")

	log := logs.NewLogger(10000)
	log.SetLogger("console", "")
	log.Warn("this is log warn")
}
