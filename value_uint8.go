package env

import "strconv"

type uint8Value uint8

func newUint8Value(val uint8, p *uint8) *uint8Value {
	*p = val

	return (*uint8Value)(p)
}

func (i *uint8Value) Set(val string) error {
	v, err := strconv.ParseInt(val, 0, 8)

	*i = uint8Value(v)

	return err
}

func (*uint8Value) Type() string { return "uint8" }

func (i *uint8Value) String() string { return strconv.FormatUint(uint64(*i), 10) }

// Uint8Var defines an uint8 environment variable with specified name, default value, and usage string.
// The argument p points to an uint8 variable in which to store the value of the environment variable.
func (s *EnvVarSet) Uint8Var(p *uint8, name string, value uint8, usage string) {
	s.Var(newUint8Value(value, p), name, usage)
}

// Uint8 defines an uint8 environment variable with specified name, default value, and usage string.
// The return value is the address of an uint8 variable that stores the value of the environment variable.
func (s *EnvVarSet) Uint8(name string, value uint8, usage string) *uint8 {
	p := new(uint8)

	s.Uint8Var(p, name, value, usage)

	return p
}

// Uint8Var defines a uint8 environment variable with specified name, default value, and usage string.
// The argument p points to a uint8 variable in which to store the value of the environment variable.
func Uint8Var(p *uint8, name string, value uint8, usage string) {
	Environment.Uint8Var(p, name, value, usage)
}

// Uint8 defines a uint8 environment variable with specified name, default value, and usage string.
// The return value is the address of a uint8 variable that stores the value of the environment variable.
func Uint8(name string, value uint8, usage string) *uint8 {
	return Environment.Uint8(name, value, usage)
}
