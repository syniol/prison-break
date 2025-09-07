package prisonbreak

import (
	"testing"
	"time"
)

func TestBreak(t *testing.T) {
	prison := &Prison{
		cells: make(map[InmateIPAddr]*PrisonInmate),
		rules: &PrisonRules{
			IsolationRedLineStrikeCount: defaultIsolationRedLineStrikeCount,
			IsolationRedLineDuration:    defaultIsolationRedLineDuration,
			PrisonBreakDuration:         defaultPrisonBreakDuration,
		},
	}

	prisonBreak(nil, prison)

	if prison.inmateCount != 0 {
		t.Error("should be empty")
	}

	prison.imprison("127.0.0.9")
	prison.imprison("127.0.0.1")
	prison.mu.RLock()
	if prison.inmateCount != 2 {
		t.Error("should be two records in memory")
	}
	prison.mu.RUnlock()

	t.Log("Waiting for Cache Clean up")
	time.Sleep(defaultPrisonBreakDuration + time.Second)

	prison.mu.RLock()
	if prison.inmateCount != 0 {
		t.Error("should be empty after cache clean-up")
	}
	prison.mu.RUnlock()

	prison.imprison("182.7.0.1")
	prison.mu.RLock()
	if len(prison.cells) != 1 {
		t.Error("should be one record in memory")
	}
	prison.mu.RUnlock()
}
