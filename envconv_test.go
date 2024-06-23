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

type sliceType interface {
	~string | ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func slicesEqual[T sliceType](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

// TestReturnValueType is a type constraint interface that is used by the run* generic
// test functions
type TestReturnValueType interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// runTest provides a generic test run for the convertor function types that
// take a single string and return a T and an error
func runTest[T TestReturnValueType, F func(string) (T, error)](
	t *testing.T,
	env string,
	value string,
	expected T,
	errExpected bool,
	handler F,
) {
	t.Run(env, func(t *testing.T) {
		os.Setenv(env, value)
		v, err := handler(env)
		if errExpected {
			assert.Error(t, err, "there should be an error")
		} else {
			assert.NoError(t, err, "there should be no error")
		}
		assert.Equal(t, expected, v, "they should be equal")
	})
}

// runEmptyTest provides a generic test run for testing an empty environment
// variable on convertor function types that retun an error.
func runEmptyTest[T TestReturnValueType, F func(string) (T, error)](t *testing.T, expected T, handler F) {
	t.Run("TEST_NON_EXISTANT does not exist", func(t *testing.T) {
		v, err := handler("TEST_NON_EXISTANT")
		assert.Error(t, err, "there should be an error")
		assert.Equal(t, expected, v, "they should be equal")
	})
}

// runWithDefaultTest provides a generic test run for the convertor function
// types that return a default value.
func runWithDefaultTest[T TestReturnValueType, F func(string, T) T](
	t *testing.T,
	env string,
	value string,
	expected T,
	defaultValue T,
	handler F,
) {
	t.Run(env, func(t *testing.T) {
		os.Setenv(env, value)
		v := handler(env, defaultValue)
		if v != expected {
			assert.Equal(t, defaultValue, v, "they should be equal")
		} else {
			assert.Equal(t, expected, v, "they should be equal")
		}
	})
}

// runWithDefaultEmptyTest provides a generic test run for testing an empty
// environment variable on convertor function types that retun a default
// value.
func runWithDefaultEmptyTest[T TestReturnValueType, F func(string, T) T](t *testing.T, expected T, handler F) {
	t.Run("TEST_NON_EXISTANT does not exist", func(t *testing.T) {
		v := handler("TEST_NON_EXISTANT", expected)
		assert.Equal(t, expected, v, "they should be equal")
	})
}

// runWithDefaultTest provides a generic test run for the convertor function
// types that return a default value.
func runSliceTest[T TestReturnValueType, F func(string, string) ([]T, error)](
	t *testing.T,
	env string,
	value string,
	separator string,
	expected []T,
	errExpected bool,
	handler F,
) {
	t.Run(env, func(t *testing.T) {
		os.Setenv(env, value)
		v, err := handler(env, separator)
		if errExpected {
			assert.Error(t, err, "there should be an error")
		} else {
			assert.NoError(t, err, "there should be no error")
		}
		assert.Equal(t, expected, v, "they should be equal")
	})
}

// runSliceEmptyTest provides a generic test run for testing an empty environment
// variable on convertor function types that return a an error.
func runSliceEmptyTest[T TestReturnValueType, F func(string, string) ([]T, error)](t *testing.T, separator string, expected []T, handler F) {
	t.Run("TEST_NON_EXISTANT does not exist", func(t *testing.T) {
		v, err := handler("TEST_NON_EXISTANT", separator)
		assert.Error(t, err, "there should be an error")
		assert.Equal(t, expected, v, "they should be equal")
	})
}

// runSliceWithDefaultTest provides a generic test run for the convertor function
// types that return a default value.
func runSliceWithDefaultTest[T TestReturnValueType, F func(string, string, []T) []T](
	t *testing.T,
	env string,
	value string,
	separator string,
	expected []T,
	defaultValue []T,
	handler F,
) {
	t.Run(env, func(t *testing.T) {
		os.Setenv(env, value)
		v := handler(env, separator, defaultValue)
		if !slicesEqual(v, expected) {
			assert.Equal(t, defaultValue, v, "they should be equal")
		} else {
			assert.Equal(t, expected, v, "they should be equal")
		}
	})
}

// runSliceWithDefaultEmptyTest provides a generic test run for testing an empty
// environment variable on convertor function types that retun a default
// value.
func runSliceWithDefaultEmptyTest[T TestReturnValueType, F func(string, string, []T) []T](
	t *testing.T,
	separator string,
	expected []T,
	handler F,
) {
	t.Run("TEST_NON_EXISTANT does not exist", func(t *testing.T) {
		v := handler("TEST_NON_EXISTANT", separator, expected)
		assert.Equal(t, expected, v, "they should be equal")
	})
}
