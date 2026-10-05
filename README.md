# e2e-go-fixture

A minimal Go module kanman uses to check team environments end to end
(plan 3.4.29): the Go toolchain exists only in the team image
(`golang:1.25`), never on the machine that runs the coding agent.

`make test` runs `go vet` and `go test`.
