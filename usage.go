package env

import (
	"bytes"
	"fmt"
	"strings"
)

// PrintDefaults prints, to standard error unless configured
// otherwise, the default values of all defined environment variables in the set.
func (s *EnvVarSet) PrintDefaults() {
	usages := s.EnvVarUsages()

	fmt.Fprint(s.out(), usages)
}

// EnvVarUsages returns a string containing the usage information for all environment variables
// in the set.
func (s *EnvVarSet) EnvVarUsages() string {
	return s.EnvVarUsagesWrapped(0)
}

// EnvVarUsages returns a string containing the usage information for all environment variables
// in the set.
// Wrapped to `cols` columns (0 for no wrapping)
func (s *EnvVarSet) EnvVarUsagesWrapped(cols int) string {
	buf := new(bytes.Buffer)

	lines := make([]string, 0, len(s.vars))

	maxlen := 0

	s.VisitAll(func(v *EnvVar) {
		name := s.normalizeVarName(v.Name)

		line := fmt.Sprintf("      %s", name)

		varname, usage := UnquoteUsage(v)
		if varname != "" {
			line += " " + varname
		}

		// This special character will be replaced with spacing once the
		// correct alignment is calculated
		line += "\x00"
		if len(line) > maxlen {
			maxlen = len(line)
		}

		line += usage
		if !v.defaultIsZeroValue() {
			if v.Value.Type() == "string" {
				line += fmt.Sprintf(" (default %q)", v.DefaultValue)
			} else {
				line += fmt.Sprintf(" (default %s)", v.DefaultValue)
			}
		}

		lines = append(lines, line)
	})

	for _, line := range lines {
		sidx := strings.Index(line, "\x00")
		spacing := strings.Repeat(" ", maxlen-sidx)
		// maxlen + 2 comes from + 1 for the \x00 and + 1 for the (deliberate) off-by-one in maxlen-sidx
		fmt.Fprintln(buf, line[:sidx], spacing, wrap(maxlen+2, cols, line[sidx+1:]))
	}

	return buf.String()
}

// UnquoteUsage extracts a back-quoted name from the usage
// string for an environment variable and returns it and the un-quoted usage.
// Given "a `name` to show" it returns ("name", "a name to show").
// If there are no back quotes, the name is an educated guess of the
// type of the environment variable's value.
func UnquoteUsage(envVar *EnvVar) (name string, usage string) {
	// Look for a back-quoted name, but avoid the strings package.
	usage = envVar.Usage
	for i := 0; i < len(usage); i++ {
		if usage[i] == '`' {
			for j := i + 1; j < len(usage); j++ {
				if usage[j] == '`' {
					name = usage[i+1: j]
					usage = usage[:i] + name + usage[j+1:]
					return name, usage
				}
			}
			break // Only one back quote; use type name.
		}
	}

	name = envVar.Value.Type()
	switch name {
	case "bool":
		name = "bool"
	case "float64":
		name = "float"
	case "int64":
		name = "int"
	case "uint64":
		name = "uint"
	case "stringSlice":
		name = "strings"
	case "intSlice":
		name = "ints"
	case "uintSlice":
		name = "uints"
	case "boolSlice":
		name = "bools"
	}

	return
}

// defaultIsZeroValue returns true if the default value for this environment variable
// represents a zero value.
func (s *EnvVar) defaultIsZeroValue() bool {
	switch s.Value.(type) {
	case *boolValue:
		return s.DefaultValue == "false"

	case *durationValue:
		// Beginning in Go 1.7, duration zero values are "0s"
		return s.DefaultValue == "0" || s.DefaultValue == "0s"

	case *intValue, *int8Value, *int32Value, *int64Value, *uintValue, *uint8Value, *uint16Value, *uint32Value, *uint64Value, *float32Value, *float64Value:
		return s.DefaultValue == "0"

	case *stringValue:
		return s.DefaultValue == ""

	default:
		switch s.Value.String() {
		case "false":
			return true
		case "<nil>":
			return true
		case "":
			return true
		case "0":
			return true
		}

		return false
	}
}

// Splits the string `s` on whitespace into an initial substring up to
// `i` runes in length and the remainder. Will go `slop` over `i` if
// that encompasses the entire string (which allows the caller to
// avoid short orphan words on the final line).
func wrapN(i, slop int, s string) (string, string) {
	if i+slop > len(s) {
		return s, ""
	}

	w := strings.LastIndexAny(s[:i], " \t\n")
	if w <= 0 {
		return s, ""
	}
	nlPos := strings.LastIndex(s[:i], "\n")
	if nlPos > 0 && nlPos < w {
		return s[:nlPos], s[nlPos+1:]
	}
	return s[:w], s[w+1:]
}

// Wraps the string `s` to a maximum width `w` with leading indent
// `i`. The first line is not indented (this is assumed to be done by
// caller). Pass `w` == 0 to do no wrapping
func wrap(i, w int, s string) string {
	if w == 0 {
		return strings.Replace(s, "\n", "\n"+strings.Repeat(" ", i), -1)
	}

	// space between indent i and end of line width w into which
	// we should wrap the text.
	wrap := w - i

	var r, l string

	// Not enough space for sensible wrapping. Wrap as a block on
	// the next line instead.
	if wrap < 24 {
		i = 16
		wrap = w - i
		r += "\n" + strings.Repeat(" ", i)
	}
	// If still not enough space then don't even try to wrap.
	if wrap < 24 {
		return strings.Replace(s, "\n", r, -1)
	}

	// Try to avoid short orphan words on the final line, by
	// allowing wrapN to go a bit over if that would fit in the
	// remainder of the line.
	slop := 5
	wrap = wrap - slop

	// Handle first line, which is indented by the caller (or the
	// special case above)
	l, s = wrapN(wrap, slop, s)
	r = r + strings.Replace(l, "\n", "\n"+strings.Repeat(" ", i), -1)

	// Now wrap the rest
	for s != "" {
		var t string

		t, s = wrapN(wrap, slop, s)
		r = r + "\n" + strings.Repeat(" ", i) + strings.Replace(t, "\n", "\n"+strings.Repeat(" ", i), -1)
	}

	return r
}
