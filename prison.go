package prisonbreak

import (
	"sync"
	"time"
)

type InmateIPAddr string

type PrisonInmate struct {
	IP                     InmateIPAddr
	Count                  int
	Isolated               bool
	RegistrationDateTime   time.Time
	LastInspectionDateTime time.Time
}

type Prison struct {
	cells map[string]*PrisonInmate
}

var once sync.Once
var instance *Prison

func NewPrison() *Prison {
	once.Do(func() {
		instance = &Prison{
			cells: make(map[string]*PrisonInmate),
		}

		PrisonBreak(instance)
	})

	return instance
}

func (p *Prison) findInmate(ip string) *PrisonInmate {
	val, ok := p.cells[ip]
	if ok != true {
		return nil
	}

	return val
}

func (p *Prison) imprison(ip string) *PrisonInmate {
	prospectiveInmate := p.findInmate(ip)
	if prospectiveInmate == nil {
		newInmate := &PrisonInmate{
			IP:                     InmateIPAddr(ip),
			Count:                  1,
			Isolated:               false,
			RegistrationDateTime:   time.Now(),
			LastInspectionDateTime: time.Now(),
		}

		p.cells[ip] = newInmate

		return newInmate
	}

	prospectiveInmate.Count = prospectiveInmate.Count + 1
	prospectiveInmate.LastInspectionDateTime = time.Now()

	return prospectiveInmate
}

func (p *Prison) isolationEligibility(inmate *PrisonInmate) *PrisonInmate {
	// todo: make time.Second*5 && 20 as a config
	if inmate.LastInspectionDateTime.Sub(time.Now()) <= time.Second*5 && inmate.Count >= 20 {
		inmate.Isolated = true
	}

	return inmate
}

func (p *Prison) IsIsolated(ip string) bool {
	return p.isolationEligibility(p.imprison(ip)).Isolated
}

func (p *Prison) Torture(ip string, cb func(args ...interface{}) error) error {
	if p.IsIsolated(ip) {
		return cb()
	}

	return nil
}
