# compare.R
# Reproduces the reference values that the Go program is checked against.
# Run from the repository root:  Rscript compare.R

integers <- scan("data/integers.csv", quiet = TRUE)
floats   <- scan("data/floats.csv",   quiet = TRUE)

cat(sprintf("integers untrimmed : %.6f\n", mean(integers)))
cat(sprintf("integers trim 0.05 : %.6f\n", mean(integers, trim = 0.05)))
cat(sprintf("floats   untrimmed : %.6f\n", mean(floats)))
cat(sprintf("floats   trim 0.05 : %.6f\n", mean(floats, trim = 0.05)))
