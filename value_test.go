package env_test

import (
	"testing"

	"time"

	"github.com/goph/env"
)

type valueVars struct {
	boolVar     *bool
	durationVar *time.Duration
	float32Var  *float32
	float64Var  *float64
	intVar      *int
	int16Var    *int16
	int32Var    *int32
	int64Var    *int64
	int8Var     *int8
	stringVar   *string
	uintVar     *uint
	uint16Var   *uint16
	uint32Var   *uint32
	uint64Var   *uint64
	uint8Var    *uint8
}

func TestValue(t *testing.T) {
	envvarset := env.NewEnvVarSet(env.ContinueOnError)
	vars := new(valueVars)

	vars.boolVar = envvarset.Bool("bool", false, "bool value")
	vars.durationVar = envvarset.Duration("duration", 0, "time.Duration value")
	vars.float32Var = envvarset.Float32("float32", 0, "float32 value")
	vars.float64Var = envvarset.Float64("float64", 0, "float64 value")
	vars.intVar = envvarset.Int("int", 0, "int value")
	vars.int16Var = envvarset.Int16("int16", 0, "int16 value")
	vars.int32Var = envvarset.Int32("int32", 0, "int32 value")
	vars.int64Var = envvarset.Int64("int64", 0, "int64 value")
	vars.int8Var = envvarset.Int8("int8", 0, "int8 value")
	vars.stringVar = envvarset.String("string", "", "string value")
	vars.uintVar = envvarset.Uint("uint", 0, "uint value")
	vars.uint16Var = envvarset.Uint16("uint16", 0, "uint16 value")
	vars.uint32Var = envvarset.Uint32("uint32", 0, "uint32 value")
	vars.uint64Var = envvarset.Uint64("uint64", 0, "uint64 value")
	vars.uint8Var = envvarset.Uint8("uint8", 0, "uint8 value")

	testValue(t, envvarset, vars)
}

func TestValueVar(t *testing.T) {
	envvarset := env.NewEnvVarSet(env.ContinueOnError)
	vars := &valueVars{
		boolVar:     new(bool),
		durationVar: new(time.Duration),
		float32Var:  new(float32),
		float64Var:  new(float64),
		intVar:      new(int),
		int16Var:    new(int16),
		int32Var:    new(int32),
		int64Var:    new(int64),
		int8Var:     new(int8),
		stringVar:   new(string),
		uintVar:     new(uint),
		uint16Var:   new(uint16),
		uint32Var:   new(uint32),
		uint64Var:   new(uint64),
		uint8Var:    new(uint8),
	}

	envvarset.BoolVar(vars.boolVar, "bool", false, "bool value")
	envvarset.DurationVar(vars.durationVar, "duration", 0, "time.Duration value")
	envvarset.Float32Var(vars.float32Var, "float32", 0, "float32 value")
	envvarset.Float64Var(vars.float64Var, "float64", 0, "float64 value")
	envvarset.IntVar(vars.intVar, "int", 0, "int value")
	envvarset.Int16Var(vars.int16Var, "int16", 0, "int16 value")
	envvarset.Int32Var(vars.int32Var, "int32", 0, "int32 value")
	envvarset.Int64Var(vars.int64Var, "int64", 0, "int64 value")
	envvarset.Int8Var(vars.int8Var, "int8", 0, "int8 value")
	envvarset.StringVar(vars.stringVar, "string", "", "string value")
	envvarset.UintVar(vars.uintVar, "uint", 0, "uint value")
	envvarset.Uint16Var(vars.uint16Var, "uint16", 0, "uint16 value")
	envvarset.Uint32Var(vars.uint32Var, "uint32", 0, "uint32 value")
	envvarset.Uint64Var(vars.uint64Var, "uint64", 0, "uint64 value")
	envvarset.Uint8Var(vars.uint8Var, "uint8", 0, "uint8 value")

	testValue(t, envvarset, vars)
}

func testValue(t *testing.T, envvarset *env.EnvVarSet, vars *valueVars) {
	environment := map[string]string{
		"bool":     "true",
		"duration": "1s",
		"float32":  "172e12",
		"float64":  "2718e28",
		"int":      "22",
		"int16":    "16",
		"int32":    "32",
		"int64":    "64",
		"int8":     "8",
		"string":   "string",
		"uint":     "22",
		"uint16":   "16",
		"uint32":   "32",
		"uint64":   "64",
		"uint8":    "8",
	}

	err := envvarset.Parse(environment)

	if err != nil {
		t.Fatal("Parse is expected to return a nil (non-error) value")
	}

	if *vars.boolVar != true {
		t.Error("bool var should be `true`, got: ", *vars.boolVar)
	}

	if *vars.durationVar != time.Second {
		t.Error("duration var should be `1s`, got: ", (*vars.durationVar).String())
	}

	if *vars.float32Var != 172e12 {
		t.Error("float32 var should be `172e12`, got: ", *vars.float32Var)
	}

	if *vars.float64Var != 2718e28 {
		t.Error("float64 var should be `172e12`, got: ", *vars.float64Var)
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

	if *vars.int64Var != 64 {
		t.Error("int64 var should be `64`, got: ", *vars.int64Var)
	}

	if *vars.int8Var != 8 {
		t.Error("int8 var should be `8`, got: ", *vars.int8Var)
	}

	if *vars.stringVar != "string" {
		t.Error("string var should be `string`, got: ", *vars.stringVar)
	}

	if *vars.uintVar != 22 {
		t.Error("uint var should be `22`, got: ", *vars.uintVar)
	}

	if *vars.uint16Var != 16 {
		t.Error("uint16 var should be `16`, got: ", *vars.uint16Var)
	}

	if *vars.uint32Var != 32 {
		t.Error("uint32 var should be `32`, got: ", *vars.uint32Var)
	}

	if *vars.uint64Var != 64 {
		t.Error("uint64 var should be `64`, got: ", *vars.uint64Var)
	}

	if *vars.uint8Var != 8 {
		t.Error("uint8 var should be `8`, got: ", *vars.uint8Var)
	}
}
