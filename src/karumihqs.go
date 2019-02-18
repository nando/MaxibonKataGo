package maxibonkata

type KarumiHQs struct {
  name string
  maxibons_left int
}

func NewKarumiHQs( name string ) (karumihqs KarumiHQs) {
  karumihqs.name = name
  karumihqs.maxibons_left = 10
  return
}

func ( hqs KarumiHQs ) maxibonsLeft() int {
  return hqs.maxibons_left
}
