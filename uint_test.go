package envconv_test

import (
	"testing"

	"github.com/rmhubbert/envconv"
)

func TestToUint(t *testing.T) {
	testData := []struct {
		env           string
		value         string
		expected      uint
		errorExpected bool
	}{
		{"TEST_UINT_0", "0", 0, false},
		{"TEST_UINT_-1", "-1", 0, true},
		{"TEST_UINT_18446744073709551615", "18446744073709551615", 18446744073709551615, false},
		{"TEST_UINT_18446744073709551616", "18446744073709551616", 0, true},
		{"TEST_UINT_NOTANUMBER", "notanumber", 0, true},
	}

	for _, td := range testData {
		runTest(t, td.env, td.value, td.expected, td.errorExpected, envconv.ToUint)
	}
	runEmptyTest(t, uint(0), envconv.ToUint)
}

func TestToUintSlice(t *testing.T) {
	testData := []struct {
		env           string
		value         string
		separator     string
		expected      []uint
		errorExpected bool
	}{
		{"TEST_UINT_SLICE_123_COMMA", "1,2,3", ",", []uint{1, 2, 3}, false},
		{"TEST_UINT_SLICE_123_COMMA_SPACE", "1, 2, 3", ",", []uint{1, 2, 3}, false},
		{"TEST_UINT_SLICE_123_SPACE", "1 2 3", ",", []uint{}, true},
		{"TEST_UINT_SLICE_18446744073709551615", "18446744073709551615,2,3", ",", []uint{18446744073709551615, 2, 3}, false},
		{"TEST_UINT_SLICE_18446744073709551616", "18446744073709551616,2,3", ",", []uint{}, true},
		{"TEST_UINT_SLICE_NOTANUMBER", "1,2,notanumber,4", ",", []uint{}, true},
	}

	for _, td := range testData {
		runSliceTest(t, td.env, td.value, td.separator, td.expected, td.errorExpected, envconv.ToUintSlice)
	}
	runSliceEmptyTest[uint](t, ",", []uint{}, envconv.ToUintSlice)
}

func TestToUintWithDefault(t *testing.T) {
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
		runWithDefaultTest(t, td.env, td.value, td.expected, td.defaultValue, envconv.ToUintWithDefault)
	}
	runWithDefaultEmptyTest(t, uint(0), envconv.ToUintWithDefault)
}

func TestToUintSliceWithDefault(t *testing.T) {
	defaultValue := []uint{1, 0, 5}
	testData := []struct {
		env             string
		value           string
		separator       string
		expected        []uint
		defaultValue    []uint
		defaultExpected bool
	}{
		{"TEST_UINT_SLICE_WITH_DEFAULT_123_COMMA", "1,2,3", ",", []uint{1, 2, 3}, defaultValue, false},
		{"TEST_UINT_SLICE_WITH_DEFAULT_123_COMMA_SPACE", "1, 2, 3", ", ", []uint{1, 2, 3}, defaultValue, false},
		{"TEST_UINT_SLICE_WITH_DEFAULT_123_SPACE", "1 2 3", ", ", []uint{}, defaultValue, true},
		{"TEST_UINT_SLICE_WITH_DEFAULT_18446744073709551615", "18446744073709551615,2,3", ",", []uint{18446744073709551615, 2, 3}, defaultValue, false},
		{"TEST_UINT_SLICE_WITH_DEFAULT_18446744073709551616", "18446744073709551616,2,3", ",", []uint{}, defaultValue, true},
		{"TEST_UINT_SLICE_WITH_DEFAULT_NOTANUMBER", "1,2,notanumber,4", ",", []uint{}, defaultValue, true},
	}

	for _, td := range testData {
		runSliceWithDefaultTest(t, td.env, td.value, td.separator, td.expected, td.defaultValue, td.defaultExpected, envconv.ToUintSliceWithDefault)
	}
	runSliceWithDefaultEmptyTest[uint](t, ",", defaultValue, envconv.ToUintSliceWithDefault)
}

