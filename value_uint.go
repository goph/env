package env

import "strconv"

type uintValue uint

func newUintValue(val uint, p *uint) *uintValue {
	*p = val

	return (*uintValue)(p)
}

func (i *uintValue) Set(val string) error {
	v, err := strconv.ParseInt(val, 0, 64)

	*i = uintValue(v)

	return err
}

func (*uintValue) Type() string { return "uint" }

func (i *uintValue) String() string { return strconv.FormatUint(uint64(*i), 10) }

// UintVar defines an uint environment variable with specified name, default value, and usage string.
// The argument p points to an uint variable in which to store the value of the environment variable.
func (s *EnvVarSet) UintVar(p *uint, name string, value uint, usage string) {
	s.Var(newUintValue(value, p), name, usage)
}

// Uint defines an uint environment variable with specified name, default value, and usage string.
// The return value is the address of an uint variable that stores the value of the environment variable.
func (s *EnvVarSet) Uint(name string, value uint, usage string) *uint {
	p := new(uint)

	s.UintVar(p, name, value, usage)

	return p
}

// UintVar defines a uint environment variable with specified name, default value, and usage string.
// The argument p points to a uint variable in which to store the value of the environment variable.
func UintVar(p *uint, name string, value uint, usage string) {
	Environment.UintVar(p, name, value, usage)
}

// Uint defines a uint environment variable with specified name, default value, and usage string.
// The return value is the address of a uint variable that stores the value of the environment variable.
func Uint(name string, value uint, usage string) *uint {
	return Environment.Uint(name, value, usage)
}
