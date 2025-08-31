# godsa — Data Structures & Algorithms in Go

A collection of classic data structures and algorithms implemented in Go, organized as small, focused examples for learning and experimentation.

This repository is intended as a learning resource and playground for implementing and exploring algorithms and data structures using idiomatic Go.

## Highlights

- Focused examples implemented in plain Go (no external dependencies).
- Easy-to-run examples from the repository root using `go run` or `go build`.
- Clear, tiny packages per topic (currently: `array`, `recursion`).

## Repository layout

- `main.go` — program entry. It prints a greeting and invokes example package starters.
- `go.mod` — module definition (module: `github.com/This-Is-Prince/godsa`).
- `array/` — array-focused examples and small ADT demonstrations.
	- `array.go`, `declarations.go`, `arrayadt.go`, `2darray.go`
- `recursion/` — recursion examples and related demonstrations.
	- `recursion.go`, `factorial.go`, `power.go`, `sum_of_natural_number.go`, etc.

See the source files for specific function names and short comments.

## Quickstart

Prerequisites

- Go 1.20+ installed (module uses Go 1.24 in `go.mod`).

Run the project (from repository root):

```zsh
go run main.go
```

That will print a greeting and then call the package starter functions. By default `main.go` controls which examples run by passing a boolean flag to each package's `Start` function (e.g. `recursion.Start(false)` and `array.Start(true)`).

To run only a single package for quick experimentation, you can run its files directly. For example, to run the `array` package's examples:

```zsh
go run ./array
```

Or build the project and run the binary:

```zsh
go build -o godsa
./godsa
```

## How the examples are organized

- Each top-level folder (like `array` and `recursion`) exposes a `Start(run bool)` function you can call from `main.go` to control execution.
- Individual files implement one concept each (e.g. factorial, power, 2D arrays). Look into the file names for quick orientation.

## Contributing

- Add small, well-documented examples.
- Keep functions small and focused.
- If you add new packages, include a short README or comment describing the example and expected output.

## Notes

- No tests are included yet; contributions that add tests are welcome.
- No license file is present in the repository. If you want this to be open-source, add a `LICENSE` (for example, MIT).

## Contact

If you want to reach out, the module path is `github.com/This-Is-Prince/godsa` — open an issue or PR on that repository.

Happy learning!

