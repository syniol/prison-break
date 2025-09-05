package prisonbreak

import (
	"time"
)

// prisonBreak will free the inmates based on criteria defined in configuration
func prisonBreak(prison *Prison) {
	go func(prison *Prison) {
		// clean up cache token every prison.rules.PrisonBreakDuration + time.Second
		cachePrisonCellTicker := time.NewTicker(prison.rules.PrisonBreakDuration + time.Millisecond)
		for _ = range cachePrisonCellTicker.C {
			var count = 0

			for i, v := range prison.cells {
				if v.LastInspectionDateTime.Sub(time.Now()) >= prison.rules.PrisonBreakDuration {
					count++
					delete(prison.cells, i)
				}
			}
		}
	}(prison)
}
