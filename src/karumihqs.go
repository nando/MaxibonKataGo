package maxibonkata

import( "math" )

const MIN_MAXIBONS int = 3
const BUY_MAXIBONS int = 10

type KarumiHQs struct {
	name string
	maxibons_left int
	melting_maxibons int
	melted_maxibons int
}

func NewKarumiHQs( name string ) (karumihqs KarumiHQs) {
	karumihqs.name = name
	karumihqs.maxibons_left = 10
	return
}

func ( hqs *KarumiHQs ) MaxibonsLeft() int {
	return hqs.maxibons_left
}

func ( hqs *KarumiHQs ) OpenFridge( dev Developer ) int {
	hqs.grabMaxibons( dev )

	if hqs.maxibons_left < MIN_MAXIBONS {
		hqs.buyMaxibons()
	}

	return hqs.maxibons_left
}

func ( hqs *KarumiHQs ) MeltingMaxibons() int {
	return hqs.melting_maxibons
}

func ( hqs *KarumiHQs ) MeltedMaxibons() int {
	return hqs.melted_maxibons
}

func ( hqs *KarumiHQs ) grabMaxibons( dev Developer ) int {
	hqs.maxibons_left = int( math.Max(+0,
																		float64( hqs.maxibons_left - dev.maxibonsToGrab())))
	return hqs.maxibons_left
}

func ( hqs *KarumiHQs ) buyMaxibons() int {
	hqs.melting_maxibons = hqs.maxibons_left
	hqs.maxibons_left =+ BUY_MAXIBONS
	hqs.melted_maxibons =+ hqs.melting_maxibons

	return hqs.maxibons_left
}
