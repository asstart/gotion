package utils_test

import (
	"testing"

	"github.com/asstart/gotion/utils"
	"github.com/stretchr/testify/assert"
)

func TestMinimize(t *testing.T) {
	source := [][]string{
		[]string{"", ""},
		[]string{" ", ""},
		[]string{"  ", ""},
		[]string{`{"field": "value" }`, `{"field":"value"}`},
		[]string{`{"field with spaces" : "value with spaces" } `, `{"field with spaces":"value with spaces"}`},
		[]string{"\n", ""},
		[]string{"\t", ""},
		[]string{"\v", ""},
		[]string{"\f", ""},
		[]string{"\r", ""},
		[]string{"\n", ""},
		[]string{"\n", ""},
	}

	for _, v := range source {
		assert.Equal(t, v[1], utils.Minimise(v[0]))
	}
}
