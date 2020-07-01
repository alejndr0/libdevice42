// Code generated by go-swagger; DO NOT EDIT.

package devices

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

// NewPutCustomFieldParams creates a new PutCustomFieldParams object
// with the default values initialized.
func NewPutCustomFieldParams() *PutCustomFieldParams {
	var ()
	return &PutCustomFieldParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutCustomFieldParamsWithTimeout creates a new PutCustomFieldParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutCustomFieldParamsWithTimeout(timeout time.Duration) *PutCustomFieldParams {
	var ()
	return &PutCustomFieldParams{

		timeout: timeout,
	}
}

// NewPutCustomFieldParamsWithContext creates a new PutCustomFieldParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutCustomFieldParamsWithContext(ctx context.Context) *PutCustomFieldParams {
	var ()
	return &PutCustomFieldParams{

		Context: ctx,
	}
}

// NewPutCustomFieldParamsWithHTTPClient creates a new PutCustomFieldParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutCustomFieldParamsWithHTTPClient(client *http.Client) *PutCustomFieldParams {
	var ()
	return &PutCustomFieldParams{
		HTTPClient: client,
	}
}

/*PutCustomFieldParams contains all the parameters to send to the API endpoint
for the put custom field operation typically these are written to a http.Request
*/
type PutCustomFieldParams struct {

	/*AddToPicklist
	  Comma separated values to add to picklist. If type is picklist and custom field is new, this is a required field. Duplicates will be ignored.

	*/
	AddToPicklist *string
	/*Asset
	  if there is more than 1 device with the same asset #, this will return “device not found”, or

	*/
	Asset *string
	/*BulkFields
	  comma separated key value pairs, with key and value separated by colon. e.g.key1:value1, key2:value2 (added in v6.4.1)

	*/
	BulkFields *string
	/*ClearValue
	  yes to clear existing value for that field (added in v6.1.0)

	*/
	ClearValue *string
	/*ID
	  ID of the device

	*/
	ID *int64
	/*Key
	  Can be new or existing. This is the custom field name.

	*/
	Key string
	/*Name
	  name of the device, or

	*/
	Name *string
	/*Notes*/
	Notes *string
	/*RelatedFieldName
	  Required if type = related_field. The existing field to relate this custom field to. Can be: appcomp (for application components), asset, building, certificate, circuit, cusotmer, device, endusers, hardware (for device hardware models), ip_address, natip, netport (for ports), os, part, partmodel, password, pdu (for power units), pdu_model (for power unit models), ports, purchase, purchaselineitem (for a line item on a purchase order), rack, room, software, vlan (for subnets), switch_vlan (for vlans), vrfgroup

	*/
	RelatedFieldName *string
	/*Serial
	  if there is more than 1 device with the same serial #, this will return “device not found”

	*/
	Serial *string
	/*Type
	  this is the custom field type. If left blank, default is text. Date should be formatted as YYYY-MM-DD

	*/
	Type *string
	/*Value*/
	Value *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put custom field params
func (o *PutCustomFieldParams) WithTimeout(timeout time.Duration) *PutCustomFieldParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put custom field params
func (o *PutCustomFieldParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put custom field params
func (o *PutCustomFieldParams) WithContext(ctx context.Context) *PutCustomFieldParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put custom field params
func (o *PutCustomFieldParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put custom field params
func (o *PutCustomFieldParams) WithHTTPClient(client *http.Client) *PutCustomFieldParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put custom field params
func (o *PutCustomFieldParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddToPicklist adds the addToPicklist to the put custom field params
func (o *PutCustomFieldParams) WithAddToPicklist(addToPicklist *string) *PutCustomFieldParams {
	o.SetAddToPicklist(addToPicklist)
	return o
}

// SetAddToPicklist adds the addToPicklist to the put custom field params
func (o *PutCustomFieldParams) SetAddToPicklist(addToPicklist *string) {
	o.AddToPicklist = addToPicklist
}

// WithAsset adds the asset to the put custom field params
func (o *PutCustomFieldParams) WithAsset(asset *string) *PutCustomFieldParams {
	o.SetAsset(asset)
	return o
}

// SetAsset adds the asset to the put custom field params
func (o *PutCustomFieldParams) SetAsset(asset *string) {
	o.Asset = asset
}

// WithBulkFields adds the bulkFields to the put custom field params
func (o *PutCustomFieldParams) WithBulkFields(bulkFields *string) *PutCustomFieldParams {
	o.SetBulkFields(bulkFields)
	return o
}

// SetBulkFields adds the bulkFields to the put custom field params
func (o *PutCustomFieldParams) SetBulkFields(bulkFields *string) {
	o.BulkFields = bulkFields
}

// WithClearValue adds the clearValue to the put custom field params
func (o *PutCustomFieldParams) WithClearValue(clearValue *string) *PutCustomFieldParams {
	o.SetClearValue(clearValue)
	return o
}

// SetClearValue adds the clearValue to the put custom field params
func (o *PutCustomFieldParams) SetClearValue(clearValue *string) {
	o.ClearValue = clearValue
}

// WithID adds the id to the put custom field params
func (o *PutCustomFieldParams) WithID(id *int64) *PutCustomFieldParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put custom field params
func (o *PutCustomFieldParams) SetID(id *int64) {
	o.ID = id
}

// WithKey adds the key to the put custom field params
func (o *PutCustomFieldParams) WithKey(key string) *PutCustomFieldParams {
	o.SetKey(key)
	return o
}

// SetKey adds the key to the put custom field params
func (o *PutCustomFieldParams) SetKey(key string) {
	o.Key = key
}

// WithName adds the name to the put custom field params
func (o *PutCustomFieldParams) WithName(name *string) *PutCustomFieldParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the put custom field params
func (o *PutCustomFieldParams) SetName(name *string) {
	o.Name = name
}

// WithNotes adds the notes to the put custom field params
func (o *PutCustomFieldParams) WithNotes(notes *string) *PutCustomFieldParams {
	o.SetNotes(notes)
	return o
}

// SetNotes adds the notes to the put custom field params
func (o *PutCustomFieldParams) SetNotes(notes *string) {
	o.Notes = notes
}

// WithRelatedFieldName adds the relatedFieldName to the put custom field params
func (o *PutCustomFieldParams) WithRelatedFieldName(relatedFieldName *string) *PutCustomFieldParams {
	o.SetRelatedFieldName(relatedFieldName)
	return o
}

// SetRelatedFieldName adds the relatedFieldName to the put custom field params
func (o *PutCustomFieldParams) SetRelatedFieldName(relatedFieldName *string) {
	o.RelatedFieldName = relatedFieldName
}

// WithSerial adds the serial to the put custom field params
func (o *PutCustomFieldParams) WithSerial(serial *string) *PutCustomFieldParams {
	o.SetSerial(serial)
	return o
}

// SetSerial adds the serial to the put custom field params
func (o *PutCustomFieldParams) SetSerial(serial *string) {
	o.Serial = serial
}

// WithType adds the typeVar to the put custom field params
func (o *PutCustomFieldParams) WithType(typeVar *string) *PutCustomFieldParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the put custom field params
func (o *PutCustomFieldParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WithValue adds the value to the put custom field params
func (o *PutCustomFieldParams) WithValue(value *string) *PutCustomFieldParams {
	o.SetValue(value)
	return o
}

// SetValue adds the value to the put custom field params
func (o *PutCustomFieldParams) SetValue(value *string) {
	o.Value = value
}

// WriteToRequest writes these params to a swagger request
func (o *PutCustomFieldParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Asset != nil {

		// form param asset
		var frAsset string
		if o.Asset != nil {
			frAsset = *o.Asset
		}
		fAsset := frAsset
		if fAsset != "" {
			if err := r.SetFormParam("asset", fAsset); err != nil {
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
		var frID int64
		if o.ID != nil {
			frID = *o.ID
		}
		fID := swag.FormatInt64(frID)
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

	if o.Serial != nil {

		// form param serial
		var frSerial string
		if o.Serial != nil {
			frSerial = *o.Serial
		}
		fSerial := frSerial
		if fSerial != "" {
			if err := r.SetFormParam("serial", fSerial); err != nil {
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