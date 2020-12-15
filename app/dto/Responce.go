package dto

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    interface{} `json:"data"`
}

func NewLocationResponse(data interface{}) Response {
	r := Response{}
	r.Status = "success"
	r.Data = data

	return r
}
