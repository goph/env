package env

import "strings"

// NormalizedName is an environment variable name that has been normalized according to rules
// for the EnvVarSet (e.g. making variable names upper case, adding prefixes, etc).
type NormalizedName string

// NormalizeFunc normalizes an environment variable name to the form it can be found in the environment.
type NormalizeFunc func(s *EnvVarSet, name string) NormalizedName

var defaultReplacer = strings.NewReplacer("-", "_", ".", "_")

// DefaultNormalizeFunc formats variable names to upper case and replaces "-." chars with "_".
func DefaultNormalizeFunc(_ *EnvVarSet, name string) NormalizedName {
	return NormalizedName(defaultReplacer.Replace(strings.ToUpper(name)))
}

// GetNormalizeFunc returns the previously set NormalizeFunc.
// If not set, it returns the default normalization function.
func (s *EnvVarSet) GetNormalizeFunc() NormalizeFunc {
	if s.normalizeNameFunc == nil {
		return DefaultNormalizeFunc
	}

	return s.normalizeNameFunc
}

func (s *EnvVarSet) normalizeVarName(name string) NormalizedName {
	fn := s.GetNormalizeFunc()

	return fn(s, name)
}

// SetNormalizeFunc allows you to add a function which can translate variable names.
// Environment variables added to the EnvVarSet will be translated, incoming environment variables
// will be matched against these translated names.
func (s *EnvVarSet) SetNormalizeFunc(fn NormalizeFunc) {
	s.normalizeNameFunc = fn
	s.sortedVars = s.sortedVars[:0]

	for name, v := range s.vars {
		normalizedName := s.normalizeVarName(v.Name)

		delete(s.vars, name)

		s.vars[normalizedName] = v
	}
}
