package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/metalmatze/i3blocks-go/fontawesome"
)

var volRegex = regexp.MustCompile(`\[(\d{1,3})\%\]`)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	vol, err := volume(ctx)
	if err != nil {
		// Write an error to STDERR, fallback display values
		// to STDOUT and exit with failure code.
		fmt.Fprintf(os.Stderr, "[i3blocks-volume] Failed to get volume: %v", err)
		//fmt.Fprintf(os.Stdout, "%s\n%s\n", fullText, shortText)
		return
	}

	//output := ""
	//if vol == -1 {
	//	output = "<span foreground=\"#777777\"></span>  off"
	//} else if vol == 0 {
	//	output = fmt.Sprintf(" %3d%%", vol)
	//} else if (vol > 0) && (vol < 50) {
	//	output = fmt.Sprintf(" %3d%%", vol)
	//} else {
	//	output = fmt.Sprintf(" %3d%%", vol)
	//}

	output := fmt.Sprintf("%s%d%%", fontawesome.VolumeUp, vol)

	fmt.Printf("%s\n%s\n", output, output)
}

// volume queries the amixer command and attempts to extract the volume and mute level
func volume(ctx context.Context) (int, error) {
	// -1 is muted
	vol := 0

	// Execute 'amixer sget Master' and store output of this command
	// written to either STDOUT or STDERR.
	cmd := exec.CommandContext(ctx, "amixer", "sget", "Master")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return 0, err
	}

	if err := cmd.Start(); err != nil {
		return 0, err
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		// Walk over received command output and check
		// for speaker sub-string.
		if strings.Contains(line, "Front Left: Playback") {

			// If speakers are muted, return -1.
			if strings.Contains(line, "[off]") {
				return -1, nil
			}

			// Otherwise, attempt to match above regex
			// on speaker string.
			matches := volRegex.FindStringSubmatch(line)
			if len(matches) != 2 {
				return 0, fmt.Errorf("expected two matches but found %d", len(matches))
			}

			// Convert extracted volume string to integer.
			vol, err = strconv.Atoi(matches[1])
			if err != nil {
				return 0, err
			}

			break
		}
	}

	if err := cmd.Wait(); err != nil {
		return 0, err
	}

	return vol, nil
}