func TestToUint8(t *testing.T) {
	testData := []struct {
		env           string
		value         string
		expected      uint8
		errorExpected bool
	}{
		{"TEST_UINT_0", "0", 0, false},
		{"TEST_UINT_-1", "-1", 0, true},
		{"TEST_UINT_255", "255", 255, false},
		{"TEST_UINT_256", "256", 0, true},
		{"TEST_UINT_NOTANUMBER", "notanumber", 0, true},
	}

	for _, td := range testData {
		runTest(t, td.env, td.value, td.expected, td.errorExpected, envconv.ToUint8)
	}
	runEmptyTest(t, uint8(0), envconv.ToUint8)
}

func TestToUint8Slice(t *testing.T) {
	testData := []struct {
		env           string
		value         string
		separator     string
		expected      []uint8
		errorExpected bool
	}{
		{"TEST_UINT8_SLICE_123_COMMA", "1,2,3", ",", []uint8{1, 2, 3}, false},
		{"TEST_UINT8_SLICE_123_COMMA_SPACE", "1, 2, 3", ",", []uint8{1, 2, 3}, false},
		{"TEST_UINT8_SLICE_123_SPACE", "1 2 3", ",", []uint8{}, true},
		{"TEST_UINT8_SLICE_255", "255,2,3", ",", []uint8{255, 2, 3}, false},
		{"TEST_UINT8_SLICE_256", "256,2,3", ",", []uint8{}, true},
		{"TEST_UINT8_SLICE_NOTANUMBER", "1,2,notanumber,4", ",", []uint8{}, true},
	}

	for _, td := range testData {
		runSliceTest(t, td.env, td.value, td.separator, td.expected, td.errorExpected, envconv.ToUint8Slice)
	}
	runSliceEmptyTest[uint8](t, ",", []uint8{}, envconv.ToUint8Slice)
}

func TestToUint8WithDefault(t *testing.T) {
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
		runWithDefaultTest(t, td.env, td.value, td.expected, td.defaultValue, envconv.ToUint8WithDefault)
	}
	runWithDefaultEmptyTest(t, uint8(0), envconv.ToUint8WithDefault)
}

func TestToUint8SliceWithDefault(t *testing.T) {
	defaultValue := []uint8{1, 0, 5}
	testData := []struct {
		env             string
		value           string
		separator       string
		expected        []uint8
		defaultValue    []uint8
		defaultExpected bool
	}{
		{"TEST_UINT8_SLICE_WITH_DEFAULT_123_COMMA", "1,2,3", ",", []uint8{1, 2, 3}, defaultValue, false},
		{"TEST_UINT8_SLICE_WITH_DEFAULT_123_COMMA_SPACE", "1, 2, 3", ", ", []uint8{1, 2, 3}, defaultValue, false},
		{"TEST_UINT8_SLICE_WITH_DEFAULT_123_SPACE", "1 2 3", ", ", []uint8{}, defaultValue, true},
		{"TEST_UINT8_SLICE_WITH_DEFAULT_255", "255,2,3", ",", []uint8{255, 2, 3}, defaultValue, false},
		{"TEST_UINT8_SLICE_WITH_DEFAULT_256", "256,2,3", ",", []uint8{}, defaultValue, true},
		{"TEST_UINT8_SLICE_WITH_DEFAULT_NOTANUMBER", "1,2,notanumber,4", ",", []uint8{}, defaultValue, true},
	}

	for _, td := range testData {
		runSliceWithDefaultTest(t, td.env, td.value, td.separator, td.expected, td.defaultValue, td.defaultExpected, envconv.ToUint8SliceWithDefault)
	}
	runSliceWithDefaultEmptyTest[uint8](t, ",", defaultValue, envconv.ToUint8SliceWithDefault)
}

