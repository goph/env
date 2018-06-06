package env

import "strconv"

type uint64Value uint64

func newUint64Value(val uint64, p *uint64) *uint64Value {
	*p = val

	return (*uint64Value)(p)
}

func (i *uint64Value) Set(val string) error {
	v, err := strconv.ParseInt(val, 0, 64)

	*i = uint64Value(v)

	return err
}

func (*uint64Value) Type() string { return "uint64" }

func (i *uint64Value) String() string { return strconv.FormatUint(uint64(*i), 10) }

// Uint64Var defines an uint64 environment variable with specified name, default value, and usage string.
// The argument p points to an uint64 variable in which to store the value of the environment variable.
func (s *EnvVarSet) Uint64Var(p *uint64, name string, value uint64, usage string) {
	s.Var(newUint64Value(value, p), name, usage)
}

// Uint64 defines an uint64 environment variable with specified name, default value, and usage string.
// The return value is the address of an uint64 variable that stores the value of the environment variable.
func (s *EnvVarSet) Uint64(name string, value uint64, usage string) *uint64 {
	p := new(uint64)

	s.Uint64Var(p, name, value, usage)

	return p
}

// Uint64Var defines a uint64 environment variable with specified name, default value, and usage string.
// The argument p points to a uint64 variable in which to store the value of the environment variable.
func Uint64Var(p *uint64, name string, value uint64, usage string) {
	Environment.Uint64Var(p, name, value, usage)
}

// Uint64 defines a uint64 environment variable with specified name, default value, and usage string.
// The return value is the address of a uint64 variable that stores the value of the environment variable.
func Uint64(name string, value uint64, usage string) *uint64 {
	return Environment.Uint64(name, value, usage)
}
