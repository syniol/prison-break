package prisonbreak

import "time"

// PrisonBreak will free the inmates based on criteria defined in configuration
// It will get triggered within defined timeframe to clean up the prison cells inside `init()` method
func PrisonBreak(prison *Prison) {
	// todo: create a ticker for every breaking time (30) + 1 seconds to trigger this method
	if prison == nil {
		return
	}

	for i, v := range prison.cells {
		// todo: add 30 to config
		if v.LastInspectionDateTime.Sub(time.Now()) >= time.Second*30 {
			delete(prison.cells, i)
		}
	}
}
