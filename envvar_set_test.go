package env_test

import (
	"testing"

	"github.com/goph/env"
)

type value struct {
	value string
}

func (v *value) String() string {
	return v.value
}

func (v *value) Set(value string) error {
	v.value = value

	return nil
}

func (*value) Type() string {
	return "value"
}

func TestEnvVarSet(t *testing.T) {
	environment := map[string]string{
		"value": "value",
	}

	envvarset := env.NewEnvVarSet()

	v := &value{}

	envvarset.Var(v, "value", "Value usage string")

	err := envvarset.Parse(environment)

	if err != nil {
		t.Fatal("Parse is expected to return a nil (non-error) value")
	}

	if v.value != "value" {
		t.Error("returned value is expected to be value")
	}
}
