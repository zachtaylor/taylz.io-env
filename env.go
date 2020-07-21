package env

// ParseDefault returns a new Service, loaded with `ParseDefault()`
func ParseDefault() Service { return New().ParseDefault() }

// ParseDefaultFile returns a new Service, loaded with `ParseDefaultFile()`
func ParseDefaultFile() Service { return New().ParseDefaultFile() }

// ParseFlags returns a new Service, loaded with `ParseFlags()`
func ParseFlags() Service { return New().ParseFlags() }
