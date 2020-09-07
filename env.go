package env

// Version is the version number
const Version = "v0.0.0"

// ParseDefault returns a new Service, loaded with `ParseDefault()`
func ParseDefault() Service { return New().ParseDefault() }

// ParseDefaultFile returns a new Service, loaded with `ParseDefaultFile()`
func ParseDefaultFile() Service { return New().ParseDefaultFile() }

// ParseFile returns a new Service, loaded with `ParseFile(path)`
func ParseFile(path string) (Service, error) { return New().ParseFile(path) }

// ParseArgs returns a new Service, loaded with `ParseArgs()`
func ParseArgs(args []string) Service { return New().ParseArgs(args) }
