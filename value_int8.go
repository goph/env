package env

import "strconv"

type int8Value int8

func newInt8Value(val int8, p *int8) *int8Value {
	*p = val

	return (*int8Value)(p)
}

func (i *int8Value) Set(val string) error {
	v, err := strconv.ParseInt(val, 0, 8)

	*i = int8Value(v)

	return err
}

func (*int8Value) Type() string { return "int8" }

func (i *int8Value) String() string { return strconv.FormatInt(int64(*i), 10) }

// Int8Var defines an int8 environment variable with specified name, default value, and usage string.
// The argument p points to an int8 variable in which to store the value of the environment variable.
func (s *EnvVarSet) Int8Var(p *int8, name string, value int8, usage string) {
	s.Var(newInt8Value(value, p), name, usage)
}

// Int8 defines an int8 environment variable with specified name, default value, and usage string.
// The return value is the address of an int8 variable that stores the value of the environment variable.
func (s *EnvVarSet) Int8(name string, value int8, usage string) *int8 {
	p := new(int8)

	s.Int8Var(p, name, value, usage)

	return p
}
