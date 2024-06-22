package envconv_test

import (
	"testing"

	"github.com/rmhubbert/envconv"
)

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
