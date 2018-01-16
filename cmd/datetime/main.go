package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	layout := flag.String("format", "2006-01-02 15:04:05", "format according to this layout")
	flag.Parse()

	now := time.Now().Format(*layout)
	fmt.Printf("%s\n%s\n", now, now)
}
