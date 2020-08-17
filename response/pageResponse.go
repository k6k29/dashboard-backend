package response

type PageResponse struct {
	Count   int         `json:"count"`
	Results interface{} `json:"results"`
}
