package models

type (
	Response struct {
		Data  interface{} `json:"data,omitempty"`
		Error interface{} `json:"error,omitempty"`
	}
)

func NewResponseWithError(err interface{}) Response {
	return NewResponse(nil, err)
}

func NewResponseWithData(data interface{}) Response {
	return NewResponse(data, nil)
}

func NewResponse(data interface{}, err interface{}) Response {
	return Response{
		Data:  data,
		Error: err,
	}
}
