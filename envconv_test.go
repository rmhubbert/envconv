package envconv_test

import (
	"os"
	"testing"

	"github.com/rmhubbert/envconv"
	"github.com/stretchr/testify/assert"
)

type EnvTest struct {
	env         string
	value       string
	errExpected bool
}

func TestLoadFromEnvironment(t *testing.T) {
	testData := []EnvTest{
		{"ENV_TEST_HELLO", "hello", false},
		{"ENV_TEST_HELLO_WORLD", "hello world", false},
		{"ENV_TEST_HELLO__NEWLINE_WORLD", "hello\nworld", false},
	}

	for _, td := range testData {
		t.Run(td.env, func(t *testing.T) {
			os.Setenv(td.env, td.value)
			v, err := envconv.LoadFromEnvironment(td.env, true)
			if td.errExpected {
				assert.Error(t, err, "there should be an error")
			} else {
				assert.NoError(t, err, "there should be no error")
			}
			assert.Equal(t, td.value, v, "they should be equal")
		})
	}
}
