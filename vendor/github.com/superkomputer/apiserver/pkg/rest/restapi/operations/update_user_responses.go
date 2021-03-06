// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/fanzhangio/superkomputer/pkg/rest/models"
)

// UpdateUserAcceptedCode is the HTTP code returned for type UpdateUserAccepted
const UpdateUserAcceptedCode int = 202

/*UpdateUserAccepted update user task has been accepted

swagger:response updateUserAccepted
*/
type UpdateUserAccepted struct {
}

// NewUpdateUserAccepted creates UpdateUserAccepted with default headers values
func NewUpdateUserAccepted() *UpdateUserAccepted {

	return &UpdateUserAccepted{}
}

// WriteResponse to the client
func (o *UpdateUserAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(202)
}

// UpdateUserNotModifiedCode is the HTTP code returned for type UpdateUserNotModified
const UpdateUserNotModifiedCode int = 304

/*UpdateUserNotModified no update required

swagger:response updateUserNotModified
*/
type UpdateUserNotModified struct {
}

// NewUpdateUserNotModified creates UpdateUserNotModified with default headers values
func NewUpdateUserNotModified() *UpdateUserNotModified {

	return &UpdateUserNotModified{}
}

// WriteResponse to the client
func (o *UpdateUserNotModified) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(304)
}

// UpdateUserNotFoundCode is the HTTP code returned for type UpdateUserNotFound
const UpdateUserNotFoundCode int = 404

/*UpdateUserNotFound The user was not found

swagger:response updateUserNotFound
*/
type UpdateUserNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateUserNotFound creates UpdateUserNotFound with default headers values
func NewUpdateUserNotFound() *UpdateUserNotFound {

	return &UpdateUserNotFound{}
}

// WithPayload adds the payload to the update user not found response
func (o *UpdateUserNotFound) WithPayload(payload *models.Error) *UpdateUserNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update user not found response
func (o *UpdateUserNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateUserNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*UpdateUserDefault Error

swagger:response updateUserDefault
*/
type UpdateUserDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateUserDefault creates UpdateUserDefault with default headers values
func NewUpdateUserDefault(code int) *UpdateUserDefault {
	if code <= 0 {
		code = 500
	}

	return &UpdateUserDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the update user default response
func (o *UpdateUserDefault) WithStatusCode(code int) *UpdateUserDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the update user default response
func (o *UpdateUserDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the update user default response
func (o *UpdateUserDefault) WithPayload(payload *models.Error) *UpdateUserDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update user default response
func (o *UpdateUserDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateUserDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
