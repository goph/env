package env

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

// ErrorHandling defines how to handle env var parsing errors.
type ErrorHandling int

const (
	// ContinueOnError will return an err from Parse() if an error is found
	ContinueOnError ErrorHandling = iota

	// ExitOnError will call os.Exit(2) if an error is found
	ExitOnError

	// PanicOnError will panic() if an error is found
	PanicOnError
)

// EnvVarSet is a set of defined environment variables.
type EnvVarSet struct {
	// SortVars is used to indicate, if user wants to have sorted environment variables in
	// help/usage messages.
	SortVars bool

	parsed            bool
	vars              map[NormalizedName]*EnvVar
	orderedVars       []*EnvVar
	sortedVars        []*EnvVar
	errorHandling     ErrorHandling
	output            io.Writer // nil means stderr; use out() accessor
	normalizeNameFunc NormalizeFunc
}

// EnvVar represents the state of an environment variable.
type EnvVar struct {
	// Name of the environment variable
	Name string

	// Usage message
	Usage string

	// Value as set
	Value Value

	// DefaultValue is shown in the usage message
	DefaultValue string
}

// NewEnvVarSet returns a new, empty environment variable set with the specified
// error handling property and SortFlags set to true.
func NewEnvVarSet(errorHandling ErrorHandling) *EnvVarSet {
	return &EnvVarSet{
		SortVars: true,

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
	s.VarE(value, name, usage)
}

// VarE is like Var, but returns the created EnvVar.
func (s *EnvVarSet) VarE(value Value, name string, usage string) *EnvVar {
	envVar := &EnvVar{
		Name:         name,
		Usage:        usage,
		Value:        value,
		DefaultValue: value.String(),
	}

	s.AddEnvVar(envVar)

	return envVar
}

// AddEnvVar will add the environment variable to the EnvVarSet.
func (s *EnvVarSet) AddEnvVar(envVar *EnvVar) {
	if s.vars == nil {
		s.vars = make(map[NormalizedName]*EnvVar)
	}

	name := s.normalizeVarName(envVar.Name)

	_, alreadyThere := s.vars[name]
	if alreadyThere {
		msg := fmt.Sprintf("%s environment variable redefined: %s", name, envVar.Name)

		fmt.Fprintln(s.out(), msg)

		panic(msg) // Happens only if environment variables are declared with identical names
	}

	s.vars[name] = envVar
	s.orderedVars = append(s.orderedVars, envVar)
}

func (s *EnvVarSet) out() io.Writer {
	if s.output == nil {
		return os.Stderr
	}

	return s.output
}

// SetOutput sets the destination for usage and error messages.
// If output is nil, os.Stderr is used.
func (s *EnvVarSet) SetOutput(output io.Writer) {
	s.output = output
}

// VisitAll visits the environment variables in lexicographical order or
// in primordial order if f.SortVars is false, calling fn for each.
// It visits all variables, even those not set.
func (s *EnvVarSet) VisitAll(fn func(*EnvVar)) {
	if len(s.vars) == 0 {
		return
	}

	var envVars []*EnvVar
	if s.SortVars {
		if len(s.vars) != len(s.sortedVars) {
			s.sortedVars = sortVars(s.vars)
		}
		envVars = s.sortedVars
	} else {
		envVars = s.orderedVars
	}

	for _, envVar := range envVars {
		fn(envVar)
	}
}

// sortVars returns the environment variables as a slice in lexicographical sorted order.
func sortVars(envVars map[NormalizedName]*EnvVar) []*EnvVar {
	list := make(sort.StringSlice, len(envVars))
	i := 0

	for k := range envVars {
		list[i] = string(k)
		i++
	}

	list.Sort()

	result := make([]*EnvVar, len(list))

	for i, name := range list {
		result[i] = envVars[NormalizedName(name)]
	}

	return result
}

// HasEnvVars returns a bool to indicate if the EnvVarSet has any environment variables defined.
func (s *EnvVarSet) HasEnvVars() bool {
	return len(s.vars) > 0
}

// Lookup returns the EnvVar structure of the named environment variable, returning nil if none exists.
func (s *EnvVarSet) Lookup(name string) *EnvVar {
	return s.vars[s.normalizeVarName(name)]
}

// Parse parses environment variables according to the definitions in the EnvVarSet.
// Must be called after all variables in the EnvVarSet
// are defined and before variables are accessed by the program.
func (s *EnvVarSet) Parse(environment map[string]string) error {
	s.parsed = true

	for name, value := range environment {
		if ev, ok := s.vars[NormalizedName(name)]; ok {
			err := ev.Value.Set(value)
			if err != nil {
				switch s.errorHandling {
				case ContinueOnError:
					return err
				case ExitOnError:
					fmt.Fprintln(s.out(), err)
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
