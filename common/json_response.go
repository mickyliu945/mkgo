package common

const (
	CodeSuccess = 1001

	CodeErrorInternal     = 9001
	CodeErrorTokenAuth    = 9002
	CodeErrorTokenExpired = 9003

	MessageSuccess = "success"
)

type JSONResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewJSONResponse(code int, message string, data interface{}) *JSONResponse {
	resp := &JSONResponse{}
	resp.Code = code
	resp.Message = message
	resp.Data = data
	return resp
}

func NewAuthFailedResponse(message string) *JSONResponse {
	return &JSONResponse{
		Code:    CodeErrorTokenAuth,
		Message: message,
	}
}

func NewTokenExpiredResponse() *JSONResponse {
	return &JSONResponse{
		Code:    CodeErrorTokenExpired,
		Message: "Token is expired, please login again!",
	}
}

func NewEmptyDataResponse(message string) *JSONResponse {
	resp := &JSONResponse{}
	resp.Code = CodeSuccess
	resp.Message = message
	resp.Data = nil
	return resp
}
