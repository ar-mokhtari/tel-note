package neshan

type (
	DistanceMatrix struct {
		Status string `json:"status,omitempty"`
		Rows   []struct {
			Elements []struct {
				Status   string `json:"status,omitempty"`
				Duration struct {
					Value int    `json:"value,omitempty"`
					Text  string `json:"text,omitempty"`
				} `json:"duration"`
				Distance struct {
					Value int    `json:"value,omitempty"`
					Text  string `json:"text,omitempty"`
				} `json:"distance"`
			} `json:"elements,omitempty"`
		} `json:"rows,omitempty"`
		OriginAddresses      []string `json:"origin_addresses,omitempty"`
		DestinationAddresses []string `json:"destination_addresses,omitempty"`
	}
)
