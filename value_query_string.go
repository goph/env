package env

import "strings"

type queryStringValue map[string]string

func newQueryStringValue(val map[string]string, p *map[string]string) *queryStringValue {
	*p = val

	return (*queryStringValue)(p)
}

func (p queryStringValue) Set(val string) error {
	// Clear the map from default values
	for k := range p {
		delete(p, k)
	}

	for _, v := range strings.Split(val, "&") {
		param := strings.SplitN(v, "=", 2)
		if len(param) != 2 {
			continue
		}

		p[param[0]] = param[1]
	}

	return nil
}

func (queryStringValue) Type() string {
	return "queryString"
}

func (p queryStringValue) String() string {
	var query string

	for key, value := range p {
		if query != "" {
			query += "&"
		}

		query += key + "=" + value
	}

	return query
}

// QueryStringVar defines a query string environment variable with specified name, default value, and usage string.
// The argument p points to a query string (string map) variable in which to store the value of the environment variable.
func (s *EnvVarSet) QueryStringVar(p *map[string]string, name string, value map[string]string, usage string) {
	s.Var(newQueryStringValue(value, p), name, usage)
}

// QueryString defines a query string environment variable with specified name, default value, and usage string.
// The return value is the address of a query string (string map) variable that stores the value of the environment variable.
func (s *EnvVarSet) QueryString(name string, value map[string]string, usage string) *map[string]string {
	p := new(map[string]string)

	s.QueryStringVar(p, name, value, usage)

	return p
}

// QueryStringVar defines a query string environment variable with specified name, default value, and usage string.
// The argument p points to a query string (string map) variable in which to store the value of the environment variable.
func QueryStringVar(p *map[string]string, name string, value map[string]string, usage string) {
	Environment.QueryStringVar(p, name, value, usage)
}

// QueryString defines a query string environment variable with specified name, default value, and usage string.
// The return value is the address of a query string (string map) variable that stores the value of the environment variable.
func QueryString(name string, value map[string]string, usage string) *map[string]string {
	return Environment.QueryString(name, value, usage)
}
