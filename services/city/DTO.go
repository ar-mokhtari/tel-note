package city

type DTO struct {
	IDF          uint    `json:"id"`
	NameF        string  `json:"name"`
	EnglishNameF string  `json:"english_name"`
	AriaCodeF    string  `json:"aria_code"`
	LatF         float64 `json:"lat"`
	LngF         float64 `json:"lng"`
}

func (c *DTO) ID() uint            { return c.IDF }
func (c *DTO) Name() string        { return c.NameF }
func (c *DTO) EnglishName() string { return c.EnglishNameF }
func (c *DTO) AriaCode() string    { return c.AriaCodeF }
func (c *DTO) Lat() float64        { return c.LatF }
func (c *DTO) Lng() float64        { return c.LngF }
