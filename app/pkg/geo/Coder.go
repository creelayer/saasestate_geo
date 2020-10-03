package geo

type Coder interface {

	Search(query string)
	Autocomplete(query string)
	Reverse(lat float64, lng float64) []Response
}