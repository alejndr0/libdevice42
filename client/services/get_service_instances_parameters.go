// Code generated by go-swagger; DO NOT EDIT.

package services

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
	"github.com/go-openapi/swag"
)

// NewGetServiceInstancesParams creates a new GetServiceInstancesParams object
// with the default values initialized.
func NewGetServiceInstancesParams() *GetServiceInstancesParams {
	var ()
	return &GetServiceInstancesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetServiceInstancesParamsWithTimeout creates a new GetServiceInstancesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetServiceInstancesParamsWithTimeout(timeout time.Duration) *GetServiceInstancesParams {
	var ()
	return &GetServiceInstancesParams{

		timeout: timeout,
	}
}

// NewGetServiceInstancesParamsWithContext creates a new GetServiceInstancesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetServiceInstancesParamsWithContext(ctx context.Context) *GetServiceInstancesParams {
	var ()
	return &GetServiceInstancesParams{

		Context: ctx,
	}
}

// NewGetServiceInstancesParamsWithHTTPClient creates a new GetServiceInstancesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetServiceInstancesParamsWithHTTPClient(client *http.Client) *GetServiceInstancesParams {
	var ()
	return &GetServiceInstancesParams{
		HTTPClient: client,
	}
}

/*GetServiceInstancesParams contains all the parameters to send to the API endpoint
for the get service instances operation typically these are written to a http.Request
*/
type GetServiceInstancesParams struct {

	/*Device
	  filter by name of the device

	*/
	Device *string
	/*DeviceID
	  filter by id of the device

	*/
	DeviceID *int64
	/*ServiceDetailID
	  filter by id of the service in use

	*/
	ServiceDetailID *int64
	/*ServiceID
	  filter by id of the service

	*/
	ServiceID *int64
	/*UserID
	  filter by id of the user

	*/
	UserID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get service instances params
func (o *GetServiceInstancesParams) WithTimeout(timeout time.Duration) *GetServiceInstancesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get service instances params
func (o *GetServiceInstancesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get service instances params
func (o *GetServiceInstancesParams) WithContext(ctx context.Context) *GetServiceInstancesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get service instances params
func (o *GetServiceInstancesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get service instances params
func (o *GetServiceInstancesParams) WithHTTPClient(client *http.Client) *GetServiceInstancesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get service instances params
func (o *GetServiceInstancesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDevice adds the device to the get service instances params
func (o *GetServiceInstancesParams) WithDevice(device *string) *GetServiceInstancesParams {
	o.SetDevice(device)
	return o
}

// SetDevice adds the device to the get service instances params
func (o *GetServiceInstancesParams) SetDevice(device *string) {
	o.Device = device
}

// WithDeviceID adds the deviceID to the get service instances params
func (o *GetServiceInstancesParams) WithDeviceID(deviceID *int64) *GetServiceInstancesParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the get service instances params
func (o *GetServiceInstancesParams) SetDeviceID(deviceID *int64) {
	o.DeviceID = deviceID
}

// WithServiceDetailID adds the serviceDetailID to the get service instances params
func (o *GetServiceInstancesParams) WithServiceDetailID(serviceDetailID *int64) *GetServiceInstancesParams {
	o.SetServiceDetailID(serviceDetailID)
	return o
}

// SetServiceDetailID adds the serviceDetailId to the get service instances params
func (o *GetServiceInstancesParams) SetServiceDetailID(serviceDetailID *int64) {
	o.ServiceDetailID = serviceDetailID
}

// WithServiceID adds the serviceID to the get service instances params
func (o *GetServiceInstancesParams) WithServiceID(serviceID *int64) *GetServiceInstancesParams {
	o.SetServiceID(serviceID)
	return o
}

// SetServiceID adds the serviceId to the get service instances params
func (o *GetServiceInstancesParams) SetServiceID(serviceID *int64) {
	o.ServiceID = serviceID
}

// WithUserID adds the userID to the get service instances params
func (o *GetServiceInstancesParams) WithUserID(userID *string) *GetServiceInstancesParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get service instances params
func (o *GetServiceInstancesParams) SetUserID(userID *string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetServiceInstancesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
		var qrDeviceID int64
		if o.DeviceID != nil {
			qrDeviceID = *o.DeviceID
		}
		qDeviceID := swag.FormatInt64(qrDeviceID)
		if qDeviceID != "" {
			if err := r.SetQueryParam("device_id", qDeviceID); err != nil {
				return err
			}
		}

	}

	if o.ServiceDetailID != nil {

		// query param service_detail_id
		var qrServiceDetailID int64
		if o.ServiceDetailID != nil {
			qrServiceDetailID = *o.ServiceDetailID
		}
		qServiceDetailID := swag.FormatInt64(qrServiceDetailID)
		if qServiceDetailID != "" {
			if err := r.SetQueryParam("service_detail_id", qServiceDetailID); err != nil {
				return err
			}
		}

	}

	if o.ServiceID != nil {

		// query param service_id
		var qrServiceID int64
		if o.ServiceID != nil {
			qrServiceID = *o.ServiceID
		}
		qServiceID := swag.FormatInt64(qrServiceID)
		if qServiceID != "" {
			if err := r.SetQueryParam("service_id", qServiceID); err != nil {
				return err
			}
		}

	}

	if o.UserID != nil {

		// query param user_id
		var qrUserID string
		if o.UserID != nil {
			qrUserID = *o.UserID
		}
		qUserID := qrUserID
		if qUserID != "" {
			if err := r.SetQueryParam("user_id", qUserID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
