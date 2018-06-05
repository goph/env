package env_test

import (
	"testing"

	"github.com/goph/env"
)

type valueVars struct {
	intVar    *int
	stringVar *string
}

func TestValue(t *testing.T) {
	envvarset := env.NewEnvVarSet(env.ContinueOnError)
	vars := new(valueVars)

	vars.intVar = envvarset.Int("int", 0, "int value")
	vars.stringVar = envvarset.String("string", "", "string value")

	testValue(t, envvarset, vars)
}

func TestValueVar(t *testing.T) {
	envvarset := env.NewEnvVarSet(env.ContinueOnError)
	vars := &valueVars{
		intVar:    new(int),
		stringVar: new(string),
	}

	envvarset.IntVar(vars.intVar, "int", 0, "int value")
	envvarset.StringVar(vars.stringVar, "string", "", "string value")

	testValue(t, envvarset, vars)
}

func testValue(t *testing.T, envvarset *env.EnvVarSet, vars *valueVars) {
	environment := map[string]string{
		"int":    "22",
		"string": "string",
	}

	err := envvarset.Parse(environment)

	if err != nil {
		t.Fatal("Parse is expected to return a nil (non-error) value")
	}

	if *vars.intVar != 22 {
		t.Error("int var should be `22`, got: ", *vars.intVar)
	}

	if *vars.stringVar != "string" {
		t.Error("string var should be `string`, got: ", *vars.stringVar)
	}
}
