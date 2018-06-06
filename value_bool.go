package env

import "strconv"

type boolValue bool

func newBoolValue(val bool, p *bool) *boolValue {
	*p = val

	return (*boolValue)(p)
}

func (b *boolValue) Set(val string) error {
	v, err := strconv.ParseBool(val)

	*b = boolValue(v)

	return err
}

func (*boolValue) Type() string { return "bool" }

func (b *boolValue) String() string { return strconv.FormatBool(bool(*b)) }

// BoolVar defines a bool environment variable with specified name, default value, and usage string.
// The argument p points to a bool variable in which to store the value of the environment variable.
func (s *EnvVarSet) BoolVar(p *bool, name string, value bool, usage string) {
	s.Var(newBoolValue(value, p), name, usage)
}

// Bool defines a bool environment variable with specified name, default value, and usage string.
// The return value is the address of a bool variable that stores the value of the environment variable.
func (s *EnvVarSet) Bool(name string, value bool, usage string) *bool {
	p := new(bool)

	s.BoolVar(p, name, value, usage)

	return p
}
