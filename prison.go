package prisonbreak

import (
	"context"
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
	LastUpdatedDateTime    time.Time
}

// Prison is a core domain that contains cells where key is an IP address with inmates information attached and rules
// cells is in-memory storage using map data structure with an IP Address of inmate as a key and inmates data as a value
// rules is PrisonRules that defines the behaviour of imprisonment and prison break
type Prison struct {
	cells map[InmateIPAddr]*PrisonInmate
	rules *PrisonRules
	mu    sync.Mutex
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

const defaultIsolationRedLineStrikeCount = 20
const defaultIsolationRedLineDuration time.Duration = time.Millisecond * 5
const defaultPrisonBreakDuration time.Duration = time.Millisecond * 30

// NewPrison will create a new instance which accept a configuration called PrisonRules
// rules are optional by default PrisonRules. Predefined rules are:
// IsolationRedLineStrikeCount: 20,
// IsolationRedLineDuration: time.Millisecond * 5,
// PrisonBreakDuration: time.Millisecond * 30,
func NewPrison(ctx context.Context, rules *PrisonRules) *Prison {
	once.Do(func() {
		instance = &Prison{
			cells: make(map[InmateIPAddr]*PrisonInmate),
			rules: func() *PrisonRules {
				if rules != nil {
					return rules
				}

				return &PrisonRules{
					IsolationRedLineStrikeCount: defaultIsolationRedLineStrikeCount,
					IsolationRedLineDuration:    defaultIsolationRedLineDuration,
					PrisonBreakDuration:         defaultPrisonBreakDuration,
				}
			}(),
		}

		// It will create a sub processing unit using goroutine to work in a background
		prisonBreak(ctx, instance)
	})

	return instance
}

func (p *Prison) findInmate(ip string) *PrisonInmate {
	p.mu.Lock()
	defer p.mu.Unlock()
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

		p.mu.Lock()
		p.cells[InmateIPAddr(ip)] = newInmate
		p.mu.Unlock()

		return newInmate
	}

	prospectiveInmate.StrikeCount = prospectiveInmate.StrikeCount + 1
	prospectiveInmate.LastUpdatedDateTime = prospectiveInmate.LastInspectionDateTime
	prospectiveInmate.LastInspectionDateTime = time.Now()

	return prospectiveInmate
}

func (p *Prison) isolationEligibility(inmate *PrisonInmate) *PrisonInmate {
	if inmate.StrikeCount > p.rules.IsolationRedLineStrikeCount &&
		time.Now().Sub(inmate.LastUpdatedDateTime) <= p.rules.IsolationRedLineDuration {
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
