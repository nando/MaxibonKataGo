package developer

import (
  "testing"
  "github.com/leanovate/gopter"
  "github.com/leanovate/gopter/gen"
  "github.com/leanovate/gopter/prop"
)

func Test_DeveloperPositiveGrabbing(t *testing.T) {
  parameters := gopter.DefaultTestParameters()
  parameters.MinSuccessfulTests = 10000
  properties := gopter.NewProperties(parameters)

  properties.Property("developer should always grab a positive number of maxibons", prop.ForAll(
    func(maxibons_to_grab int) bool {
      developer := Developer{ "Ada", maxibons_to_grab }

      return ( developer.maxibonsToGrab() >= 0 )
    },
    gen.Int(),
  ))

  properties.TestingRun(t)
}