func TestToUint16(t *testing.T) {
	testData := []struct {
		env           string
		value         string
		expected      uint16
		errorExpected bool
	}{
		{"TEST_UINT16_0", "0", 0, false},
		{"TEST_UINT16_-1", "-1", 0, true},
		{"TEST_UINT16_65535", "65535", 65535, false},
		{"TEST_UINT16_65536", "65536", 0, true},
		{"TEST_UINT16_NOTANUMBER", "notanumber", 0, true},
	}

	for _, td := range testData {
		runTest(t, td.env, td.value, td.expected, td.errorExpected, envconv.ToUint16)
	}
	runEmptyTest(t, uint16(0), envconv.ToUint16)
}

func TestToUint16Slice(t *testing.T) {
	testData := []struct {
		env           string
		value         string
		separator     string
		expected      []uint16
		errorExpected bool
	}{
		{"TEST_UINT16_SLICE_123_COMMA", "1,2,3", ",", []uint16{1, 2, 3}, false},
		{"TEST_UINT16_SLICE_123_COMMA_SPACE", "1, 2, 3", ",", []uint16{1, 2, 3}, false},
		{"TEST_UINT16_SLICE_123_SPACE", "1 2 3", ",", []uint16{}, true},
		{"TEST_UINT16_SLICE_65535", "65535,2,3", ",", []uint16{65535, 2, 3}, false},
		{"TEST_UINT16_SLICE_65536", "65536,2,3", ",", []uint16{}, true},
		{"TEST_UINT16_SLICE_NOTANUMBER", "1,2,notanumber,4", ",", []uint16{}, true},
	}

	for _, td := range testData {
		runSliceTest(t, td.env, td.value, td.separator, td.expected, td.errorExpected, envconv.ToUint16Slice)
	}
	runSliceEmptyTest[uint16](t, ",", []uint16{}, envconv.ToUint16Slice)
}

func TestToUint16WithDefault(t *testing.T) {
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
		runWithDefaultTest(t, td.env, td.value, td.expected, td.defaultValue, envconv.ToUint16WithDefault)
	}
	runWithDefaultEmptyTest(t, uint16(0), envconv.ToUint16WithDefault)
}

func TestToUint16SliceWithDefault(t *testing.T) {
	defaultValue := []uint16{1, 0, 5}
	testData := []struct {
		env             string
		value           string
		separator       string
		expected        []uint16
		defaultValue    []uint16
		defaultExpected bool
	}{
		{"TEST_UINT16_SLICE_WITH_DEFAULT_123_COMMA", "1,2,3", ",", []uint16{1, 2, 3}, defaultValue, false},
		{"TEST_UINT16_SLICE_WITH_DEFAULT_123_COMMA_SPACE", "1, 2, 3", ", ", []uint16{1, 2, 3}, defaultValue, false},
		{"TEST_UINT16_SLICE_WITH_DEFAULT_123_SPACE", "1 2 3", ", ", []uint16{}, defaultValue, true},
		{"TEST_UINT16_SLICE_WITH_DEFAULT_65535", "65535,2,3", ",", []uint16{65535, 2, 3}, defaultValue, false},
		{"TEST_UINT16_SLICE_WITH_DEFAULT_65536", "65536,2,3", ",", []uint16{}, defaultValue, true},
		{"TEST_UINT16_SLICE_WITH_DEFAULT_NOTANUMBER", "1,2,notanumber,4", ",", []uint16{}, defaultValue, true},
	}

	for _, td := range testData {
		runSliceWithDefaultTest(t, td.env, td.value, td.separator, td.expected, td.defaultValue, td.defaultExpected, envconv.ToUint16SliceWithDefault)
	}
	runSliceWithDefaultEmptyTest[uint16](t, ",", defaultValue, envconv.ToUint16SliceWithDefault)
}

