package model

type SuccessResponse struct {
	Found int      `json:"found"`
	Data  []Figure `json:"data"`
}

type FailResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}
