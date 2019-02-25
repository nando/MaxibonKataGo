package maxibonkata

import (
	"testing"
	"github.com/flyingmutant/rapid"
)

func TestDeveloperProperties(t *testing.T) {
	// Property #1: developer should always grab a positive number of maxibons.
	rapid.Check( t, func( t *rapid.T, maxibons_to_grab int ) {
		developer := Developer{ "Ada", maxibons_to_grab }

		if developer.maxibonsToGrab() < 0 {
			t.Fatalf("developer should always grab a positive number of maxibons.")
		}
	}, rapid.IntsRange(-100, 100))

	// Property #2: developer should assign the name of the developer in "construction".
	rapid.Check( t, func( t *rapid.T, name string ) {
		developer := Developer{ name, 1 }

		if developer.name != name {
			t.Fatalf("developer should assign the name of the developer in 'construction'")
		}
	}, rapid.Strings())
}

// Test #3: should assign the number of maxibons specified to every developer.
func Test_the_number_of_maxibons_specified_to_every_developer(t *testing.T) {
	developers := MakeDevelopers()
	maxibons := make(map[string]int)

	maxibons[ "pedro" ]  = 3
	maxibons[ "fran" ]	 = 1
	maxibons[ "davide" ] = 0
	maxibons[ "sergio" ] = 2
	maxibons[ "jorge" ]  = 1
	maxibons[ "nuria" ]  = 3
	maxibons[ "fausto" ] = 2
	maxibons[ "julia" ]  = 0
	maxibons[ "luismi" ] = 3
	maxibons[ "susana" ] = 2
	maxibons[ "sahu" ]   = 1
	maxibons[ "vero" ]   = 3
	maxibons[ "vito" ]   = 3

	for key, dev := range developers {
		if dev.maxibonsToGrab() != maxibons[ key ] {
			t.Fatalf("developer %s should grab %d maxibons.\n",
							 developers[ key ].name,
							 maxibons[ key ])
		}
	}
}
