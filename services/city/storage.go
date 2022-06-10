package city

type city struct {
	id          uint    `json:"id"`
	name        string  `json:"name"`
	englishName string  `json:"english_name"`
	ariaCode    string  `json:"aria_code"`
	lat         float64 `json:"lat"`
	lng         float64 `json:"lng"`
}

func (c *city) ID() uint            { return c.id }
func (c *city) Name() string        { return c.name }
func (c *city) EnglishName() string { return c.englishName }
func (c *city) AriaCode() string    { return c.ariaCode }
func (c *city) Lat() float64        { return c.lat }
func (c *city) Lng() float64        { return c.lng }
