package envconv_test

import (
	"os"
	"testing"

	"github.com/rmhubbert/envconv"
	"github.com/stretchr/testify/assert"
)

type TestReturnValueType interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

// runTest provides a generic test run for the convertor function types that
// return an error
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

func TestToInt(t *testing.T) {
	testData := []struct {
		env         string
		value       string
		expected    int
		errExpected bool
	}{
		{"TEST_INT_0", "0", 0, false},
		{"TEST_INT_9223372036854775807", "9223372036854775807", 9223372036854775807, false},
		{"TEST_INT_9223372036854775808", "9223372036854775808", 0, true},
		{"TEST_INT_-9223372036854775808", "-9223372036854775808", -9223372036854775808, false},
		{"TEST_INT_-9223372036854775809", "-9223372036854775809", 0, true},
		{"TEST_INT_NOTANUMBER", "notanumber", 0, true},
	}

	for _, td := range testData {
		runTest(t, td.env, td.value, td.expected, td.errExpected, envconv.ToInt)
	}
	runEmptyTest(t, int(0), envconv.ToInt)
}

func TestToIntSlice(t *testing.T) {
	testData := []struct {
		env         string
		value       string
		separator   string
		expected    []int
		errExpected bool
	}{
		{"TEST_INT_SLICE_123_COMMA", "1,2,3", ",", []int{1, 2, 3}, false},
		{"TEST_INT_SLICE_123_COMMA_SPACE", "1, 2, 3", ",", []int{1, 2, 3}, false},
		{"TEST_INT_SLICE_123_SPACE", "1 2 3", ",", []int{}, true},
		{"TEST_INT_SLICE_9223372036854775807", "9223372036854775807,2,3", ",", []int{9223372036854775807, 2, 3}, false},
		{"TEST_INT_SLICE_9223372036854775808", "9223372036854775808,2,3", ",", []int{}, true},
		{"TEST_INT_SLICE_-9223372036854775808", "-9223372036854775808,2,3", ",", []int{-9223372036854775808, 2, 3}, false},
		{"TEST_INT_SLICE_-9223372036854775809", "-9223372036854775809,2,3", ",", []int{}, true},
		{"TEST_INT_SLICE_NOTANUMBER", "1,2,notanumber,4", ",", []int{}, true},
	}

	for _, td := range testData {
		runSliceTest(t, td.env, td.value, td.separator, td.expected, td.errExpected, envconv.ToIntSlice)
	}
	runSliceEmptyTest[int](t, ",", []int{}, envconv.ToIntSlice)
}

func TestToIntWithDefault(t *testing.T) {
	testData := []struct {
		env          string
		value        string
		expected     int
		defaultValue int
	}{
		{"TEST_INT_WITH_DEFAULT_0", "0", 0, 105},
		{"TEST_INT_WITH_DEFAULT_9223372036854775807", "9223372036854775807", 9223372036854775807, 105},
		{"TEST_INT_WITH_DEFAULT_9223372036854775808", "9223372036854775808", 0, 105},
		{"TEST_INT_WITH_DEFAULT_-9223372036854775808", "-9223372036854775808", -9223372036854775808, 105},
		{"TEST_INT_WITH_DEFAULT_-9223372036854775809", "-9223372036854775809", 0, 105},
		{"TEST_INT_WITH_DEFAULT_NOTANUMBER", "notanumber", 0, 105},
	}

	for _, td := range testData {
		runWithDefaultTest(t, td.env, td.value, td.expected, td.defaultValue, envconv.ToIntWithDefault)
	}
	runWithDefaultEmptyTest(t, int(0), envconv.ToIntWithDefault)
}

