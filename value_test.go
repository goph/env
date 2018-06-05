package env_test

import (
	"testing"

	"github.com/goph/env"
)

type valueVars struct {
	intVar    *int
	int16Var  *int16
	int32Var  *int32
	int8Var   *int8
	stringVar *string
}

func TestValue(t *testing.T) {
	envvarset := env.NewEnvVarSet(env.ContinueOnError)
	vars := new(valueVars)

	vars.intVar = envvarset.Int("int", 0, "int value")
	vars.int16Var = envvarset.Int16("int16", 0, "int16 value")
	vars.int32Var = envvarset.Int32("int32", 0, "int32 value")
	vars.int8Var = envvarset.Int8("int8", 0, "int8 value")
	vars.stringVar = envvarset.String("string", "", "string value")

	testValue(t, envvarset, vars)
}

func TestValueVar(t *testing.T) {
	envvarset := env.NewEnvVarSet(env.ContinueOnError)
	vars := &valueVars{
		intVar:    new(int),
		int16Var:  new(int16),
		int32Var:  new(int32),
		int8Var:   new(int8),
		stringVar: new(string),
	}

	envvarset.IntVar(vars.intVar, "int", 0, "int value")
	envvarset.Int16Var(vars.int16Var, "int16", 0, "int16 value")
	envvarset.Int32Var(vars.int32Var, "int32", 0, "int32 value")
	envvarset.Int8Var(vars.int8Var, "int8", 0, "int8 value")
	envvarset.StringVar(vars.stringVar, "string", "", "string value")

	testValue(t, envvarset, vars)
}

func testValue(t *testing.T, envvarset *env.EnvVarSet, vars *valueVars) {
	environment := map[string]string{
		"int":    "22",
		"int16":  "16",
		"int32":  "32",
		"int8":   "8",
		"string": "string",
	}

	err := envvarset.Parse(environment)

	if err != nil {
		t.Fatal("Parse is expected to return a nil (non-error) value")
	}

	if *vars.intVar != 22 {
		t.Error("int var should be `22`, got: ", *vars.intVar)
	}

	if *vars.int16Var != 16 {
		t.Error("int16 var should be `16`, got: ", *vars.int16Var)
	}

	if *vars.int32Var != 32 {
		t.Error("int32 var should be `32`, got: ", *vars.int32Var)
	}

	if *vars.int8Var != 8 {
		t.Error("int8 var should be `8`, got: ", *vars.int8Var)
	}

	if *vars.stringVar != "string" {
		t.Error("string var should be `string`, got: ", *vars.stringVar)
	}
}
