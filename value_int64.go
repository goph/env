package env

import "strconv"

type int64Value int64

func newInt64Value(val int64, p *int64) *int64Value {
	*p = val

	return (*int64Value)(p)
}

func (i *int64Value) Set(val string) error {
	v, err := strconv.ParseInt(val, 0, 64)

	*i = int64Value(v)

	return err
}

func (*int64Value) Type() string { return "int64" }

func (i *int64Value) String() string { return strconv.FormatInt(int64(*i), 10) }

// Int64Var defines an int64 environment variable with specified name, default value, and usage string.
// The argument p points to an int64 variable in which to store the value of the environment variable.
func (s *EnvVarSet) Int64Var(p *int64, name string, value int64, usage string) {
	s.Var(newInt64Value(value, p), name, usage)
}

// Int64 defines an int64 environment variable with specified name, default value, and usage string.
// The return value is the address of an int64 variable that stores the value of the environment variable.
func (s *EnvVarSet) Int64(name string, value int64, usage string) *int64 {
	p := new(int64)

	s.Int64Var(p, name, value, usage)

	return p
}

// Int64Var defines an int64 environment variable with specified name, default value, and usage string.
// The argument p points to an int64 variable in which to store the value of the environment variable.
func Int64Var(p *int64, name string, value int64, usage string) {
	Environment.Int64Var(p, name, value, usage)
}

// Int64 defines an int64 environment variable with specified name, default value, and usage string.
// The return value is the address of an int64 variable that stores the value of the environment variable.
func Int64(name string, value int64, usage string) *int64 {
	return Environment.Int64(name, value, usage)
}
