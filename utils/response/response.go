package response

import "errors"

/****Add the Response JSON model structs here***/
type Success struct {
	Code   int         `json:"code"`
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
}
type Error struct {
	Code   int    `json:"code"`
	Status bool   `json:"status"`
	Error  string `json:"error"`
}

func SuccessResponse(data interface{}) Success {
	result := Success{
		Code:   200,
		Status: true,
		Data:   data,
	}
	return result
}
func ErrorMessage(code int, e error) Error {

	errMsg := Error{
		Status: false,
		Code:   code,
		Error:  e.Error(),
	}

	return errMsg
}

func CustomErrorMessage(code int, e string) Error {

	errMsg := Error{
		Status: false,
		Code:   code,
		Error:  e,
	}

	return errMsg
}
func CustomError(e string) error {
	return errors.New(e)
}
