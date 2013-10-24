package agent

import (
	"github.com/brandscreen/logutils"
	"io/ioutil"
)

// LevelFilter returns a LevelFilter that is configured with the log
// levels that we use.
func LevelFilter() *logutils.LevelFilter {
	return &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"DEBUG", "INFO", "WARN", "ERR"},
		MinLevel: "INFO",
		Writer:   ioutil.Discard,
	}
}

// ValidateLevelFilter verifies that the log levels within the filter
// are valid.
func ValidateLevelFilter(filter *logutils.LevelFilter) bool {
	for _, level := range filter.Levels {
		if level == filter.MinLevel {
			return true
		}
	}

	return false
}
