package env

// EnvVarSet is a set of defined environment variables.
type EnvVarSet struct {
}

// NewEnvVarSet returns a new, empty environment variable set.
func NewEnvVarSet() *EnvVarSet {
	return &EnvVarSet{}
}

// Parse parses environment variables according to the definitions in the EnvVarSet.
// Must be called after all variables in the EnvVarSet
// are defined and before variables are accessed by the program.
func (s *EnvVarSet) Parse() error {
	return nil
}
