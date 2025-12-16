package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

// Version as printed with -version option
var Version = "UNKNOWN"

const (
	// Help as printed with -help option
	Help = `iso2unix [-h] [-v] iso
Convert ISO 8601 date to UNIX timestamp:
-h     To print this help
-v     To print version
iso    ISO 8601 date to convert
Accepted ISO formats:
- 2006-01-02T15:04:05MST
- 2006-01-02T15:04:05Z
- 2006-01-02 15:04:05 MST
- 2006-01-02 15:04:05 Z
- 2006-01-02T15:04:05
- 2006-01-02 15:04:05
If no timezone is provided, local timezone is assumed`
)

var (
	// TimeFormats defines the accepted ISO time formats with timezone
	TimeFormats = []string{
		"2006-01-02T15:04:05MST",
		"2006-01-02T15:04:05Z",
		"2006-01-02 15:04:05 MST",
		"2006-01-02 15:04:05 Z",
	}
	// TimeFormatsNoTimeZone defines the accepted ISO time formats without timezone
	TimeFormatsNoTimeZone = []string{
		"2006-01-02T15:04:05",
		"2006-01-02 15:04:05",
	}
)

func main() {
	help := flag.Bool("h", false, "Print help")
	version := flag.Bool("v", false, "Print version")
	flag.Parse()
	if *help {
		fmt.Println(Help)
		os.Exit(0)
	}
	if *version {
		fmt.Println(Version)
		os.Exit(0)
	}
	if flag.NArg() < 1 {
		fmt.Println("ERROR missing ISO timestamp argument")
		fmt.Println(Help)
		os.Exit(1)
	}
	iso := strings.TrimSpace(strings.Join(flag.Args(), " "))
	unix, err := Iso2Unix(iso)
	if err != nil {
		fmt.Printf("ERROR '%s' is an invalid ISO timestamp\n", iso)
		os.Exit(1)
	}
	fmt.Println(unix)
}

func Iso2Unix(iso string) (int64, error) {
	for _, format := range TimeFormats {
		timestamp, err := time.Parse(format, iso)
		if err == nil {
			return timestamp.Unix(), nil
		}
	}
	for _, format := range TimeFormatsNoTimeZone {
		timestamp, err := time.ParseInLocation(format, iso, time.Local)
		if err == nil {
			return timestamp.Unix(), nil
		}
	}
	return 0, fmt.Errorf("could not parse ISO timestamp")
}
