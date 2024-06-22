package envconv_test

import (
	"testing"

	"github.com/rmhubbert/envconv"
)

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

func TestToInt8Slice(t *testing.T) {
	testData := []struct {
		env         string
		value       string
		separator   string
		expected    []int8
		errExpected bool
	}{
		{"TEST_INT_SLICE_123_COMMA", "1,2,3", ",", []int8{1, 2, 3}, false},
		{"TEST_INT_SLICE_123_COMMA_SPACE", "1, 2, 3", ",", []int8{1, 2, 3}, false},
		{"TEST_INT_SLICE_123_SPACE", "1 2 3", ",", []int8{}, true},
		{"TEST_INT_SLICE_127", "127,2,3", ",", []int8{127, 2, 3}, false},
		{"TEST_INT_SLICE_128", "128,2,3", ",", []int8{}, true},
		{"TEST_INT_SLICE_-128", "-128,2,3", ",", []int8{-128, 2, 3}, false},
		{"TEST_INT_SLICE_-129", "-129,2,3", ",", []int8{}, true},
		{"TEST_INT_SLICE_NOTANUMBER", "1,2,notanumber,4", ",", []int8{}, true},
	}

	for _, td := range testData {
		runSliceTest(t, td.env, td.value, td.separator, td.expected, td.errExpected, envconv.ToInt8Slice)
	}
	runSliceEmptyTest[int8](t, ",", []int8{}, envconv.ToInt8Slice)
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

func TestToInt8SliceWithDefault(t *testing.T) {
	testData := []struct {
		env          string
		value        string
		separator    string
		expected     []int8
		defaultValue []int8
	}{
		{"TEST_INT_SLICE_WITH_DEFAULT_123_COMMA", "1,2,3", ",", []int8{1, 2, 3}, []int8{1, 0, 5}},
		{"TEST_INT_SLICE_WITH_DEFAULT_123_COMMA_SPACE", "1, 2, 3", ",", []int8{1, 2, 3}, []int8{1, 0, 5}},
		{"TEST_INT_SLICE_WITH_DEFAULT_123_SPACE", "1 2 3", ",", []int8{}, []int8{1, 0, 5}},
		{"TEST_INT_SLICE_WITH_DEFAULT_127", "127,2,3", ",", []int8{127, 2, 3}, []int8{1, 0, 5}},
		{"TEST_INT_SLICE_WITH_DEFAULT_128", "128,2,3", ",", []int8{}, []int8{1, 0, 5}},
		{"TEST_INT_SLICE_WITH_DEFAULT_-128", "-128,2,3", ",", []int8{-128, 2, 3}, []int8{1, 0, 5}},
		{"TEST_INT_SLICE_WITH_DEFAULT_-129", "-129,2,3", ",", []int8{}, []int8{1, 0, 5}},
		{"TEST_INT_SLICE_WITH_DEFAULT_NOTANUMBER", "1,2,notanumber,4", ",", []int8{}, []int8{1, 0, 5}},
	}

	for _, td := range testData {
		runSliceWithDefaultTest(t, td.env, td.value, td.separator, td.expected, td.defaultValue, envconv.ToInt8SliceWithDefault)
	}
	runSliceWithDefaultEmptyTest[int8](t, ",", []int8{1, 0, 5}, envconv.ToInt8SliceWithDefault)
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
