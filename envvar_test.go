package env_test

import (
	"bytes"
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
		"VALUE": "value",
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

func TestEnvVarSet_ParseEnviron(t *testing.T) {
	environment := []string{
		"VALUE=value=value",
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

	err := envvarset.ParseEnviron(environment)

	if err != nil {
		t.Fatal("Parse is expected to return a nil (non-error) value")
	}

	if !envvarset.Parsed() {
		t.Error("not parsed after Parse is called")
	}

	if v.value != "value=value" {
		t.Error("returned value is expected to be value=value")
	}
}

func TestEnvVarSet_ErrorHandling_ContinueOnError(t *testing.T) {
	environment := map[string]string{
		"VALUE": "value",
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
		"VALUE": "value",
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
			t.Error("Parse is expected to panic")
		}
	}()

	envvarset.Parse(environment) // nolint:errcheck
}

func TestRedeclare(t *testing.T) {
	var buf bytes.Buffer

	envvarset := env.NewEnvVarSet(env.PanicOnError)

	envvarset.SetOutput(&buf)

	v := &valueStub{
		typ: "valueStub",
		err: errors.New("error"),
	}

	envvarset.Var(v, "value", "Value usage string")

	defer func() {
		err := recover()

		if err == nil {
			t.Error("Var is expected to panic when redeclaring a variable")
		}

		errString := buf.String()

		if errString != "VALUE environment variable redefined: value\n" {
			t.Error("expected error output")
		}
	}()

	envvarset.Var(v, "value", "Value usage string")
}

func TestEnvVarSet_VisitAll(t *testing.T) {
	envvarset := env.NewEnvVarSet(env.ContinueOnError)

	names := map[string]bool{
		"A": false,
		"B": false,
		"C": false,
		"D": false,
	}
	for name := range names {
		envvarset.Bool(name, false, "")
	}

	envvarset.VisitAll(func(e *env.EnvVar) {
		names[e.Name] = true
	})

	for name, visited := range names {
		if !visited {
			t.Errorf("variable %q is expected to be visited", name)
		}
	}
}

func TestEnvVarSet_VisitAll_Order(t *testing.T) {
	envvarset := env.NewEnvVarSet(env.ContinueOnError)

	names := []string{"C", "B", "A", "D"}
	for _, name := range names {
		envvarset.Bool(name, false, "")
	}

	i := 0
	envvarset.VisitAll(func(v *env.EnvVar) {
		if names[i] != v.Name {
			t.Errorf("Incorrect order. Expected %v, got %v", names[i], v.Name)
		}
		i++
	})
}

func TestEnvVarSet_HasEnvVars(t *testing.T) {
	envvarset := env.NewEnvVarSet(env.ContinueOnError)

	v := &valueStub{}

	envvarset.Var(v, "value", "value usage string")

	if !envvarset.HasEnvVars() {
		t.Error("EnvVarSet is expected to have at least one environment variable defined")
	}
}
