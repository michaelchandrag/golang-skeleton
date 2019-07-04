package utilhttp

import (
	"encoding/json"
	"log"
	"net/http"
)

//HandlerFunc for http handling
type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

func (fn HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.Close = true
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if err := fn(w, r); err != nil {
		log.Println(err)
		apiObject := ConstructAPIError(http.StatusInternalServerError, ErrGeneral, SysMsgErrGeneral, MsgErrGeneral)
		SendAPIObject(w, apiObject)
		return
	}
}

//ConstructErrorObject for construct ErrorObject
func ConstructErrorObject(httpStatus int, code string, sysMsg string, msg string) *ErrorObject {
	if httpStatus == 0 {
		httpStatus = http.StatusInternalServerError
	}

	return &ErrorObject{
		HTTPStatus: httpStatus,
		Code:       code,
		SysMessage: sysMsg,
		Message:    msg,
	}
}

func WrapErrorObject(errObject *ErrorObject) *APIObject {
	return &APIObject{
		Error: errObject,
	}
}

//ConstructAPIError for construct APIObject with ErrorObject
func ConstructAPIError(httpStatus int, code string, sysMsg string, msg string) *APIObject {
	return &APIObject{
		Error: ConstructErrorObject(httpStatus, code, sysMsg, msg),
	}
}

//SendAPIObject for sending data to writer
func SendAPIObject(w http.ResponseWriter, apiObject *APIObject) {
	if apiObject == nil {
		apiObject = ConstructAPIError(http.StatusInternalServerError, ErrAPIObject, SysMsgErrAPIObject, MsgErrGeneral)
	}

	jsonByte, err := json.Marshal(apiObject)
	if err != nil {
		log.Println(err)
		apiObject = ConstructAPIError(http.StatusInternalServerError, ErrMarshal, SysMsgErrMarshal, MsgErrGeneral)
	}

	// please WriteHeader first, because Write() will WriteHeader 200 as default
	if apiObject.Error != nil && apiObject.Error.HTTPStatus != 0 {
		w.WriteHeader(apiObject.Error.HTTPStatus)
	}

	w.Write([]byte(jsonByte))
}
