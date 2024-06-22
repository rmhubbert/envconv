package envconv

import (
	"strconv"
	"strings"
)

// Type intType is a convenience interface to wrap all of the possible int types
type IntType interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

// toIntType returns the value of the requested environment variable
// converted to type T. An error will be returned if the
// environment variable is not found or the conversion to
// type T fails.
func toIntType[T IntType, RT int64 | uint64](varName string, bitSize int, conversionFunc func(string, int, int) (RT, error)) (T, error) {
	value, err := LoadFromEnvironment(varName, true)
	if err != nil {
		return T(0), err
	}

	convertedValue, err := conversionFunc(value, 10, bitSize)
	if err != nil {
		return T(0), err
	}

	return T(convertedValue), err
}

// toIntSliceType returns the value of the requested environment variable
// converted to type []T. An error will be returned if the
// environment variable is not found or the conversion to
// type []T fails.
func toIntSliceType[T IntType, RT int64 | uint64](varName string, separator string, bitSize int, conversionFunc func(string, int, int) (RT, error)) ([]T, error) {
	value, err := LoadFromEnvironment(varName, true)
	if err != nil {
		return []T{}, err
	}

	valueSlice := strings.Split(value, separator)
	var convertedValues []T
	for _, v := range valueSlice {
		convertedValue, err := conversionFunc(strings.TrimSpace(v), 10, bitSize)
		if err != nil {
			return []T{}, err
		}
		convertedValues = append(convertedValues, T(convertedValue))
	}

	return convertedValues, err
}

// TointTypeWithDefault returns the value of the requested environment
// variable converted to type T. The default value passed as the
// second parameter will be returned if the environmentvariable
// is not found or the conversion to type T fails.
func toIntTypeWithDefault[T IntType, RT int64 | uint64](varName string, defaultValue T, bitSize int, conversionFunc func(string, int, int) (RT, error)) T {
	value, err := LoadFromEnvironment(varName, true)
	if err != nil {
		return defaultValue
	}

	convertedValue, err := conversionFunc(value, 10, bitSize)
	if err != nil {
		return defaultValue
	}

	return T(convertedValue)
}

// toIntSliceType returns the value of the requested environment variable
// converted to type []T. The default value passed as the second
// parameter will be returned if the environment variable is
// not found or the conversion to type []T fails.
func toIntSliceTypeWithDefault[T IntType, RT int64 | uint64](varName string, separator string, defaultValue []T, bitSize int, conversionFunc func(string, int, int) (RT, error)) []T {
	value, err := LoadFromEnvironment(varName, true)
	if err != nil {
		return defaultValue
	}

	valueSlice := strings.Split(value, separator)
	var convertedValues []T
	for _, v := range valueSlice {
		convertedValue, err := conversionFunc(strings.TrimSpace(v), 10, bitSize)
		if err != nil {
			return defaultValue
		}
		convertedValues = append(convertedValues, T(convertedValue))
	}

	return convertedValues
}

// ToInt returns the value of the requested environment variable
// converted to an int. An error will be returned if the
// environment variable is not found or the conversion to
// int fails.
func ToInt(varName string) (int, error) {
	return toIntType[int](varName, 64, strconv.ParseInt)
}

// ToIntSlice returns the value of the requested environment variable
// converted to a slice of int. An error will be returned if the
// environment variable is not found or the conversion to
// slice of ints fails.
func ToIntSlice(varName string, separator string) ([]int, error) {
	return toIntSliceType[int](varName, separator, 64, strconv.ParseInt)
}

// ToIntWithDefault returns the value of the requested environment
// variable converted to an ints. The default value passed as
// the second parameter will be returned if the environment
// variable is not found or the conversion to int fails.
func ToIntWithDefault(varName string, defaultValue int) int {
	return toIntTypeWithDefault[int](varName, defaultValue, 64, strconv.ParseInt)
}

// ToIntSliceWithDefault returns the value of the requested environment
// variable converted to a slice of ints. The default value passed as
// the second parameter will be returned if the environment
// variable is not found or the conversion to a slice of ints fails.
func ToIntSliceWithDefault(varName string, separator string, defaultValue []int) []int {
	return toIntSliceTypeWithDefault[int](varName, separator, defaultValue, 64, strconv.ParseInt)
}

// ToInt8 returns the value of the requested environment variable
// converted to an int8. An error will be returned if the
// environment variable is not found or the conversion to
// int8 fails.
func ToInt8(varName string) (int8, error) {
	return toIntType[int8](varName, 8, strconv.ParseInt)
}

// ToIntSlice returns the value of the requested environment variable
// converted to a slice of int8s. An error will be returned if the
// environment variable is not found or the conversion to
// slice of int8s fails.
func ToInt8Slice(varName string, separator string) ([]int8, error) {
	return toIntSliceType[int8](varName, separator, 8, strconv.ParseInt)
}

