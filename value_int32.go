package env

import "strconv"

type int32Value int32

func newInt32Value(val int32, p *int32) *int32Value {
	*p = val

	return (*int32Value)(p)
}

func (i *int32Value) Set(val string) error {
	v, err := strconv.ParseInt(val, 0, 32)

	*i = int32Value(v)

	return err
}

func (*int32Value) Type() string { return "int32" }

func (i *int32Value) String() string { return strconv.FormatInt(int64(*i), 10) }

// Int32Var defines an int32 environment variable with specified name, default value, and usage string.
// The argument p points to an int32 variable in which to store the value of the environment variable.
func (s *EnvVarSet) Int32Var(p *int32, name string, value int32, usage string) {
	s.Var(newInt32Value(value, p), name, usage)
}

// Int32 defines an int32 environment variable with specified name, default value, and usage string.
// The return value is the address of an int32 variable that stores the value of the environment variable.
func (s *EnvVarSet) Int32(name string, value int32, usage string) *int32 {
	p := new(int32)

	s.Int32Var(p, name, value, usage)

	return p
}

// Int32Var defines an int32 environment variable with specified name, default value, and usage string.
// The argument p points to an int32 variable in which to store the value of the environment variable.
func Int32Var(p *int32, name string, value int32, usage string) {
	Environment.Int32Var(p, name, value, usage)
}

// Int32 defines an int32 environment variable with specified name, default value, and usage string.
// The return value is the address of an int32 variable that stores the value of the environment variable.
func Int32(name string, value int32, usage string) *int32 {
	return Environment.Int32(name, value, usage)
}
