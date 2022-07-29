package city

type City struct {
	id          uint    `json:"id"`
	name        string  `json:"name"`
	englishName string  `json:"english_name"`
	ariaCode    string  `json:"aria_code"`
	lat         float64 `json:"lat"`
	lng         float64 `json:"lng"`
}

func (c City) ID() uint            { return c.id }
func (c City) Name() string        { return c.name }
func (c City) EnglishName() string { return c.englishName }
func (c City) AriaCode() string    { return c.ariaCode }
func (c City) Lat() float64        { return c.lat }
func (c City) Lng() float64        { return c.lng }

func SetCity(inputCity City) (res NewCityRequest) {
	return struct{ city }{city: inputCity}
}
