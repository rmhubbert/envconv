package envconv

import (
	"time"
)

// convertduration converts the passed string to a time.Duration. It will
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
// to time.Duration fails.
func ToDuration(varName string) (time.Duration, error) {
	value, err := LoadFromEnvironment(varName, true)
	if err != nil {
		return 0, err
	}
	return convertDuration(value)
}

// ToDurationWithDefault returns the value of the requested environment
// variable converted to a time.Duration. The default value passed as
// the second parameter will be returned if the environment
// variable is not found or the conversion to int fails.
func ToDurationWithDefault(varName string, defaultValue time.Duration) time.Duration {
	value, err := LoadFromEnvironment(varName, false)
	if err != nil {
		return defaultValue
	}

	var convertedValue time.Duration
	convertedValue, err = convertDuration(value)
	if err != nil {
		return defaultValue
	}
	return convertedValue
}
