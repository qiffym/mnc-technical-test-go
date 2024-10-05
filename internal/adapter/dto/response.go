package dto

type ResponseMeta struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	ResponseMeta
	Data   any `json:"data"`
	Errors any `json:"errors,omitempty"`
}

func DefaultErrorResponse() *ErrorResponse {
	return DefaultErrorResponseWithMessage("")
}

func DefaultErrorResponseWithMessage(msg string) *ErrorResponse {
	return &ErrorResponse{
		ResponseMeta: ResponseMeta{
			Success: false,
			Message: msg,
		},
		Data: nil,
	}
}

func DefaultErrorInvalidDataWithMessage(msg string) *ErrorResponse {
	return &ErrorResponse{
		ResponseMeta: ResponseMeta{
			Success: false,
			Message: msg,
		},
	}
}

func DefaultDataInvalidResponse(validationErrors any) *ErrorResponse {
	return &ErrorResponse{
		ResponseMeta: ResponseMeta{
			Message: "Data invalid.",
		},
		Errors: validationErrors,
	}
}

func DefaultBadRequestResponse() *ErrorResponse {
	return DefaultErrorResponseWithMessage("Bad request")
}

type Response struct {
	ResponseMeta
	Data any `json:"data"`
}

func DefaultInvalidInputFormResponse(errs map[string][]string) *Response {
	var msg string
	for _, val := range errs {
		msg = val[0]
		break
	}

	return &Response{
		ResponseMeta: ResponseMeta{
			Success: false,
			Message: msg,
		},

		Data: errs,
	}
}

func NewSuccessResponse(data any, msg string, resTime string) *Response {
	return &Response{
		ResponseMeta: ResponseMeta{
			Success: true,
			Message: msg,
		},
		Data: data,
	}
}
