package envconv_test

import (
	"os"
	"testing"

	"github.com/rmhubbert/envconv"
	"github.com/stretchr/testify/assert"
)

func TestToString(t *testing.T) {
	testData := []struct {
		env         string
		value       string
		expected    string
		errExpected bool
	}{
		{"TEST_STRING_HELLO_WORLD", "Hello World", "Hello World", false},
	}

	for _, td := range testData {
		t.Run(td.env, func(t *testing.T) {
			os.Setenv(td.env, td.value)
			v, err := envconv.ToString(td.env)
			if td.errExpected {
				assert.Error(t, err, "there should be an error")
			} else {
				assert.NoError(t, err, "there should be no error")
			}
			assert.Equal(t, td.expected, v, "they should be equal")
		})
	}

	t.Run("TEST_NON_EXISTANT does not exist", func(t *testing.T) {
		v, err := envconv.ToString("TEST_NON_EXISTANT")
		assert.Error(t, err, "there should be an error")
		assert.Equal(t, "", v, "they should be equal")
	})
}

func TestToStringSlice(t *testing.T) {
	testData := []struct {
		env         string
		value       string
		separator   string
		expected    []string
		errExpected bool
	}{
		{"TEST_STRING_SLICE_HELLO_WORLD_SPACE", "Hello World", " ", []string{"Hello", "World"}, false},
		{"TEST_STRING_SLICE_HELLO_WORLD_COMMA", "Hello World", ",", []string{"Hello World"}, false},
	}

	for _, td := range testData {
		t.Run(td.env, func(t *testing.T) {
			os.Setenv(td.env, td.value)
			v, err := envconv.ToStringSlice(td.env, td.separator)
			if td.errExpected {
				assert.Error(t, err, "there should be an error")
			} else {
				assert.NoError(t, err, "there should be no error")
			}
			assert.Equal(t, td.expected, v, "they should be equal")
		})
	}

	t.Run("TEST_NON_EXISTANT does not exist", func(t *testing.T) {
		v, err := envconv.ToStringSlice("TEST_NON_EXISTANT", " ")
		assert.Error(t, err, "there should be an error")
		assert.Equal(t, []string{}, v, "they should be equal")
	})
}

func TestToStringWithDefault(t *testing.T) {
	testData := []struct {
		env          string
		value        string
		expected     string
		defaultValue string
	}{
		{"TEST_STRING_HELLO_WORLD", "Hello World", "Hello World", "Default"},
		{"TEST_STRING_EMPTY", "", "", "Default"},
	}

	for _, td := range testData {
		t.Run(td.env, func(t *testing.T) {
			os.Setenv(td.env, td.value)
			v := envconv.ToStringWithDefault(td.env, td.defaultValue)
			if v != td.expected {
				assert.Equal(t, td.defaultValue, v, "they should be equal")
			} else {
				assert.Equal(t, td.expected, v, "they should be equal")
			}

		})
	}

	t.Run("TEST_NON_EXISTANT does not exist", func(t *testing.T) {
		v := envconv.ToStringWithDefault("TEST_NON_EXISTANT", "")
		assert.Equal(t, "", v, "they should be equal")
	})
}

func TestToStringSliceWithDefault(t *testing.T) {
	testData := []struct {
		env          string
		value        string
		separator    string
		expected     []string
		defaultValue []string
	}{
		{"TEST_STRING_SLICE_WITH_DEFAUTL_HELLO_WORLD", "Hello World", " ", []string{"Hello", "World"}, []string{"105"}},
		{"TEST_STRING_SLICE_WITH_DEFAULT_EMPTY", "", " ", []string{""}, []string{"105"}},
	}

	for _, td := range testData {
		t.Run(td.env, func(t *testing.T) {
			os.Setenv(td.env, td.value)
			v := envconv.ToStringSliceWithDefault(td.env, td.separator, td.defaultValue)
			if !slicesEqual(v, td.expected) {
				assert.Equal(t, td.defaultValue, v, "they should be equal")
			} else {
				assert.Equal(t, td.expected, v, "they should be equal")
			}
		})
	}

	t.Run("TEST_NON_EXISTANT does not exist", func(t *testing.T) {
		v := envconv.ToStringSliceWithDefault("TEST_NON_EXISTANT", " ", []string{})
		assert.Equal(t, []string{}, v, "they should be equal")
	})
}
