package env_test

import (
	"testing"

	"time"

	"github.com/goph/env"
)

type valueVars struct {
	boolVar        *bool
	durationVar    *time.Duration
	float32Var     *float32
	float64Var     *float64
	intVar         *int
	int16Var       *int16
	int32Var       *int32
	int64Var       *int64
	int8Var        *int8
	queryStringVar *map[string]string
	stringVar      *string
	stringSliceVar *[]string
	uintVar        *uint
	uint16Var      *uint16
	uint32Var      *uint32
	uint64Var      *uint64
	uint8Var       *uint8
}

func newValueVars() *valueVars {
	return &valueVars{
		boolVar:        new(bool),
		durationVar:    new(time.Duration),
		float32Var:     new(float32),
		float64Var:     new(float64),
		intVar:         new(int),
		int16Var:       new(int16),
		int32Var:       new(int32),
		int64Var:       new(int64),
		int8Var:        new(int8),
		queryStringVar: new(map[string]string),
		stringVar:      new(string),
		stringSliceVar: new([]string),
		uintVar:        new(uint),
		uint16Var:      new(uint16),
		uint32Var:      new(uint32),
		uint64Var:      new(uint64),
		uint8Var:       new(uint8),
	}
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
	vars.queryStringVar = envvarset.QueryString("query-string", map[string]string{}, "query string value")
	vars.stringVar = envvarset.String("string", "", "string value")
	vars.stringSliceVar = envvarset.StringSlice("string-slice", []string{}, "string slice value")
	vars.uintVar = envvarset.Uint("uint", 0, "uint value")
	vars.uint16Var = envvarset.Uint16("uint16", 0, "uint16 value")
	vars.uint32Var = envvarset.Uint32("uint32", 0, "uint32 value")
	vars.uint64Var = envvarset.Uint64("uint64", 0, "uint64 value")
	vars.uint8Var = envvarset.Uint8("uint8", 0, "uint8 value")

	testValue(t, envvarset, vars)
}

func TestValueVar(t *testing.T) {
	envvarset := env.NewEnvVarSet(env.ContinueOnError)
	vars := newValueVars()

	envvarset.BoolVar(vars.boolVar, "bool", false, "bool value")
	envvarset.DurationVar(vars.durationVar, "duration", 0, "time.Duration value")
	envvarset.Float32Var(vars.float32Var, "float32", 0, "float32 value")
	envvarset.Float64Var(vars.float64Var, "float64", 0, "float64 value")
	envvarset.IntVar(vars.intVar, "int", 0, "int value")
	envvarset.Int16Var(vars.int16Var, "int16", 0, "int16 value")
	envvarset.Int32Var(vars.int32Var, "int32", 0, "int32 value")
	envvarset.Int64Var(vars.int64Var, "int64", 0, "int64 value")
	envvarset.Int8Var(vars.int8Var, "int8", 0, "int8 value")
	envvarset.QueryStringVar(vars.queryStringVar, "query-string", map[string]string{}, "string value")
	envvarset.StringVar(vars.stringVar, "string", "", "string value")
	envvarset.StringSliceVar(vars.stringSliceVar, "string-slice", []string{}, "string slice value")
	envvarset.UintVar(vars.uintVar, "uint", 0, "uint value")
	envvarset.Uint16Var(vars.uint16Var, "uint16", 0, "uint16 value")
	envvarset.Uint32Var(vars.uint32Var, "uint32", 0, "uint32 value")
	envvarset.Uint64Var(vars.uint64Var, "uint64", 0, "uint64 value")
	envvarset.Uint8Var(vars.uint8Var, "uint8", 0, "uint8 value")

	testValue(t, envvarset, vars)
}

