package resolver

import "fmt"

type customError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e customError) Error() string {
	return fmt.Sprintf("error [%s]: %s", e.Code, e.Message)
}

func (e customError) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"code":    e.Code,
		"message": e.Message,
	}
}