func TestToIntSliceWithDefault(t *testing.T) {
	testData := []struct {
		env          string
		value        string
		separator    string
		expected     []int
		defaultValue []int
	}{
		{"TEST_INT_SLICE_WITH_DEFAULT_123_COMMA", "1,2,3", ", ", []int{1, 2, 3}, []int{1, 0, 5}},
		{"TEST_INT_SLICE_WITH_DEFAULT_123_COMMA_SPACE", "1, 2, 3", ", ", []int{1, 2, 3}, []int{1, 0, 5}},
		{"TEST_INT_SLICE_WITH_DEFAULT_123_SPACE", "1 2 3", ", ", []int{}, []int{1, 0, 5}},
		{"TEST_INT_SLICE_WITH_DEFAULT_9223372036854775807", "9223372036854775807,2,3", ", ", []int{9223372036854775807, 2, 3}, []int{1, 0, 5}},
		{"TEST_INT_SLICE_WITH_DEFAULT_9223372036854775808", "9223372036854775808,2,3", ", ", []int{}, []int{1, 0, 5}},
		{"TEST_INT_SLICE_WITH_DEFAULT_-9223372036854775808", "-9223372036854775808,2,3", ", ", []int{-9223372036854775808, 2, 3}, []int{1, 0, 5}},
		{"TEST_INT_SLICE_WITH_DEFAULT_-9223372036854775809", "-9223372036854775809,2,3", ", ", []int{}, []int{1, 0, 5}},
		{"TEST_INT_SLICE_WITH_DEFAULT_NOTANUMBER", "1,2,notanumber,4", ", ", []int{}, []int{1, 0, 5}},
	}

	for _, td := range testData {
		runSliceWithDefaultTest(t, td.env, td.value, td.separator, td.expected, td.defaultValue, envconv.ToIntSliceWithDefault)
	}
	runSliceWithDefaultEmptyTest[int](t, ",", []int{1, 0, 5}, envconv.ToIntSliceWithDefault)
}

func TestToInt8(t *testing.T) {
	testData := []struct {
		env         string
		value       string
		expected    int8
		errExpected bool
	}{
		{"TEST_INT8_0", "0", 0, false},
		{"TEST_INT8_127", "127", 127, false},
		{"TEST_INT8_128", "128", 0, true},
		{"TEST_INT8_-128", "-128", -128, false},
		{"TEST_INT8_-129", "-129", 0, true},
		{"TEST_INT8_NOTANUMBER", "notanumber", 0, true},
	}

	for _, td := range testData {
		runTest(t, td.env, td.value, td.expected, td.errExpected, envconv.ToInt8)
	}
	runEmptyTest(t, int8(0), envconv.ToInt8)
}

func TestToInt8WithDefault(t *testing.T) {
	testData := []struct {
		env          string
		value        string
		expected     int8
		defaultValue int8
	}{
		{"TEST_INT8_WITH_DEFAULT_0", "0", 0, 105},
		{"TEST_INT8_WITH_DEFAULT_127", "127", 127, 105},
		{"TEST_INT8_WITH_DEFAULT_128", "128", 0, 105},
		{"TEST_INT8_WITH_DEFAULT_-128", "-128", -128, 105},
		{"TEST_INT8_WITH_DEFAULT_-129", "-129", 0, 105},
		{"TEST_INT8_WITH_DEFAULT_NOTANUMBER", "notanumber", 1, 105},
	}

	for _, td := range testData {
		runWithDefaultTest(t, td.env, td.value, td.expected, td.defaultValue, envconv.ToInt8WithDefault)
	}
	runWithDefaultEmptyTest(t, int8(0), envconv.ToInt8WithDefault)
}

func TestToInt16(t *testing.T) {
	testData := []struct {
		env         string
		value       string
		expected    int16
		errExpected bool
	}{
		{"TEST_INT16_0", "0", 0, false},
		{"TEST_INT16_32767", "32767", 32767, false},
		{"TEST_INT16_32768", "32768", 0, true},
		{"TEST_INT16_-32768", "-32768", -32768, false},
		{"TEST_INT16_-32769", "-32769", 0, true},
		{"TEST_INT16_NOTANUMBER", "notanumber", 0, true},
	}

	for _, td := range testData {
		runTest(t, td.env, td.value, td.expected, td.errExpected, envconv.ToInt16)
	}
	runEmptyTest(t, int16(0), envconv.ToInt16)
}

