package env

import "strconv"

type uint32Value uint32

func newUint32Value(val uint32, p *uint32) *uint32Value {
	*p = val

	return (*uint32Value)(p)
}

func (i *uint32Value) Set(val string) error {
	v, err := strconv.ParseInt(val, 0, 32)

	*i = uint32Value(v)

	return err
}

func (*uint32Value) Type() string { return "uint32" }

func (i *uint32Value) String() string { return strconv.FormatUint(uint64(*i), 10) }

// Uint32Var defines an uint32 environment variable with specified name, default value, and usage string.
// The argument p points to an uint32 variable in which to store the value of the environment variable.
func (s *EnvVarSet) Uint32Var(p *uint32, name string, value uint32, usage string) {
	s.Var(newUint32Value(value, p), name, usage)
}

// Uint32 defines an uint32 environment variable with specified name, default value, and usage string.
// The return value is the address of an uint32 variable that stores the value of the environment variable.
func (s *EnvVarSet) Uint32(name string, value uint32, usage string) *uint32 {
	p := new(uint32)

	s.Uint32Var(p, name, value, usage)

	return p
}

// Uint32Var defines a uint32 environment variable with specified name, default value, and usage string.
// The argument p points to a uint32 variable in which to store the value of the environment variable.
func Uint32Var(p *uint32, name string, value uint32, usage string) {
	Environment.Uint32Var(p, name, value, usage)
}

// Uint32 defines a uint32 environment variable with specified name, default value, and usage string.
// The return value is the address of a uint32 variable that stores the value of the environment variable.
func Uint32(name string, value uint32, usage string) *uint32 {
	return Environment.Uint32(name, value, usage)
}
