package envconv_test

import (
	"os"
	"testing"

	"github.com/rmhubbert/envconv"
	"github.com/stretchr/testify/assert"
)

func TestToBool(t *testing.T) {
	testData := []struct {
		env           string
		value         string
		expected      bool
		errorExpected bool
	}{
		{"TEST_BOOL_1", "1", true, false},
		{"TEST_BOOL_true", "true", true, false},
		{"TEST_BOOL_TRUE", "TRUE", true, false},
		{"TEST_BOOL_tRuE", "tRuE", true, false},
		{"TEST_BOOL_0", "0", false, false},
		{"TEST_BOOL_false", "false", false, false},
		{"TEST_BOOL_FALSE", "FALSE", false, false},
		{"TEST_BOOL_FaLsE", "FaLsE", false, false},
		{"TEST_BOOL_NOTABOOL", "notabool", false, true},
	}

	for _, td := range testData {
		t.Run(td.env, func(t *testing.T) {
			os.Setenv(td.env, td.value)
			v, err := envconv.ToBool(td.env)
			if td.errorExpected {
				assert.Error(t, err, "there should be an error")
			} else {
				assert.NoError(t, err, "there should be no error")
			}
			assert.Equal(t, td.expected, v, "they should be equal")
		})
	}

	t.Run("TEST_NON_EXISTANT does not exist", func(t *testing.T) {
		v, err := envconv.ToBool("TEST_NON_EXISTANT")
		assert.Error(t, err, "there should be an error")
		assert.Equal(t, false, v, "they should be equal")
	})
}

func TestToBoolSlice(t *testing.T) {
	testData := []struct {
		env           string
		value         string
		separator     string
		expected      []bool
		errorExpected bool
	}{
		{"TEST_BOOL_SLICE_TRUE_FALSE_TRUE_SPACE", "true false true", " ", []bool{true, false, true}, false},
		{"TEST_BOOL_SLICE_1_0_T_COMMA", "1,0,t", ",", []bool{true, false, true}, false},
		{"TEST_BOOL_SLICE_105", "1,0,2", ",", []bool{}, true},
		{"TEST_BOOL_SLICE_NOTABOOL", "TRUE,notabool,false", ",", []bool{}, true},
	}

	for _, td := range testData {
		runSliceTest(t, td.env, td.value, td.separator, td.expected, td.errorExpected, envconv.ToBoolSlice)
	}
	runSliceEmptyTest[bool](t, ",", []bool{}, envconv.ToBoolSlice)
}

func TestToBoolWithDefault(t *testing.T) {
	testData := []struct {
		env          string
		value        string
		expected     bool
		defaultValue bool
	}{
		{"TEST_BOOL_WITH_DEFAULT_1", "1", true, false},
		{"TEST_BOOL_WITH_DEFAULT_true", "true", true, false},
		{"TEST_BOOL_WITH_DEFAULT_TRUE", "TRUE", true, false},
		{"TEST_BOOL_WITH_DEFAULT_tRuE", "tRuE", true, false},
		{"TEST_BOOL_WITH_DEFAULT_0", "0", false, true},
		{"TEST_BOOL_WITH_DEFAULT_false", "false", false, true},
		{"TEST_BOOL_WITH_DEFAULT_FALSE", "FALSE", false, true},
		{"TEST_BOOL_WITH_DEFAULT_FaLsE", "FaLsE", false, true},
		{"TEST_BOOL_WITH_DEFAULT_NOTABOOL", "notabool", false, true},
	}

	for _, td := range testData {
		t.Run(td.env, func(t *testing.T) {
			os.Setenv(td.env, td.value)
			v := envconv.ToBoolWithDefault(td.env, td.defaultValue)
			if v != td.expected {
				assert.Equal(t, td.defaultValue, v, "they should be equal")
			} else {
				assert.Equal(t, td.expected, v, "they should be equal")
			}
		})
	}

	t.Run("TEST_NON_EXISTANT does not exist", func(t *testing.T) {
		v := envconv.ToBoolWithDefault("TEST_NON_EXISTANT", false)
		assert.Equal(t, false, v, "they should be equal")
	})
}

func TestToBoolSliceWithDefault(t *testing.T) {
	def := []bool{true, false, true}
	testData := []struct {
		env             string
		value           string
		separator       string
		expected        []bool
		defaultValue    []bool
		defaultExpected bool
	}{
		{"TEST_DURATION_SLICE_TRUE_FALSE_TRUE_SPACE", "true false true", " ", []bool{true, false, true}, def, false},
		{"TEST_DURATION_SLICE_1_0_T_COMMA", "1,0,T", ",", []bool{true, false, true}, def, false},
		{"TEST_DURATION_SLICE_105", "105", ",", []bool{}, def, true},
		{"TEST_DURATION_SLICE_NOTABOOL", "true,notaduration,0", ",", []bool{}, def, true},
	}

	for _, td := range testData {
		runSliceWithDefaultTest(t, td.env, td.value, td.separator, td.expected, td.defaultValue, td.defaultExpected, envconv.ToBoolSliceWithDefault)
	}
	runSliceWithDefaultEmptyTest[bool](t, ",", def, envconv.ToBoolSliceWithDefault)
}
