package env

// EnvVarSet is a set of defined environment variables.
type EnvVarSet struct {
	vars map[string]Value
}

// NewEnvVarSet returns a new, empty environment variable set.
func NewEnvVarSet() *EnvVarSet {
	return &EnvVarSet{}
}

// Parse parses environment variables according to the definitions in the EnvVarSet.
// Must be called after all variables in the EnvVarSet
// are defined and before variables are accessed by the program.
func (s *EnvVarSet) Parse(environment map[string]string) error {
	for key, value := range environment {
		if _var, ok := s.vars[key]; ok {
			err := _var.Set(value)
			if err != nil {
				return err
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
