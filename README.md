# Duplicate Contact Finder

A command-line tool written in Go that detects duplicate contact records within a CSV file using a field-level string similarity algorithm. Records are compared across five fields and flagged as duplicates when their average similarity exceeds a configurable threshold.

---

## How It Works

The tool reads a CSV file containing contact records and performs an all-pairs comparison to identify near-duplicate entries.

### Similarity Algorithm

**`PercentageMatch(str1, str2 string) int`**

Compares two strings character by character from the beginning (case-insensitive). The match score is expressed as a percentage of the longer string's length. An exact match always returns 100 regardless of length.

**`HasMatch(record1, record2 Record) int`**

Computes the average similarity across the five compared fields:

| Field       | Description             |
|-------------|-------------------------|
| `Name`      | First name              |
| `Name1`     | Last name               |
| `Email`     | Email address           |
| `PostalZip` | Postal / ZIP code       |
| `Address`   | Street address          |

**`FindDuplicates(records []Record) ([]Response, error)`**

Iterates over all record pairs. If the average field similarity exceeds 20%, the pair is recorded as a match and the duplicate is removed from further comparisons to avoid redundant reporting.

### Input Format

The tool expects a CSV file with the following column order (header row is read as a record and should be excluded or handled accordingly):

```
contactID,name,name1,email,postalZip,address
```

### Output Format

Each detected duplicate is reported as a `Response` value containing:

| Field             | Description                                     |
|-------------------|-------------------------------------------------|
| `ContactIDSource` | ID of the original (first-seen) record          |
| `ContactIDMatch`  | ID of the record identified as a duplicate      |
| `Acurracy`        | Similarity score as a percentage (0 to 100)     |

---

## Project Structure

```
.
├── main.go         # Entry point, CSV loading, core duplicate-detection logic
├── models.go       # Data structures: Record and Response
├── main_test.go    # Unit and integration tests
├── go.mod          # Module definition and dependencies
└── go.sum          # Dependency checksums
```

---

## Requirements

- [Go](https://go.dev/dl/) 1.21 or later

---

## Installation

**1. Clone the repository**

```bash
git clone <repository-url>
cd compass_tech
```

**2. Install dependencies**

```bash
go mod download
```

---

## Running the Program

The program reads from the CSV file hardcoded in `main.go`. Place your input file in the project root and update the filename in `main()` if needed, then run:

```bash
go run .
```

The tool will print each detected duplicate pair along with its similarity score to standard output.

**Example output:**

```
Duplicates found: [{1 3 96} {1 6 92} {2 11 100}]
```

Each entry represents `{ContactIDSource ContactIDMatch Accuracy}`.

---

## Running the Tests

The test suite covers `PercentageMatch`, `HasMatch`, and `FindDuplicates` across no-duplicate, single-duplicate, and multi-duplicate scenarios.

```bash
go test ./...
```

To run with verbose output:

```bash
go test -v ./...
```

---

## Dependencies

| Package                        | Purpose                         |
|--------------------------------|---------------------------------|
| `github.com/stretchr/testify` | Assertion helpers in unit tests |
