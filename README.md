# Go Luftdaten API

simple server to collect data from your sensors for the luftdaten.info-firmware
 written in golang. All the data gets written to a csv-file.

## Installation

To install this package it is recommended to use `go get`.

```https://github.com/craftamap/go-luftdaten-api```

Binaries for various platforms will be provided in the future on the releases
 page.

## Usage/Examples

To launch the server, simply run the application.

```go-luftdaten-api```

Use `-o` to set the csv-output file, and `--host` (defaults to 0.0.0.0)
and `-p` (or `--port`,defaults to 8080) to set those.

If you want to use this as a systemd-service, the unit file will follow soon™.

## To-Do's / Roadmap

- [ ] Commenting the Code / Cleaning up the codebase
- [ ] SQLite Support
- [ ] config file
