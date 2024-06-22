package envconv

import "strconv"

// ToUint returns the value of the requested environment variable
// converted to an int. An error will be returned if the
// environment variable is not found or the conversion to
// int fails.
func ToUint(varName string) (uint, error) {
	return toIntType[uint](varName, 64, strconv.ParseUint)
}

// ToUintSlice returns the value of the requested environment variable
// converted to a slice of uints. An error will be returned if the
// environment variable is not found or the conversion to
// slice of ints fails.
func ToUintSlice(varName string, separator string) ([]uint, error) {
	return toIntSliceType[uint](varName, separator, 64, strconv.ParseUint)
}

// ToUintWithDefault returns the value of the requested environment
// variable converted to an int. The default value passed as
// the second parameter will be returned if the environment
// variable is not found or the conversion to int fails.
func ToUintWithDefault(varName string, defaultValue uint) uint {
	return toIntTypeWithDefault[uint](varName, defaultValue, 64, strconv.ParseUint)
}

// ToUintSliceWithDefault returns the value of the requested environment
// variable converted to a slice of uints. The default value passed as
// the second parameter will be returned if the environment
// variable is not found or the conversion to a slice of uints fails.
func ToUintSliceWithDefault(varName string, separator string, defaultValue []uint) []uint {
	return toIntSliceTypeWithDefault[uint](varName, separator, defaultValue, 64, strconv.ParseUint)
}

// ToUint8 returns the value of the requested environment variable
// converted to an uint8. An error will be returned if the
// environment variable is not found or the conversion to
// uint8 fails.
func ToUint8(varName string) (uint8, error) {
	return toIntType[uint8](varName, 8, strconv.ParseUint)
}

// ToUint8Slice returns the value of the requested environment variable
// converted to a slice of uint8s. An error will be returned if the
// environment variable is not found or the conversion to
// slice of int8s fails.
func ToUint8Slice(varName string, separator string) ([]uint8, error) {
	return toIntSliceType[uint8](varName, separator, 8, strconv.ParseUint)
}

// ToUint8WithDefault returns the value of the requested environment
// variable converted to an uint8. The default value passed as
// the second parameter will be returned if the environment
// variable is not found or the conversion to uint8 fails.
func ToUint8WithDefault(varName string, defaultValue uint8) uint8 {
	return toIntTypeWithDefault[uint8](varName, defaultValue, 8, strconv.ParseUint)
}

// ToUint8SliceWithDefault returns the value of the requested environment
// variable converted to a slice of uint8s. The default value passed as
// the second parameter will be returned if the environment
// variable is not found or the conversion to a slice of uint8s fails.
func ToUint8SliceWithDefault(varName string, separator string, defaultValue []uint8) []uint8 {
	return toIntSliceTypeWithDefault[uint8](varName, separator, defaultValue, 8, strconv.ParseUint)
}

// ToUint16 returns the value of the requested environment variable
// converted to an uint16. An error will be returned if the
// environment variable is not found or the conversion to
// uint16 fails.
func ToUint16(varName string) (uint16, error) {
	return toIntType[uint16](varName, 16, strconv.ParseUint)
}

// ToUint16Slice returns the value of the requested environment variable
// converted to a slice of uint16s. An error will be returned if the
// environment variable is not found or the conversion to
// slice of int16s fails.
func ToUint16Slice(varName string, separator string) ([]uint16, error) {
	return toIntSliceType[uint16](varName, separator, 16, strconv.ParseUint)
}

// ToUint16WithDefault returns the value of the requested environment
// variable converted to an uint16. The default value passed as
// the second parameter will be returned if the environment
// variable is not found or the conversion to uint16 fails.
func ToUint16WithDefault(varName string, defaultValue uint16) uint16 {
	return toIntTypeWithDefault[uint16](varName, defaultValue, 16, strconv.ParseUint)
}

// ToUint16SliceWithDefault returns the value of the requested environment
// variable converted to a slice of uint16s. The default value passed as
// the second parameter will be returned if the environment
// variable is not found or the conversion to a slice of uint16s fails.
func ToUint16SliceWithDefault(varName string, separator string, defaultValue []uint16) []uint16 {
	return toIntSliceTypeWithDefault[uint16](varName, separator, defaultValue, 16, strconv.ParseUint)
}

// ToUint32 returns the value of the requested environment variable
// converted to an uint32. An error will be returned if the
// environment variable is not found or the conversion to
// uint32 fails.
func ToUint32(varName string) (uint32, error) {
	return toIntType[uint32](varName, 32, strconv.ParseUint)
}

// ToUint32Slice returns the value of the requested environment variable
// converted to a slice of uint32s. An error will be returned if the
// environment variable is not found or the conversion to
// slice of int32s fails.
func ToUint32Slice(varName string, separator string) ([]uint32, error) {
	return toIntSliceType[uint32](varName, separator, 32, strconv.ParseUint)
}

// ToUint32WithDefault returns the value of the requested environment
// variable converted to an uint32. The default value passed as
// the second parameter will be returned if the environment
// variable is not found or the conversion to uint32 fails.
func ToUint32WithDefault(varName string, defaultValue uint32) uint32 {
	return toIntTypeWithDefault[uint32](varName, defaultValue, 32, strconv.ParseUint)
}

// ToUint32SliceWithDefault returns the value of the requested environment
// variable converted to a slice of uint32s. The default value passed as
// the second parameter will be returned if the environment
// variable is not found or the conversion to a slice of uint32s fails.
func ToUint32SliceWithDefault(varName string, separator string, defaultValue []uint32) []uint32 {
	return toIntSliceTypeWithDefault[uint32](varName, separator, defaultValue, 32, strconv.ParseUint)
}

// ToUint64 returns the value of the requested environment variable
// converted to an uint64. An error will be returned if the
// environment variable is not found or the conversion to
// uint64 fails.
func ToUint64(varName string) (uint64, error) {
	return toIntType[uint64](varName, 64, strconv.ParseUint)
}

// ToUint64Slice returns the value of the requested environment variable
// converted to a slice of uint64s. An error will be returned if the
// environment variable is not found or the conversion to
// slice of int64s fails.
func ToUint64Slice(varName string, separator string) ([]uint64, error) {
	return toIntSliceType[uint64](varName, separator, 64, strconv.ParseUint)
}

// ToUint64WithDefault returns the value of the requested environment
// variable converted to an uint64. The default value passed as
// the second parameter will be returned if the environment
// variable is not found or the conversion to uint64 fails.
func ToUint64WithDefault(varName string, defaultValue uint64) uint64 {
	return toIntTypeWithDefault[uint64](varName, defaultValue, 64, strconv.ParseUint)
}

// ToUint64SliceWithDefault returns the value of the requested environment
// variable converted to a slice of uint64s. The default value passed as
// the second parameter will be returned if the environment
// variable is not found or the conversion to a slice of uint64s fails.
func ToUint64SliceWithDefault(varName string, separator string, defaultValue []uint64) []uint64 {
	return toIntSliceTypeWithDefault[uint64](varName, separator, defaultValue, 64, strconv.ParseUint)
}
