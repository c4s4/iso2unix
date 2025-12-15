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
iso    ISO 8601 date to convert`
	// Time format in ISO 8601 format
	IsoFormat = "2006-01-02T15:04:05Z"
	// Time format in human readable format
	HumanFormat = "2006-01-02 15:04:05 UTC"
)

func main() {
	help := flag.Bool("h", false, "Print help")
	version := flag.Bool("v", false, "Print version")
	human := flag.Bool("r", false, "Print human readable format")
	flag.Parse()
	if *help {
		fmt.Println(Help)
		os.Exit(0)
	}
	if *version {
		fmt.Println(Version)
		os.Exit(0)
	}
	if flag.NArg() != 1 {
		fmt.Println("ERROR missing ISO timestamp argument")
		fmt.Println(Help)
		os.Exit(1)
	}
	unix, err := Iso2Unix(flag.Arg(0), *human)
	if err != nil {
		fmt.Printf("ERROR invalid ISO timestamp: %v\n", err)
		fmt.Println(Help)
		os.Exit(1)
	}
	fmt.Println(unix)
}

func Iso2Unix(iso string, human bool) (int64, error) {
	iso = strings.TrimSpace(iso)
	var t time.Time
	var err error
	if human {
		t, err = time.Parse(HumanFormat, iso)
		if err != nil {
			return 0, fmt.Errorf("cannot parse human readable time: %v", err)
		}
	} else {
		t, err = time.Parse(IsoFormat, iso)
		if err != nil {
			return 0, fmt.Errorf("cannot parse ISO time: %v", err)
		}
	}
	return t.Unix(), nil
}
