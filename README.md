# Unix Ti ISO

Tool to parse a ISO date and print it as a Unix timestamp.

## Installation

### Unix users (Linux and MacOSX)

Unix users may download and install latest *iso2unix* release with command:

```bash
sh -c "$(curl https://sweetohm.net/dist/iso2unix/install)"
```

If *curl* is not installed on you system, you might run:

```bash
sh -c "$(wget -O - https://sweetohm.net/dist/iso2unix/install)"
```

**Note:** Some directories are protected, even as *root*, on **MacOSX** (since *El Capitan* release), thus you can't install *iso2unix* in */usr/bin* for instance.

### Binary package

Otherwise, you can download latest binary archive at <https://github.com/c4s4/iso2unix/releases>. Unzip the archive, put the binary of your platform somewhere in your *PATH* and rename it *iso2unix*.

## Usage

To convert ISO date *2021-01-01T00:00:00Z* to Unix timestamp, run:

```bash
$ iso2unix 2021-01-01T00:00:00Z
1609459200
```

You can get help with `-h` flag:

```bash
$ iso2unix -h
iso2unix [-h] [-v] iso
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
If no timezone is provided, local timezone is assumed
```

And version with `-v` flag:

```bash
$ iso2unix -v
1.0.0
```

*Enjoy!*
