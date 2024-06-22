package envconv_test

import (
	"os"
	"testing"

	"github.com/rmhubbert/envconv"
	"github.com/stretchr/testify/assert"
)

type ToByteSliceTest struct {
	env         string
	value       string
	expected    []byte
	errExpected bool
}

type ToByteSliceWithDefaultTest struct {
	env          string
	value        string
	expected     []byte
	defaultValue []byte
}

func TestToByteSlice(t *testing.T) {
	testData := []ToByteSliceTest{
		{"TEST_BYTE_SLICE_HELLO_WORLD", "Hello World", []byte{72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100}, false},
	}

	for _, td := range testData {
		t.Run(td.env, func(t *testing.T) {
			os.Setenv(td.env, td.value)
			v, err := envconv.ToByteSlice(td.env)
			if td.errExpected {
				assert.Error(t, err, "there should be an error")
			} else {
				assert.NoError(t, err, "there should be no error")
			}
			assert.Equal(t, td.expected, v, "they should be equal")
		})
	}

	t.Run("TEST_NON_EXISTANT does not exist", func(t *testing.T) {
		v, err := envconv.ToByteSlice("TEST_NON_EXISTANT")
		assert.Error(t, err, "there should be an error")
		assert.Equal(t, []byte{}, v, "they should be equal")
	})
}

func TestToByteSliceWithDefault(t *testing.T) {
	testData := []ToByteSliceWithDefaultTest{
		{"TEST_BOOL_WITH_DEFAULT_HELLO_WORLD", "Hello World", []byte{72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100}, []byte{72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100}},
	}

	for _, td := range testData {
		t.Run(td.env, func(t *testing.T) {
			os.Setenv(td.env, td.value)
			v := envconv.ToByteSliceWithDefault(td.env, td.defaultValue)
			assert.Equal(t, td.expected, v, "they should be equal")
		})
	}

	t.Run("TEST_NON_EXISTANT does not exist", func(t *testing.T) {
		v := envconv.ToByteSliceWithDefault("TEST_NON_EXISTANT", []byte{})
		assert.Equal(t, []byte{}, v, "they should be equal")
	})
}
