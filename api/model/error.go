package model

// ItemNotFoundErrorResponse means can not find item
func ItemNotFoundErrorResponse() Response {
	return Response{Status: "not found", Result: ""}
}

// ItemExistErrorResponse means already existed
func ItemExistErrorResponse(result interface{}) Response {
	return Response{Status: "existed", Result: result}
}

func BadParametersErrorResponse() Response {
	return Response{Status: "bad parameters", Result: ""}
}
