package developer

import( "math" )

type Developer struct {
    name string
    maxibons_to_grab int
}

func (dev Developer) maxibonsToGrab() int {
    return int( math.Max( +0, float64(dev.maxibons_to_grab )))
}

func Developers() (developers map[ string ]Developer) {
  developers = make( map[string]Developer )
  developers[ "pedro" ]  = Developer{ "Pedro", 3 }
  developers[ "fran" ]   = Developer{ "Fran", 1 }
  developers[ "davide" ] = Developer{ "Davide", 0 }
  developers[ "sergio" ] = Developer{ "Sergio", 2 }
  developers[ "jorge" ]  = Developer{ "Jorge", 1 }
  return
}
