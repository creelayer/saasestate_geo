package google

type Response struct {
	Results      []Results
	Status       string
	ErrorMessage string `json:"error_message"`
}

type Results struct {
	AddressComponents []AddressComponents `json:"address_components"`
	FormattedAddress  string `json:"formatted_address"`
	Geometry          Geometry
	PlaceId           string `json:"place_id"`
	Types             []string
}

type AddressComponents struct {
	LongName  string `json:"long_name"`
	ShortName string `json:"short_name"`
	Types     []string
}

type Geometry struct {
	Location     Location
	LocationType string `json:"location_type"`
	Viewport     Viewport
}

type Location struct {
	Lat float64
	Lng float64
}

type Viewport struct {
	Northeast Location
	Southwest Location
}