package envconv_test

import (
	"testing"

	"github.com/rmhubbert/envconv"
)

func TestToFloat32(t *testing.T) {
	testData := []struct {
		env           string
		value         string
		expected      float32
		errorExpected bool
	}{
		{"TEST_FLOAT32_0", "0", 0, false},
		{"TEST_FLOAT32_3.40282346638528859811704183484516925440000000000000e+38", "3.40282346638528859811704183484516925440000000000000e+38", 3.40282346638528859811704183484516925440000000000000e+38, false},
		{"TEST_FLOAT32_3.50282346638528859811704183484516925440000000000000e+38", "3.50282346638528859811704183484516925440000000000000e+38", 0, true},
		{"TEST_FLOAT32_1.40129846432481707092372958328991613128026194187652e-45", "1.40129846432481707092372958328991613128026194187652e-45", 1.40129846432481707092372958328991613128026194187652e-45, false},
		{"TEST_FLOAT32_NOTANUMBER", "notanumber", 0, true},
	}

	for _, td := range testData {
		runTest(t, td.env, td.value, td.expected, td.errorExpected, envconv.ToFloat32)
	}
	runEmptyTest(t, float32(0), envconv.ToFloat32)
}

func TestToFloat32Slice(t *testing.T) {
	testData := []struct {
		env           string
		value         string
		separator     string
		expected      []float32
		errorExpected bool
	}{
		{"TEST_FLOAT32_SLICE_123_COMMA", "1.05,2.123,3", ",", []float32{1.05, 2.123, 3}, false},
		{"TEST_FLOAT32_SLICE_123_COMMA_SPACE", "1.05, 2.123, 3", ",", []float32{1.05, 2.123, 3}, false},
		{"TEST_FLOAT32_SLICE_123_SPACE", "1.05 2.123 3", ",", []float32{}, true},
		{"TEST_FLOAT32_SLICE_3.40282346638528859811704183484516925440000000000000e+38", "3.40282346638528859811704183484516925440000000000000e+38,2.123,3", ",", []float32{3.40282346638528859811704183484516925440000000000000e+38, 2.123, 3}, false},
		{"TEST_FLOAT32_SLICE_3.50282346638528859811704183484516925440000000000000e+38", "3.50282346638528859811704183484516925440000000000000e+38,2.123,3", ",", []float32{}, true},
		{"TEST_FLOAT32_SLICE_1.40129846432481707092372958328991613128026194187652e-45", "1.40129846432481707092372958328991613128026194187652e-45,2.123,3", ",", []float32{1.40129846432481707092372958328991613128026194187652e-45, 2.123, 3}, false},
		{"TEST_FLOAT32_SLICE_NOTANUMBER", "1,2,notanumber,4", ",", []float32{}, true},
	}

	for _, td := range testData {
		runSliceTest(t, td.env, td.value, td.separator, td.expected, td.errorExpected, envconv.ToFloat32Slice)
	}
	runSliceEmptyTest[float32](t, ",", []float32{}, envconv.ToFloat32Slice)
}

func TestToFloat32WithDefault(t *testing.T) {
	testData := []struct {
		env          string
		value        string
		expected     float32
		defaultValue float32
	}{
		{"TEST_FLOAT32_WITH_DEFAULT_0", "0", 0, 105},
		{"TEST_FLOAT32_WITH_DEFAULT_3.40282346638528859811704183484516925440000000000000e+38", "3.40282346638528859811704183484516925440000000000000e+38", 3.40282346638528859811704183484516925440000000000000e+38, 105},
		{"TEST_FLOAT32_WITH_DEFAULT_3.40282346638528859811704183484516925440000000000000e+39", "3.40282346638528859811704183484516925440000000000000e+39", 0, 105},
		{"TEST_FLOAT32_WITH_DEFAULT_1.40129846432481707092372958328991613128026194187652e-45", "1.40129846432481707092372958328991613128026194187652e-45", 1.40129846432481707092372958328991613128026194187652e-45, 105},
		{"TEST_FLOAT32_WITH_DEFAULT_NOTANUMBER", "notanumber", 0, 105},
	}

	for _, td := range testData {
		runWithDefaultTest(t, td.env, td.value, td.expected, td.defaultValue, envconv.ToFloat32WithDefault)
	}
	runWithDefaultEmptyTest(t, float32(0), envconv.ToFloat32WithDefault)
}

