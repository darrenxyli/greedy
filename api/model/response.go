package model

// Response is valid response for api
type Response struct {
	Status string      `json:"status"`
	Result interface{} `json:"result"`
}

// CreateSuccessResponse response
func CreateSuccessResponse(result interface{}) Response {
	return Response{Status: "success", Result: result}
}
