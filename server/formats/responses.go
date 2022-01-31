package formats

type JSONResult struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

type JSONError struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}
