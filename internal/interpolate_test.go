package interpolate

import (
	"testing"
)

func TestInterpolation(t *testing.T) {
	envMap := map[string]string{
		"testKey": "testValue",
	}

	Execute(envMap, "test {{ .testKey }}")
}
