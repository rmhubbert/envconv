// Package envconv implements utility functions for retrieving and
// converting an environment variable to the specified type in a
// single step.
//
// Each implemented conversion type has two kinds of functions.
// The first kind will return an error when the environment
// variable is missing or the conversion fails. The second
// kind, which will be named with a "WithDefault" suffix,
// will accept and return a default value instead of
// returning an error. This second kind is ideal
// for use when initialising configuration
// structs.
package envconv

import (
	"errors"
	"os"
)

// LoadFromEvironment returns the value of the requested environment variable.
// An error is returned  if that variable is not set or (assuming the
// allowEmpty parameter is set to false), the loaded environment
// variable  is empty.
func LoadFromEnvironment(varName string, allowEmpty bool) (string, error) {
	val, ok := os.LookupEnv(varName)
	if !ok {
		return "", errors.New("unknown environment variable")
	}
	if !allowEmpty && val == "" {
		return "", errors.New("empty environment variable")
	}
	return val, nil
}

// LoadFromEvironmentWithDefault returns the value of the requested environment variable.
// A default value is returned if that variable is not set or the loaded environment
// variable  is empty.
func LoadFromEnvironmentWithDefault(varName string, defaultValue string) string {
	val, err := LoadFromEnvironment(varName, true)
	if err != nil {
		return defaultValue
	}
	return val
}
