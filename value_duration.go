package env

import "time"

type durationValue time.Duration

func newDurationValue(val time.Duration, p *time.Duration) *durationValue {
	*p = val

	return (*durationValue)(p)
}

func (d *durationValue) Set(val string) error {
	v, err := time.ParseDuration(val)

	*d = durationValue(v)

	return err
}

func (*durationValue) Type() string { return "duration" }

func (d *durationValue) String() string { return (*time.Duration)(d).String() }

// DurationVar defines a duration environment variable with specified name, default value, and usage string.
// The argument p points to a duration variable in which to store the value of the environment variable.
func (s *EnvVarSet) DurationVar(p *time.Duration, name string, value time.Duration, usage string) {
	s.Var(newDurationValue(value, p), name, usage)
}

// Duration defines a duration environment variable with specified name, default value, and usage string.
// The return value is the address of a duration variable that stores the value of the environment variable.
func (s *EnvVarSet) Duration(name string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)

	s.DurationVar(p, name, value, usage)

	return p
}
