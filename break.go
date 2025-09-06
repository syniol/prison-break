package prisonbreak

import (
	"context"
	"time"
)

// prisonBreak will free the inmates based on criteria defined in configuration
func prisonBreak(ctx context.Context, prison *Prison) {
	cachePrisonCellTicker := time.NewTicker(prison.rules.PrisonBreakDuration + time.Millisecond)

	go func(ctx context.Context, prison *Prison, cachePrisonCellTicker *time.Ticker) {
		select {
		case <-ctx.Done():
			prison.cells = make(map[InmateIPAddr]*PrisonInmate)
			cachePrisonCellTicker.Stop()
			return
		default:
			for range cachePrisonCellTicker.C {
				prison.mu.Lock()
				for i, v := range prison.cells {
					prison.mu.Lock()
					if time.Now().Sub(v.LastInspectionDateTime) >= prison.rules.PrisonBreakDuration {
						delete(prison.cells, i)
					}
					prison.mu.Unlock()
				}
				prison.mu.Unlock()
			}
		}
	}(ctx, prison, cachePrisonCellTicker)
}