func TestToInt16WithDefault(t *testing.T) {
	testData := []struct {
		env          string
		value        string
		expected     int16
		defaultValue int16
	}{
		{"TEST_INT16_WITH_DEFAULT_0", "0", 0, 105},
		{"TEST_INT16_WITH_DEFAULT_32767", "32767", 32767, 105},
		{"TEST_INT16_WITH_DEFAULT_32768", "32768", 0, 105},
		{"TEST_INT16_WITH_DEFAULT_-32768", "-32767", -32767, 105},
		{"TEST_INT16_WITH_DEFAULT_-32769", "-32769", 0, 105},
		{"TEST_INT16_WITH_DEFAULT_NOTANUMBER", "notanumber", 1, 105},
	}

	for _, td := range testData {
		runWithDefaultTest(t, td.env, td.value, td.expected, td.defaultValue, envconv.ToInt16WithDefault)
	}
	runWithDefaultEmptyTest(t, int16(0), envconv.ToInt16WithDefault)
}

func TestToInt32(t *testing.T) {
	testData := []struct {
		env         string
		value       string
		expected    int32
		errExpected bool
	}{
		{"TEST_INT32_0", "0", 0, false},
		{"TEST_INT32_2147483647", "2147483647", 2147483647, false},
		{"TEST_INT32_2147483648", "2147483648", 0, true},
		{"TEST_INT32_-2147483648", "-2147483648", -2147483648, false},
		{"TEST_INT32_-2147483649", "-2147483649", 0, true},
		{"TEST_INT32_NOTANUMBER", "notanumber", 0, true},
	}

	for _, td := range testData {
		runTest(t, td.env, td.value, td.expected, td.errExpected, envconv.ToInt32)
	}
	runEmptyTest(t, int32(0), envconv.ToInt32)
}

func TestToInt32WithDefault(t *testing.T) {
	testData := []struct {
		env          string
		value        string
		expected     int32
		defaultValue int32
	}{
		{"TEST_INT32_WITH_DEFAULT_0", "0", 0, 105},
		{"TEST_INT32_WITH_DEFAULT_2147483647", "2147483647", 2147483647, 105},
		{"TEST_INT32_WITH_DEFAULT_2147483648", "2147483648", 0, 105},
		{"TEST_INT32_WITH_DEFAULT_-2147483648", "-2147483648", -2147483648, 105},
		{"TEST_INT32_WITH_DEFAULT_-2147483649", "-2147483649", 0, 105},
		{"TEST_INT32_WITH_DEFAULT_NOTANUMBER", "notanumber", 0, 105},
	}

	for _, td := range testData {
		runWithDefaultTest(t, td.env, td.value, td.expected, td.defaultValue, envconv.ToInt32WithDefault)
	}
	runWithDefaultEmptyTest(t, int32(0), envconv.ToInt32WithDefault)
}

func TestToInt64(t *testing.T) {
	testData := []struct {
		env         string
		value       string
		expected    int64
		errExpected bool
	}{
		{"TEST_INT64_0", "0", 0, false},
		{"TEST_INT64_9223372036854775807", "9223372036854775807", 9223372036854775807, false},
		{"TEST_INT64_9223372036854775808", "9223372036854775808", 0, true},
		{"TEST_INT64_-9223372036854775808", "-9223372036854775808", -9223372036854775808, false},
		{"TEST_INT64_-9223372036854775809", "-9223372036854775809", 0, true},
		{"TEST_INT64_NOTANUMBER", "notanumber", 0, true},
	}

	for _, td := range testData {
		runTest(t, td.env, td.value, td.expected, td.errExpected, envconv.ToInt64)
	}
	runEmptyTest(t, int64(0), envconv.ToInt64)
}

func TestToInt64WithDefault(t *testing.T) {
	testData := []struct {
		env          string
		value        string
		expected     int64
		defaultValue int64
	}{
		{"TEST_INT64_WITH_DEFAULT_0", "0", 0, 105},
		{"TEST_INT64_WITH_DEFAULT_9223372036854775807", "9223372036854775807", 9223372036854775807, 105},
		{"TEST_INT64_WITH_DEFAULT_9223372036854775808", "9223372036854775808", 0, 105},
		{"TEST_INT64_WITH_DEFAULT_-9223372036854775808", "-9223372036854775808", -9223372036854775808, 105},
		{"TEST_INT64_WITH_DEFAULT_-9223372036854775809", "-9223372036854775809", 0, 105},
		{"TEST_INT64_WITH_DEFAULT_NOTANUMBER", "notanumber", 0, 105},
	}

	for _, td := range testData {
		runWithDefaultTest(t, td.env, td.value, td.expected, td.defaultValue, envconv.ToInt64WithDefault)
	}
	runWithDefaultEmptyTest(t, int64(0), envconv.ToInt64WithDefault)
}

