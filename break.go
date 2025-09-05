package prison_break

import "time"

func PrisonBreak(prison *Prison) {
	if prison == nil {
		return
	}

	for i, v := range prison.Cells {
		// todo: add 30 to config
		if v.LastInspectionDateTime.Sub(time.Now()) >= time.Second*30 {
			delete(prison.Cells, i)
		}
	}
}
