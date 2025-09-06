package prisonbreak

import (
	"context"
	"sync"
	"time"
)

// prisonBreak will free the inmates based on criteria defined in configuration
func prisonBreak(ctx context.Context, prison *Prison) {
	go func() {
		cachePrisonCellTicker := time.NewTicker(prison.rules.PrisonBreakDuration + time.Millisecond)
		var xx sync.Mutex
		xx.Lock()
		for range cachePrisonCellTicker.C {
			prison.mu.Lock()
			for inmateIP, inmate := range prison.cells {
				prison.freeInmate(inmateIP, inmate)
			}
			prison.mu.Unlock()
		}
		xx.Unlock()
	}()

	//go func(ctx context.Context, prison *Prison) {
	//	select {
	//	case <-ctx.Done():
	//		return
	//	default:
	//		cachePrisonCellTicker := time.NewTicker(prison.rules.PrisonBreakDuration + time.Millisecond)
	//		for range cachePrisonCellTicker.C {
	//			for inmateIP, inmate := range prison.cells {
	//				prison.freeInmate(inmateIP, inmate)
	//			}
	//		}
	//	}
	//}(ctx, prison)
}
