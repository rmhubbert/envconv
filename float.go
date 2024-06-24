package envconv

import (
	"strconv"
	"strings"
)

// Type intType is a convenience interface to wrap all of the possible int types
type floatType interface {
	float32 | float64
}

// toFloatType returns the value of the requested environment variable
// converted to type T. An error will be returned if the
// environment variable is not found or the conversion to
// type T fails.
func toFloatType[T floatType](varName string, bitSize int, conversionFunc func(string, int) (float64, error)) (T, error) {
	value, err := LoadFromEnvironment(varName, true)
	if err != nil {
		return T(0), err
	}

	convertedValue, err := conversionFunc(value, bitSize)
	if err != nil {
		return T(0), err
	}

	return T(convertedValue), err
}

// toFloatSliceType returns the value of the requested environment variable
// converted to type []T. An error will be returned if the
// environment variable is not found or the conversion to
// type []T fails.
func toFloatSliceType[T floatType](varName string, separator string, bitSize int, conversionFunc func(string, int) (float64, error)) ([]T, error) {
	value, err := LoadFromEnvironment(varName, true)
	if err != nil {
		return []T{}, err
	}

	valueSlice := strings.Split(value, separator)
	var convertedValues []T
	for _, v := range valueSlice {
		convertedValue, err := conversionFunc(strings.TrimSpace(v), bitSize)
		if err != nil {
			return []T{}, err
		}
		convertedValues = append(convertedValues, T(convertedValue))
	}

	return convertedValues, err
}

// ToFloatTypeWithDefault returns the value of the requested environment
// variable converted to type T. The default value passed as the
// second parameter will be returned if the environmentvariable
// is not found or the conversion to type T fails.
func toFloatTypeWithDefault[T floatType](varName string, defaultValue T, bitSize int, conversionFunc func(string, int) (float64, error)) T {
	value, err := toFloatType[T](varName, bitSize, conversionFunc)
	if err != nil {
		return defaultValue
	}
	return value
}

// toFloatSliceType returns the value of the requested environment variable
// converted to type []T. The default value passed as the second
// parameter will be returned if the environment variable is
// not found or the conversion to type []T fails.
func toFloatSliceTypeWithDefault[T floatType](varName string, separator string, defaultValue []T, bitSize int, conversionFunc func(string, int) (float64, error)) []T {
	value, err := toFloatSliceType[T](varName, separator, bitSize, conversionFunc)
	if err != nil {
		return defaultValue
	}
	return value
}

// ToFloat32 returns the value of the requested environment variable
// converted to an float32. An error will be returned if the
// environment variable is not found or the conversion to
// float32 fails.
func ToFloat32(varName string) (float32, error) {
	return toFloatType[float32](varName, 32, strconv.ParseFloat)
}

// ToFloat32Slice returns the value of the requested environment variable
// converted to a slice of float32s. An error will be returned if the
// environment variable is not found or the conversion to
// slice of float32s fails.
func ToFloat32Slice(varName string, separator string) ([]float32, error) {
	return toFloatSliceType[float32](varName, separator, 32, strconv.ParseFloat)
}

// ToFloat32WithDefault returns the value of the requested environment
// variable converted to an float32. The default value passed as
// the second parameter will be returned if the environment
// variable is not found or the conversion to float32 fails.
func ToFloat32WithDefault(varName string, defaultValue float32) float32 {
	return toFloatTypeWithDefault[float32](varName, defaultValue, 32, strconv.ParseFloat)
}

// ToFloat32SliceWithDefault returns the value of the requested environment
// variable converted to a slice of float32s. The default value passed as
// the second parameter will be returned if the environment
// variable is not found or the conversion to a slice of float32s fails.
func ToFloat32SliceWithDefault(varName string, separator string, defaultValue []float32) []float32 {
	return toFloatSliceTypeWithDefault[float32](varName, separator, defaultValue, 32, strconv.ParseFloat)
}

// ToFloat64 returns the value of the requested environment variable
// converted to an float64. An error will be returned if the
// environment variable is not found or the conversion to
// float64 fails.
func ToFloat64(varName string) (float64, error) {
	return toFloatType[float64](varName, 64, strconv.ParseFloat)
}

// ToFloat64Slice returns the value of the requested environment variable
// converted to a slice of float64s. An error will be returned if the
// environment variable is not found or the conversion to
// slice of float64s fails.
func ToFloat64Slice(varName string, separator string) ([]float64, error) {
	return toFloatSliceType[float64](varName, separator, 64, strconv.ParseFloat)
}

// ToFloat64WithDefault returns the value of the requested environment
// variable converted to an float64. The default value passed as
// the second parameter will be returned if the environment
// variable is not found or the conversion to float64 fails.
func ToFloat64WithDefault(varName string, defaultValue float64) float64 {
	return toFloatTypeWithDefault[float64](varName, defaultValue, 64, strconv.ParseFloat)
}

// ToFloat64SliceWithDefault returns the value of the requested environment
// variable converted to a slice of float64s. The default value passed as
// the second parameter will be returned if the environment
// variable is not found or the conversion to a slice of float64s fails.
func ToFloat64SliceWithDefault(varName string, separator string, defaultValue []float64) []float64 {
	return toFloatSliceTypeWithDefault[float64](varName, separator, defaultValue, 64, strconv.ParseFloat)
}
