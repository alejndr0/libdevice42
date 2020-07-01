// Code generated by go-swagger; DO NOT EDIT.

package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostDeviceRackReader is a Reader for the PostDeviceRack structure.
type PostDeviceRackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostDeviceRackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostDeviceRackOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostDeviceRackBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostDeviceRackUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostDeviceRackForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostDeviceRackNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPostDeviceRackMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPostDeviceRackGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostDeviceRackInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostDeviceRackServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostDeviceRackOK creates a PostDeviceRackOK with default headers values
func NewPostDeviceRackOK() *PostDeviceRackOK {
	return &PostDeviceRackOK{}
}

/*PostDeviceRackOK handles this case with default header values.

The above command returns results like this:
*/
type PostDeviceRackOK struct {
	Payload *PostDeviceRackOKBody
}

func (o *PostDeviceRackOK) Error() string {
	return fmt.Sprintf("[POST /api/1.0/device/rack/][%d] postDeviceRackOK  %+v", 200, o.Payload)
}

func (o *PostDeviceRackOK) GetPayload() *PostDeviceRackOKBody {
	return o.Payload
}

func (o *PostDeviceRackOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostDeviceRackOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDeviceRackBadRequest creates a PostDeviceRackBadRequest with default headers values
func NewPostDeviceRackBadRequest() *PostDeviceRackBadRequest {
	return &PostDeviceRackBadRequest{}
}

/*PostDeviceRackBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type PostDeviceRackBadRequest struct {
}

func (o *PostDeviceRackBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/1.0/device/rack/][%d] postDeviceRackBadRequest ", 400)
}

func (o *PostDeviceRackBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostDeviceRackUnauthorized creates a PostDeviceRackUnauthorized with default headers values
func NewPostDeviceRackUnauthorized() *PostDeviceRackUnauthorized {
	return &PostDeviceRackUnauthorized{}
}

/*PostDeviceRackUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type PostDeviceRackUnauthorized struct {
}

func (o *PostDeviceRackUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/1.0/device/rack/][%d] postDeviceRackUnauthorized ", 401)
}

func (o *PostDeviceRackUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostDeviceRackForbidden creates a PostDeviceRackForbidden with default headers values
func NewPostDeviceRackForbidden() *PostDeviceRackForbidden {
	return &PostDeviceRackForbidden{}
}

/*PostDeviceRackForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type PostDeviceRackForbidden struct {
}

func (o *PostDeviceRackForbidden) Error() string {
	return fmt.Sprintf("[POST /api/1.0/device/rack/][%d] postDeviceRackForbidden ", 403)
}

func (o *PostDeviceRackForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostDeviceRackNotFound creates a PostDeviceRackNotFound with default headers values
func NewPostDeviceRackNotFound() *PostDeviceRackNotFound {
	return &PostDeviceRackNotFound{}
}

/*PostDeviceRackNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type PostDeviceRackNotFound struct {
}

func (o *PostDeviceRackNotFound) Error() string {
	return fmt.Sprintf("[POST /api/1.0/device/rack/][%d] postDeviceRackNotFound ", 404)
}

func (o *PostDeviceRackNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostDeviceRackMethodNotAllowed creates a PostDeviceRackMethodNotAllowed with default headers values
func NewPostDeviceRackMethodNotAllowed() *PostDeviceRackMethodNotAllowed {
	return &PostDeviceRackMethodNotAllowed{}
}

/*PostDeviceRackMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type PostDeviceRackMethodNotAllowed struct {
}

func (o *PostDeviceRackMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /api/1.0/device/rack/][%d] postDeviceRackMethodNotAllowed ", 405)
}

func (o *PostDeviceRackMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostDeviceRackGone creates a PostDeviceRackGone with default headers values
func NewPostDeviceRackGone() *PostDeviceRackGone {
	return &PostDeviceRackGone{}
}

/*PostDeviceRackGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type PostDeviceRackGone struct {
}

func (o *PostDeviceRackGone) Error() string {
	return fmt.Sprintf("[POST /api/1.0/device/rack/][%d] postDeviceRackGone ", 410)
}

func (o *PostDeviceRackGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostDeviceRackInternalServerError creates a PostDeviceRackInternalServerError with default headers values
func NewPostDeviceRackInternalServerError() *PostDeviceRackInternalServerError {
	return &PostDeviceRackInternalServerError{}
}

/*PostDeviceRackInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned “msg” from the call.)
*/
type PostDeviceRackInternalServerError struct {
}

func (o *PostDeviceRackInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/1.0/device/rack/][%d] postDeviceRackInternalServerError ", 500)
}

func (o *PostDeviceRackInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostDeviceRackServiceUnavailable creates a PostDeviceRackServiceUnavailable with default headers values
func NewPostDeviceRackServiceUnavailable() *PostDeviceRackServiceUnavailable {
	return &PostDeviceRackServiceUnavailable{}
}

/*PostDeviceRackServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type PostDeviceRackServiceUnavailable struct {
}

func (o *PostDeviceRackServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/1.0/device/rack/][%d] postDeviceRackServiceUnavailable ", 503)
}

func (o *PostDeviceRackServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PostDeviceRackOKBody post device rack o k body
swagger:model PostDeviceRackOKBody
*/
type PostDeviceRackOKBody struct {

	// code
	Code interface{} `json:"code,omitempty"`

	// msg
	Msg interface{} `json:"msg,omitempty"`
}

// Validate validates this post device rack o k body
func (o *PostDeviceRackOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostDeviceRackOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostDeviceRackOKBody) UnmarshalBinary(b []byte) error {
	var res PostDeviceRackOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}