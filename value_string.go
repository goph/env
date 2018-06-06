package env

type stringValue string

func newStringValue(val string, p *string) *stringValue {
	*p = val

	return (*stringValue)(p)
}

func (s *stringValue) Set(val string) error {
	*s = stringValue(val)

	return nil
}

func (*stringValue) Type() string { return "string" }

func (s *stringValue) String() string { return string(*s) }

// StringVar defines a string environment variable with specified name, default value, and usage string.
// The argument p points to a string variable in which to store the value of the environment variable.
func (s *EnvVarSet) StringVar(p *string, name string, value string, usage string) {
	s.Var(newStringValue(value, p), name, usage)
}

// String defines a string environment variable with specified name, default value, and usage string.
// The return value is the address of a string variable that stores the value of the environment variable.
func (s *EnvVarSet) String(name string, value string, usage string) *string {
	p := new(string)

	s.StringVar(p, name, value, usage)

	return p
}

// StringVar defines a string environment variable with specified name, default value, and usage string.
// The argument p points to a string variable in which to store the value of the environment variable.
func StringVar(p *string, name string, value string, usage string) {
	Environment.StringVar(p, name, value, usage)
}

// String defines a string environment variable with specified name, default value, and usage string.
// The return value is the address of a string variable that stores the value of the environment variable.
func String(name string, value string, usage string) *string {
	return Environment.String(name, value, usage)
}