func TestGlobalValue(t *testing.T) {
	env.Environment = env.NewEnvVarSet(env.ContinueOnError)
	vars := new(valueVars)

	vars.boolVar = env.Bool("bool", false, "bool value")
	vars.durationVar = env.Duration("duration", 0, "time.Duration value")
	vars.float32Var = env.Float32("float32", 0, "float32 value")
	vars.float64Var = env.Float64("float64", 0, "float64 value")
	vars.intVar = env.Int("int", 0, "int value")
	vars.int16Var = env.Int16("int16", 0, "int16 value")
	vars.int32Var = env.Int32("int32", 0, "int32 value")
	vars.int64Var = env.Int64("int64", 0, "int64 value")
	vars.int8Var = env.Int8("int8", 0, "int8 value")
	vars.queryStringVar = env.QueryString("query-string", map[string]string{}, "query string value")
	vars.stringVar = env.String("string", "", "string value")
	vars.stringSliceVar = env.StringSlice("string-slice", []string{}, "string slice value")
	vars.uintVar = env.Uint("uint", 0, "uint value")
	vars.uint16Var = env.Uint16("uint16", 0, "uint16 value")
	vars.uint32Var = env.Uint32("uint32", 0, "uint32 value")
	vars.uint64Var = env.Uint64("uint64", 0, "uint64 value")
	vars.uint8Var = env.Uint8("uint8", 0, "uint8 value")

	testValue(t, env.Environment, vars)
}

func TestGlobalValueVar(t *testing.T) {
	env.Environment = env.NewEnvVarSet(env.ContinueOnError)
	vars := newValueVars()

	env.BoolVar(vars.boolVar, "bool", false, "bool value")
	env.DurationVar(vars.durationVar, "duration", 0, "time.Duration value")
	env.Float32Var(vars.float32Var, "float32", 0, "float32 value")
	env.Float64Var(vars.float64Var, "float64", 0, "float64 value")
	env.IntVar(vars.intVar, "int", 0, "int value")
	env.Int16Var(vars.int16Var, "int16", 0, "int16 value")
	env.Int32Var(vars.int32Var, "int32", 0, "int32 value")
	env.Int64Var(vars.int64Var, "int64", 0, "int64 value")
	env.Int8Var(vars.int8Var, "int8", 0, "int8 value")
	env.QueryStringVar(vars.queryStringVar, "query-string", map[string]string{}, "string value")
	env.StringVar(vars.stringVar, "string", "", "string value")
	env.StringSliceVar(vars.stringSliceVar, "string-slice", []string{}, "string slice value")
	env.UintVar(vars.uintVar, "uint", 0, "uint value")
	env.Uint16Var(vars.uint16Var, "uint16", 0, "uint16 value")
	env.Uint32Var(vars.uint32Var, "uint32", 0, "uint32 value")
	env.Uint64Var(vars.uint64Var, "uint64", 0, "uint64 value")
	env.Uint8Var(vars.uint8Var, "uint8", 0, "uint8 value")

	testValue(t, env.Environment, vars)
}

func testValue(t *testing.T, envvarset *env.EnvVarSet, vars *valueVars) {
	environment := map[string]string{
		"BOOL":         "true",
		"DURATION":     "1s",
		"FLOAT32":      "172e12",
		"FLOAT64":      "2718e28",
		"INT":          "22",
		"INT16":        "16",
		"INT32":        "32",
		"INT64":        "64",
		"INT8":         "8",
		"QUERY_STRING": "key=value&key2=value2",
		"STRING":       "string",
		"STRING_SLICE": "one,two,three,four",
		"UINT":         "22",
		"UINT16":       "16",
		"UINT32":       "32",
		"UINT64":       "64",
		"UINT8":        "8",
	}

	err := envvarset.Parse(environment)

	if err != nil {
		t.Fatal("Parse is expected to return a nil (non-error) value")
	}

	if !*vars.boolVar {
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

	if (*vars.queryStringVar)["key"] != "value" || (*vars.queryStringVar)["key2"] != "value2" {
		t.Error("query string var should be `key=value&key2=value2`, got: ", *vars.queryStringVar)
	}

	if *vars.stringVar != "string" {
		t.Error("string var should be `string`, got: ", *vars.stringVar)
	}

	for key, value := range []string{"one", "two", "three", "four"} {
		if (*vars.stringSliceVar)[key] != value {
			t.Errorf("string slice var[%d] should be %s, got: %s", key, value, (*vars.stringSliceVar)[key])
		}
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
