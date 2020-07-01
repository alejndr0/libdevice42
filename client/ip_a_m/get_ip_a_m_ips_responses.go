// Code generated by go-swagger; DO NOT EDIT.

package ip_a_m

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

	"github.com/mjpitz/libdevice42/models"
)

// GetIPAMIpsReader is a Reader for the GetIPAMIps structure.
type GetIPAMIpsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIPAMIpsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIPAMIpsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetIPAMIpsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetIPAMIpsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetIPAMIpsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetIPAMIpsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetIPAMIpsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewGetIPAMIpsGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetIPAMIpsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetIPAMIpsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetIPAMIpsOK creates a GetIPAMIpsOK with default headers values
func NewGetIPAMIpsOK() *GetIPAMIpsOK {
	return &GetIPAMIpsOK{}
}

/*GetIPAMIpsOK handles this case with default header values.

The above command returns results like this:
*/
type GetIPAMIpsOK struct {
	Payload *GetIPAMIpsOKBody
}

func (o *GetIPAMIpsOK) Error() string {
	return fmt.Sprintf("[GET /api/1.0/ips/][%d] getIpAMIpsOK  %+v", 200, o.Payload)
}

func (o *GetIPAMIpsOK) GetPayload() *GetIPAMIpsOKBody {
	return o.Payload
}

func (o *GetIPAMIpsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetIPAMIpsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIPAMIpsBadRequest creates a GetIPAMIpsBadRequest with default headers values
func NewGetIPAMIpsBadRequest() *GetIPAMIpsBadRequest {
	return &GetIPAMIpsBadRequest{}
}

/*GetIPAMIpsBadRequest handles this case with default header values.

Bad Request (A validation exception has occurred.)
*/
type GetIPAMIpsBadRequest struct {
}

func (o *GetIPAMIpsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/1.0/ips/][%d] getIpAMIpsBadRequest ", 400)
}

func (o *GetIPAMIpsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIPAMIpsUnauthorized creates a GetIPAMIpsUnauthorized with default headers values
func NewGetIPAMIpsUnauthorized() *GetIPAMIpsUnauthorized {
	return &GetIPAMIpsUnauthorized{}
}

/*GetIPAMIpsUnauthorized handles this case with default header values.

Unauthorized (Your credentials suck)
*/
type GetIPAMIpsUnauthorized struct {
}

func (o *GetIPAMIpsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/1.0/ips/][%d] getIpAMIpsUnauthorized ", 401)
}

func (o *GetIPAMIpsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIPAMIpsForbidden creates a GetIPAMIpsForbidden with default headers values
func NewGetIPAMIpsForbidden() *GetIPAMIpsForbidden {
	return &GetIPAMIpsForbidden{}
}

/*GetIPAMIpsForbidden handles this case with default header values.

Forbidden (The resource requested is hidden)
*/
type GetIPAMIpsForbidden struct {
}

func (o *GetIPAMIpsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/1.0/ips/][%d] getIpAMIpsForbidden ", 403)
}

func (o *GetIPAMIpsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIPAMIpsNotFound creates a GetIPAMIpsNotFound with default headers values
func NewGetIPAMIpsNotFound() *GetIPAMIpsNotFound {
	return &GetIPAMIpsNotFound{}
}

/*GetIPAMIpsNotFound handles this case with default header values.

Not Found (The specified resource could not be found)
*/
type GetIPAMIpsNotFound struct {
}

func (o *GetIPAMIpsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/1.0/ips/][%d] getIpAMIpsNotFound ", 404)
}

func (o *GetIPAMIpsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIPAMIpsMethodNotAllowed creates a GetIPAMIpsMethodNotAllowed with default headers values
func NewGetIPAMIpsMethodNotAllowed() *GetIPAMIpsMethodNotAllowed {
	return &GetIPAMIpsMethodNotAllowed{}
}

/*GetIPAMIpsMethodNotAllowed handles this case with default header values.

Method Not Allowed (You tried to access a resource with an invalid method)
*/
type GetIPAMIpsMethodNotAllowed struct {
}

func (o *GetIPAMIpsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/1.0/ips/][%d] getIpAMIpsMethodNotAllowed ", 405)
}

func (o *GetIPAMIpsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIPAMIpsGone creates a GetIPAMIpsGone with default headers values
func NewGetIPAMIpsGone() *GetIPAMIpsGone {
	return &GetIPAMIpsGone{}
}

/*GetIPAMIpsGone handles this case with default header values.

Gone (The resource requested has been removed from our servers)
*/
type GetIPAMIpsGone struct {
}

func (o *GetIPAMIpsGone) Error() string {
	return fmt.Sprintf("[GET /api/1.0/ips/][%d] getIpAMIpsGone ", 410)
}

func (o *GetIPAMIpsGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIPAMIpsInternalServerError creates a GetIPAMIpsInternalServerError with default headers values
func NewGetIPAMIpsInternalServerError() *GetIPAMIpsInternalServerError {
	return &GetIPAMIpsInternalServerError{}
}

/*GetIPAMIpsInternalServerError handles this case with default header values.

Internal Server Error (Some parameter missing or issue with the server. Check with returned “msg” from the call.)
*/
type GetIPAMIpsInternalServerError struct {
}

func (o *GetIPAMIpsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/1.0/ips/][%d] getIpAMIpsInternalServerError ", 500)
}

func (o *GetIPAMIpsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIPAMIpsServiceUnavailable creates a GetIPAMIpsServiceUnavailable with default headers values
func NewGetIPAMIpsServiceUnavailable() *GetIPAMIpsServiceUnavailable {
	return &GetIPAMIpsServiceUnavailable{}
}

/*GetIPAMIpsServiceUnavailable handles this case with default header values.

Service Unavailable (Please check if your Device42 instance is working normally.)
*/
type GetIPAMIpsServiceUnavailable struct {
}

func (o *GetIPAMIpsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/1.0/ips/][%d] getIpAMIpsServiceUnavailable ", 503)
}

func (o *GetIPAMIpsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetIPAMIpsOKBody get IP a m ips o k body
swagger:model GetIPAMIpsOKBody
*/
type GetIPAMIpsOKBody struct {

	// ips
	Ips []*models.IPAMips `json:"ips"`

	// limit
	Limit interface{} `json:"limit,omitempty"`

	// offset
	Offset interface{} `json:"offset,omitempty"`

	// total count
	TotalCount interface{} `json:"total_count,omitempty"`
}

// Validate validates this get IP a m ips o k body
func (o *GetIPAMIpsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateIps(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetIPAMIpsOKBody) validateIps(formats strfmt.Registry) error {

	if swag.IsZero(o.Ips) { // not required
		return nil
	}

	for i := 0; i < len(o.Ips); i++ {
		if swag.IsZero(o.Ips[i]) { // not required
			continue
		}

		if o.Ips[i] != nil {
			if err := o.Ips[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getIpAMIpsOK" + "." + "ips" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetIPAMIpsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetIPAMIpsOKBody) UnmarshalBinary(b []byte) error {
	var res GetIPAMIpsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}