func TestToUInt(t *testing.T) {
	testData := []struct {
		env         string
		value       string
		expected    uint
		errExpected bool
	}{
		{"TEST_UINT_0", "0", 0, false},
		{"TEST_UINT_-1", "-1", 0, true},
		{"TEST_UINT_18446744073709551615", "18446744073709551615", 18446744073709551615, false},
		{"TEST_UINT_18446744073709551616", "18446744073709551616", 0, true},
		{"TEST_UINT_NOTANUMBER", "notanumber", 0, true},
	}

	for _, td := range testData {
		runTest(t, td.env, td.value, td.expected, td.errExpected, envconv.ToUInt)
	}
	runEmptyTest(t, uint(0), envconv.ToUInt)
}

func TestToUIntWithDefault(t *testing.T) {
	testData := []struct {
		env          string
		value        string
		expected     uint
		defaultValue uint
	}{
		{"TEST_UINT_WITH_DEFAULT_0", "0", 0, 105},
		{"TEST_UINT_WITH_DEFAULT_18446744073709551615", "18446744073709551615", 18446744073709551615, 105},
		{"TEST_UINT_WITH_DEFAULT_18446744073709551616", "18446744073709551616", 0, 105},
		{"TEST_UINT_WITH_DEFAULT_NOTANUMBER", "notanumber", 0, 105},
	}

	for _, td := range testData {
		runWithDefaultTest(t, td.env, td.value, td.expected, td.defaultValue, envconv.ToUIntWithDefault)
	}
	runWithDefaultEmptyTest(t, uint(0), envconv.ToUIntWithDefault)
}

func TestToUInt8(t *testing.T) {
	testData := []struct {
		env         string
		value       string
		expected    uint8
		errExpected bool
	}{
		{"TEST_UINT_0", "0", 0, false},
		{"TEST_UINT_-1", "-1", 0, true},
		{"TEST_UINT_255", "255", 255, false},
		{"TEST_UINT_256", "256", 0, true},
		{"TEST_UINT_NOTANUMBER", "notanumber", 0, true},
	}

	for _, td := range testData {
		runTest(t, td.env, td.value, td.expected, td.errExpected, envconv.ToUInt8)
	}
	runEmptyTest(t, uint8(0), envconv.ToUInt8)
}

func TestToUInt8WithDefault(t *testing.T) {
	testData := []struct {
		env          string
		value        string
		expected     uint8
		defaultValue uint8
	}{
		{"TEST_UINT_WITH_DEFAULT_0", "0", 0, 105},
		{"TEST_UINT_WITH_DEFAULT_255", "255", 255, 105},
		{"TEST_UINT_WITH_DEFAULT_256", "256", 0, 105},
		{"TEST_UINT_WITH_DEFAULT_NOTANUMBER", "notanumber", 0, 105},
	}

	for _, td := range testData {
		runWithDefaultTest(t, td.env, td.value, td.expected, td.defaultValue, envconv.ToUInt8WithDefault)
	}
	runWithDefaultEmptyTest(t, uint8(0), envconv.ToUInt8WithDefault)
}

func TestToUInt16(t *testing.T) {
	testData := []struct {
		env         string
		value       string
		expected    uint16
		errExpected bool
	}{
		{"TEST_UINT16_0", "0", 0, false},
		{"TEST_UINT16_-1", "-1", 0, true},
		{"TEST_UINT16_65535", "65535", 65535, false},
		{"TEST_UINT16_65536", "65536", 0, true},
		{"TEST_UINT16_NOTANUMBER", "notanumber", 0, true},
	}

	for _, td := range testData {
		runTest(t, td.env, td.value, td.expected, td.errExpected, envconv.ToUInt16)
	}
	runEmptyTest(t, uint16(0), envconv.ToUInt16)
}

