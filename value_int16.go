package env

import "strconv"

type int16Value int16

func newInt16Value(val int16, p *int16) *int16Value {
	*p = val

	return (*int16Value)(p)
}

func (i *int16Value) Set(val string) error {
	v, err := strconv.ParseInt(val, 0, 16)

	*i = int16Value(v)

	return err
}

func (*int16Value) Type() string { return "int16" }

func (i *int16Value) String() string { return strconv.FormatInt(int64(*i), 10) }

// Int16Var defines an int16 environment variable with specified name, default value, and usage string.
// The argument p points to an int16 variable in which to store the value of the environment variable.
func (s *EnvVarSet) Int16Var(p *int16, name string, value int16, usage string) {
	s.Var(newInt16Value(value, p), name, usage)
}

// Int16 defines an int16 environment variable with specified name, default value, and usage string.
// The return value is the address of an int16 variable that stores the value of the environment variable.
func (s *EnvVarSet) Int16(name string, value int16, usage string) *int16 {
	p := new(int16)

	s.Int16Var(p, name, value, usage)

	return p
}