// ToInt8WithDefault returns the value of the requested environment
// variable converted to an int8. The default value passed as
// the second parameter will be returned if the environment
// variable is not found or the conversion to int8 fails.
func ToInt8WithDefault(varName string, defaultValue int8) int8 {
	return toIntTypeWithDefault[int8](varName, defaultValue, 8, strconv.ParseInt)
}

// ToInt8SliceWithDefault returns the value of the requested environment
// variable converted to a slice of int8s. The default value passed as
// the second parameter will be returned if the environment
// variable is not found or the conversion to a slice of int8s fails.
func ToInt8SliceWithDefault(varName string, separator string, defaultValue []int8) []int8 {
	return toIntSliceTypeWithDefault[int8](varName, separator, defaultValue, 8, strconv.ParseInt)
}

// ToInt16 returns the value of the requested environment variable
// converted to an int16. An error will be returned if the
// environment variable is not found or the conversion to
// int16 fails.
func ToInt16(varName string) (int16, error) {
	return toIntType[int16](varName, 16, strconv.ParseInt)
}

// ToInt16Slice returns the value of the requested environment variable
// converted to a slice of int16s. An error will be returned if the
// environment variable is not found or the conversion to
// slice of int16s fails.
func ToInt16Slice(varName string, separator string) ([]int16, error) {
	return toIntSliceType[int16](varName, separator, 16, strconv.ParseInt)
}

// ToInt16WithDefault returns the value of the requested environment
// variable converted to an int16. The default value passed as
// the second parameter will be returned if the environment
// variable is not found or the conversion to int16 fails.
func ToInt16WithDefault(varName string, defaultValue int16) int16 {
	return toIntTypeWithDefault[int16](varName, defaultValue, 16, strconv.ParseInt)
}

// ToInt16SliceWithDefault returns the value of the requested environment
// variable converted to a slice of int16s. The default value passed as
// the second parameter will be returned if the environment
// variable is not found or the conversion to a slice of int16s fails.
func ToInt16SliceWithDefault(varName string, separator string, defaultValue []int16) []int16 {
	return toIntSliceTypeWithDefault[int16](varName, separator, defaultValue, 16, strconv.ParseInt)
}

// ToInt32 returns the value of the requested environment variable
// converted to an int32. An error will be returned if the
// environment variable is not found or the conversion to
// int32 fails.
func ToInt32(varName string) (int32, error) {
	return toIntType[int32](varName, 32, strconv.ParseInt)
}

// ToInt32Slice returns the value of the requested environment variable
// converted to a slice of int32s. An error will be returned if the
// environment variable is not found or the conversion to
// slice of int32s fails.
func ToInt32Slice(varName string, separator string) ([]int32, error) {
	return toIntSliceType[int32](varName, separator, 32, strconv.ParseInt)
}

// ToInt32WithDefault returns the value of the requested environment
// variable converted to an int32. The default value passed as
// the second parameter will be returned if the environment
// variable is not found or the conversion to int32 fails.
func ToInt32WithDefault(varName string, defaultValue int32) int32 {
	return toIntTypeWithDefault[int32](varName, defaultValue, 32, strconv.ParseInt)
}

// ToInt32SliceWithDefault returns the value of the requested environment
// variable converted to a slice of int32s. The default value passed as
// the second parameter will be returned if the environment
// variable is not found or the conversion to a slice of int32s fails.
func ToInt32SliceWithDefault(varName string, separator string, defaultValue []int32) []int32 {
	return toIntSliceTypeWithDefault[int32](varName, separator, defaultValue, 32, strconv.ParseInt)
}

// ToInt64 returns the value of the requested environment variable
// converted to an int64. An error will be returned if the
// environment variable is not found or the conversion to
// int64 fails.
func ToInt64(varName string) (int64, error) {
	return toIntType[int64](varName, 64, strconv.ParseInt)
}

// ToInt64Slice returns the value of the requested environment variable
// converted to a slice of int64s. An error will be returned if the
// environment variable is not found or the conversion to
// slice of int64s fails.
func ToInt64Slice(varName string, separator string) ([]int64, error) {
	return toIntSliceType[int64](varName, separator, 64, strconv.ParseInt)
}

// ToInt64WithDefault returns the value of the requested environment
// variable converted to an int64. The default value passed as
// the second parameter will be returned if the environment
// variable is not found or the conversion to int64 fails.
func ToInt64WithDefault(varName string, defaultValue int64) int64 {
	return toIntTypeWithDefault[int64](varName, defaultValue, 64, strconv.ParseInt)
}

// ToInt64SliceWithDefault returns the value of the requested environment
// variable converted to a slice of int64s. The default value passed as
// the second parameter will be returned if the environment
// variable is not found or the conversion to a slice of int64s fails.
func ToInt64SliceWithDefault(varName string, separator string, defaultValue []int64) []int64 {
	return toIntSliceTypeWithDefault[int64](varName, separator, defaultValue, 64, strconv.ParseInt)
}
