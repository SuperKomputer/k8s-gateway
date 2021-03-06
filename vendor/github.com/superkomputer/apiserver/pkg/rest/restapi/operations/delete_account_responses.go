// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/fanzhangio/superkomputer/pkg/rest/models"
)

// DeleteAccountAcceptedCode is the HTTP code returned for type DeleteAccountAccepted
const DeleteAccountAcceptedCode int = 202

/*DeleteAccountAccepted delete account task has been accepted

swagger:response deleteAccountAccepted
*/
type DeleteAccountAccepted struct {
}

// NewDeleteAccountAccepted creates DeleteAccountAccepted with default headers values
func NewDeleteAccountAccepted() *DeleteAccountAccepted {

	return &DeleteAccountAccepted{}
}

// WriteResponse to the client
func (o *DeleteAccountAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(202)
}

// DeleteAccountNotFoundCode is the HTTP code returned for type DeleteAccountNotFound
const DeleteAccountNotFoundCode int = 404

/*DeleteAccountNotFound The account was not found

swagger:response deleteAccountNotFound
*/
type DeleteAccountNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteAccountNotFound creates DeleteAccountNotFound with default headers values
func NewDeleteAccountNotFound() *DeleteAccountNotFound {

	return &DeleteAccountNotFound{}
}

// WithPayload adds the payload to the delete account not found response
func (o *DeleteAccountNotFound) WithPayload(payload *models.Error) *DeleteAccountNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete account not found response
func (o *DeleteAccountNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAccountNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*DeleteAccountDefault Error

swagger:response deleteAccountDefault
*/
type DeleteAccountDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteAccountDefault creates DeleteAccountDefault with default headers values
func NewDeleteAccountDefault(code int) *DeleteAccountDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteAccountDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete account default response
func (o *DeleteAccountDefault) WithStatusCode(code int) *DeleteAccountDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete account default response
func (o *DeleteAccountDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete account default response
func (o *DeleteAccountDefault) WithPayload(payload *models.Error) *DeleteAccountDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete account default response
func (o *DeleteAccountDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAccountDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
