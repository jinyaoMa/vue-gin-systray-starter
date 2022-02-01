package formats

type JSONResult struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

type JSONError struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type UserDataResult struct {
	ID         uint   `json:"id"`
	Account    string `json:"account"`
	Password   string `json:"password"`
	Name       string `json:"name"`
	Permission string `json:"permission"`
	Active     bool   `json:"active"`
}
