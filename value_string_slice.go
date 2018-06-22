package env

import (
	"bytes"
	"encoding/csv"
	"strings"
)

type stringSliceValue struct {
	value   *[]string
	changed bool
}

func newStringSliceValue(val []string, p *[]string) *stringSliceValue {
	ssv := new(stringSliceValue)

	ssv.value = p
	*ssv.value = val

	return ssv
}

func readAsCSV(val string) ([]string, error) {
	if val == "" {
		return []string{}, nil
	}

	csvReader := csv.NewReader(strings.NewReader(val))

	return csvReader.Read()
}

func writeAsCSV(vals []string) (string, error) {
	b := &bytes.Buffer{}
	w := csv.NewWriter(b)

	err := w.Write(vals)
	if err != nil {
		return "", err
	}

	w.Flush()

	return strings.TrimSuffix(b.String(), "\n"), nil
}

func (s *stringSliceValue) Set(val string) error {
	v, err := readAsCSV(val)
	if err != nil {
		return err
	}

	if !s.changed {
		*s.value = v
	} else {
		*s.value = append(*s.value, v...)
	}

	s.changed = true

	return nil
}

func (*stringSliceValue) Type() string { return "stringSlice" }

func (s *stringSliceValue) String() string {
	str, _ := writeAsCSV(*s.value)

	return "[" + str + "]"
}

// StringSliceVar defines a string environment variable with specified name, default value, and usage string.
// The argument p points to a []string variable in which to store the value of the environment variable.
// For example:
//   STRING_SLICE="v1,v2"
// will result in
//   []string{"v1", "v2"}
func (s *EnvVarSet) StringSliceVar(p *[]string, name string, value []string, usage string) {
	s.Var(newStringSliceValue(value, p), name, usage)
}

// StringSlice defines a string environment variable with specified name, default value, and usage string.
// The return value is the address of a []string variable that stores the value of the environment variable.
// For example:
//   STRING_SLICE="v1,v2"
// will result in
//   []string{"v1", "v2"}
func (s *EnvVarSet) StringSlice(name string, value []string, usage string) *[]string {
	p := new([]string)

	s.StringSliceVar(p, name, value, usage)

	return p
}

// StringSliceVar defines a string environment variable with specified name, default value, and usage string.
// The argument p points to a []string variable in which to store the value of the environment variable.
// For example:
//   STRING_SLICE="v1,v2"
// will result in
//   []string{"v1", "v2"}
func StringSliceVar(p *[]string, name string, value []string, usage string) {
	Environment.StringSliceVar(p, name, value, usage)
}

// StringSlice defines a string environment variable with specified name, default value, and usage string.
// The return value is the address of a []string variable that stores the value of the environment variable.
// For example:
//   STRING_SLICE="v1,v2"
// will result in
//   []string{"v1", "v2"}
func StringSlice(name string, value []string, usage string) *[]string {
	return Environment.StringSlice(name, value, usage)
}
