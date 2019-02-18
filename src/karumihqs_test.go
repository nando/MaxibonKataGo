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
}
