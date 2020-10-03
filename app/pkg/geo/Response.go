package geo

type Response struct {
	Id int64
	PlaceId string
	Name string
	Lat float64
	Lng float64
	Components []Component
	Data []byte `json:"-"`
}

type Component struct {
	LongName string
	ShortName string
	Types []string
}