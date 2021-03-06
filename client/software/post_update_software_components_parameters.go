// Code generated by go-swagger; DO NOT EDIT.

package software

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

// NewPostUpdateSoftwareComponentsParams creates a new PostUpdateSoftwareComponentsParams object
// with the default values initialized.
func NewPostUpdateSoftwareComponentsParams() *PostUpdateSoftwareComponentsParams {
	var (
		licensingModelDefault = string("Individual - Device/Perpetual")
		softwareTypeDefault   = string("Unmanaged")
	)
	return &PostUpdateSoftwareComponentsParams{
		LicensingModel: licensingModelDefault,
		SoftwareType:   softwareTypeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewPostUpdateSoftwareComponentsParamsWithTimeout creates a new PostUpdateSoftwareComponentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostUpdateSoftwareComponentsParamsWithTimeout(timeout time.Duration) *PostUpdateSoftwareComponentsParams {
	var (
		licensingModelDefault = string("Individual - Device/Perpetual")
		softwareTypeDefault   = string("Unmanaged")
	)
	return &PostUpdateSoftwareComponentsParams{
		LicensingModel: licensingModelDefault,
		SoftwareType:   softwareTypeDefault,

		timeout: timeout,
	}
}

// NewPostUpdateSoftwareComponentsParamsWithContext creates a new PostUpdateSoftwareComponentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostUpdateSoftwareComponentsParamsWithContext(ctx context.Context) *PostUpdateSoftwareComponentsParams {
	var (
		licensingModelDefault = string("Individual - Device/Perpetual")
		softwareTypeDefault   = string("Unmanaged")
	)
	return &PostUpdateSoftwareComponentsParams{
		LicensingModel: licensingModelDefault,
		SoftwareType:   softwareTypeDefault,

		Context: ctx,
	}
}

// NewPostUpdateSoftwareComponentsParamsWithHTTPClient creates a new PostUpdateSoftwareComponentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostUpdateSoftwareComponentsParamsWithHTTPClient(client *http.Client) *PostUpdateSoftwareComponentsParams {
	var (
		licensingModelDefault = string("Individual - Device/Perpetual")
		softwareTypeDefault   = string("Unmanaged")
	)
	return &PostUpdateSoftwareComponentsParams{
		LicensingModel: licensingModelDefault,
		SoftwareType:   softwareTypeDefault,
		HTTPClient:     client,
	}
}

