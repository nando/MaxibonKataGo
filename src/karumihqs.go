package maxibonkata

import( "math" )

func KarumiHQs_MIN_MAXIBONS() int { return 3 }
func KarumiHQs_BUY_MAXIBONS() int { return 10 }

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

func ( hqs *KarumiHQs ) maxibonsLeft() int {
	return hqs.maxibons_left
}

func ( hqs *KarumiHQs ) openFridge( dev Developer ) int {
	hqs.grabMaxibons( dev )

	if hqs.maxibons_left < KarumiHQs_MIN_MAXIBONS() {
		hqs.buyMaxibons()
	}

	return hqs.maxibons_left
}

func ( hqs *KarumiHQs ) meltingMaxibons() int {
	return hqs.melting_maxibons
}

func ( hqs *KarumiHQs ) meltedMaxibons() int {
	return hqs.melted_maxibons
}

func ( hqs *KarumiHQs ) grabMaxibons( dev Developer ) int {
	hqs.maxibons_left = int( math.Max(+0,
																		float64( hqs.maxibons_left - dev.maxibonsToGrab())))
	return hqs.maxibons_left
}

func ( hqs *KarumiHQs ) buyMaxibons() int {
	hqs.melting_maxibons = hqs.maxibons_left
	hqs.maxibons_left =+ KarumiHQs_BUY_MAXIBONS()

	return hqs.maxibons_left
}
