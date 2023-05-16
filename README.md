Practical Go for Developers

This repo contains the material I produced via the "Practical Go Foundations" Ardan Labs course ran by Miki Tebeka. Great tutor and would highly recommend to anyone jumping into Go from another language they've learned well.

---

## Day 1 Overview

### Agenda

- Strings & formatted output
    - What is a string?
    - Unicode basics
    - Using fmt package for formatted output
- Calling REST APIs
    - Making HTTP calls with net/http
    - Defining structs
    - Serializing JSON
- Working with files
    - Handling errors
    - Using defer to manage resources
    - Working with io.Reader & io.Writer interfaces

### Code

- [hw.go](hw/hw.go) - Hello World
    - `GOOS=drawin go build` (also `GOARCH`)
- [banner.go](banner/banner.go) - Strings & printing
- [github.go](github/github.go) - Calling REST APIs
- [sha1.go](sha1/sha1.go) - Working with `io.Reader` & `io.Writer`

---

## Day 2 Overview

### Agenda

- Sorting
    - Working with slices
    - Writing methods
    - Understanding interfaces
- Catching panics
    - The built-in recover function
    - Named return values
- Processing text
    - Reading line by line with bufio.Scanner
    - Using regular expressions
    - Working with maps

### Code

- [slices.go](slices/slices.go) - Working with slices
- [game.go](game/game.go) - Structs, methods & interfaces
- [empty.go](empty/empty.go) - The empty interface, type assertions
- [div.go](div/div.go) - Catching panics
- [freq.go](freq/freq.go) - Most common word (files, regular expressions, maps)

### Exercises

- Read and understand the [sort package examples](https://pkg.go.dev/sort/#pkg-examples)
- Implement `sortByDistance(players []Player, x, y int)` in `game.go`
- Change `mostCommon` to return the most common `n` words (e.g. `func mostCommon(r io.Reader, n int) ([]string, error)`)

---

## Day 3 Overview

### Agenda

- Distributing work
    - Using goroutines & channels
    - Using the sync package to coordinate work
- Timeouts & cancellation
    - Working with multiple channels using select
    - Using context for timeouts & cancellations
    - Standard library support for context

### Code

- [go_chan.go](go_chan/go_chan.go) - Goroutines & channels
    - [sleep_sort.sh](go_chan/sleep_sort.sh) - Sleep sort in bash
- [taxi_check.go](taxi/taxi_check.go) - Turn sequential code to parallel
- [sites_time.go](sites_time/sites_time.go) - Using sync.WaitGroup
- [payment.go](payment/payment.go) - Using sync.Once & sync.WaitGroup
- [counter.go](counter/counter.go) - Using the race detector, sync.Mutex and sync/atomic
- [select.go](select/select.go) - Using `select`
- [rtb.go](rtb/rtb.go) - Using `context` for cancellations

### Exercise

In `taxi_check.go`
- Limit the number of goroutines to "n". Which "n" yields the best results?
- Cancel all goroutines once there's an error or mismatch in signature

---

## Day 4 Overview

### Agenda

- Testing your code
    - Working with the testing package
    - Using testify
    - Managing dependencies with go mod
- Structuring your code
    - Writing sub-packages
- Writing an HTTP server
    - Writing handlers
    - Using gorilla/mux for routing
      Adding metrics & logging
    - Using expvar for metrics
    - Using the log package and a look at user/zap
- Configuration patterns
    - Reading environment variables and a look at external packages
    - Using the flag package for command line processing

### Code

`nlp` project

<pre>
├── <a href="nlp/go.mod">go.mod</a> - Project & dependencies
├── <a href="nlp/nlp.go">nlp.go</a> - Package code
├── <a href="nlp/doc.go">doc.go</a> - Package level documentation
├── <a href="nlp/nlp_test.go">nlp_test.go</a> - Test & benchmark file
├── <a href="nlp/example_test.go">example_test.go</a> - Testable example
├── stemmer - Sub module
│   ├── <a href="nlp/stemmer/stemmer.go">stemmer.go</a>
│   └── <a href="nlp/stemmer/stemmer_test.go">stemmer_test.go</a>
├── testdata - Test data
│      └── <a href="nlp/testdata/tokenize_cases.toml">tokenize_cases.toml</a> - Test cases
└── cmd  - Executables
    └── nlpd - HTTP server
        ├── <a href="nlp/cmd/nlpd/main.go">main.go</a>
        └── <a href="nlp/cmd/nlpd/main_test.go">main_test.go</a>
</pre>