/*PostUpdateSoftwareComponentsParams contains all the parameters to send to the API endpoint
for the post update software components operation typically these are written to a http.Request
*/
type PostUpdateSoftwareComponentsParams struct {

	/*Aliases
	  any software aliases

	*/
	Aliases *string
	/*Category
	  Filter by user-defined software categories (new or existing)

	*/
	Category *string
	/*ID
	  The ID of the software, required if not using NAME

	*/
	ID *string
	/*LicensingModel
	  Custom models can be made via UI. Existing values include: Free, Trial, Individual - User/Perpetual, Individual - User/Subscription, Named User / Perpetual, Volume - User/Perpetual, Concurrent - User/Perpetual, Individual - Device/Perpetual, Individual - Device/Subscription, Node Locked / Perpetual, Volume - Device/Perpetual, OEM, CAL / Per Seat Device, CAL / Per Seat User, CAL / Per Server, CAL / Per Processor, CAL / Per Mailbox

	*/
	LicensingModel string
	/*Name
	  The name of the software (new or existing)

	*/
	Name *string
	/*Notes
	  Any additional notes

	*/
	Notes *string
	/*SoftwareType*/
	SoftwareType string
	/*Tags*/
	Tags *string
	/*TagsRemove
	  remove tags from component

	*/
	TagsRemove *string
	/*TrackLicensedCountByKeys
	  whether or not to track software by discovered count

	*/
	TrackLicensedCountByKeys *string
	/*Vendor
	  Software Vendor (new or existing)

	*/
	Vendor *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post update software components params
func (o *PostUpdateSoftwareComponentsParams) WithTimeout(timeout time.Duration) *PostUpdateSoftwareComponentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post update software components params
func (o *PostUpdateSoftwareComponentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post update software components params
func (o *PostUpdateSoftwareComponentsParams) WithContext(ctx context.Context) *PostUpdateSoftwareComponentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post update software components params
func (o *PostUpdateSoftwareComponentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post update software components params
func (o *PostUpdateSoftwareComponentsParams) WithHTTPClient(client *http.Client) *PostUpdateSoftwareComponentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post update software components params
func (o *PostUpdateSoftwareComponentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAliases adds the aliases to the post update software components params
func (o *PostUpdateSoftwareComponentsParams) WithAliases(aliases *string) *PostUpdateSoftwareComponentsParams {
	o.SetAliases(aliases)
	return o
}

// SetAliases adds the aliases to the post update software components params
func (o *PostUpdateSoftwareComponentsParams) SetAliases(aliases *string) {
	o.Aliases = aliases
}

// WithCategory adds the category to the post update software components params
func (o *PostUpdateSoftwareComponentsParams) WithCategory(category *string) *PostUpdateSoftwareComponentsParams {
	o.SetCategory(category)
	return o
}

// SetCategory adds the category to the post update software components params
func (o *PostUpdateSoftwareComponentsParams) SetCategory(category *string) {
	o.Category = category
}

// WithID adds the id to the post update software components params
func (o *PostUpdateSoftwareComponentsParams) WithID(id *string) *PostUpdateSoftwareComponentsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post update software components params
func (o *PostUpdateSoftwareComponentsParams) SetID(id *string) {
	o.ID = id
}

// WithLicensingModel adds the licensingModel to the post update software components params
func (o *PostUpdateSoftwareComponentsParams) WithLicensingModel(licensingModel string) *PostUpdateSoftwareComponentsParams {
	o.SetLicensingModel(licensingModel)
	return o
}

// SetLicensingModel adds the licensingModel to the post update software components params
func (o *PostUpdateSoftwareComponentsParams) SetLicensingModel(licensingModel string) {
	o.LicensingModel = licensingModel
}

// WithName adds the name to the post update software components params
func (o *PostUpdateSoftwareComponentsParams) WithName(name *string) *PostUpdateSoftwareComponentsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the post update software components params
func (o *PostUpdateSoftwareComponentsParams) SetName(name *string) {
	o.Name = name
}

// WithNotes adds the notes to the post update software components params
func (o *PostUpdateSoftwareComponentsParams) WithNotes(notes *string) *PostUpdateSoftwareComponentsParams {
	o.SetNotes(notes)
	return o
}

// SetNotes adds the notes to the post update software components params
func (o *PostUpdateSoftwareComponentsParams) SetNotes(notes *string) {
	o.Notes = notes
}

// WithSoftwareType adds the softwareType to the post update software components params
func (o *PostUpdateSoftwareComponentsParams) WithSoftwareType(softwareType string) *PostUpdateSoftwareComponentsParams {
	o.SetSoftwareType(softwareType)
	return o
}

// SetSoftwareType adds the softwareType to the post update software components params
func (o *PostUpdateSoftwareComponentsParams) SetSoftwareType(softwareType string) {
	o.SoftwareType = softwareType
}

// WithTags adds the tags to the post update software components params
func (o *PostUpdateSoftwareComponentsParams) WithTags(tags *string) *PostUpdateSoftwareComponentsParams {
	o.SetTags(tags)
	return o
}

// SetTags adds the tags to the post update software components params
func (o *PostUpdateSoftwareComponentsParams) SetTags(tags *string) {
	o.Tags = tags
}

// WithTagsRemove adds the tagsRemove to the post update software components params
func (o *PostUpdateSoftwareComponentsParams) WithTagsRemove(tagsRemove *string) *PostUpdateSoftwareComponentsParams {
	o.SetTagsRemove(tagsRemove)
	return o
}

// SetTagsRemove adds the tagsRemove to the post update software components params
func (o *PostUpdateSoftwareComponentsParams) SetTagsRemove(tagsRemove *string) {
	o.TagsRemove = tagsRemove
}

// WithTrackLicensedCountByKeys adds the trackLicensedCountByKeys to the post update software components params
func (o *PostUpdateSoftwareComponentsParams) WithTrackLicensedCountByKeys(trackLicensedCountByKeys *string) *PostUpdateSoftwareComponentsParams {
	o.SetTrackLicensedCountByKeys(trackLicensedCountByKeys)
	return o
}

// SetTrackLicensedCountByKeys adds the trackLicensedCountByKeys to the post update software components params
func (o *PostUpdateSoftwareComponentsParams) SetTrackLicensedCountByKeys(trackLicensedCountByKeys *string) {
	o.TrackLicensedCountByKeys = trackLicensedCountByKeys
}

// WithVendor adds the vendor to the post update software components params
func (o *PostUpdateSoftwareComponentsParams) WithVendor(vendor *string) *PostUpdateSoftwareComponentsParams {
	o.SetVendor(vendor)
	return o
}

// SetVendor adds the vendor to the post update software components params
func (o *PostUpdateSoftwareComponentsParams) SetVendor(vendor *string) {
	o.Vendor = vendor
}

// WriteToRequest writes these params to a swagger request
func (o *PostUpdateSoftwareComponentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Aliases != nil {

		// form param aliases
		var frAliases string
		if o.Aliases != nil {
			frAliases = *o.Aliases
		}
		fAliases := frAliases
		if fAliases != "" {
			if err := r.SetFormParam("aliases", fAliases); err != nil {
				return err
			}
		}

	}

	if o.Category != nil {

		// form param category
		var frCategory string
		if o.Category != nil {
			frCategory = *o.Category
		}
		fCategory := frCategory
		if fCategory != "" {
			if err := r.SetFormParam("category", fCategory); err != nil {
				return err
			}
		}

	}

	if o.ID != nil {

		// form param id
		var frID string
		if o.ID != nil {
			frID = *o.ID
		}
		fID := frID
		if fID != "" {
			if err := r.SetFormParam("id", fID); err != nil {
				return err
			}
		}

	}

	// form param licensing_model
	frLicensingModel := o.LicensingModel
	fLicensingModel := frLicensingModel
	if fLicensingModel != "" {
		if err := r.SetFormParam("licensing_model", fLicensingModel); err != nil {
			return err
		}
	}

	if o.Name != nil {

		// form param name
		var frName string
		if o.Name != nil {
			frName = *o.Name
		}
		fName := frName
		if fName != "" {
			if err := r.SetFormParam("name", fName); err != nil {
				return err
			}
		}

	}

	if o.Notes != nil {

		// form param notes
		var frNotes string
		if o.Notes != nil {
			frNotes = *o.Notes
		}
		fNotes := frNotes
		if fNotes != "" {
			if err := r.SetFormParam("notes", fNotes); err != nil {
				return err
			}
		}

	}

	// form param software_type
	frSoftwareType := o.SoftwareType
	fSoftwareType := frSoftwareType
	if fSoftwareType != "" {
		if err := r.SetFormParam("software_type", fSoftwareType); err != nil {
			return err
		}
	}

	if o.Tags != nil {

		// form param tags
		var frTags string
		if o.Tags != nil {
			frTags = *o.Tags
		}
		fTags := frTags
		if fTags != "" {
			if err := r.SetFormParam("tags", fTags); err != nil {
				return err
			}
		}

	}

	if o.TagsRemove != nil {

		// form param tags_remove
		var frTagsRemove string
		if o.TagsRemove != nil {
			frTagsRemove = *o.TagsRemove
		}
		fTagsRemove := frTagsRemove
		if fTagsRemove != "" {
			if err := r.SetFormParam("tags_remove", fTagsRemove); err != nil {
				return err
			}
		}

	}

	if o.TrackLicensedCountByKeys != nil {

		// form param track_licensed_count_by_keys
		var frTrackLicensedCountByKeys string
		if o.TrackLicensedCountByKeys != nil {
			frTrackLicensedCountByKeys = *o.TrackLicensedCountByKeys
		}
		fTrackLicensedCountByKeys := frTrackLicensedCountByKeys
		if fTrackLicensedCountByKeys != "" {
			if err := r.SetFormParam("track_licensed_count_by_keys", fTrackLicensedCountByKeys); err != nil {
				return err
			}
		}

	}

	if o.Vendor != nil {

		// form param vendor
		var frVendor string
		if o.Vendor != nil {
			frVendor = *o.Vendor
		}
		fVendor := frVendor
		if fVendor != "" {
			if err := r.SetFormParam("vendor", fVendor); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
