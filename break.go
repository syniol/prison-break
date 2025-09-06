package prisonbreak

import (
	"context"
	"time"
)

// prisonBreak will free the inmates based on criteria defined in configuration
func prisonBreak(ctx context.Context, prison *Prison) {
	go func(ctx context.Context, prison *Prison) {
		select {
		case <-ctx.Done():
			return
		default:
			cachePrisonCellTicker := time.NewTicker(prison.rules.PrisonBreakDuration + time.Millisecond)
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
	}(ctx, prison)
}
