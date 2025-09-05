package prisonbreak

import (
	"sync"
	"time"
)

// InmateIPAddr is a specific string type for Clients (Prisoners) IP Address
type InmateIPAddr string

// PrisonInmate is definition of prisoner in each cell defined in Prison
type PrisonInmate struct {
	StrikeCount            int
	Isolated               bool
	LastInspectionDateTime time.Time
}

// Prison is a core domain that contains cells where key is an IP address with inmates information attached and rules
// cells is in-memory storage using map data structure with an IP Address of inmate as a key and inmates data as a value
// rules is PrisonRules that defines the behaviour of imprisonment and prison break
type Prison struct {
	cells map[InmateIPAddr]*PrisonInmate
	rules *PrisonRules
}

// PrisonRules defines the set of rules to be utilised for: isolation eligibility and prison cells clean up
// IsolationRedLineStrikeCount is the maximum number of strikes to reach for an isolation (Solidarity Containment)
// IsolationRedLineDuration is the maximum duration used for an isolation (Solidarity Containment)
// PrisonBreakDuration is the minimum duration for breaking prison for a well-behaved inmates
type PrisonRules struct {
	IsolationRedLineStrikeCount int
	IsolationRedLineDuration    time.Duration
	PrisonBreakDuration         time.Duration
}

var once sync.Once
var instance *Prison

// NewPrison will create a new instance which accept a configuration called PrisonRules
// rules are optional by default PrisonRules are:
// IsolationRedLineStrikeCount:    20,
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
					IsolationRedLineStrikeCount: 20,
					IsolationRedLineDuration:    time.Second * 5,
					PrisonBreakDuration:         time.Second * 10,
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
			StrikeCount:            1,
			Isolated:               false,
			LastInspectionDateTime: time.Now(),
		}

		p.cells[InmateIPAddr(ip)] = newInmate

		return newInmate
	}

	prospectiveInmate.StrikeCount = prospectiveInmate.StrikeCount + 1
	prospectiveInmate.LastInspectionDateTime = time.Now()

	return prospectiveInmate
}

func (p *Prison) isolationEligibility(inmate *PrisonInmate) *PrisonInmate {
	if inmate.LastInspectionDateTime.Sub(time.Now()) <= p.rules.IsolationRedLineDuration &&
		inmate.StrikeCount >= p.rules.IsolationRedLineStrikeCount {
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
