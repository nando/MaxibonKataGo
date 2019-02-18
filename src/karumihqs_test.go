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
}
