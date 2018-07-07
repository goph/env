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

// VarE is like Var, but returns the created EnvVar.
func VarE(value Value, name string, usage string) *EnvVar {
	return Environment.VarE(value, name, usage)
}

// Parse parses environment variables from os.Environ() according to the definitions in the EnvVarSet.
// Must be called after all variables in the EnvVarSet
// are defined and before variables are accessed by the program.
func Parse() {
	Environment.ParseEnviron(os.Environ()) // nolint:errcheck
}

// Parsed returns true if the environment variables have been parsed.
func Parsed() bool {
	return Environment.Parsed()
}

// PrintDefaults prints to standard error the default values of all defined environment variables.
func PrintDefaults() {
	Environment.PrintDefaults()
}

// VisitAll visits the environment variables in lexicographical order or
// in primordial order if f.SortVars is false, calling fn for each.
// It visits all variables, even those not set.
func VisitAll(fn func(*EnvVar)) {
	Environment.VisitAll(fn)
}

// Lookup returns the EnvVar structure of the named environment variable, returning nil if none exists.
func Lookup(name string) *EnvVar {
	return Environment.Lookup(name)
}
