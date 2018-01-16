package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"io/ioutil"

	"github.com/metalmatze/i3blocks-go/fontawesome"
)

func main() {

	// Allow to specify high and critical thresholds.
	highFlag := flag.Int("highTemp", 72, "Specify which temperature threshold in Celsius is considered high.")
	criticalFlag := flag.Int("criticalTemp", 80, "Specify which temperature threshold in Celsius is considered critical.")
	flag.Parse()

	// Gather temperature thresholds.
	criticalTemp := *criticalFlag
	highTemp := *highFlag
	diffTemp := criticalTemp - highTemp

	// Set display texts to defaults.
	var icon string
	var output string
	var fullText string = "error"
	var shortText string = "error"

	// Read CPU temperature information from kernel
	// pseudo-file-system mounted at /sys.
	tempRaw, err := ioutil.ReadFile("/sys/class/hwmon/hwmon0/temp1_input")
	if err != nil {

		// Write an error to STDERR, fallback display values
		// to STDOUT and exit with failure code.
		fmt.Fprintf(os.Stderr, "[i3blocks-go] Failed to read CPU temperature file: %s", err.Error())
		fmt.Fprintf(os.Stdout, "%s\n%s\n", fullText, shortText)
		os.Exit(0)
	}

	// Trim spaces.
	tempString := strings.TrimSpace(string(tempRaw))

	// Convert temperature string to integer.
	temp, err := strconv.Atoi(tempString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[i3blocks-go] Could not convert temperature value: %s", err.Error())
		fmt.Fprintf(os.Stdout, "%s\n%s\n", fullText, shortText)
		os.Exit(0)
	}

	// Normalize temperature value.
	temp = temp / 1000

	// Define temperature values in dependence on
	// specified high and critical values.
	mediumTemp := highTemp - diffTemp
	lowTemp := mediumTemp - diffTemp

	// Depending on current temperature value,
	// set appropriate thermometer icon.
	if temp <= lowTemp {
		icon = fontawesome.ThermometerEmpty
	} else if (temp > lowTemp) && (temp <= mediumTemp) {
		icon = fontawesome.ThermometerQuarter
	} else if (temp > mediumTemp) && (temp <= highTemp) {
		icon = fontawesome.ThermometerHalf
	} else if (temp > highTemp) && (temp <= criticalTemp) {
		icon = "<span foreground=\"#ffae00\">" + fontawesome.ThermometerThreeQuarters + "</span>"
	} else {
		icon = "<span foreground=\"#ff0000\">" + fontawesome.ThermometerFull + "</span>"
	}

	// Build final output string.
	output = fmt.Sprintf("%s%d°C", icon, temp)

	fullText = output
	shortText = output

	// Write out gathered information to STDOUT.
	fmt.Fprintf(os.Stdout, "%s\n%s\n", fullText, shortText)
	os.Exit(0)
}
