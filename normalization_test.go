package env_test

import (
	"testing"

	"github.com/goph/env"
)

func TestEnvVarSet_SetNormalizeFunc(t *testing.T) {
	envvarset := env.NewEnvVarSet(env.ContinueOnError)

	environment := map[string]string{
		"prefix_value": "value",
	}

	v := &valueStub{
		typ: "valueStub",
		err: nil,
	}

	envvarset.Var(v, "value", "Value usage string")

	envvarset.SetNormalizeFunc(func(s *env.EnvVarSet, name string) env.NormalizedName {
		return env.NormalizedName("prefix_" + name)
	})

	err := envvarset.Parse(environment)

	if err != nil {
		t.Fatal("Parse is expected to return a nil (non-error) value")
	}

	if v.value != "value" {
		t.Error("returned value is expected to be value")
	}
}
