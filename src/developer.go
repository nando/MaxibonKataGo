package maxibonkata

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
	// Karumi's
	developers[ "pedro" ]  = Developer{ "Pedro", 3 }
	developers[ "fran" ]   = Developer{ "Fran", 1 }
	developers[ "davide" ] = Developer{ "Davide", 0 }
	developers[ "sergio" ] = Developer{ "Sergio", 2 }
	developers[ "jorge" ]  = Developer{ "Jorge", 1 }
	// Tck's
	developers[ "nuria" ]  = Developer{ "Nuria", 3 }
	developers[ "fausto" ] = Developer{ "Fausto", 2 }
	developers[ "julia" ]  = Developer{ "Julia",	0 }
	developers[ "luismi" ] = Developer{ "Luismi", 3 }
	developers[ "susana" ] = Developer{ "Susana", 2 }
	developers[ "sahu" ]   = Developer{ "Sahu", 1 }
	developers[ "vero" ]   = Developer{ "Vero", 3 }
	developers[ "vito" ]   = Developer{ "Vito", 3 }
	return
}
