// Code generated by go-swagger; DO NOT EDIT.

package application_components

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetAppcompsParams creates a new GetAppcompsParams object
// with the default values initialized.
func NewGetAppcompsParams() *GetAppcompsParams {
	var ()
	return &GetAppcompsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAppcompsParamsWithTimeout creates a new GetAppcompsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAppcompsParamsWithTimeout(timeout time.Duration) *GetAppcompsParams {
	var ()
	return &GetAppcompsParams{

		timeout: timeout,
	}
}

// NewGetAppcompsParamsWithContext creates a new GetAppcompsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAppcompsParamsWithContext(ctx context.Context) *GetAppcompsParams {
	var ()
	return &GetAppcompsParams{

		Context: ctx,
	}
}

// NewGetAppcompsParamsWithHTTPClient creates a new GetAppcompsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAppcompsParamsWithHTTPClient(client *http.Client) *GetAppcompsParams {
	var ()
	return &GetAppcompsParams{
		HTTPClient: client,
	}
}

/*GetAppcompsParams contains all the parameters to send to the API endpoint
for the get appcomps operation typically these are written to a http.Request
*/
type GetAppcompsParams struct {

	/*Device
	  Device name

	*/
	Device *string
	/*DeviceID
	  filter by id of device

	*/
	DeviceID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get appcomps params
func (o *GetAppcompsParams) WithTimeout(timeout time.Duration) *GetAppcompsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get appcomps params
func (o *GetAppcompsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get appcomps params
func (o *GetAppcompsParams) WithContext(ctx context.Context) *GetAppcompsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get appcomps params
func (o *GetAppcompsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get appcomps params
func (o *GetAppcompsParams) WithHTTPClient(client *http.Client) *GetAppcompsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get appcomps params
func (o *GetAppcompsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDevice adds the device to the get appcomps params
func (o *GetAppcompsParams) WithDevice(device *string) *GetAppcompsParams {
	o.SetDevice(device)
	return o
}

// SetDevice adds the device to the get appcomps params
func (o *GetAppcompsParams) SetDevice(device *string) {
	o.Device = device
}

// WithDeviceID adds the deviceID to the get appcomps params
func (o *GetAppcompsParams) WithDeviceID(deviceID *string) *GetAppcompsParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the get appcomps params
func (o *GetAppcompsParams) SetDeviceID(deviceID *string) {
	o.DeviceID = deviceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAppcompsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Device != nil {

		// query param device
		var qrDevice string
		if o.Device != nil {
			qrDevice = *o.Device
		}
		qDevice := qrDevice
		if qDevice != "" {
			if err := r.SetQueryParam("device", qDevice); err != nil {
				return err
			}
		}

	}

	if o.DeviceID != nil {

		// query param device_id
		var qrDeviceID string
		if o.DeviceID != nil {
			qrDeviceID = *o.DeviceID
		}
		qDeviceID := qrDeviceID
		if qDeviceID != "" {
			if err := r.SetQueryParam("device_id", qDeviceID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
