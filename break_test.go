package prisonbreak

import (
	"testing"
	"time"
)

func TestPrisonBreak(t *testing.T) {
	prison := NewPrison(nil)

	t.Run("when duration between LastInspectedDateTime and PrisonBreakDuration is greater", func(t *testing.T) {
		if len(prison.cells) != 0 {
			t.Errorf("prisonBreak should be empty")
		}

		prison.imprison("127.0.0.1")
		if len(prison.cells) != 1 {
			t.Errorf("prisonBreak should have a new inmate")
		}

		time.Sleep(defaultPrisonBreakDuration + time.Millisecond*2)
		if len(prison.cells) != 0 {
			t.Errorf("prisonBreak should be empty")
		}

		prison.imprison("127.0.0.1")
		if len(prison.cells) != 1 {
			t.Errorf("prisonBreak should have a new inmate")
		}

		time.Sleep(defaultPrisonBreakDuration - time.Nanosecond)
		if len(prison.cells) != 1 {
			t.Errorf("prisonBreak should have a new inmate")
		}
	})

}
