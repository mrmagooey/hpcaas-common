package common

// JSONResponse the json response structure
//
// This is what the daemon will respond with.
// The Status is one of "success", "fail", "error"
// The Data is dependent on response
type JSONResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
