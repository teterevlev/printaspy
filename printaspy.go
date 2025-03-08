package printaspy

import (
	"fmt"
	"strings"
)
type PrintConfig struct {
	sep string
	end string
}

type Option func(*PrintConfig)

func WithSep(sep string) Option {
	return func(c *PrintConfig) {
		c.sep = sep
	}
}

func WithEnd(end string) Option {
	return func(c *PrintConfig) {
		c.end = end
	}
}

func Print(values ...interface{}) {
	var options []Option
	var args []interface{}

	// Look for `Option`, remaining ones go to args
	for _, v := range values {
		if opt, ok := v.(Option); ok {
			options = append(options, opt)
		} else {
			args = append(args, v)
		}
	}

	// Default
	config := PrintConfig{
		sep: " ",
		end: "\n",
	}

	// Apply options
	for _, opt := range options {
		opt(&config)
	}

	// To strings
	strValues := make([]string, len(args))
	for i, v := range args {
		strValues[i] = fmt.Sprint(v)
	}

	// Output
	fmt.Print(strings.Join(strValues, config.sep) + config.end)
}