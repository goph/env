package env

import "strconv"

type float64Value float64

func newFloat64Value(val float64, p *float64) *float64Value {
	*p = val

	return (*float64Value)(p)
}

func (f *float64Value) Set(val string) error {
	v, err := strconv.ParseFloat(val, 64)

	*f = float64Value(v)

	return err
}

func (*float64Value) Type() string { return "float64" }

func (f *float64Value) String() string { return strconv.FormatFloat(float64(*f), 'g', -1, 64) }

// Float64Var defines a float64 environment variable with specified name, default value, and usage string.
// The argument p points to a float64 variable in which to store the value of the environment variable.
func (s *EnvVarSet) Float64Var(p *float64, name string, value float64, usage string) {
	s.Var(newFloat64Value(value, p), name, usage)
}

// Float64 defines a float64 environment variable with specified name, default value, and usage string.
// The return value is the address of a float64 variable that stores the value of the environment variable.
func (s *EnvVarSet) Float64(name string, value float64, usage string) *float64 {
	p := new(float64)

	s.Float64Var(p, name, value, usage)

	return p
}
