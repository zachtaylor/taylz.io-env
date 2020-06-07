package env

// Default returns a new Service, loaded with `ParseDefaultFile()`, `ParseEnv()`, and `ParseAllFlags()`
func Default() Service {
	service := NewService()
	service.ParseDefault()
	return service
}

// DefaultFile returns a new Service, loaded with `ParseDefaultFile()`
func DefaultFile() Service {
	service := NewService()
	service.ParseDefaultFile()
	return service
}

// File returns a new Service with `ParseFile(path)` loaded
func File(path string) (Service, error) {
	service := NewService()
	if err := service.ParseFile(path); err != nil {
		return nil, err
	}
	return service, nil
}

// Flags returns a new Service with `ParseAllFlags()` loaded
func Flags() Service {
	service := NewService()
	service.ParseAllFlags()
	return service
}