func TestToFloat32SliceWithDefault(t *testing.T) {
	defaultValue := []float32{1, 0, 5}
	testData := []struct {
		env             string
		value           string
		separator       string
		expected        []float32
		defaultValue    []float32
		defaultExpected bool
	}{
		{"TEST_FLOAT32_SLICE_WITH_DEFAULT_123_COMMA", "1,2,3", ",", []float32{1, 2, 3}, defaultValue, false},
		{"TEST_FLOAT32_SLICE_WITH_DEFAULT_123_COMMA_SPACE", "1, 2, 3", ",", []float32{1, 2, 3}, defaultValue, false},
		{"TEST_FLOAT32_SLICE_WITH_DEFAULT_123_SPACE", "1 2 3", ",", []float32{}, defaultValue, true},
		{"TEST_FLOAT32_SLICE_WITH_DEFAULT_3.40282346638528859811704183484516925440000000000000e+38", "3.40282346638528859811704183484516925440000000000000e+38,2,3", ",", []float32{3.40282346638528859811704183484516925440000000000000e+38, 2, 3}, defaultValue, false},
		{"TEST_FLOAT32_SLICE_WITH_DEFAULT_3.40282346638528859811704183484516925440000000000000e+39", "3.40282346638528859811704183484516925440000000000000e+39,2,3", ",", []float32{}, defaultValue, true},
		{"TEST_FLOAT32_SLICE_WITH_DEFAULT_1.40129846432481707092372958328991613128026194187652e-45", "1.40129846432481707092372958328991613128026194187652e-45,2,3", ",", []float32{1.40129846432481707092372958328991613128026194187652e-45, 2, 3}, defaultValue, false},
		{"TEST_FLOAT32_SLICE_WITH_DEFAULT_NOTANUMBER", "1,2,notanumber,4", ",", []float32{}, defaultValue, true},
	}

	for _, td := range testData {
		runSliceWithDefaultTest(t, td.env, td.value, td.separator, td.expected, td.defaultValue, td.defaultExpected, envconv.ToFloat32SliceWithDefault)
	}
	runSliceWithDefaultEmptyTest[float32](t, ",", defaultValue, envconv.ToFloat32SliceWithDefault)
}

func TestToFloat64(t *testing.T) {
	testData := []struct {
		env           string
		value         string
		expected      float64
		errorExpected bool
	}{
		{"TEST_FLOAT64_0", "0", 0, false},
		{"TEST_FLOAT64_1.79769313486231570814527423731704356798070567525845e+308", "1.79769313486231570814527423731704356798070567525845e+308", 1.79769313486231570814527423731704356798070567525845e+308, false},
		{"TEST_FLOAT64_1.79769313486231570814527423731704356798070567525845e+309", "1.79769313486231570814527423731704356798070567525845e+309", 0, true},
		{"TEST_FLOAT64_4.94065645841246544176568792868221372365059802614325e-324", "4.94065645841246544176568792868221372365059802614325e-324", 4.94065645841246544176568792868221372365059802614325e-324, false},
		{"TEST_FLOAT64_NOTANUMBER", "notanumber", 0, true},
	}

	for _, td := range testData {
		runTest(t, td.env, td.value, td.expected, td.errorExpected, envconv.ToFloat64)
	}
	runEmptyTest(t, float64(0), envconv.ToFloat64)
}

func TestToFloat64Slice(t *testing.T) {
	testData := []struct {
		env           string
		value         string
		separator     string
		expected      []float64
		errorExpected bool
	}{
		{"TEST_FLOAT64_SLICE_123_COMMA", "1.05,2.123,3", ",", []float64{1.05, 2.123, 3}, false},
		{"TEST_FLOAT64_SLICE_123_COMMA_SPACE", "1.05, 2.123, 3", ",", []float64{1.05, 2.123, 3}, false},
		{"TEST_FLOAT64_SLICE_123_SPACE", "1.05 2.123 3", ",", []float64{}, true},
		{"TEST_FLOAT64_SLICE_1.79769313486231570814527423731704356798070567525845e+308", "1.79769313486231570814527423731704356798070567525845e+308,2.123,3", ",", []float64{1.79769313486231570814527423731704356798070567525845e+308, 2.123, 3}, false},
		{"TEST_FLOAT64_SLICE_1.79769313486231570814527423731704356798070567525845e+309", "1.79769313486231570814527423731704356798070567525845e+309,2.123,3", ",", []float64{}, true},
		{"TEST_FLOAT64_SLICE_4.94065645841246544176568792868221372365059802614325e-324", "4.94065645841246544176568792868221372365059802614325e-324,2.123,3", ",", []float64{4.94065645841246544176568792868221372365059802614325e-324, 2.123, 3}, false},
		{"TEST_FLOAT64_SLICE_NOTANUMBER", "1,2,notanumber,4", ",", []float64{}, true},
	}

	for _, td := range testData {
		runSliceTest(t, td.env, td.value, td.separator, td.expected, td.errorExpected, envconv.ToFloat64Slice)
	}
	runSliceEmptyTest[float64](t, ",", []float64{}, envconv.ToFloat64Slice)
}

