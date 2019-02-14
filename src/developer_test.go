package developer

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
}
