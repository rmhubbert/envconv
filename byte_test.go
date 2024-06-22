package envconv_test

import (
	"os"
	"testing"

	"github.com/rmhubbert/envconv"
	"github.com/stretchr/testify/assert"
)

type ToByteTest struct {
	env         string
	value       string
	expected    byte
	errExpected bool
}

type ToByteWithDefaultTest struct {
	env          string
	value        string
	expected     byte
	defaultValue byte
}

func TestToByte(t *testing.T) {
	testData := []ToByteTest{
		{"TEST_BYTE_1", "-128", 128, false},
		{"TEST_BYTE_1", "-129", 0, true},
		{"TEST_BYTE_1", "-1", 255, false},
		{"TEST_BYTE_0", "0", 0, false},
		{"TEST_BYTE_1", "1", 1, false},
		{"TEST_BYTE_127", "127", 127, false},
		{"TEST_BYTE_128", "128", 0, true},
	}

	for _, td := range testData {
		t.Run(td.env, func(t *testing.T) {
			os.Setenv(td.env, td.value)
			v, err := envconv.ToByte(td.env)
			if td.errExpected {
				assert.Error(t, err, "there should be an error")
			} else {
				assert.NoError(t, err, "there should be no error")
			}
			assert.Equal(t, td.expected, v, "they should be equal")
		})
	}

	t.Run("TEST_NON_EXISTANT does not exist", func(t *testing.T) {
		v, err := envconv.ToByte("TEST_NON_EXISTANT")
		assert.Error(t, err, "there should be an error")
		assert.Equal(t, byte(0), v, "they should be equal")
	})
}

func TestToByteWithDefault(t *testing.T) {
	testData := []ToByteWithDefaultTest{
		{"TEST_BYTE_1", "-128", 128, 105},
		{"TEST_BYTE_1", "-129", 0, 105},
		{"TEST_BYTE_1", "-1", 255, 105},
		{"TEST_BYTE_0", "0", 0, 105},
		{"TEST_BYTE_1", "1", 1, 105},
		{"TEST_BYTE_127", "127", 127, 105},
		{"TEST_BYTE_128", "128", 0, 105},
	}

	for _, td := range testData {
		t.Run(td.env, func(t *testing.T) {
			os.Setenv(td.env, td.value)
			v := envconv.ToByteWithDefault(td.env, td.defaultValue)
			if v != td.expected {
				assert.Equal(t, td.defaultValue, v, "they should be equal")
			} else {
				assert.Equal(t, td.expected, v, "they should be equal")
			}
		})
	}

	t.Run("TEST_NON_EXISTANT does not exist", func(t *testing.T) {
		v := envconv.ToByteWithDefault("TEST_NON_EXISTANT", byte(0))
		assert.Equal(t, byte(0), v, "they should be equal")
	})
}
