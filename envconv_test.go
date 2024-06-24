package envconv_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// testReturnValueType is a type constraint interface that is used by the run* generic
// test functions
type testReturnValueType interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64 | bool
}

// runTest provides a generic test run for the convertor function types that
// take a single string and return a T and an error
func runTest[T testReturnValueType, F func(string) (T, error)](
	t *testing.T,
	env string,
	value string,
	expected T,
	errorExpected bool,
	handler F,
) {
	t.Run(env, func(t *testing.T) {
		os.Setenv(env, value)
		v, err := handler(env)
		if errorExpected {
			assert.Error(t, err, "there should be an error")
		} else {
			assert.NoError(t, err, "there should be no error")
		}
		assert.Equal(t, expected, v, "they should be equal")
	})
}

// runEmptyTest provides a generic test run for testing an empty environment
// variable on convertor function types that retun an error.
func runEmptyTest[T testReturnValueType, F func(string) (T, error)](t *testing.T, expected T, handler F) {
	t.Run("TEST_NON_EXISTANT does not exist", func(t *testing.T) {
		v, err := handler("TEST_NON_EXISTANT")
		assert.Error(t, err, "there should be an error")
		assert.Equal(t, expected, v, "they should be equal")
	})
}

// runWithDefaultTest provides a generic test run for the convertor function
// types that return a default value.
func runWithDefaultTest[T testReturnValueType, F func(string, T) T](
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
func runWithDefaultEmptyTest[T testReturnValueType, F func(string, T) T](t *testing.T, expected T, handler F) {
	t.Run("TEST_NON_EXISTANT does not exist", func(t *testing.T) {
		v := handler("TEST_NON_EXISTANT", expected)
		assert.Equal(t, expected, v, "they should be equal")
	})
}

// runWithDefaultTest provides a generic test run for the convertor function
// types that return a default value.
func runSliceTest[T testReturnValueType, F func(string, string) ([]T, error)](
	t *testing.T,
	env string,
	value string,
	separator string,
	expected []T,
	errorExpected bool,
	handler F,
) {
	t.Run(env, func(t *testing.T) {
		os.Setenv(env, value)
		v, err := handler(env, separator)
		if errorExpected {
			assert.Error(t, err, "there should be an error")
		} else {
			assert.NoError(t, err, "there should be no error")
		}
		assert.Equal(t, expected, v, "they should be equal")
	})
}

// runSliceEmptyTest provides a generic test run for testing an empty environment
// variable on convertor function types that return a an error.
func runSliceEmptyTest[T testReturnValueType, F func(string, string) ([]T, error)](t *testing.T, separator string, expected []T, handler F) {
	t.Run("TEST_NON_EXISTANT does not exist", func(t *testing.T) {
		v, err := handler("TEST_NON_EXISTANT", separator)
		assert.Error(t, err, "there should be an error")
		assert.Equal(t, expected, v, "they should be equal")
	})
}

// runSliceWithDefaultTest provides a generic test run for the convertor function
// types that return a default value.
func runSliceWithDefaultTest[T testReturnValueType, F func(string, string, []T) []T](
	t *testing.T,
	env string,
	value string,
	separator string,
	expected []T,
	defaultValue []T,
	defaultExpected bool,
	handler F,
) {
	t.Run(env, func(t *testing.T) {
		os.Setenv(env, value)
		v := handler(env, separator, defaultValue)
		if defaultExpected {
			assert.Equal(t, defaultValue, v, "they should be equal")
		} else {
			assert.Equal(t, expected, v, "they should be equal")
		}
	})
}

// runSliceWithDefaultEmptyTest provides a generic test run for testing an empty
// environment variable on convertor function types that retun a default
// value.
func runSliceWithDefaultEmptyTest[T testReturnValueType, F func(string, string, []T) []T](
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
