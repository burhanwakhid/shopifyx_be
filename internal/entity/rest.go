package entity

type RestResponse struct {
	Data interface{} `json:"data"`
	Meta interface{} `json:"meta"`

	// Success true if the request is success, false otherwise.
	Success bool        `json:"success"`
	Errors  []RestError `json:"errors"`
}

func NewRestResponse(
	success bool,
	data interface{},
	meta interface{},
	errors ...RestError,
) RestResponse {
	res := RestResponse{
		Data:    data,
		Meta:    meta,
		Success: success,
		Errors:  errors,
	}

	// Prevent null pointer exception in FE side.
	if res.Errors == nil {
		res.Errors = []RestError{}
	}

	return res
}

type RestError struct {
	// Mandatory fields.
	Code            string `json:"code" example:"Sd-%d"`
	Message         string `json:"message"`
	MessageTitle    string `json:"message_title"`
	MessageSeverity string `json:"message_severity" example:"error"`

	// Additional fields.
	Entity string `json:"entity" example:"shopifyx-backend"`
	Cause  string `json:"cause"`
}
