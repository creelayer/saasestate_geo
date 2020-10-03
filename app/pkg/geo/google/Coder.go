package google

import (
	"app/pkg/geo"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

type Coder struct {
	api        string
	key        string
	additional map[string]string
}

func NewGoogleCoder(key string) *Coder {
	c := &Coder{}
	c.api = "https://maps.googleapis.com"
	c.key = key
	c.additional = make(map[string]string)
	c.additional["result_type"] = "street_address"
	return c
}

func (c *Coder) Search(query string) {

}

func (c *Coder) Autocomplete(query string) {

}

func (c *Coder) Reverse(lat float64, lng float64) []geo.Response {

	latLng := strconv.FormatFloat(lat, 'f', -1, 64) + "," + strconv.FormatFloat(lng, 'f', -1, 64)
	q := url.Values{"latlng": {latLng}}

	r, err := http.Get(c.createUrl("/maps/api/geocode/json", q))

	if err != nil {
		return nil
	}

	defer r.Body.Close()

	response := Response{}
	err = json.NewDecoder(r.Body).Decode(&response)

	return c.mapToResponse(&response.Results)

}

func (c *Coder) mapToResponse(results *[]Results) []geo.Response {

	locations := make([]geo.Response, len(*results))

	for i, r := range *results {

		location := geo.Response{
			PlaceId: r.PlaceId,
			Name:    r.FormattedAddress,
			Lat:     r.Geometry.Location.Lat,
			Lng:     r.Geometry.Location.Lng,
		}

		location.Components = make([]geo.Component, len(r.AddressComponents))

		for j, c := range r.AddressComponents {
			location.Components[j] = geo.Component{
				LongName:  c.LongName,
				ShortName: c.ShortName,
				Types:     c.Types,
			}
		}

		location.Data, _ = json.Marshal(r)
		locations[i] = location
	}

	return locations

}

func (c *Coder) createUrl(url string, q url.Values) string {

	q.Add("key", c.key)

	for k, v := range c.additional {
		q.Add(k, v)
	}

	return c.api + url + "?" + q.Encode()
}
