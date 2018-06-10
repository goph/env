package env

import "strconv"

type intValue int

func newIntValue(val int, p *int) *intValue {
	*p = val

	return (*intValue)(p)
}

func (i *intValue) Set(val string) error {
	v, err := strconv.ParseInt(val, 0, 64)

	*i = intValue(v)

	return err
}

func (*intValue) Type() string { return "int" }

func (i *intValue) String() string { return strconv.FormatInt(int64(*i), 10) }

// IntVar defines an int environment variable with specified name, default value, and usage string.
// The argument p points to an int variable in which to store the value of the environment variable.
func (s *EnvVarSet) IntVar(p *int, name string, value int, usage string) {
	s.Var(newIntValue(value, p), name, usage)
}

// Int defines an int environment variable with specified name, default value, and usage string.
// The return value is the address of an int variable that stores the value of the environment variable.
func (s *EnvVarSet) Int(name string, value int, usage string) *int {
	p := new(int)

	s.IntVar(p, name, value, usage)

	return p
}

// IntVar defines an int environment variable with specified name, default value, and usage string.
// The argument p points to an int variable in which to store the value of the environment variable.
func IntVar(p *int, name string, value int, usage string) {
	Environment.IntVar(p, name, value, usage)
}

// Int defines an int environment variable with specified name, default value, and usage string.
// The return value is the address of an int variable that stores the value of the environment variable.
func Int(name string, value int, usage string) *int {
	return Environment.Int(name, value, usage)
}
