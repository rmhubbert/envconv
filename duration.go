package envconv

import (
	"strings"
	"time"
)

// convertDuration converts the passed string to a time.Duration. It will
// also return an error, if applicable.
func convertDuration(value string) (time.Duration, error) {
	convertedValue, err := time.ParseDuration(value)
	if err != nil {
		return 0, err
	}
	return convertedValue, nil
}

// ToDuration returns the value of the requested environment variable
// converted to a time.Duration. An error will be returned if the
// environment variable is not found or the conversion to
// time.Duration fails.
func ToDuration(varName string) (time.Duration, error) {
	value, err := loadFromEnvironment(varName, true)
	if err != nil {
		return 0, err
	}
	return convertDuration(value)
}

// ToDuration returns the value of the requested environment variable
// converted to a slice of time.Durations. An error will be returned
// if the environment variable is not found or the conversion to
// time.Duration fails.
func ToDurationSlice(varName string, separator string) ([]time.Duration, error) {
	value, err := loadFromEnvironment(varName, true)
	if err != nil {
		return []time.Duration{}, err
	}

	durationStrings := strings.Split(value, separator)
	var durations = []time.Duration{}
	for _, duration := range durationStrings {
		convertedDuration, err := convertDuration(duration)
		if err != nil {
			return []time.Duration{}, err
		}
		durations = append(durations, convertedDuration)
	}
	return durations, nil
}

// ToDurationWithDefault returns the value of the requested environment
// variable converted to a time.Duration. The default value passed as
// the second parameter will be returned if the environment
// variable is not found or the conversion to
// time.Duration fails.
func ToDurationWithDefault(varName string, defaultValue time.Duration) time.Duration {
	value, err := ToDuration(varName)
	if err != nil {
		return defaultValue
	}
	return value
}

// ToDurationSliceWithDefault returns the value of the requested environment
// variable converted to a slice of time.Durations. The default value
// passed as the second parameter will be returned if the environment
// variable is not found or the conversion to time.Duration fails.
func ToDurationSliceWithDefault(varName string, separator string, defaultValue []time.Duration) []time.Duration {
	value, err := ToDurationSlice(varName, separator)
	if err != nil {
		return defaultValue
	}
	return value
}
