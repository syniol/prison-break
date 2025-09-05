package prison_break

import (
	"sync"
	"time"
)

var once sync.Once
var instance *Prison

func init() {
	// todo: create a ticker for every breaking time (30) + 1 seconds to trigger this method
	PrisonBreak(instance)
}

type InmateIPAddr string

type PrisonInmate struct {
	IP                     InmateIPAddr
	Count                  int
	Isolated               bool
	RegistrationDateTime   time.Time
	LastInspectionDateTime time.Time
}

type Prison struct {
	Cells map[string]*PrisonInmate
}

func NewPrison() *Prison {
	once.Do(func() {
		instance = &Prison{
			Cells: make(map[string]*PrisonInmate),
		}
	})

	return instance
}

func (p *Prison) findInmate(ip string) *PrisonInmate {
	val, ok := p.Cells[ip]
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

		p.Cells[ip] = newInmate

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

func (p *Prison) Torture(ip string, cb func(args ...any) error) error {
	if p.IsIsolated(ip) {
		return cb()
	}

	return nil
}
