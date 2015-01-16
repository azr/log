package log_test

import (
	"testing"

	"github.com/azr/log"
)

func TestScopes(t *testing.T) {
	log.SetLevel(log.CRITICAL)

	ld := log.NewLogger("test one")
	ld.SetLevel(log.DEBUG)

	le := log.NewLogger("test two")
	le.SetLevel(log.ERROR)

	ld.Debug("FAKE DEBUG")
	le.Debug("FAKE DEBUG")
	le.Error("FAKE ERROR")
	log.Critical("FAKE CRITICAL")

	if le.GetLevel() == ld.GetLevel() {
		t.Error("Crap ! different loggers share the same level")
	}
	if ld.GetLevel() != log.DEBUG {
		t.Error("Crap ! different loggers share the same level")
	}
	if le.GetLevel() != log.ERROR {
		t.Error("Crap ! different loggers share the same level")
	}
}
