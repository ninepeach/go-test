package must

import (
    "github.com/google/go-cmp/cmp"
)

// Settings holds cmp.Options to customize test assertions.
type Settings struct {
    cmpOptions []cmp.Option
}

// Setting modifies the Settings configuration.
type Setting func(*Settings)

// Cmp adds custom cmp.Option values for cmp.Equal behavior.
func Cmp(options ...cmp.Option) Setting {
    return func(s *Settings) {
        s.cmpOptions = append(s.cmpOptions, options...)
    }
}

// options aggregates and returns all cmp.Options from the settings.
func options(settings ...Setting) []cmp.Option {
    s := new(Settings)
    for _, setting := range settings {
        setting(s)
    }
    return s.cmpOptions
}