package env

import "strconv"

type float32Value float32

func newFloat32Value(val float32, p *float32) *float32Value {
	*p = val

	return (*float32Value)(p)
}

func (f *float32Value) Set(val string) error {
	v, err := strconv.ParseFloat(val, 32)

	*f = float32Value(v)

	return err
}

func (*float32Value) Type() string { return "float32" }

func (f *float32Value) String() string { return strconv.FormatFloat(float64(*f), 'g', -1, 32) }

// Float32Var defines a float32 environment variable with specified name, default value, and usage string.
// The argument p points to a float32 variable in which to store the value of the environment variable.
func (s *EnvVarSet) Float32Var(p *float32, name string, value float32, usage string) {
	s.Var(newFloat32Value(value, p), name, usage)
}

// Float32 defines a float32 environment variable with specified name, default value, and usage string.
// The return value is the address of a float32 variable that stores the value of the environment variable.
func (s *EnvVarSet) Float32(name string, value float32, usage string) *float32 {
	p := new(float32)

	s.Float32Var(p, name, value, usage)

	return p
}

// Float32Var defines a float32 environment variable with specified name, default value, and usage string.
// The argument p points to a float32 variable in which to store the value of the environment variable.
func Float32Var(p *float32, name string, value float32, usage string) {
	Environment.Float32Var(p, name, value, usage)
}

// Float32 defines a float32 environment variable with specified name, default value, and usage string.
// The return value is the address of a float32 variable that stores the value of the environment variable.
func Float32(name string, value float32, usage string) *float32 {
	return Environment.Float32(name, value, usage)
}
