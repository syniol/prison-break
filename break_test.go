package prisonbreak

import (
	"testing"
	"time"
)

func TestNewPrison(t *testing.T) {
	t.Run("when duration between UpdatedDateTime and PrisonBreakDuration is less", func(t *testing.T) {
		prison := &Prison{
			cells: make(map[InmateIPAddr]*PrisonInmate),
			rules: func() *PrisonRules {
				return &PrisonRules{
					IsolationRedLineStrikeCount: defaultIsolationRedLineStrikeCount,
					IsolationRedLineDuration:    defaultIsolationRedLineDuration,
					PrisonBreakDuration:         defaultPrisonBreakDuration,
				}
			}(),
		}

		for i := 1; i <= defaultIsolationRedLineStrikeCount+10; i++ {
			result := prison.IsIsolated("166.187.0.2")

			if i > defaultIsolationRedLineStrikeCount {
				if result != true {
					t.Error("Isolated RedLine Strike should be true", i, result)
				}
			}

			if i < defaultIsolationRedLineStrikeCount {
				if result != false {
					t.Error("Isolated RedLine Strike should be false", i, result)
				}
			}
		}
	})

	t.Run("when duration between UpdatedDateTime and PrisonBreakDuration is higher", func(t *testing.T) {
		prison := &Prison{
			cells: make(map[InmateIPAddr]*PrisonInmate),
			rules: func() *PrisonRules {
				return &PrisonRules{
					IsolationRedLineStrikeCount: defaultIsolationRedLineStrikeCount,
					IsolationRedLineDuration:    defaultIsolationRedLineDuration,
					PrisonBreakDuration:         defaultPrisonBreakDuration,
				}
			}(),
		}

		for i := 1; i <= defaultIsolationRedLineStrikeCount+1; i++ {
			time.Sleep(defaultIsolationRedLineDuration + time.Nanosecond)
			result := prison.IsIsolated("166.187.0.2")

			if result != false {
				t.Error("Isolated RedLine Strike should be false", i, result)
			}
		}
	})
}
