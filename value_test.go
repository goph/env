package env_test

import (
	"testing"

	"github.com/goph/env"
)

type valueVars struct {
	stringVar *string
}

func TestValue(t *testing.T) {
	envvarset := env.NewEnvVarSet(env.ContinueOnError)
	vars := new(valueVars)

	vars.stringVar = envvarset.String("string", "", "string value")

	testValue(t, envvarset, vars)
}

func TestValueVar(t *testing.T) {
	envvarset := env.NewEnvVarSet(env.ContinueOnError)
	vars := &valueVars{
		stringVar: new(string),
	}

	envvarset.StringVar(vars.stringVar, "string", "", "string value")

	testValue(t, envvarset, vars)
}

func testValue(t *testing.T, envvarset *env.EnvVarSet, vars *valueVars) {
	environment := map[string]string{
		"string": "string",
	}

	err := envvarset.Parse(environment)

	if err != nil {
		t.Fatal("Parse is expected to return a nil (non-error) value")
	}

	if *vars.stringVar != "string" {
		t.Error("string var should be `string`, got: ", *vars.stringVar)
	}
}
