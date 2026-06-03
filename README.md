# trimmedmean-demo

An example program that imports the
[`trimmedmean`](https://github.com/Yinan920/trimmedmean) package, computes
trimmed means on samples of 100 integers and 100 floats, and verifies the
results against R's `mean(x, trim = 0.05)`.


## What is here

| Path                  | Role                                                        |
| --------------------- | ----------------------------------------------------------- |
| `main.go`             | reads the samples, prints untrimmed / symmetric / asymmetric trimmed means |
| `bootstrap.go`        | optional bootstrap study of estimator stability             |
| `data/integers.csv`   | 100 integers (one per line) from a contaminated distribution |
| `data/floats.csv`     | 100 floats (one per line) from a contaminated normal mixture |
| `compare.R`           | reproduces the reference values in R                        |
| `bin/`                | pre-built executables (macOS arm64, Windows `.exe`)         |

The samples are deliberately drawn from mixtures with a heavy right tail, so a
few large outliers pull the ordinary mean upward while the trimmed mean stays
near the bulk of the data.

## Getting the package

This program depends on the `trimmedmean` package. Once that package is pushed
to GitHub, fetch it with:

```bash
go get github.com/Yinan920/trimmedmean
```

To develop both repositories side by side **before** publishing, clone them
into sibling folders and add a local `replace` directive to this module's
`go.mod`:

```
replace github.com/Yinan920/trimmedmean => ../trimmedmean
```

Remove that line (`go mod edit -dropreplace=github.com/Yinan920/trimmedmean`)
once the package is public.

## Running

```bash
go run .                      # uses ./data/integers.csv and ./data/floats.csv
go run . -ints x.csv -floats y.csv   # custom data files
go run . -bootstrap 2000      # also run the optional bootstrap study
```

Or run a pre-built executable:

```bash
./bin/trimdemo-macos-arm64    # macOS (Apple Silicon)
bin\trimdemo.exe              # Windows
```

The program validates its input: a missing file or a non-numeric line prints a
clear message and exits with a non-zero status instead of crashing.

## Building an executable

```bash
go build -o bin/trimdemo .                                   # for the current machine
GOOS=darwin  GOARCH=arm64 go build -o bin/trimdemo-macos .   # macOS Apple Silicon
GOOS=windows GOARCH=amd64 go build -o bin/trimdemo.exe .      # Windows 64-bit
```

## Results

Program output:

```
Integer sample (n = 100)
  untrimmed mean             : 67.810000
  symmetric trim 0.05        : 54.122222   <- R mean(x, trim = 0.05)
  asymmetric trim 0.05/0.10  : 50.905882

Float sample (n = 100)
  untrimmed mean             : 113.825629
  symmetric trim 0.05        : 105.168190   <- R mean(x, trim = 0.05)
  asymmetric trim 0.05/0.10  : 100.190718
```

### Comparison with R

Running `Rscript compare.R` on the same data files gives identical values:

| Statistic                 | Go (this package) | R                 |
| ------------------------- | ----------------- | ----------------- |
| integers, untrimmed       | 67.810000         | 67.810000         |
| integers, `trim = 0.05`   | **54.122222**     | **54.122222**     |
| floats, untrimmed         | 113.825629        | 113.825629        |
| floats, `trim = 0.05`     | **105.168190**    | **105.168190**    |

The match is exact because the package removes `floor(n * trim)` observations
from each end, which is precisely R's rule in `mean.default`. On both samples
the symmetric trimmed mean is well below the ordinary mean, confirming that
trimming has removed the influence of the large outliers.

## Optional bootstrap study

`go run . -bootstrap 2000` resamples each dataset with replacement and reports
the bootstrap standard error of three estimators. Lower is more stable:

```
Bootstrap standard errors for integer sample (2000 replications)
  mean         : 6.764931
  trimmed 0.05 : 5.121360
  median       : 0.783511
```

The trimmed mean has a smaller standard error than the ordinary mean, so it is
the more stable estimator of central tendency for this outlier-prone data. (The
median is even more stable here but discards far more information; the trimmed
mean is a middle ground, as discussed in Efron and Tibshirani, 1993.)

## GenAI Tools

This project was developed with assistance from Anthropic's Claude.  A log of the conversation is included in
`docs/genai-log.md`.

## References

Efron, B. and Tibshirani, R. J. (1993). *An Introduction to the Bootstrap.*
Chapman and Hall/CRC.