func TestToUInt16WithDefault(t *testing.T) {
	testData := []struct {
		env          string
		value        string
		expected     uint16
		defaultValue uint16
	}{
		{"TEST_UINT16_WITH_DEFAULT_0", "0", 0, 105},
		{"TEST_UINT16_WITH_DEFAULT_65535", "65535", 65535, 105},
		{"TEST_UINT16_WITH_DEFAULT_65536", "65536", 0, 105},
		{"TEST_UINT16_WITH_DEFAULT_NOTANUMBER", "notanumber", 0, 105},
	}

	for _, td := range testData {
		runWithDefaultTest(t, td.env, td.value, td.expected, td.defaultValue, envconv.ToUInt16WithDefault)
	}
	runWithDefaultEmptyTest(t, uint16(0), envconv.ToUInt16WithDefault)
}

func TestToUInt32(t *testing.T) {
	testData := []struct {
		env         string
		value       string
		expected    uint32
		errExpected bool
	}{
		{"TEST_UINT32_0", "0", 0, false},
		{"TEST_UINT32_-1", "-1", 0, true},
		{"TEST_UINT32_4294967295", "4294967295", 4294967295, false},
		{"TEST_UINT32_4294967296", "4294967296", 0, true},
		{"TEST_UINT32_NOTANUMBER", "notanumber", 0, true},
	}

	for _, td := range testData {
		runTest(t, td.env, td.value, td.expected, td.errExpected, envconv.ToUInt32)
	}
	runEmptyTest(t, uint32(0), envconv.ToUInt32)
}

func TestToUInt32WithDefault(t *testing.T) {
	testData := []struct {
		env          string
		value        string
		expected     uint32
		defaultValue uint32
	}{
		{"TEST_UINT32_WITH_DEFAULT_0", "0", 0, 105},
		{"TEST_UINT32_WITH_DEFAULT_4294967295", "4294967295", 4294967295, 105},
		{"TEST_UINT32_WITH_DEFAULT_4294967296", "4294967296", 0, 105},
		{"TEST_UINT32_WITH_DEFAULT_NOTANUMBER", "notanumber", 0, 105},
	}

	for _, td := range testData {
		runWithDefaultTest(t, td.env, td.value, td.expected, td.defaultValue, envconv.ToUInt32WithDefault)
	}
	runWithDefaultEmptyTest(t, uint32(0), envconv.ToUInt32WithDefault)
}

func TestToUInt64(t *testing.T) {
	testData := []struct {
		env         string
		value       string
		expected    uint64
		errExpected bool
	}{
		{"TEST_UINT64_0", "0", 0, false},
		{"TEST_UINT64_-1", "-1", 0, true},
		{"TEST_UINT64_4294967295", "18446744073709551615", 18446744073709551615, false},
		{"TEST_UINT64_4294967296", "18446744073709551616", 0, true},
		{"TEST_UINT64_NOTANUMBER", "notanumber", 0, true},
	}

	for _, td := range testData {
		runTest(t, td.env, td.value, td.expected, td.errExpected, envconv.ToUInt64)
	}
	runEmptyTest(t, uint64(0), envconv.ToUInt64)
}

func TestToUInt64WithDefault(t *testing.T) {
	testData := []struct {
		env          string
		value        string
		expected     uint64
		defaultValue uint64
	}{
		{"TEST_UINT64_WITH_DEFAULT_0", "0", 0, 105},
		{"TEST_UINT64_WITH_DEFAULT_18446744073709551615", "18446744073709551615", 18446744073709551615, 105},
		{"TEST_UINT64_WITH_DEFAULT_18446744073709551616", "18446744073709551616", 0, 105},
		{"TEST_UINT64_WITH_DEFAULT_NOTANUMBER", "notanumber", 0, 105},
	}

	for _, td := range testData {
		runWithDefaultTest(t, td.env, td.value, td.expected, td.defaultValue, envconv.ToUInt64WithDefault)
	}
	runWithDefaultEmptyTest(t, uint64(0), envconv.ToUInt64WithDefault)
}
