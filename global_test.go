package env_test

import (
	"os"
	"testing"

	"github.com/goph/env"
)

func TestEnvironment(t *testing.T) {
	os.Clearenv()

	os.Setenv("VALUE", "value")

	if env.Parsed() {
		t.Fatal("parsed before Parse is called")
	}

	v := &valueStub{
		typ: "valueStub",
		err: nil,
	}

	env.Var(v, "value", "Value usage string")

	env.Parse()

	os.Clearenv()

	if !env.Parsed() {
		t.Error("not parsed after Parse is called")
	}

	if v.value != "value" {
		t.Error("returned value is expected to be value")
	}
}
