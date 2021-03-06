// Code generated by go-swagger; DO NOT EDIT.

package ip_a_m

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

// NewGetIPAMCustomerIDParams creates a new GetIPAMCustomerIDParams object
// with the default values initialized.
func NewGetIPAMCustomerIDParams() *GetIPAMCustomerIDParams {
	var ()
	return &GetIPAMCustomerIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIPAMCustomerIDParamsWithTimeout creates a new GetIPAMCustomerIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIPAMCustomerIDParamsWithTimeout(timeout time.Duration) *GetIPAMCustomerIDParams {
	var ()
	return &GetIPAMCustomerIDParams{

		timeout: timeout,
	}
}

// NewGetIPAMCustomerIDParamsWithContext creates a new GetIPAMCustomerIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIPAMCustomerIDParamsWithContext(ctx context.Context) *GetIPAMCustomerIDParams {
	var ()
	return &GetIPAMCustomerIDParams{

		Context: ctx,
	}
}

// NewGetIPAMCustomerIDParamsWithHTTPClient creates a new GetIPAMCustomerIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIPAMCustomerIDParamsWithHTTPClient(client *http.Client) *GetIPAMCustomerIDParams {
	var ()
	return &GetIPAMCustomerIDParams{
		HTTPClient: client,
	}
}

/*GetIPAMCustomerIDParams contains all the parameters to send to the API endpoint
for the get IP a m customer id operation typically these are written to a http.Request
*/
type GetIPAMCustomerIDParams struct {

	/*CustomerID
	  Customer Id

	*/
	CustomerID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get IP a m customer id params
func (o *GetIPAMCustomerIDParams) WithTimeout(timeout time.Duration) *GetIPAMCustomerIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get IP a m customer id params
func (o *GetIPAMCustomerIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get IP a m customer id params
func (o *GetIPAMCustomerIDParams) WithContext(ctx context.Context) *GetIPAMCustomerIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get IP a m customer id params
func (o *GetIPAMCustomerIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get IP a m customer id params
func (o *GetIPAMCustomerIDParams) WithHTTPClient(client *http.Client) *GetIPAMCustomerIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get IP a m customer id params
func (o *GetIPAMCustomerIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerID adds the customerID to the get IP a m customer id params
func (o *GetIPAMCustomerIDParams) WithCustomerID(customerID int64) *GetIPAMCustomerIDParams {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the get IP a m customer id params
func (o *GetIPAMCustomerIDParams) SetCustomerID(customerID int64) {
	o.CustomerID = customerID
}

// WriteToRequest writes these params to a swagger request
func (o *GetIPAMCustomerIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customer_id
	if err := r.SetPathParam("customer_id", swag.FormatInt64(o.CustomerID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
