package main

import (
	"fmt"
	"time"

	jtime "github.com/whatever/time"
)

func main() {
	ticker := jtime.NewJitterTicker(300*time.Millisecond, 1*time.Second)
	last := time.Now()
	for t := range ticker.C {
		fmt.Println(t.Sub(last))
		last = t
	}
}
