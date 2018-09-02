package env

import "strconv"

type uint16Value uint16

func newUint16Value(val uint16, p *uint16) *uint16Value {
	*p = val

	return (*uint16Value)(p)
}

func (i *uint16Value) Set(val string) error {
	v, err := strconv.ParseInt(val, 0, 16)

	*i = uint16Value(v)

	return err
}

func (*uint16Value) Type() string { return "uint16" }

func (i *uint16Value) String() string { return strconv.FormatUint(uint64(*i), 10) }

// Uint16Var defines a uint16 environment variable with specified name, default value, and usage string.
// The argument p points to a uint16 variable in which to store the value of the environment variable.
func (s *EnvVarSet) Uint16Var(p *uint16, name string, value uint16, usage string) {
	s.Var(newUint16Value(value, p), name, usage)
}

// Uint16 defines a uint16 environment variable with specified name, default value, and usage string.
// The return value is the address of a uint16 variable that stores the value of the environment variable.
func (s *EnvVarSet) Uint16(name string, value uint16, usage string) *uint16 {
	p := new(uint16)

	s.Uint16Var(p, name, value, usage)

	return p
}

// Uint16Var defines a uint16 environment variable with specified name, default value, and usage string.
// The argument p points to a uint16 variable in which to store the value of the environment variable.
func Uint16Var(p *uint16, name string, value uint16, usage string) {
	Environment.Uint16Var(p, name, value, usage)
}

// Uint16 defines a uint16 environment variable with specified name, default value, and usage string.
// The return value is the address of a uint16 variable that stores the value of the environment variable.
func Uint16(name string, value uint16, usage string) *uint16 {
	return Environment.Uint16(name, value, usage)
}
