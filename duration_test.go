package envconv_test

import (
	"testing"
	"time"

	"github.com/rmhubbert/envconv"
)

func TestToDuration(t *testing.T) {
	testData := []struct {
		env         string
		value       string
		expected    time.Duration
		errExpected bool
	}{
		{"TEST_DURATION_1s", "1s", 1000000000, false},
		{"TEST_DURATION_1m", "1m", 60000000000, false},
		{"TEST_DURATION_1h", "1h", 3600000000000, false},
		{"TEST_DURATION_105", "105", 0, true},
		{"TEST_DURATION_NOTADURATION", "notaduration", time.Duration(0), true},
	}

	for _, td := range testData {
		runTest(t, td.env, td.value, td.expected, td.errExpected, envconv.ToDuration)
	}
	runEmptyTest(t, time.Duration(0), envconv.ToDuration)
}

func TestToDurationSlice(t *testing.T) {
	testData := []struct {
		env         string
		value       string
		separator   string
		expected    []time.Duration
		errExpected bool
	}{
		{"TEST_DURATION_SLICE_1s_1m_1h_SPACE", "1s 1m 1h", " ", []time.Duration{1000000000, 60000000000, 3600000000000}, false},
		{"TEST_DURATION_SLICE_1s_1m_1h_COMMA", "1s,1m,1h", ",", []time.Duration{1000000000, 60000000000, 3600000000000}, false},
		{"TEST_DURATION_SLICE_1s_1m_1h_COMMA_SPACE", "1s, 1m, 1h", ", ", []time.Duration{1000000000, 60000000000, 3600000000000}, false},
		{"TEST_DURATION_SLICE_1s", "1s", ", ", []time.Duration{1000000000}, false},
		{"TEST_DURATION_SLICE_105", "105", ", ", []time.Duration{}, true},
		{"TEST_DURATION_SLICE_NOTADURATION", "notaduration", ",", []time.Duration{}, true},
	}

	for _, td := range testData {
		runSliceTest(t, td.env, td.value, td.separator, td.expected, td.errExpected, envconv.ToDurationSlice)
	}
	runSliceEmptyTest[time.Duration](t, ",", []time.Duration{}, envconv.ToDurationSlice)
}

func TestToDurationWithDefault(t *testing.T) {
	def, _ := time.ParseDuration("1m")
	testData := []struct {
		env          string
		value        string
		expected     time.Duration
		defaultValue time.Duration
	}{
		{"TEST_DURATION_WITH_DEFAULT_1s", "1s", 1000000000, def},
		{"TEST_DURATION_WITH_DEFAULT_1m", "1m", 60000000000, def},
		{"TEST_DURATION_WITH_DEFAULT_1h", "1h", 3600000000000, def},
		{"TEST_DURATION_WITH_DEFAULT_105", "105", def, def},
		{"TEST_DURATION_WITH_DEFAULT_NOTADURATION", "notaduration", def, def},
	}

	for _, td := range testData {
		runWithDefaultTest(t, td.env, td.value, td.expected, td.defaultValue, envconv.ToDurationWithDefault)
	}
	runWithDefaultEmptyTest(t, time.Duration(0), envconv.ToDurationWithDefault)
}

func TestToDurationSliceWithDefault(t *testing.T) {
	hour, _ := time.ParseDuration("1m")
	def := []time.Duration{hour}

	testData := []struct {
		env          string
		value        string
		separator    string
		expected     []time.Duration
		defaultValue []time.Duration
	}{
		{"TEST_DURATION_SLICE_1s_1m_1h_SPACE", "1s 1m 1h", " ", []time.Duration{1000000000, 60000000000, 3600000000000}, def},
		{"TEST_DURATION_SLICE_1s_1m_1h_COMMA", "1s,1m,1h", ",", []time.Duration{1000000000, 60000000000, 3600000000000}, def},
		{"TEST_DURATION_SLICE_1s_1m_1h_COMMA_SPACE", "1s, 1m, 1h", ", ", []time.Duration{1000000000, 60000000000, 3600000000000}, def},
		{"TEST_DURATION_SLICE_1s", "1s", ", ", []time.Duration{1000000000}, def},
		{"TEST_DURATION_SLICE_105", "105", ", ", []time.Duration{}, def},
		{"TEST_DURATION_SLICE_NOTADURATION", "notaduration", ",", []time.Duration{}, def},
	}

	for _, td := range testData {
		runSliceWithDefaultTest(t, td.env, td.value, td.separator, td.expected, td.defaultValue, envconv.ToDurationSliceWithDefault)
	}
	runSliceWithDefaultEmptyTest[time.Duration](t, ",", def, envconv.ToDurationSliceWithDefault)
}
