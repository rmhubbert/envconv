package envconv_test

import (
	"fmt"
	"log"
	"os"

	"github.com/rmhubbert/envconv"
)

func setEnvironmentVariables() {
	os.Setenv("VALID_INT", "105")
	os.Setenv("INVALID_INT", "notanumber")
	os.Setenv("VALID_INT_SLICE", "12,5,74")
	os.Setenv("INVALID_INT_SLICE", "12,five,74")
}

func Example() {
	// We only need this step for the sake of this example. We'd expect any
	// environment variables to either be injected into the container or
	// loaded from a .env file prior to using envconv
	setEnvironmentVariables()

	// ToInt returns the value of the requested environment variable
	// converted to an int. An error will be returned if the
	// environment variable is not found or the conversion to
	// int fails.
	validInt, err := envconv.ToInt("VALID_INT")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("VALID_INT = %d\n", validInt)

	// ToIntWithDefault returns the value of the requested environment
	// variable converted to an ints. The default value passed as
	// the second parameter will be returned if the environment
	// variable is not found or the conversion to int fails.
	defaultValidInt := envconv.ToIntWithDefault("INVALID_INT", 123)
	fmt.Printf("INVALID_INT = %d\n", defaultValidInt)

	// ToIntSlice returns the value of the requested environment variable
	// converted to a slice of int. An error will be returned if the
	// environment variable is not found or the conversion to
	// slice of ints fails.
	validIntSlice, err := envconv.ToIntSlice("VALID_INT_SLICE", ",")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("VALID_INT_SLICE = %v\n", validIntSlice)

	// ToIntSliceWithDefault returns the value of the requested environment
	// variable converted to a slice of ints. The default value passed as
	// the second parameter will be returned if the environment
	// variable is not found or the conversion to a slice of ints fails.
	defaultIntSlice := envconv.ToIntSliceWithDefault("INVALID_INT_SLICE", ",", []int{1, 2, 3})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("INVALID_INT_SLICE = %v\n", defaultIntSlice)

	// Output:
	// VALID_INT = 105
	// INVALID_INT = 123
	// VALID_INT_SLICE = [12 5 74]
	// INVALID_INT_SLICE = [1 2 3]
}
