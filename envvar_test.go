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

func TestEnvVarSet(t *testing.T) {
	environment := map[string]string{
		"value": "value",
	}

	envvarset := env.NewEnvVarSet(env.ContinueOnError)

	if envvarset.Parsed() {
		t.Fatal("parsed before Parse is called")
	}

	v := &valueStub{
		typ: "valueStub",
		err: nil,
	}

	envvarset.Var(v, "value", "Value usage string")

	err := envvarset.Parse(environment)

	if err != nil {
		t.Fatal("Parse is expected to return a nil (non-error) value")
	}

	if !envvarset.Parsed() {
		t.Error("not parsed after Parse is called")
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
