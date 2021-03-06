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
)

// NewPutCustomFieldsAssetParams creates a new PutCustomFieldsAssetParams object
// with the default values initialized.
func NewPutCustomFieldsAssetParams() *PutCustomFieldsAssetParams {
	var ()
	return &PutCustomFieldsAssetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutCustomFieldsAssetParamsWithTimeout creates a new PutCustomFieldsAssetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutCustomFieldsAssetParamsWithTimeout(timeout time.Duration) *PutCustomFieldsAssetParams {
	var ()
	return &PutCustomFieldsAssetParams{

		timeout: timeout,
	}
}

// NewPutCustomFieldsAssetParamsWithContext creates a new PutCustomFieldsAssetParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutCustomFieldsAssetParamsWithContext(ctx context.Context) *PutCustomFieldsAssetParams {
	var ()
	return &PutCustomFieldsAssetParams{

		Context: ctx,
	}
}

// NewPutCustomFieldsAssetParamsWithHTTPClient creates a new PutCustomFieldsAssetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutCustomFieldsAssetParamsWithHTTPClient(client *http.Client) *PutCustomFieldsAssetParams {
	var ()
	return &PutCustomFieldsAssetParams{
		HTTPClient: client,
	}
}

/*PutCustomFieldsAssetParams contains all the parameters to send to the API endpoint
for the put custom fields asset operation typically these are written to a http.Request
*/
type PutCustomFieldsAssetParams struct {

	/*AddToPicklist
	  Comma separated values to add to picklist. If type is picklist and custom field is new, this is a required field. Duplicates will be ignored.

	*/
	AddToPicklist *string
	/*BulkFields
	  comma separated key value pairs, with key and value separated by colon. e.g.key1:value1, key2:value2

	*/
	BulkFields *string
	/*ClearValue
	  yes to clear existing value for that field

	*/
	ClearValue *string
	/*ID
	  ID of asset

	*/
	ID *string
	/*Key
	  Can be new or existing. This is the custom field name.

	*/
	Key string
	/*Name
	  Name of asset

	*/
	Name *string
	/*Notes
	  Any additional notes

	*/
	Notes *string
	/*RelatedFieldName
	  Required if type = related field.

	*/
	RelatedFieldName *string
	/*Type
	  this is the custom field type. If left blank, default is text. Date should be formatted as YYYY-MM-DD

	*/
	Type *string
	/*Value
	  This will set the value of the custom field for the specific object.

	*/
	Value *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put custom fields asset params
func (o *PutCustomFieldsAssetParams) WithTimeout(timeout time.Duration) *PutCustomFieldsAssetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put custom fields asset params
func (o *PutCustomFieldsAssetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put custom fields asset params
func (o *PutCustomFieldsAssetParams) WithContext(ctx context.Context) *PutCustomFieldsAssetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put custom fields asset params
func (o *PutCustomFieldsAssetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put custom fields asset params
func (o *PutCustomFieldsAssetParams) WithHTTPClient(client *http.Client) *PutCustomFieldsAssetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put custom fields asset params
func (o *PutCustomFieldsAssetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddToPicklist adds the addToPicklist to the put custom fields asset params
func (o *PutCustomFieldsAssetParams) WithAddToPicklist(addToPicklist *string) *PutCustomFieldsAssetParams {
	o.SetAddToPicklist(addToPicklist)
	return o
}

// SetAddToPicklist adds the addToPicklist to the put custom fields asset params
func (o *PutCustomFieldsAssetParams) SetAddToPicklist(addToPicklist *string) {
	o.AddToPicklist = addToPicklist
}

// WithBulkFields adds the bulkFields to the put custom fields asset params
func (o *PutCustomFieldsAssetParams) WithBulkFields(bulkFields *string) *PutCustomFieldsAssetParams {
	o.SetBulkFields(bulkFields)
	return o
}

// SetBulkFields adds the bulkFields to the put custom fields asset params
func (o *PutCustomFieldsAssetParams) SetBulkFields(bulkFields *string) {
	o.BulkFields = bulkFields
}

// WithClearValue adds the clearValue to the put custom fields asset params
func (o *PutCustomFieldsAssetParams) WithClearValue(clearValue *string) *PutCustomFieldsAssetParams {
	o.SetClearValue(clearValue)
	return o
}

// SetClearValue adds the clearValue to the put custom fields asset params
func (o *PutCustomFieldsAssetParams) SetClearValue(clearValue *string) {
	o.ClearValue = clearValue
}

// WithID adds the id to the put custom fields asset params
func (o *PutCustomFieldsAssetParams) WithID(id *string) *PutCustomFieldsAssetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put custom fields asset params
func (o *PutCustomFieldsAssetParams) SetID(id *string) {
	o.ID = id
}

// WithKey adds the key to the put custom fields asset params
func (o *PutCustomFieldsAssetParams) WithKey(key string) *PutCustomFieldsAssetParams {
	o.SetKey(key)
	return o
}

// SetKey adds the key to the put custom fields asset params
func (o *PutCustomFieldsAssetParams) SetKey(key string) {
	o.Key = key
}

// WithName adds the name to the put custom fields asset params
func (o *PutCustomFieldsAssetParams) WithName(name *string) *PutCustomFieldsAssetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the put custom fields asset params
func (o *PutCustomFieldsAssetParams) SetName(name *string) {
	o.Name = name
}

// WithNotes adds the notes to the put custom fields asset params
func (o *PutCustomFieldsAssetParams) WithNotes(notes *string) *PutCustomFieldsAssetParams {
	o.SetNotes(notes)
	return o
}

// SetNotes adds the notes to the put custom fields asset params
func (o *PutCustomFieldsAssetParams) SetNotes(notes *string) {
	o.Notes = notes
}

// WithRelatedFieldName adds the relatedFieldName to the put custom fields asset params
func (o *PutCustomFieldsAssetParams) WithRelatedFieldName(relatedFieldName *string) *PutCustomFieldsAssetParams {
	o.SetRelatedFieldName(relatedFieldName)
	return o
}

// SetRelatedFieldName adds the relatedFieldName to the put custom fields asset params
func (o *PutCustomFieldsAssetParams) SetRelatedFieldName(relatedFieldName *string) {
	o.RelatedFieldName = relatedFieldName
}

// WithType adds the typeVar to the put custom fields asset params
func (o *PutCustomFieldsAssetParams) WithType(typeVar *string) *PutCustomFieldsAssetParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the put custom fields asset params
func (o *PutCustomFieldsAssetParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WithValue adds the value to the put custom fields asset params
func (o *PutCustomFieldsAssetParams) WithValue(value *string) *PutCustomFieldsAssetParams {
	o.SetValue(value)
	return o
}

// SetValue adds the value to the put custom fields asset params
func (o *PutCustomFieldsAssetParams) SetValue(value *string) {
	o.Value = value
}

// WriteToRequest writes these params to a swagger request
func (o *PutCustomFieldsAssetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AddToPicklist != nil {

		// form param add_to_picklist
		var frAddToPicklist string
		if o.AddToPicklist != nil {
			frAddToPicklist = *o.AddToPicklist
		}
		fAddToPicklist := frAddToPicklist
		if fAddToPicklist != "" {
			if err := r.SetFormParam("add_to_picklist", fAddToPicklist); err != nil {
				return err
			}
		}

	}

	if o.BulkFields != nil {

		// form param bulk_fields
		var frBulkFields string
		if o.BulkFields != nil {
			frBulkFields = *o.BulkFields
		}
		fBulkFields := frBulkFields
		if fBulkFields != "" {
			if err := r.SetFormParam("bulk_fields", fBulkFields); err != nil {
				return err
			}
		}

	}

	if o.ClearValue != nil {

		// form param clear_value
		var frClearValue string
		if o.ClearValue != nil {
			frClearValue = *o.ClearValue
		}
		fClearValue := frClearValue
		if fClearValue != "" {
			if err := r.SetFormParam("clear_value", fClearValue); err != nil {
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

	// form param key
	frKey := o.Key
	fKey := frKey
	if fKey != "" {
		if err := r.SetFormParam("key", fKey); err != nil {
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

	if o.RelatedFieldName != nil {

		// form param related_field_name
		var frRelatedFieldName string
		if o.RelatedFieldName != nil {
			frRelatedFieldName = *o.RelatedFieldName
		}
		fRelatedFieldName := frRelatedFieldName
		if fRelatedFieldName != "" {
			if err := r.SetFormParam("related_field_name", fRelatedFieldName); err != nil {
				return err
			}
		}

	}

	if o.Type != nil {

		// form param type
		var frType string
		if o.Type != nil {
			frType = *o.Type
		}
		fType := frType
		if fType != "" {
			if err := r.SetFormParam("type", fType); err != nil {
				return err
			}
		}

	}

	if o.Value != nil {

		// form param value
		var frValue string
		if o.Value != nil {
			frValue = *o.Value
		}
		fValue := frValue
		if fValue != "" {
			if err := r.SetFormParam("value", fValue); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
