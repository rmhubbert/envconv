package envconv_test

import (
	"os"
	"testing"

	"github.com/rmhubbert/envconv"
	"github.com/stretchr/testify/assert"
)

func TestToByte(t *testing.T) {
	testData := []struct {
		env           string
		value         string
		expected      byte
		errorExpected bool
	}{
		{"TEST_BYTE_1", "-1", 0, true},
		{"TEST_BYTE_0", "0", 0, false},
		{"TEST_BYTE_1", "1", 1, false},
		{"TEST_BYTE_127", "255", 255, false},
		{"TEST_BYTE_128", "256", 0, true},
	}

	for _, td := range testData {
		t.Run(td.env, func(t *testing.T) {
			os.Setenv(td.env, td.value)
			v, err := envconv.ToByte(td.env)
			if td.errorExpected {
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

func TestToByteSlice(t *testing.T) {
	testData := []struct {
		env           string
		value         string
		expected      []byte
		errorExpected bool
	}{
		{"TEST_BYTE_SLICE_HELLO_WORLD", "Hello World", []byte{72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100}, false},
	}

	for _, td := range testData {
		t.Run(td.env, func(t *testing.T) {
			os.Setenv(td.env, td.value)
			v, err := envconv.ToByteSlice(td.env)
			if td.errorExpected {
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

func TestToByteWithDefault(t *testing.T) {
	testData := []struct {
		env          string
		value        string
		expected     byte
		defaultValue byte
	}{
		{"TEST_BYTE_WITH_DEFAULT_1", "-1", 0, 105},
		{"TEST_BYTE_WITH_DEFAULT_0", "0", 0, 105},
		{"TEST_BYTE_WITH_DEFAULT_1", "1", 1, 105},
		{"TEST_BYTE_WITH_DEFAULT_255", "255", 255, 105},
		{"TEST_BYTE_WITH_DEFAULT_256", "256", 0, 105},
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

func TestToByteSliceWithDefault(t *testing.T) {
	defaultValue := []byte{72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100}
	testData := []struct {
		env             string
		value           string
		expected        []byte
		defaultValue    []byte
		defaultExpected bool
	}{
		{"TEST_BYTE_SLICE_WITH_DEFAULT_HELLO_WORLD", "Hello World", []byte{72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100}, defaultValue, false},
	}

	for _, td := range testData {
		t.Run(td.env, func(t *testing.T) {
			os.Setenv(td.env, td.value)
			v := envconv.ToByteSliceWithDefault(td.env, td.defaultValue)
			if td.defaultExpected {
				assert.Equal(t, td.defaultValue, v, "they should be equal")
			} else {
				assert.Equal(t, td.expected, v, "they should be equal")
			}
		})
	}

	t.Run("TEST_NON_EXISTANT does not exist", func(t *testing.T) {
		v := envconv.ToByteSliceWithDefault("TEST_NON_EXISTANT", []byte{})
		assert.Equal(t, []byte{}, v, "they should be equal")
	})
}
