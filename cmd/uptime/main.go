package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/metalmatze/i3blocks-go/fontawesome"
)

func main() {
	raw, err := ioutil.ReadFile("/proc/uptime")
	if err != nil {
		// TODO
		os.Exit(1)
	}
	raw = bytes.TrimSpace(raw)            // Remove trailing newline
	raws := bytes.Split(raw, []byte(" ")) // Split up & idle

	up, err := strconv.ParseFloat(string(raws[0]), 64)
	if err != nil {
		// TODO
		os.Exit(2)
	}

	uptime := time.Duration(up) * time.Second

	out := fmt.Sprintf("%s%d:%d",
		fontawesome.ArrowCircleUp,
		int(uptime.Hours())%24,
		int(uptime.Minutes())%60,
	)

	fmt.Printf("%s\n%s\n", out, out)
}
