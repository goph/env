package env_test

import (
	"errors"
	"testing"

	"github.com/goph/env"
)

type valueStub struct {
	err error
	typ string

	value string
}

func (v *valueStub) String() string {
	return v.value
}

func (v *valueStub) Set(value string) error {
	v.value = value

	return v.err
}

func (v *valueStub) Type() string {
	return v.typ
}

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

	envvarset := env.NewEnvVarSet(env.ContinueOnError)

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

func TestEnvVarSet_ErrorHandling_ContinueOnError(t *testing.T) {
	environment := map[string]string{
		"value": "value",
	}

	envvarset := env.NewEnvVarSet(env.ContinueOnError)

	v := &valueStub{
		typ: "valueStub",
		err: errors.New("error"),
	}

	envvarset.Var(v, "value", "Value usage string")

	err := envvarset.Parse(environment)

	if err == nil {
		t.Fatal("Parse is expected to return a non-nil error value")
	}
}

func TestEnvVarSet_ErrorHandling_Panic(t *testing.T) {
	environment := map[string]string{
		"value": "value",
	}

	envvarset := env.NewEnvVarSet(env.PanicOnError)

	v := &valueStub{
		typ: "valueStub",
		err: errors.New("error"),
	}

	envvarset.Var(v, "value", "Value usage string")

	defer func() {
		err := recover()

		if err == nil {
			t.Fatal("Parse is expected to panic")
		}
	}()

	envvarset.Parse(environment)
}
