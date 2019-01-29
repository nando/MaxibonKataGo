package developer

import( "math" )

type Developer struct {
    name string
    maxibons_to_grab int
}

func (dev Developer) maxibonsToGrab() int {
    return int( math.Max( +0, float64(dev.maxibons_to_grab )))
}
