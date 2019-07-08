package utilhttp

type (
	//APIObject is root for response body
	APIObject struct {
		Type       string       `json:"type,omitempty"`
		ID         string       `json:"id,omitempty"`
		Attributes interface{}  `json:"attributes,omitempty"`
		Error      *ErrorObject `json:"error,omitempty"`
	}

	//ErrorObject is component of APIObject
	ErrorObject struct {
		HTTPStatus int    `json:"-"`
		Code       string `json:"code,omitempty"`
		SysMessage string `json:"sys_message,omitempty"`
		Message    string `json:"message,omitempty"`
	}
)

const (
	//ErrGeneral for unknown / default error
	ErrGeneral       = "0"
	SysMsgErrGeneral = "Unknown error."
	MsgErrGeneral    = "Terjadi kendala teknis."

	//ErrCtxDeadlineExceeded if response too late
	ErrCtxDeadlineExceeded       = "mr_1001"
	SysMsgErrCtxDeadlineExceeded = "408 - Request timeout."

	ErrAPIObject       = "mr_2001"
	SysMsgErrAPIObject = "Error, APIObject is nil"

	ErrMarshal       = "mr_2002"
	SysMsgErrMarshal = "Error to marshal object."
)