func TestToUint32(t *testing.T) {
	testData := []struct {
		env           string
		value         string
		expected      uint32
		errorExpected bool
	}{
		{"TEST_UINT32_0", "0", 0, false},
		{"TEST_UINT32_-1", "-1", 0, true},
		{"TEST_UINT32_4294967295", "4294967295", 4294967295, false},
		{"TEST_UINT32_4294967296", "4294967296", 0, true},
		{"TEST_UINT32_NOTANUMBER", "notanumber", 0, true},
	}

	for _, td := range testData {
		runTest(t, td.env, td.value, td.expected, td.errorExpected, envconv.ToUint32)
	}
	runEmptyTest(t, uint32(0), envconv.ToUint32)
}

func TestToUint32Slice(t *testing.T) {
	testData := []struct {
		env           string
		value         string
		separator     string
		expected      []uint32
		errorExpected bool
	}{
		{"TEST_UINT32_SLICE_123_COMMA", "1,2,3", ",", []uint32{1, 2, 3}, false},
		{"TEST_UINT32_SLICE_123_COMMA_SPACE", "1, 2, 3", ",", []uint32{1, 2, 3}, false},
		{"TEST_UINT32_SLICE_123_SPACE", "1 2 3", ",", []uint32{}, true},
		{"TEST_UINT32_SLICE_4294967295", "4294967295,2,3", ",", []uint32{4294967295, 2, 3}, false},
		{"TEST_UINT32_SLICE_4294967296", "4294967296,2,3", ",", []uint32{}, true},
		{"TEST_UINT32_SLICE_NOTANUMBER", "1,2,notanumber,4", ",", []uint32{}, true},
	}

	for _, td := range testData {
		runSliceTest(t, td.env, td.value, td.separator, td.expected, td.errorExpected, envconv.ToUint32Slice)
	}
	runSliceEmptyTest[uint32](t, ",", []uint32{}, envconv.ToUint32Slice)
}

func TestToUint32WithDefault(t *testing.T) {
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
		runWithDefaultTest(t, td.env, td.value, td.expected, td.defaultValue, envconv.ToUint32WithDefault)
	}
	runWithDefaultEmptyTest(t, uint32(0), envconv.ToUint32WithDefault)
}

func TestToUint32SliceWithDefault(t *testing.T) {
	defaultValue := []uint32{1, 0, 5}
	testData := []struct {
		env             string
		value           string
		separator       string
		expected        []uint32
		defaultValue    []uint32
		defaultExpected bool
	}{
		{"TEST_UINT32_SLICE_WITH_DEFAULT_123_COMMA", "1,2,3", ",", []uint32{1, 2, 3}, defaultValue, false},
		{"TEST_UINT32_SLICE_WITH_DEFAULT_123_COMMA_SPACE", "1, 2, 3", ", ", []uint32{1, 2, 3}, defaultValue, false},
		{"TEST_UINT32_SLICE_WITH_DEFAULT_123_SPACE", "1 2 3", ", ", []uint32{}, defaultValue, true},
		{"TEST_UINT32_SLICE_WITH_DEFAULT_4294967295", "4294967295,2,3", ",", []uint32{4294967295, 2, 3}, defaultValue, false},
		{"TEST_UINT32_SLICE_WITH_DEFAULT_4294967296", "4294967296,2,3", ",", []uint32{}, defaultValue, true},
		{"TEST_UINT32_SLICE_WITH_DEFAULT_NOTANUMBER", "1,2,notanumber,4", ",", []uint32{}, defaultValue, true},
	}

	for _, td := range testData {
		runSliceWithDefaultTest(t, td.env, td.value, td.separator, td.expected, td.defaultValue, td.defaultExpected, envconv.ToUint32SliceWithDefault)
	}
	runSliceWithDefaultEmptyTest[uint32](t, ",", defaultValue, envconv.ToUint32SliceWithDefault)
}

