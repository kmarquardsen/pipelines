// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

/*LogoutUserDefault successful operation

swagger:response logoutUserDefault
*/
type LogoutUserDefault struct {
	_statusCode int
}

// NewLogoutUserDefault creates LogoutUserDefault with default headers values
func NewLogoutUserDefault(code int) *LogoutUserDefault {
	if code <= 0 {
		code = 500
	}

	return &LogoutUserDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the logout user default response
func (o *LogoutUserDefault) WithStatusCode(code int) *LogoutUserDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the logout user default response
func (o *LogoutUserDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WriteResponse to the client
func (o *LogoutUserDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(o._statusCode)
}