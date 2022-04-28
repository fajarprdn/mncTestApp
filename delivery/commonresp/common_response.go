package commonresp

import "encoding/json"

type ResponseMessage struct {
	Status       string      `json:"status"`
	ResponseCode string      `json:"response_code"`
	Description  string      `json:"description"`
	Data         interface{} `json:"data"`
}

type ErrorMessage struct {
	HttpCode int `json:"http_code"`
	ErrorDescription
}

type ErrorDescription struct {
	Status       string `json:"status"`
	Service      string `json:"service"`
	ResponseCode string `json:"response_code"`
	Description  string `json:"description"`
}

func (e ErrorMessage) ToJson() string {
	b, err := json.Marshal(e)
	if err != nil {
		return ""
	}
	return string(b)
}

func NewResponseMessage(respCode string, description string, data interface{}) ResponseMessage {
	return ResponseMessage{
		"Success", respCode, description, data,
	}
}
func NewErrorMessage(httpCode int, service, respCode, description string) ErrorMessage {
	return ErrorMessage{
		httpCode, ErrorDescription{
			Status:       "Error",
			Service:      service,
			ResponseCode: respCode,
			Description:  description,
		},
	}
}
