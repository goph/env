package env

import (
	"fmt"
	"os"
	"strings"
)

// ErrorHandling defines how to handle env var parsing errors.
type ErrorHandling int

const (
	// ContinueOnError will return an err from Parse() if an error is found
	ContinueOnError ErrorHandling = iota

	// ExitOnError will call os.Exit(2) if an error is found when parsing
	ExitOnError

	// PanicOnError will panic() if an error is found when parsing flags
	PanicOnError
)

// Value is the interface to the dynamic value stored in an environment variable.
// (The default value is represented as a string.)
type Value interface {
	String() string
	Set(string) error
	Type() string
}

// EnvVarSet is a set of defined environment variables.
type EnvVarSet struct {
	parsed        bool
	vars          map[string]Value
	errorHandling ErrorHandling
}

// NewEnvVarSet returns a new, empty environment variable set.
func NewEnvVarSet(errorHandling ErrorHandling) *EnvVarSet {
	return &EnvVarSet{
		errorHandling: errorHandling,
	}
}

// Var defines an environment variable with the specified name and usage string.
// The type and value of the variable are represented by the first argument,
// of type Value, which typically holds a user-defined implementation of Value.
// For instance, the caller could create an environment variable
// that turns a comma-separated string into a slice of strings by giving the slice the methods of Value;
// in particular, Set would decompose the comma-separated string into the slice.
func (s *EnvVarSet) Var(value Value, name string, usage string) {
	if s.vars == nil {
		s.vars = make(map[string]Value)
	}

	s.vars[name] = value
}

// Parse parses environment variables according to the definitions in the EnvVarSet.
// Must be called after all variables in the EnvVarSet
// are defined and before variables are accessed by the program.
func (s *EnvVarSet) Parse(environment map[string]string) error {
	s.parsed = true

	for key, value := range environment {
		if _var, ok := s.vars[key]; ok {
			err := _var.Set(value)
			if err != nil {
				switch s.errorHandling {
				case ContinueOnError:
					return err
				case ExitOnError:
					fmt.Println(err)
					os.Exit(2)
				case PanicOnError:
					panic(err)
				}
			}
		}
	}

	return nil
}

// ParseEnviron accepts a list of environment variables in the "key=value" format
// returned by os.Environ(), transforms it into a string map and calls Parse.
func (s *EnvVarSet) ParseEnviron(environ []string) error {
	env := map[string]string{}

	for _, value := range environ {
		v := strings.SplitN(value, "=", 2)

		env[v[0]] = v[1]
	}

	return s.Parse(env)
}

// Parsed reports whether Parse has been called on EnvVarSet.
func (s *EnvVarSet) Parsed() bool {
	return s.parsed
}

