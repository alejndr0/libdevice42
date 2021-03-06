// Code generated by go-swagger; DO NOT EDIT.

package rooms

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

// NewPostAssetsRoomParams creates a new PostAssetsRoomParams object
// with the default values initialized.
func NewPostAssetsRoomParams() *PostAssetsRoomParams {
	var ()
	return &PostAssetsRoomParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAssetsRoomParamsWithTimeout creates a new PostAssetsRoomParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAssetsRoomParamsWithTimeout(timeout time.Duration) *PostAssetsRoomParams {
	var ()
	return &PostAssetsRoomParams{

		timeout: timeout,
	}
}

// NewPostAssetsRoomParamsWithContext creates a new PostAssetsRoomParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAssetsRoomParamsWithContext(ctx context.Context) *PostAssetsRoomParams {
	var ()
	return &PostAssetsRoomParams{

		Context: ctx,
	}
}

// NewPostAssetsRoomParamsWithHTTPClient creates a new PostAssetsRoomParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAssetsRoomParamsWithHTTPClient(client *http.Client) *PostAssetsRoomParams {
	var ()
	return &PostAssetsRoomParams{
		HTTPClient: client,
	}
}

/*PostAssetsRoomParams contains all the parameters to send to the API endpoint
for the post assets room operation typically these are written to a http.Request
*/
type PostAssetsRoomParams struct {

	/*Asset
	  The name of the asset

	*/
	Asset string
	/*RoomID
	  Room ID or UI > Tools > Export > Room

	*/
	RoomID string
	/*Type*/
	Type string
	/*Wall
	  Choose 'middle' if you do not want the object to be placed along one of the 4 walls.

	*/
	Wall *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post assets room params
func (o *PostAssetsRoomParams) WithTimeout(timeout time.Duration) *PostAssetsRoomParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post assets room params
func (o *PostAssetsRoomParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post assets room params
func (o *PostAssetsRoomParams) WithContext(ctx context.Context) *PostAssetsRoomParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post assets room params
func (o *PostAssetsRoomParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post assets room params
func (o *PostAssetsRoomParams) WithHTTPClient(client *http.Client) *PostAssetsRoomParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post assets room params
func (o *PostAssetsRoomParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAsset adds the asset to the post assets room params
func (o *PostAssetsRoomParams) WithAsset(asset string) *PostAssetsRoomParams {
	o.SetAsset(asset)
	return o
}

// SetAsset adds the asset to the post assets room params
func (o *PostAssetsRoomParams) SetAsset(asset string) {
	o.Asset = asset
}

// WithRoomID adds the roomID to the post assets room params
func (o *PostAssetsRoomParams) WithRoomID(roomID string) *PostAssetsRoomParams {
	o.SetRoomID(roomID)
	return o
}

// SetRoomID adds the roomId to the post assets room params
func (o *PostAssetsRoomParams) SetRoomID(roomID string) {
	o.RoomID = roomID
}

// WithType adds the typeVar to the post assets room params
func (o *PostAssetsRoomParams) WithType(typeVar string) *PostAssetsRoomParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the post assets room params
func (o *PostAssetsRoomParams) SetType(typeVar string) {
	o.Type = typeVar
}

// WithWall adds the wall to the post assets room params
func (o *PostAssetsRoomParams) WithWall(wall *string) *PostAssetsRoomParams {
	o.SetWall(wall)
	return o
}

// SetWall adds the wall to the post assets room params
func (o *PostAssetsRoomParams) SetWall(wall *string) {
	o.Wall = wall
}

// WriteToRequest writes these params to a swagger request
func (o *PostAssetsRoomParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form param asset
	frAsset := o.Asset
	fAsset := frAsset
	if fAsset != "" {
		if err := r.SetFormParam("asset", fAsset); err != nil {
			return err
		}
	}

	// form param room_id
	frRoomID := o.RoomID
	fRoomID := frRoomID
	if fRoomID != "" {
		if err := r.SetFormParam("room_id", fRoomID); err != nil {
			return err
		}
	}

	// form param type
	frType := o.Type
	fType := frType
	if fType != "" {
		if err := r.SetFormParam("type", fType); err != nil {
			return err
		}
	}

	if o.Wall != nil {

		// form param wall
		var frWall string
		if o.Wall != nil {
			frWall = *o.Wall
		}
		fWall := frWall
		if fWall != "" {
			if err := r.SetFormParam("wall", fWall); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
