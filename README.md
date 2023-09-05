# bp-csv

Formats exports from
[Blood Pressure Monitor](https://f-droid.org/en/packages/com.derdilla.bloodPressureApp/)
to a more human-readable format.

### Pre-Reqs

[Go 1.18+](https://go.dev/)

## Install

```bash
make install
# or
go install
```

## Build

```bash
make
# or
go build -o ./bin/bp-csv
```

## Usage

```bash
bp-csv ./path/to/original/csv/export.csv
```
