package prisonbreak_test

import (
	"context"
	"fmt"
	"net/http"

	prisonbreak "github.com/syniol/prison-break"
)

func ExampleNewPrison() {
	prison := prisonbreak.NewPrison(context.TODO(), nil)

	// Checking initial status without count increment to be false
	// This will determine Isolation (Solidarity Confinement) status of inmate
	fmt.Println(prison.IsIsolated("127.0.0.11"))

	// Flexibility to parse as many argument as you need
	var req *http.Request
	tortureChamber := func() error {
		req.Close = true
		defer req.Context().Done()

		return nil
	}

	err := prison.Torture("127.0.0.11", tortureChamber)
	fmt.Println(err)

	// Output:
	// false
	// <nil>
}
