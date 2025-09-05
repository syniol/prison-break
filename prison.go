package prisonbreak

import (
	"sync"
	"time"
)

// InmateIPAddr is a specific string type for Clients (Prisoners) IP Address
type InmateIPAddr string

// PrisonInmate is definition of prisoner in each cell defined in Prison
type PrisonInmate struct {
	Count                  int
	Isolated               bool
	LastInspectionDateTime time.Time
}

// Prison is a core domain that contains cells where key is an IP address with inmates information attached and rules
// cells is
// rules is
type Prison struct {
	cells map[InmateIPAddr]*PrisonInmate
	rules *PrisonRules
}

// PrisonRules defines the set of rules to be utilised for: isolation eligibility and prison cells clean up
// IsolationRedLineCount is
// IsolationRedLineDuration is
// PrisonBreakDuration is
type PrisonRules struct {
	IsolationRedLineCount    int
	IsolationRedLineDuration time.Duration
	PrisonBreakDuration      time.Duration
}

var once sync.Once
var instance *Prison

// NewPrison will create a new instance which accept a configuration called PrisonRules
// rules are optional by default PrisonRules are:
// IsolationRedLineCount:    20,
// IsolationRedLineDuration: time.Second * 5,
// PrisonBreakDuration:      time.Second * 10,
func NewPrison(rules *PrisonRules) *Prison {
	once.Do(func() {
		instance = &Prison{
			cells: make(map[InmateIPAddr]*PrisonInmate),
			rules: func() *PrisonRules {
				if rules != nil {
					return rules
				}

				return &PrisonRules{
					IsolationRedLineCount:    20,
					IsolationRedLineDuration: time.Second * 5,
					PrisonBreakDuration:      time.Second * 10,
				}
			}(),
		}

		prisonBreak(instance)
	})

	return instance
}

func (p *Prison) findInmate(ip string) *PrisonInmate {
	val, ok := p.cells[InmateIPAddr(ip)]
	if ok != true {
		return nil
	}

	return val
}

func (p *Prison) imprison(ip string) *PrisonInmate {
	prospectiveInmate := p.findInmate(ip)
	if prospectiveInmate == nil {
		newInmate := &PrisonInmate{
			Count:                  1,
			Isolated:               false,
			LastInspectionDateTime: time.Now(),
		}

		p.cells[InmateIPAddr(ip)] = newInmate

		return newInmate
	}

	prospectiveInmate.Count = prospectiveInmate.Count + 1
	prospectiveInmate.LastInspectionDateTime = time.Now()

	return prospectiveInmate
}

func (p *Prison) isolationEligibility(inmate *PrisonInmate) *PrisonInmate {
	if inmate.LastInspectionDateTime.Sub(time.Now()) <= p.rules.IsolationRedLineDuration &&
		inmate.Count >= p.rules.IsolationRedLineCount {
		inmate.Isolated = true
	}

	return inmate
}

// IsIsolated examines the criteria against PrisonRules
// It determines the eligibility of Solidarity Confinement
func (p *Prison) IsIsolated(ip string) bool {
	return p.isolationEligibility(p.imprison(ip)).Isolated
}

// Torture executes IsIsolated and the call back method defined as argument
// You can find an example of this inside the documentation and the example_test.go
func (p *Prison) Torture(ip string, cb func() error) error {
	if p.IsIsolated(ip) {
		return cb()
	}

	return nil
}
