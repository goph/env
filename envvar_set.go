package env

import (
	"fmt"
	"os"
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

// EnvVarSet is a set of defined environment variables.
type EnvVarSet struct {
	vars          map[string]Value
	errorHandling ErrorHandling
}

// NewEnvVarSet returns a new, empty environment variable set.
func NewEnvVarSet(errorHandling ErrorHandling) *EnvVarSet {
	return &EnvVarSet{
		errorHandling: errorHandling,
	}
}

// Parse parses environment variables according to the definitions in the EnvVarSet.
// Must be called after all variables in the EnvVarSet
// are defined and before variables are accessed by the program.
func (s *EnvVarSet) Parse(environment map[string]string) error {
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

func (s *EnvVarSet) Var(value Value, name string, usage string) {
	if s.vars == nil {
		s.vars = make(map[string]Value)
	}

	s.vars[name] = value
}