func TestToFloat64WithDefault(t *testing.T) {
	testData := []struct {
		env          string
		value        string
		expected     float64
		defaultValue float64
	}{
		{"TEST_FLOAT64_WITH_DEFAULT_0", "0", 0, 105},
		{"TEST_FLOAT64_WITH_DEFAULT_1.79769313486231570814527423731704356798070567525845e+308", "1.79769313486231570814527423731704356798070567525845e+308", 1.79769313486231570814527423731704356798070567525845e+308, 105},
		{"TEST_FLOAT64_WITH_DEFAULT_1.79769313486231570814527423731704356798070567525845e+309", "1.79769313486231570814527423731704356798070567525845e+309", 0, 105},
		{"TEST_FLOAT64_WITH_DEFAULT_4.94065645841246544176568792868221372365059802614325e-324", "4.94065645841246544176568792868221372365059802614325e-324", 4.94065645841246544176568792868221372365059802614325e-324, 105},
		{"TEST_FLOAT64_WITH_DEFAULT_NOTANUMBER", "notanumber", 0, 105},
	}

	for _, td := range testData {
		runWithDefaultTest(t, td.env, td.value, td.expected, td.defaultValue, envconv.ToFloat64WithDefault)
	}
	runWithDefaultEmptyTest(t, float64(0), envconv.ToFloat64WithDefault)
}

func TestToFloat64SliceWithDefault(t *testing.T) {
	defaultValue := []float64{1, 0, 5}
	testData := []struct {
		env             string
		value           string
		separator       string
		expected        []float64
		defaultValue    []float64
		defaultExpected bool
	}{
		{"TEST_FLOAT64_SLICE_WITH_DEFAULT_123_COMMA", "1,2,3", ",", []float64{1, 2, 3}, defaultValue, false},
		{"TEST_FLOAT64_SLICE_WITH_DEFAULT_123_COMMA_SPACE", "1, 2, 3", ",", []float64{1, 2, 3}, defaultValue, false},
		{"TEST_FLOAT64_SLICE_WITH_DEFAULT_123_SPACE", "1 2 3", ",", []float64{}, defaultValue, true},
		{"TEST_FLOAT64_SLICE_WITH_DEFAULT_1.79769313486231570814527423731704356798070567525845e+308", "1.79769313486231570814527423731704356798070567525845e+308,2,3", ",", []float64{1.79769313486231570814527423731704356798070567525845e+308, 2, 3}, defaultValue, false},
		{"TEST_FLOAT64_SLICE_WITH_DEFAULT_1.79769313486231570814527423731704356798070567525845e+309", "1.79769313486231570814527423731704356798070567525845e+309,2,3", ",", []float64{}, defaultValue, true},
		{"TEST_FLOAT64_SLICE_WITH_DEFAULT_4.94065645841246544176568792868221372365059802614325e-324", "4.94065645841246544176568792868221372365059802614325e-324,2,3", ",", []float64{4.94065645841246544176568792868221372365059802614325e-324, 2, 3}, defaultValue, false},
		{"TEST_FLOAT64_SLICE_WITH_DEFAULT_NOTANUMBER", "1,2,notanumber,4", ",", []float64{}, defaultValue, true},
	}

	for _, td := range testData {
		runSliceWithDefaultTest(t, td.env, td.value, td.separator, td.expected, td.defaultValue, td.defaultExpected, envconv.ToFloat64SliceWithDefault)
	}
	runSliceWithDefaultEmptyTest[float64](t, ",", defaultValue, envconv.ToFloat64SliceWithDefault)
}
