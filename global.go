package env

import "os"

// Environment is the default set of environment variables, parsed from os.Environ().
var Environment = NewEnvVarSet(ExitOnError)

// Var defines an environment variable with the specified name and usage string.
// The type and value of the variable are represented by the first argument,
// of type Value, which typically holds a user-defined implementation of Value.
// For instance, the caller could create an environment variable
// that turns a comma-separated string into a slice of strings by giving the slice the methods of Value;
// in particular, Set would decompose the comma-separated string into the slice.
func Var(value Value, name string, usage string) {
	Environment.Var(value, name, usage)
}

// Parse parses environment variables from os.Environ() according to the definitions in the EnvVarSet.
// Must be called after all variables in the EnvVarSet
// are defined and before variables are accessed by the program.
func Parse() error {
	return Environment.ParseEnviron(os.Environ())
}

// Parsed returns true if the environment variables have been parsed.
func Parsed() bool {
	return Environment.Parsed()
}
