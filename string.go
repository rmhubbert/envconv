package envconv

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
