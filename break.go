package prisonbreak

import "time"

// prisonBreak will free the inmates based on criteria defined in configuration
func prisonBreak(prison *Prison) {
	cachePrisonCellTicker := time.NewTicker(prison.rules.PrisonBreakDuration + time.Millisecond)
	for range cachePrisonCellTicker.C {
		for i, v := range prison.cells {
			if time.Now().Sub(v.LastInspectionDateTime) >= prison.rules.PrisonBreakDuration {
				delete(prison.cells, i)
			}
		}
	}
}
