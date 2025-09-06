package prisonbreak

import (
	"context"
	"time"
)

// prisonBreak will free the inmates based on criteria defined in configuration
func prisonBreak(ctx context.Context, prison *Prison) {
	cachePrisonCellTicker := time.NewTicker(prison.rules.PrisonBreakDuration + time.Millisecond)

	go func(prison *Prison, ticker *time.Ticker, ctx context.Context) {
		select {
		case <-ctx.Done():
			return
		default:
			for range ticker.C {
				for i, v := range prison.cells {
					if time.Now().Sub(v.LastInspectionDateTime) >= prison.rules.PrisonBreakDuration {
						delete(prison.cells, i)
					}
				}
			}
		}
	}(prison, cachePrisonCellTicker, ctx)
}
