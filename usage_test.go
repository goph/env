package env_test

import (
	"testing"

	"github.com/goph/env"
)

func TestEnvVarSet_EnvVarUsages(t *testing.T) {
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
	vars.uintVar = envvarset.Uint("uint", 0, "uint value")
	vars.uint16Var = envvarset.Uint16("uint16", 0, "uint16 value")
	vars.uint32Var = envvarset.Uint32("uint32", 0, "uint32 value")
	vars.uint64Var = envvarset.Uint64("uint64", 0, "uint64 value")
	vars.uint8Var = envvarset.Uint8("uint8", 0, "uint8 value")

	expected := `      BOOL bool                  bool value
      DURATION duration          time.Duration value
      FLOAT32 float32            float32 value
      FLOAT64 float              float64 value
      INT int                    int value
      INT16 int16                int16 value
      INT32 int32                int32 value
      INT64 int                  int64 value
      INT8 int8                  int8 value
      QUERY_STRING queryString   query string value
      STRING string              string value
      UINT uint                  uint value
      UINT16 uint16              uint16 value
      UINT32 uint32              uint32 value
      UINT64 uint                uint64 value
      UINT8 uint8                uint8 value
`

	actual := envvarset.EnvVarUsages()

	if actual != expected {
		t.Errorf("expected the following output: \n%s\n\nactual: \n%s", expected, actual)
	}
}
