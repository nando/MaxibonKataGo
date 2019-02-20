package maxibonkata

import (
	"testing"
	"github.com/flyingmutant/rapid"
)

func TestKarumiHQsProperties(t *testing.T) {
	// Property #1: office should start the day with 10 maxibons.
	rapid.Check( t, func( t *rapid.T, name string ) {
		hqs := NewKarumiHQs( name )

		if hqs.maxibonsLeft() != 10 {
			t.Fatalf("new KarumiHQs instances should start the day with 10 maxibons.")
		}
	}, rapid.Strings())

	// Property #2: should always has more than two maxibons in the fridge.
	rapid.Check( t, func( t *rapid.T, maxibons_to_grab int ) {
		hqs := NewKarumiHQs( "The Cocktail Oviedo" )
		developer := Developer{ "Ada", maxibons_to_grab }

		hqs.openFridge( developer )

		if hqs.maxibonsLeft() < 3 {
			t.Fatalf("should always has more than two maxibons in the fridge.")
		}
	}, rapid.IntsRange(0, 42))

	// Property #3: should buy 10 more maxibons if there are less than 3 in the fridge
	rapid.Check( t, func( t *rapid.T, maxibons_to_grab int ) {
		hqs := NewKarumiHQs( "The Cocktail DF" )
		developer := Developer{ "Frida", maxibons_to_grab }

		hqs.openFridge( developer )

		if hqs.maxibonsLeft() < 10 {
			t.Fatalf("should buy 10 more maxibons if there are less than 3 in the fridge.")
		}
	}, rapid.IntsRange(8, 42))

	// Property #4: maxibons above BUY_MAXIBONS are melting maxibons
	rapid.Check( t, func( t *rapid.T, maxibons_to_grab int ) {
		hqs := NewKarumiHQs( "The Cocktail Santiago" )
		developer := Developer{ "Violeta", maxibons_to_grab }

		hqs.openFridge( developer )

		if hqs.meltingMaxibons() == 0 {
			t.Fatalf("maxibons above BUY_MAXIBONS should be melting maxibons.")
		}
	}, rapid.IntsRange(8, 9)) // => buy_maxibons() ==> 1 or 2 melting...
}
