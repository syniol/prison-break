package prisonbreak

import (
	"context"
	"time"
)

// prisonBreak will free the inmates based on criteria defined in configuration
func prisonBreak(ctx context.Context, prison *Prison) {
	if ctx == nil {
		go func(prison *Prison) {
			cachePrisonCellTicker := time.NewTicker(prison.rules.PrisonBreakDuration + time.Millisecond)
			for range cachePrisonCellTicker.C {
				for inmateIP, inmate := range prison.cells {
					prison.freeInmate(inmateIP, inmate)
				}
			}
		}(prison)

		return
	}

	go func(ctx context.Context, prison *Prison) {
		select {
		case <-ctx.Done():
			return
		default:
			cachePrisonCellTicker := time.NewTicker(prison.rules.PrisonBreakDuration + time.Millisecond)
			for range cachePrisonCellTicker.C {
				for inmateIP, inmate := range prison.cells {
					prison.freeInmate(inmateIP, inmate)
				}
			}
		}
	}(ctx, prison)
}
