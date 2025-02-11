# Go Integer Division by Zero Panic

This repository demonstrates a common error in Go: an integer division by zero that leads to a panic, rather than a gracefully handled error.

The `bug.go` file contains the buggy code. The `bugSolution.go` file presents a corrected version.

**Bug:** The function `myFunc` doesn't correctly handle the case where `x` is zero, leading to a runtime panic.

**Solution:** The improved function uses proper error checking to avoid the panic.