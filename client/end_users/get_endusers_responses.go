// Code generated by go-swagger; DO NOT EDIT.

package end_users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetEndusersReader is a Reader for the GetEndusers structure.
type GetEndusersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEndusersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEndusersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetEndusersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetEndusersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetEndusersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetEndusersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetEndusersMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewGetEndusersGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetEndusersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetEndusersServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetEndusersOK creates a GetEndusersOK with default headers values
func NewGetEndusersOK() *GetEndusersOK {
	return &GetEndusersOK{}
}

/*GetEndusersOK handles this case with default header values.

The above command returns results like this:
*/
type GetEndusersOK struct {
	Payload *GetEndusersOKBody
}

func (o *GetEndusersOK) Error() string {
	return fmt.Sprintf("[GET /api/1.0/endusers/][%d] getEndusersOK  %+v", 200, o.Payload)
}

func (o *GetEndusersOK) GetPayload() *GetEndusersOKBody {
	return o.Payload
}

func (o *GetEndusersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetEndusersOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEndusersBadRequest creates a GetEndusersBadRequest with default headers values
func NewGetEndusersBadRequest() *GetEndusersBadRequest {
	return &GetEndusersBadRequest{}
}

/*GetEndusersBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type GetEndusersBadRequest struct {
}

func (o *GetEndusersBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/1.0/endusers/][%d] getEndusersBadRequest ", 400)
}

func (o *GetEndusersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEndusersUnauthorized creates a GetEndusersUnauthorized with default headers values
func NewGetEndusersUnauthorized() *GetEndusersUnauthorized {
	return &GetEndusersUnauthorized{}
}

/*GetEndusersUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type GetEndusersUnauthorized struct {
}

func (o *GetEndusersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/1.0/endusers/][%d] getEndusersUnauthorized ", 401)
}

func (o *GetEndusersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEndusersForbidden creates a GetEndusersForbidden with default headers values
func NewGetEndusersForbidden() *GetEndusersForbidden {
	return &GetEndusersForbidden{}
}

/*GetEndusersForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type GetEndusersForbidden struct {
}

func (o *GetEndusersForbidden) Error() string {
	return fmt.Sprintf("[GET /api/1.0/endusers/][%d] getEndusersForbidden ", 403)
}

func (o *GetEndusersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEndusersNotFound creates a GetEndusersNotFound with default headers values
func NewGetEndusersNotFound() *GetEndusersNotFound {
	return &GetEndusersNotFound{}
}

/*GetEndusersNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type GetEndusersNotFound struct {
}

func (o *GetEndusersNotFound) Error() string {
	return fmt.Sprintf("[GET /api/1.0/endusers/][%d] getEndusersNotFound ", 404)
}

func (o *GetEndusersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEndusersMethodNotAllowed creates a GetEndusersMethodNotAllowed with default headers values
func NewGetEndusersMethodNotAllowed() *GetEndusersMethodNotAllowed {
	return &GetEndusersMethodNotAllowed{}
}

/*GetEndusersMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type GetEndusersMethodNotAllowed struct {
}

func (o *GetEndusersMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/1.0/endusers/][%d] getEndusersMethodNotAllowed ", 405)
}

func (o *GetEndusersMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEndusersGone creates a GetEndusersGone with default headers values
func NewGetEndusersGone() *GetEndusersGone {
	return &GetEndusersGone{}
}

/*GetEndusersGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type GetEndusersGone struct {
}

func (o *GetEndusersGone) Error() string {
	return fmt.Sprintf("[GET /api/1.0/endusers/][%d] getEndusersGone ", 410)
}

func (o *GetEndusersGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEndusersInternalServerError creates a GetEndusersInternalServerError with default headers values
func NewGetEndusersInternalServerError() *GetEndusersInternalServerError {
	return &GetEndusersInternalServerError{}
}

/*GetEndusersInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned “msg” from the call.)
*/
type GetEndusersInternalServerError struct {
}

func (o *GetEndusersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/1.0/endusers/][%d] getEndusersInternalServerError ", 500)
}

func (o *GetEndusersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEndusersServiceUnavailable creates a GetEndusersServiceUnavailable with default headers values
func NewGetEndusersServiceUnavailable() *GetEndusersServiceUnavailable {
	return &GetEndusersServiceUnavailable{}
}

/*GetEndusersServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type GetEndusersServiceUnavailable struct {
}

func (o *GetEndusersServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/1.0/endusers/][%d] getEndusersServiceUnavailable ", 503)
}

func (o *GetEndusersServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetEndusersOKBody get endusers o k body
swagger:model GetEndusersOKBody
*/
type GetEndusersOKBody struct {

	// values
	Values []*ValuesItems0 `json:"values"`
}

// Validate validates this get endusers o k body
func (o *GetEndusersOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateValues(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetEndusersOKBody) validateValues(formats strfmt.Registry) error {

	if swag.IsZero(o.Values) { // not required
		return nil
	}

	for i := 0; i < len(o.Values); i++ {
		if swag.IsZero(o.Values[i]) { // not required
			continue
		}

		if o.Values[i] != nil {
			if err := o.Values[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getEndusersOK" + "." + "values" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetEndusersOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetEndusersOKBody) UnmarshalBinary(b []byte) error {
	var res GetEndusersOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ValuesItems0 values items0
swagger:model ValuesItems0
*/
type ValuesItems0 struct {

	// adusername
	Adusername interface{} `json:"adusername,omitempty"`

	// contact
	Contact interface{} `json:"contact,omitempty"`

	// domain
	Domain interface{} `json:"domain,omitempty"`

	// email
	Email interface{} `json:"email,omitempty"`

	// id
	ID interface{} `json:"id,omitempty"`

	// location
	Location interface{} `json:"location,omitempty"`

	// name
	Name interface{} `json:"name,omitempty"`

	// notes
	Notes interface{} `json:"notes,omitempty"`
}

// Validate validates this values items0
func (o *ValuesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ValuesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ValuesItems0) UnmarshalBinary(b []byte) error {
	var res ValuesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}