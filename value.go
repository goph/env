package env

// Value is the interface to the dynamic value stored in an environment variable.
// (The default value is represented as a string.)
type Value interface {
	// Set is responsible for converting the string value to the actual type
	// and setting it in the target variable.
	//
	// When the conversion fails an error should be returned.
	Set(string) error

	// Type returns the name of the type the value represents.
	Type() string

	// String returns the actual value as string.
	//
	// Before any environment parsing happens, the returned value will be the default value.
	String() string
}
