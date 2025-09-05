package prisonbreak_test

import (
	prisonbreak "github.com/syniol/prison-break"
)

type ServiceForTorture interface {
	execute()
	// ...any other method
}

func ExampleNewPrison() {
	prisonbreak.NewPrison().IsIsolated("127.0.0.11")

	// Flexibility to parse as many argument as you need
	tortureChamber := func(args ...any) error {
		ss := args[0].(ServiceForTorture)
		ss.execute()

		return nil
	}

	_ = prisonbreak.NewPrison().Torture("127.0.0.11", tortureChamber)
	// Output:
	//
}
