package env

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Service is a basic k/v map
type Service map[string]string

// New creates a Service
func New() Service { return Service{} }

// Keys returns a new slice of all named Service settings
func (s Service) Keys() []string {
	keys := make([]string, len(s))
	i := 0
	for k := range s {
		keys[i] = k
		i++
	}
	return keys
}

// Match returns a new Service containing values of settings prefixes matched, with the prefix removed
func (s Service) Match(pre string) Service {
	service, lpre := Service{}, len(pre)
	for k, v := range s {
		if len(k) > lpre && k[:lpre] == pre {
			service[k[lpre:]] = v
		}
	}
	return service
}

// Merge writes another Service's settings into this Service under the prefix
//
// This method returns itself for chaining sub Service settings
func (s Service) Merge(pre string, sub Service) Service {
	for k, v := range sub {
		s[pre+k] = v
	}
	return s
}

// Parse uses "x=y" format to write a value to Service
//
// In a setting, `=y` is optional, and this defaults to `[setting]="true"`
func (s Service) Parse(setting string) {
	if kv := strings.Split(setting, "="); len(kv) == 1 {
		s[kv[0]] = "true"
	} else if len(kv) == 2 {
		s[kv[0]] = strings.Trim(kv[1], ` 	"`)
	}
}

// ParseDefault is a macro for `ParseDefaultFile()`, `ParseEnv()`, and `ParseFlags()`
func (s Service) ParseDefault() Service { return s.ParseDefaultFile().ParseEnv().ParseFlags() }

// ParseEnv scans `os.Getenv` for available updates to this Service
func (s Service) ParseEnv() Service {
	for _, k := range s.Keys() {
		if v := os.Getenv(k); len(v) > 1 {
			s[k] = v
		}
	}
	return s
}

// ParseFile uses Parse to write file contents to Service
//
// Files can have comments (`#` style)
func (s Service) ParseFile(path string) error {
	file, e := ioutil.ReadFile(path)
	if e != nil {
		return e
	}
	for _, line := range strings.Split(string(file), "\n") {
		line = strings.Trim(strings.Split(line, "#")[0], " \r")
		if line != "" {
			s.Parse(line)
		}
	}
	return nil
}

// ParseDefaultFile calls `ParseFile` with `".env"`, logs and clears file error
func (s Service) ParseDefaultFile() Service {
	if err := s.ParseFile(".env"); err != nil {
		fmt.Println("env: " + err.Error())
	}
	return s
}

// ParseFlags returns a new Service, loaded with `os.Args[1:]`
func (s Service) ParseFlags() Service {
	for _, arg := range os.Args[1:] {
		if len(arg) > 1 && arg[0] == '-' {
			s.Parse(arg[1:])
		}
	}
	return s
}
