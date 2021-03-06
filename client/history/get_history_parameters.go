// Code generated by go-swagger; DO NOT EDIT.

package history

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

// NewGetHistoryParams creates a new GetHistoryParams object
// with the default values initialized.
func NewGetHistoryParams() *GetHistoryParams {

	return &GetHistoryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetHistoryParamsWithTimeout creates a new GetHistoryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetHistoryParamsWithTimeout(timeout time.Duration) *GetHistoryParams {

	return &GetHistoryParams{

		timeout: timeout,
	}
}

// NewGetHistoryParamsWithContext creates a new GetHistoryParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetHistoryParamsWithContext(ctx context.Context) *GetHistoryParams {

	return &GetHistoryParams{

		Context: ctx,
	}
}

// NewGetHistoryParamsWithHTTPClient creates a new GetHistoryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetHistoryParamsWithHTTPClient(client *http.Client) *GetHistoryParams {

	return &GetHistoryParams{
		HTTPClient: client,
	}
}

/*GetHistoryParams contains all the parameters to send to the API endpoint
for the get history operation typically these are written to a http.Request
*/
type GetHistoryParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get history params
func (o *GetHistoryParams) WithTimeout(timeout time.Duration) *GetHistoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get history params
func (o *GetHistoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get history params
func (o *GetHistoryParams) WithContext(ctx context.Context) *GetHistoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get history params
func (o *GetHistoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get history params
func (o *GetHistoryParams) WithHTTPClient(client *http.Client) *GetHistoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get history params
func (o *GetHistoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetHistoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
