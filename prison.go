package prison_break

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
	Cells map[string]*PrisonInmate
}

var once sync.Once
var instance *Prison

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

func (p *Prison) isolation(inmate *PrisonInmate) *PrisonInmate {
	if inmate.LastInspectionDateTime.Before(time.Now()) && inmate.Count > 10 {
		inmate.Isolated = true
	}

	return inmate
}

func (p *Prison) IsIsolated(ip string) bool {
	return p.isolation(p.imprison(ip)).Isolated
}

func (p *Prison) Torture(ip string, cb func(args ...any) error) error {
	if p.IsIsolated(ip) {
		return cb()
	}

	return nil
}
