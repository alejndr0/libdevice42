// Code generated by go-swagger; DO NOT EDIT.

package assets

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

// NewDeleteAssetsParams creates a new DeleteAssetsParams object
// with the default values initialized.
func NewDeleteAssetsParams() *DeleteAssetsParams {
	var ()
	return &DeleteAssetsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAssetsParamsWithTimeout creates a new DeleteAssetsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAssetsParamsWithTimeout(timeout time.Duration) *DeleteAssetsParams {
	var ()
	return &DeleteAssetsParams{

		timeout: timeout,
	}
}

// NewDeleteAssetsParamsWithContext creates a new DeleteAssetsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAssetsParamsWithContext(ctx context.Context) *DeleteAssetsParams {
	var ()
	return &DeleteAssetsParams{

		Context: ctx,
	}
}

// NewDeleteAssetsParamsWithHTTPClient creates a new DeleteAssetsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAssetsParamsWithHTTPClient(client *http.Client) *DeleteAssetsParams {
	var ()
	return &DeleteAssetsParams{
		HTTPClient: client,
	}
}

/*DeleteAssetsParams contains all the parameters to send to the API endpoint
for the delete assets operation typically these are written to a http.Request
*/
type DeleteAssetsParams struct {

	/*AssetID
	  asset id

	*/
	AssetID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete assets params
func (o *DeleteAssetsParams) WithTimeout(timeout time.Duration) *DeleteAssetsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete assets params
func (o *DeleteAssetsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete assets params
func (o *DeleteAssetsParams) WithContext(ctx context.Context) *DeleteAssetsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete assets params
func (o *DeleteAssetsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete assets params
func (o *DeleteAssetsParams) WithHTTPClient(client *http.Client) *DeleteAssetsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete assets params
func (o *DeleteAssetsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAssetID adds the assetID to the delete assets params
func (o *DeleteAssetsParams) WithAssetID(assetID int64) *DeleteAssetsParams {
	o.SetAssetID(assetID)
	return o
}

// SetAssetID adds the assetId to the delete assets params
func (o *DeleteAssetsParams) SetAssetID(assetID int64) {
	o.AssetID = assetID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAssetsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param asset-id
	if err := r.SetPathParam("asset-id", swag.FormatInt64(o.AssetID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
