package env_test

import (
	"testing"

	"github.com/goph/env"
)

func TestEnvVarSet(t *testing.T) {
	envvarset := env.NewEnvVarSet()

	err := envvarset.Parse()

	if err != nil {
		t.Error("Parse is expected to return a nil (non-error) value")
	}
}