func TestToUint64(t *testing.T) {
	testData := []struct {
		env           string
		value         string
		expected      uint64
		errorExpected bool
	}{
		{"TEST_UINT64_0", "0", 0, false},
		{"TEST_UINT64_-1", "-1", 0, true},
		{"TEST_UINT64_18446744073709551615", "18446744073709551615", 18446744073709551615, false},
		{"TEST_UINT64_18446744073709551616", "18446744073709551616", 0, true},
		{"TEST_UINT64_NOTANUMBER", "notanumber", 0, true},
	}

	for _, td := range testData {
		runTest(t, td.env, td.value, td.expected, td.errorExpected, envconv.ToUint64)
	}
	runEmptyTest(t, uint64(0), envconv.ToUint64)
}

func TestToUint64Slice(t *testing.T) {
	testData := []struct {
		env           string
		value         string
		separator     string
		expected      []uint64
		errorExpected bool
	}{
		{"TEST_UINT64_SLICE_123_COMMA", "1,2,3", ",", []uint64{1, 2, 3}, false},
		{"TEST_UINT64_SLICE_123_COMMA_SPACE", "1, 2, 3", ",", []uint64{1, 2, 3}, false},
		{"TEST_UINT64_SLICE_123_SPACE", "1 2 3", ",", []uint64{}, true},
		{"TEST_UINT64_SLICE_18446744073709551615", "18446744073709551615,2,3", ",", []uint64{18446744073709551615, 2, 3}, false},
		{"TEST_UINT64_SLICE_18446744073709551616", "18446744073709551616,2,3", ",", []uint64{}, true},
		{"TEST_UINT64_SLICE_NOTANUMBER", "1,2,notanumber,4", ",", []uint64{}, true},
	}

	for _, td := range testData {
		runSliceTest(t, td.env, td.value, td.separator, td.expected, td.errorExpected, envconv.ToUint64Slice)
	}
	runSliceEmptyTest[uint64](t, ",", []uint64{}, envconv.ToUint64Slice)
}

func TestToUint64WithDefault(t *testing.T) {
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
		runWithDefaultTest(t, td.env, td.value, td.expected, td.defaultValue, envconv.ToUint64WithDefault)
	}
	runWithDefaultEmptyTest(t, uint64(0), envconv.ToUint64WithDefault)
}

func TestToUint64SliceWithDefault(t *testing.T) {
	defaultValue := []uint64{1, 0, 5}
	testData := []struct {
		env             string
		value           string
		separator       string
		expected        []uint64
		defaultValue    []uint64
		defaultExpected bool
	}{
		{"TEST_UINT64_SLICE_WITH_DEFAULT_123_COMMA", "1,2,3", ",", []uint64{1, 2, 3}, defaultValue, false},
		{"TEST_UINT64_SLICE_WITH_DEFAULT_123_COMMA_SPACE", "1, 2, 3", ", ", []uint64{1, 2, 3}, defaultValue, false},
		{"TEST_UINT64_SLICE_WITH_DEFAULT_123_SPACE", "1 2 3", ", ", []uint64{}, defaultValue, true},
		{"TEST_UINT64_SLICE_WITH_DEFAULT_18446744073709551615", "18446744073709551615,2,3", ",", []uint64{18446744073709551615, 2, 3}, defaultValue, false},
		{"TEST_UINT64_SLICE_WITH_DEFAULT_18446744073709551616", "18446744073709551616,2,3", ",", []uint64{}, defaultValue, true},
		{"TEST_UINT64_SLICE_WITH_DEFAULT_NOTANUMBER", "1,2,notanumber,4", ",", []uint64{}, defaultValue, true},
	}

	for _, td := range testData {
		runSliceWithDefaultTest(t, td.env, td.value, td.separator, td.expected, td.defaultValue, td.defaultExpected, envconv.ToUint64SliceWithDefault)
	}
	runSliceWithDefaultEmptyTest[uint64](t, ",", defaultValue, envconv.ToUint64SliceWithDefault)
}
