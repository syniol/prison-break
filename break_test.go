package prisonbreak

import (
	"testing"
	"time"
)

func _TestBreak(t *testing.T) {
	prison := &Prison{
		cells: make(map[InmateIPAddr]*PrisonInmate),
		rules: &PrisonRules{
			IsolationRedLineStrikeCount: defaultIsolationRedLineStrikeCount,
			IsolationRedLineDuration:    defaultIsolationRedLineDuration,
			PrisonBreakDuration:         defaultPrisonBreakDuration,
		},
	}

	prisonBreak(nil, prison)

	if len(prison.cells) != 0 {
		t.Error("should be empty")
	}

	prison.IsIsolated("127.0.0.9")
	prison.IsIsolated("127.0.0.1")
	if len(prison.cells) != 2 {
		t.Error("should be two records in memory")
	}

	t.Log("Waiting for Cache Clean up")
	time.Sleep(defaultPrisonBreakDuration + time.Second)

	if len(prison.cells) != 0 {
		t.Error("should be empty after cache clean-up")
	}

	prison.IsIsolated("182.7.0.1")
	if len(prison.cells) != 1 {
		t.Error("should be one record in memory")
	}
}
