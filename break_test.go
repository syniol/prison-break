package prisonbreak

import (
	"testing"
	"time"
)

func TestPrisonBreak(t *testing.T) {
	prison := NewPrison(nil)

	t.Run("when duration between LastInspectedDateTime and PrisonBreakDuration is greater", func(t *testing.T) {
		t.Log(len(prison.cells))
		prison.imprison("127.0.0.1")
		t.Log(len(prison.cells))
		time.Sleep(defaultPrisonBreakDuration + time.Millisecond*2)
		t.Log(len(prison.cells))
	})

	//t.Run("when duration between LastInspectedDateTime and PrisonBreakDuration is less", func(t *testing.T) {
	//	t.Log(len(prison.cells))
	//	prison.imprison("127.0.0.2")
	//	t.Log(len(prison.cells))
	//})
}
