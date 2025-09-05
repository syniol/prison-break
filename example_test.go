package prisonbreak_test

import (
	"fmt"

	prisonbreak "github.com/syniol/prison-break"
)

type ServiceForTorture interface {
	execute()
	// ...any other method
}

func ExampleNewPrison() {
	prison := prisonbreak.NewPrison()

	// Checking initial status without count increment to be false
	// This will determine Isolation (Solidarity Confinement) status of inmate
	fmt.Println(prison.IsIsolated("127.0.0.11"))

	// Flexibility to parse as many argument as you need
	tortureChamber := func(args ...interface{}) error {
		ss := args[0].(ServiceForTorture)
		ss.execute()

		return nil
	}

	err := prison.Torture("127.0.0.11", tortureChamber)
	fmt.Println(err)

	// Output:
	// false
	// <nil>
}
