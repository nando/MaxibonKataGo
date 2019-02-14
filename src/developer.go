package developer

type Developer struct {
    name string
    maxibons_to_grab int
}

func (dev Developer) maxibonsToGrab() int {
    return dev.maxibons_to_grab
}
