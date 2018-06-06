package env

type {{type}}Value {{type}}

func new{{Type}}Value(val {{type}}, p *{{type}}) *{{type}}Value {
	*p = val

	return (*{{type}}Value)(p)
}

func ({{t}} *{{type}}Value) Set(val string) error {
	*{{t}} = {{type}}Value(val)

	return nil
}

func (*{{type}}Value) Type() string	 { return "{{type}}" }

func ({{t}} *{{type}}Value) String() string { return {{type}}(*{{t}}) }

// {{Type}}Var defines a {{type}} environment variable with specified name, default value, and usage string.
// The argument p points to a {{type}} variable in which to store the value of the environment variable.
func (s *EnvVarSet) {{Type}}Var(p *{{type}}, name string, value {{type}}, usage string) {
	s.Var(new{{Type}}Value(value, p), name, usage)
}

// {{Type}} defines a {{type}} environment variable with specified name, default value, and usage string.
// The return value is the address of a {{type}} variable that stores the value of the environment variable.
func (s *EnvVarSet) {{Type}}(name string, value {{type}}, usage string) *{{type}} {
	p := new({{type}})

	s.{{Type}}Var(p, name, value, usage)

	return p
}

// {{Type}}Var defines a {{type}} environment variable with specified name, default value, and usage string.
// The argument p points to a {{type}} variable in which to store the value of the environment variable.
func {{Type}}Var(p *{{type}}, name string, value {{type}}, usage string) {
	Environment.{{Type}}Var(p, name, value, usage)
}

// {{Type}} defines a {{type}} environment variable with specified name, default value, and usage string.
// The return value is the address of a {{type}} variable that stores the value of the environment variable.
func {{Type}}(name string, value {{type}}, usage string) *{{type}} {
	return Environment.{{Type}}(name, value, usage)
}
