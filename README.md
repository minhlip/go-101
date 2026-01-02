# Go 101

Learn Go programming language from basics with hands-on lessons.

## Table of Contents

- [Introduction](#introduction)
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
- [Project Structure](#project-structure)
- [Lessons](#lessons)
- [Configuration](#configuration)
- [Code Formatting](#code-formatting)

## Introduction

This project contains basic Go programming lessons, organized as separate lesson functions in the `lesson` package. Each lesson demonstrates fundamental Go concepts with practical examples.

## Requirements

- Go 1.24.2 or higher
- Terminal/Command line

## Installation

1. Clone the repository or download the source code
2. Navigate to the project directory:
```bash
cd go-101
```

3. Initialize Go module (if not already done):
```bash
go mod init go-101
```

4. Download dependencies (if any):
```bash
go mod tidy
```

## Usage

### Run all enabled lessons

```bash
go run main.go
```

### Run a specific file

```bash
go run lesson/hello-world.go
```

### Build binary

```bash
go build
./go-101  # or ./main on Windows
```

## Project Structure

```
go-101/
├── main.go              # Entry point, runs lessons
├── go.mod              # Go module file
├── README.md           # This file
└── lesson/             # Directory containing lessons
    ├── hello-world.go  # Lesson 1: Hello World
    ├── values.go       # Lesson 2: Values
    └── variables.go    # Lesson 3: Variables
```

## Lessons

### Lesson 1: Hello World
- **File**: `lesson/hello-world.go`
- **Description**: Print "Hello World" to the console
- **Function**: `lesson.HelloWorld()`
- **Output**: `Hello World`

### Lesson 2: Values
- **File**: `lesson/values.go`
- **Description**: Learn about basic value types in Go (strings, integers, floats, booleans)
- **Function**: `lesson.Values()`
- **Output**:
  - `golang`
  - `1+1 = 2`
  - `7.0/3.0 = 2.3333333333333335`
  - `false`
  - `true`
  - `false`

### Lesson 3: Variables
- **File**: `lesson/variables.go`
- **Description**: Learn about variable declaration and usage in Go
- **Function**: `lesson.Variables()`
- **Topics Covered**:
  - Variable declaration with `var`
  - Multiple variable declaration
  - Type inference
  - Zero values
  - Short variable declaration with `:=`
- **Output**:
  - `initial`
  - `1 2`
  - `true`
  - `0`
  - `""`
  - `apple`

## Configuration

In `main.go`, you can enable or disable lessons by setting `Enabled: true/false`:

```go
lessons := []Lesson{
    {
        "Hello World",
        lesson.HelloWorld,
        true,  // Set true to run, false to skip
    },
    {
        "Values",
        lesson.Values,
        false,  // This lesson will be skipped
    },
    // ...
}
```

The lesson index will automatically adjust based on enabled lessons only.

## Code Formatting

This project uses `gofmt` for code formatting:

```bash
# Format a single file
gofmt -w main.go

# Format entire project
gofmt -w .

# Or use go fmt
go fmt ./...
```

For automatic import organization, you can use `goimports`:

```bash
goimports -w main.go
```

## References

- [Go Official Documentation](https://go.dev/doc/)
- [A Tour of Go](https://go.dev/tour/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go Blog](https://go.dev/blog/)

## License

MIT

---

Happy coding!

