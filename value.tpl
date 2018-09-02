package env

import "strconv"

type {{ .Type }}Value {{ .Type }}

func new{{ .TypeName }}Value(val {{ .Type }}, p *{{ .Type }}) *{{ .Type }}Value {
	*p = val

	return (*{{ .Type }}Value)(p)
}

func ({{ .Receiver }} *{{ .Type }}Value) Set(val string) error {
	v, err := {{ .ParseFunc }}

	*{{ .Receiver }} = {{ .Type }}Value(v)

	return err
}

func (*{{ .Type }}Value) Type() string { return "{{ .Type }}" }

func ({{ .Receiver }} *{{ .Type }}Value) String() string { return {{ .FormatFunc }} }

// {{ .TypeName }}Var defines {{ .Article }} {{ .Type }} environment variable with specified name, default value, and usage string.
// The argument p points to {{ .Article }} {{ .Type }} variable in which to store the value of the environment variable.
func (s *EnvVarSet) {{ .TypeName }}Var(p *{{ .Type }}, name string, value {{ .Type }}, usage string) {
	s.Var(new{{ .TypeName }}Value(value, p), name, usage)
}

// {{ .TypeName }} defines {{ .Article }} {{ .Type }} environment variable with specified name, default value, and usage string.
// The return value is the address of {{ .Article }} {{ .Type }} variable that stores the value of the environment variable.
func (s *EnvVarSet) {{ .TypeName }}(name string, value {{ .Type }}, usage string) *{{ .Type }} {
	p := new({{ .Type }})

	s.{{ .TypeName }}Var(p, name, value, usage)

	return p
}

// {{ .TypeName }}Var defines {{ .Article }} {{ .Type }} environment variable with specified name, default value, and usage string.
// The argument p points to {{ .Article }} {{ .Type }} variable in which to store the value of the environment variable.
func {{ .TypeName }}Var(p *{{ .Type }}, name string, value {{ .Type }}, usage string) {
	Environment.{{ .TypeName }}Var(p, name, value, usage)
}

// {{ .TypeName }} defines {{ .Article }} {{ .Type }} environment variable with specified name, default value, and usage string.
// The return value is the address of {{ .Article }} {{ .Type }} variable that stores the value of the environment variable.
func {{ .TypeName }}(name string, value {{ .Type }}, usage string) *{{ .Type }} {
	return Environment.{{ .TypeName }}(name, value, usage)
}
