package response

type PageResponse struct {
	Count   int64         `json:"count"`
	Results interface{} `json:"results"`
}
