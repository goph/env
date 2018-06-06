# Env

[![Build Status](https://img.shields.io/travis/goph/env.svg?style=flat-square)](https://travis-ci.org/goph/env)
[![Go Report Card](https://goreportcard.com/badge/github.com/goph/env?style=flat-square)](https://goreportcard.com/report/github.com/goph/env)
[![GolangCI](https://golangci.com/badges/github.com/goph/env.svg)](https://golangci.com)
[![GoDoc](http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/github.com/goph/env)

**Env provides a similar interface for environment variables as flag/pflag package.**

## Usage

The interface follows closely the [flag](https://golang.org/pkg/flag)/[pflag](https://github.com/spf13/pflag) packages.

Define environment variables using env.String(), Bool(), Int(), etc.

```go
var intVar *int = env.Int("int", 1234, "help message for int")
```

If you like, you can bind the environment variable to a variable using the Var() functions.

```go
var intVar int

func init() {
    env.IntVar(&intVar, "int", 1234, "help message for int")
}
```

Or you can create custom variables that satisfy the Value interface (with pointer receivers) and couple them to variable parsing by

```go
env.Var(&intVal, "int", "help message for int")
```

For such environment variables, the default value is just the initial value of the variable.

After all variables are defined, call

```go
env.Parse()
```

to parse the environment into the defined variables.


## Development

When all coding and testing is done, please run the test suite:

``` bash
$ go test
```

For linting we use [GolangCI](https://golangci.com/). You can run the linter locally using it's [binary version](https://github.com/golangci/golangci-lint#ci-installation):

```bash
$ golangci-lint run
```


## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.

This project is heavily inspired by [flag](https://golang.org/pkg/flag)/[pflag](https://github.com/spf13/pflag) packages.
