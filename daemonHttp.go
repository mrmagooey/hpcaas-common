package common

// JSONResponse the json response structure
type JSONResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
