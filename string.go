package envconv

import "strings"

// ToString returns the value of the requested environment variable
// without converting from the original string. An error will be
// returned if the environment variable is not found.
//
// This function is a convenience wrapper around
// LoadFromEnvironment, for the sake of pattern
// consistency.
func ToString(varName string) (string, error) {
	return LoadFromEnvironment(varName, true)
}

// ToString returns the value of the requested environment variable
// as a slice of strings that have been split by the passed
// separator. An error will be returned if the
// environment variable is not found or the conversion to a
// slice of strings fails.
func ToStringSlice(varName string, separator string) ([]string, error) {
	value, err := LoadFromEnvironment(varName, true)
	if err != nil {
		return []string{}, err
	}

	return strings.Split(value, separator), nil
}

// ToStringWithDefault returns the value of the requested environment
// variable without converting from the original string. The default
// value passed as the second parameter will be returned if the
// environment variable is not found.
//
// This function is a convenience wrapper around
// LoadFromEnvironmentWithDefault, for the sake
// of pattern consistency.
func ToStringWithDefault(varName string, defaultValue string) string {
	return LoadFromEnvironmentWithDefault(varName, defaultValue)
}

// ToString returns the value of the requested environment variable
// as a slice of strings that have been split by the passed
// separator. The passed default slice will be returned if
// the environment variable is not found or the conversion
// to a slice of strings fails.
func ToStringSliceWithDefault(varName string, separator string, defaultValue []string) []string {
	value, err := ToStringSlice(varName, separator)
	if err != nil {
		return defaultValue
	}

	return value
}
