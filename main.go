// Command trimdemo reads samples of integers and floating-point numbers from
// the data directory, computes untrimmed, symmetric, and asymmetric trimmed
// means with the trimmedmean package, and prints the results in a form that
// can be checked against R's mean(x, trim = 0.05).
//
// Usage:
//
//	trimdemo                      # reads ./data/integers.csv and ./data/floats.csv
//	trimdemo -ints a.csv -floats b.csv
package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/Yinan920/trimmedmean"
)

func main() {
	intsPath := flag.String("ints", "data/integers.csv", "path to the integer sample (one value per line)")
	floatsPath := flag.String("floats", "data/floats.csv", "path to the float sample (one value per line)")
	bootstrap := flag.Int("bootstrap", 0, "if > 0, run this many bootstrap replications to compare estimator stability")
	flag.Parse()

	integers, err := readIntegers(*intsPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "cannot read integer sample:", err)
		os.Exit(1)
	}
	floats, err := readFloats(*floatsPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "cannot read float sample:", err)
		os.Exit(1)
	}

	fmt.Println("Trimmed-mean demo (compare the 0.05 column to R's mean(x, trim = 0.05))")
	report("Integer sample", integers)
	report("Float sample", floats)

	if *bootstrap > 0 {
		rng := rand.New(rand.NewSource(2025)) // fixed seed for reproducible results
		floatInts := make([]float64, len(integers))
		for i, v := range integers {
			floatInts[i] = float64(v)
		}
		runBootstrap("integer sample", floatInts, *bootstrap, rng)
		runBootstrap("float sample", floats, *bootstrap, rng)
	}
}

// report prints the untrimmed mean alongside a symmetric and an asymmetric
// trimmed mean for one sample. The generic type parameter lets the same
// function serve both the integer and the float slice.
func report[T trimmedmean.Number](label string, data []T) {
	untrimmed, err := trimmedmean.Compute(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", label, err)
		return
	}
	symmetric, _ := trimmedmean.Compute(data, 0.05)
	asymmetric, _ := trimmedmean.Compute(data, 0.05, 0.10)

	fmt.Printf("\n%s (n = %d)\n", label, len(data))
	fmt.Printf("  untrimmed mean             : %.6f\n", untrimmed)
	fmt.Printf("  symmetric trim 0.05        : %.6f   <- R mean(x, trim = 0.05)\n", symmetric)
	fmt.Printf("  asymmetric trim 0.05/0.10  : %.6f\n", asymmetric)
}

func readIntegers(path string) ([]int, error) {
	lines, err := readLines(path)
	if err != nil {
		return nil, err
	}
	values := make([]int, 0, len(lines))
	for _, line := range lines {
		value, err := strconv.Atoi(line)
		if err != nil {
			return nil, fmt.Errorf("invalid integer %q: %w", line, err)
		}
		values = append(values, value)
	}
	return values, nil
}

func readFloats(path string) ([]float64, error) {
	lines, err := readLines(path)
	if err != nil {
		return nil, err
	}
	values := make([]float64, 0, len(lines))
	for _, line := range lines {
		value, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid float %q: %w", line, err)
		}
		values = append(values, value)
	}
	return values, nil
}

// readLines returns the non-empty, whitespace-trimmed lines of a text file.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if line := strings.TrimSpace(scanner.Text()); line != "" {
			lines = append(lines, line)
		}
	}
	return lines, scanner.Err()
}
