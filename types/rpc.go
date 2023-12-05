package types

type ResponseError struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Result interface{}   `json:"result"`
	Error  ResponseError `json:"error"`
}
