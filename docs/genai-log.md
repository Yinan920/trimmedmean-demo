# GenAI Tools

**Author:** Yinan He
**Course:** Week 9 Assignment — Create a Go Package
**Tool used:** Anthropic's Claude (chat assistant)

This section documents how I used a generative-AI assistant on this assignment,
what I did and verified myself, and what I learned in the process. I treated the
assistant as a tutor and pair-programmer, not as a black box: every piece of
code was run, checked, and understood before it became part of my submission.

## How I used the assistant

I started from the assignment requirements and my own setup (MacBook M4, VS Code)
and asked the assistant to walk me through the project step by step. Over the
conversation I used it to:

- Draft an initial version of the `trimmedmean` package, the demo program, the
  unit tests, and the two README files.
- Generate the synthetic sample data (100 integers and 100 floats from a
  distribution with outliers) so that trimming would have a visible effect.
- Explain Go concepts I was less familiar with: generics with a custom `Number`
  constraint, Go modules, the local `replace` directive, and cross-compiling.

I then drove the rest of the workflow myself and made the assistant explain
anything I did not understand.

## What I did and verified myself

- **Set up the toolchain and project** on my own machine: installed Go, placed
  the two repositories as sibling folders so the demo could resolve the package
  locally, and opened the project in VS Code.
- **Ran the tests myself.** `go test -v ./...` in the package returned all
  PASS, including the test that checks the result against R's floor rule.
- **Ran the demo myself** with `go run .` and reproduced the reported values.
- **Verified the Go results against R.** I confirmed that the symmetric
  trimmed means (integers 54.122222, floats 105.168190) match what R's
  `mean(x, trim = 0.05)` produces on the same data files. I understood *why*
  they match: R removes `floor(n * trim)` observations from each end, and the
  package uses the same rule, so the comparison is exact rather than approximate.
- **Reviewed the code before accepting it.** I read through the package to make
  sure the logic was clear: sort a copy of the data, floor the trim counts,
  average the middle slice, and return sentinel errors for invalid input.
- **Made the project-specific decisions:** confirming the two-repository
  structure required by the assignment, replacing the placeholder module path
  with my own GitHub username, and deciding what belonged in each README.

## What I learned

- **The trimmed mean as a robust estimator.** Trimming a small proportion from
  each end removes the influence of outliers. The optional bootstrap study made
  this concrete: on the outlier-heavy samples, the trimmed mean had a smaller
  bootstrap standard error than the ordinary mean, so it is the more stable
  estimate of central tendency.
- **Matching R exactly.** The key detail is the floor rule; without it the Go
  result would drift from R whenever `n * trim` is not a whole number.
- **Idiomatic Go.** Using generics lets one `Compute` function serve both
  integer and float samples; using sentinel errors with `errors.Is` lets callers
  test for specific failure causes; and a local `replace` directive lets two
  separate repositories be developed together before publishing.

## Verification summary

| Check                          | Command                         | Result            |
| ------------------------------ | ------------------------------- | ----------------- |
| Unit tests                     | `go test ./...`                 | all PASS          |
| Static analysis                | `go vet ./...`                  | clean             |
| Demo runs                      | `go run .`                      | expected output   |
| Matches R `mean(x, trim=0.05)` | `Rscript compare.R`             | identical values  |

## Conversation log

The full transcript of my conversation with the assistant is available here:

<!-- Paste the exported chat text below, or replace this line with a link. -->

