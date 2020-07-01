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
)

// NewPostIPAMsubnetsParams creates a new PostIPAMsubnetsParams object
// with the default values initialized.
func NewPostIPAMsubnetsParams() *PostIPAMsubnetsParams {
	var ()
	return &PostIPAMsubnetsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostIPAMsubnetsParamsWithTimeout creates a new PostIPAMsubnetsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostIPAMsubnetsParamsWithTimeout(timeout time.Duration) *PostIPAMsubnetsParams {
	var ()
	return &PostIPAMsubnetsParams{

		timeout: timeout,
	}
}

// NewPostIPAMsubnetsParamsWithContext creates a new PostIPAMsubnetsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostIPAMsubnetsParamsWithContext(ctx context.Context) *PostIPAMsubnetsParams {
	var ()
	return &PostIPAMsubnetsParams{

		Context: ctx,
	}
}

// NewPostIPAMsubnetsParamsWithHTTPClient creates a new PostIPAMsubnetsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostIPAMsubnetsParamsWithHTTPClient(client *http.Client) *PostIPAMsubnetsParams {
	var ()
	return &PostIPAMsubnetsParams{
		HTTPClient: client,
	}
}

/*PostIPAMsubnetsParams contains all the parameters to send to the API endpoint
for the post IP a msubnets operation typically these are written to a http.Request
*/
type PostIPAMsubnetsParams struct {

	/*Allocated
	  ‘yes’ if allocated. ‘no’ (default) if unallocated.

	*/
	Allocated *string
	/*Assigned
	  ‘yes’ if assigned. ‘no’ (default) if unassigned.

	*/
	Assigned *string
	/*AutoAddIps
	  If ‘yes’, addresses within subnet will be automatically added to Device42. (Only available in POST)

	*/
	AutoAddIps *string
	/*Category
	  If multitenancy is on, admin groups that have access to this object are specified here, e.g. Prod_East:no,Corp:yes specifies that the admin groups for this object are Prod_East with view only permission and Corp with change permission.

	*/
	Category *string
	/*CategoryID
	  ID of the category

	*/
	CategoryID *string
	/*CustomerID*/
	CustomerID *string
	/*Description*/
	Description *string
	/*Gateway
	  Gateway (added in v7.2.0)

	*/
	Gateway *string
	/*MaskBits
	  Cannot be modified after subnet creation

	*/
	MaskBits string
	/*Name*/
	Name *string
	/*Network
	  Required for creation, cannot be modified after subnet creation.

	*/
	Network string
	/*Notes
	  Any additional notes

	*/
	Notes *string
	/*ParentSubnetID
	  Change the parent subnet of the subnet. Note: must be valid parent.

	*/
	ParentSubnetID *string
	/*ParentVlanID*/
	ParentVlanID *string
	/*RangeBegin
	  Range Begin (added in v7.2.0)

	*/
	RangeBegin *string
	/*RangeEnd
	  Range End (added in v7.2.0)

	*/
	RangeEnd *string
	/*ServiceLevel
	  Must already exist

	*/
	ServiceLevel *string
	/*VrfGroup
	  VRF group name

	*/
	VrfGroup *string
	/*VrfGroupID
	  ID of the VRF group

	*/
	VrfGroupID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post IP a msubnets params
func (o *PostIPAMsubnetsParams) WithTimeout(timeout time.Duration) *PostIPAMsubnetsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post IP a msubnets params
func (o *PostIPAMsubnetsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post IP a msubnets params
func (o *PostIPAMsubnetsParams) WithContext(ctx context.Context) *PostIPAMsubnetsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post IP a msubnets params
func (o *PostIPAMsubnetsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post IP a msubnets params
func (o *PostIPAMsubnetsParams) WithHTTPClient(client *http.Client) *PostIPAMsubnetsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post IP a msubnets params
func (o *PostIPAMsubnetsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllocated adds the allocated to the post IP a msubnets params
func (o *PostIPAMsubnetsParams) WithAllocated(allocated *string) *PostIPAMsubnetsParams {
	o.SetAllocated(allocated)
	return o
}

// SetAllocated adds the allocated to the post IP a msubnets params
func (o *PostIPAMsubnetsParams) SetAllocated(allocated *string) {
	o.Allocated = allocated
}

// WithAssigned adds the assigned to the post IP a msubnets params
func (o *PostIPAMsubnetsParams) WithAssigned(assigned *string) *PostIPAMsubnetsParams {
	o.SetAssigned(assigned)
	return o
}

// SetAssigned adds the assigned to the post IP a msubnets params
func (o *PostIPAMsubnetsParams) SetAssigned(assigned *string) {
	o.Assigned = assigned
}

// WithAutoAddIps adds the autoAddIps to the post IP a msubnets params
func (o *PostIPAMsubnetsParams) WithAutoAddIps(autoAddIps *string) *PostIPAMsubnetsParams {
	o.SetAutoAddIps(autoAddIps)
	return o
}

// SetAutoAddIps adds the autoAddIps to the post IP a msubnets params
func (o *PostIPAMsubnetsParams) SetAutoAddIps(autoAddIps *string) {
	o.AutoAddIps = autoAddIps
}

// WithCategory adds the category to the post IP a msubnets params
func (o *PostIPAMsubnetsParams) WithCategory(category *string) *PostIPAMsubnetsParams {
	o.SetCategory(category)
	return o
}

// SetCategory adds the category to the post IP a msubnets params
func (o *PostIPAMsubnetsParams) SetCategory(category *string) {
	o.Category = category
}

// WithCategoryID adds the categoryID to the post IP a msubnets params
func (o *PostIPAMsubnetsParams) WithCategoryID(categoryID *string) *PostIPAMsubnetsParams {
	o.SetCategoryID(categoryID)
	return o
}

// SetCategoryID adds the categoryId to the post IP a msubnets params
func (o *PostIPAMsubnetsParams) SetCategoryID(categoryID *string) {
	o.CategoryID = categoryID
}

// WithCustomerID adds the customerID to the post IP a msubnets params
func (o *PostIPAMsubnetsParams) WithCustomerID(customerID *string) *PostIPAMsubnetsParams {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the post IP a msubnets params
func (o *PostIPAMsubnetsParams) SetCustomerID(customerID *string) {
	o.CustomerID = customerID
}

// WithDescription adds the description to the post IP a msubnets params
func (o *PostIPAMsubnetsParams) WithDescription(description *string) *PostIPAMsubnetsParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the post IP a msubnets params
func (o *PostIPAMsubnetsParams) SetDescription(description *string) {
	o.Description = description
}

// WithGateway adds the gateway to the post IP a msubnets params
func (o *PostIPAMsubnetsParams) WithGateway(gateway *string) *PostIPAMsubnetsParams {
	o.SetGateway(gateway)
	return o
}

// SetGateway adds the gateway to the post IP a msubnets params
func (o *PostIPAMsubnetsParams) SetGateway(gateway *string) {
	o.Gateway = gateway
}

// WithMaskBits adds the maskBits to the post IP a msubnets params
func (o *PostIPAMsubnetsParams) WithMaskBits(maskBits string) *PostIPAMsubnetsParams {
	o.SetMaskBits(maskBits)
	return o
}

// SetMaskBits adds the maskBits to the post IP a msubnets params
func (o *PostIPAMsubnetsParams) SetMaskBits(maskBits string) {
	o.MaskBits = maskBits
}

// WithName adds the name to the post IP a msubnets params
func (o *PostIPAMsubnetsParams) WithName(name *string) *PostIPAMsubnetsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the post IP a msubnets params
func (o *PostIPAMsubnetsParams) SetName(name *string) {
	o.Name = name
}

// WithNetwork adds the network to the post IP a msubnets params
func (o *PostIPAMsubnetsParams) WithNetwork(network string) *PostIPAMsubnetsParams {
	o.SetNetwork(network)
	return o
}

// SetNetwork adds the network to the post IP a msubnets params
func (o *PostIPAMsubnetsParams) SetNetwork(network string) {
	o.Network = network
}

// WithNotes adds the notes to the post IP a msubnets params
func (o *PostIPAMsubnetsParams) WithNotes(notes *string) *PostIPAMsubnetsParams {
	o.SetNotes(notes)
	return o
}

// SetNotes adds the notes to the post IP a msubnets params
func (o *PostIPAMsubnetsParams) SetNotes(notes *string) {
	o.Notes = notes
}

// WithParentSubnetID adds the parentSubnetID to the post IP a msubnets params
func (o *PostIPAMsubnetsParams) WithParentSubnetID(parentSubnetID *string) *PostIPAMsubnetsParams {
	o.SetParentSubnetID(parentSubnetID)
	return o
}

// SetParentSubnetID adds the parentSubnetId to the post IP a msubnets params
func (o *PostIPAMsubnetsParams) SetParentSubnetID(parentSubnetID *string) {
	o.ParentSubnetID = parentSubnetID
}

// WithParentVlanID adds the parentVlanID to the post IP a msubnets params
func (o *PostIPAMsubnetsParams) WithParentVlanID(parentVlanID *string) *PostIPAMsubnetsParams {
	o.SetParentVlanID(parentVlanID)
	return o
}

// SetParentVlanID adds the parentVlanId to the post IP a msubnets params
func (o *PostIPAMsubnetsParams) SetParentVlanID(parentVlanID *string) {
	o.ParentVlanID = parentVlanID
}

// WithRangeBegin adds the rangeBegin to the post IP a msubnets params
func (o *PostIPAMsubnetsParams) WithRangeBegin(rangeBegin *string) *PostIPAMsubnetsParams {
	o.SetRangeBegin(rangeBegin)
	return o
}

// SetRangeBegin adds the rangeBegin to the post IP a msubnets params
func (o *PostIPAMsubnetsParams) SetRangeBegin(rangeBegin *string) {
	o.RangeBegin = rangeBegin
}

// WithRangeEnd adds the rangeEnd to the post IP a msubnets params
func (o *PostIPAMsubnetsParams) WithRangeEnd(rangeEnd *string) *PostIPAMsubnetsParams {
	o.SetRangeEnd(rangeEnd)
	return o
}

// SetRangeEnd adds the rangeEnd to the post IP a msubnets params
func (o *PostIPAMsubnetsParams) SetRangeEnd(rangeEnd *string) {
	o.RangeEnd = rangeEnd
}

// WithServiceLevel adds the serviceLevel to the post IP a msubnets params
func (o *PostIPAMsubnetsParams) WithServiceLevel(serviceLevel *string) *PostIPAMsubnetsParams {
	o.SetServiceLevel(serviceLevel)
	return o
}

// SetServiceLevel adds the serviceLevel to the post IP a msubnets params
func (o *PostIPAMsubnetsParams) SetServiceLevel(serviceLevel *string) {
	o.ServiceLevel = serviceLevel
}

// WithVrfGroup adds the vrfGroup to the post IP a msubnets params
func (o *PostIPAMsubnetsParams) WithVrfGroup(vrfGroup *string) *PostIPAMsubnetsParams {
	o.SetVrfGroup(vrfGroup)
	return o
}

// SetVrfGroup adds the vrfGroup to the post IP a msubnets params
func (o *PostIPAMsubnetsParams) SetVrfGroup(vrfGroup *string) {
	o.VrfGroup = vrfGroup
}

// WithVrfGroupID adds the vrfGroupID to the post IP a msubnets params
func (o *PostIPAMsubnetsParams) WithVrfGroupID(vrfGroupID *string) *PostIPAMsubnetsParams {
	o.SetVrfGroupID(vrfGroupID)
	return o
}

// SetVrfGroupID adds the vrfGroupId to the post IP a msubnets params
func (o *PostIPAMsubnetsParams) SetVrfGroupID(vrfGroupID *string) {
	o.VrfGroupID = vrfGroupID
}

// WriteToRequest writes these params to a swagger request
func (o *PostIPAMsubnetsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Allocated != nil {

		// form param allocated
		var frAllocated string
		if o.Allocated != nil {
			frAllocated = *o.Allocated
		}
		fAllocated := frAllocated
		if fAllocated != "" {
			if err := r.SetFormParam("allocated", fAllocated); err != nil {
				return err
			}
		}

	}

	if o.Assigned != nil {

		// form param assigned
		var frAssigned string
		if o.Assigned != nil {
			frAssigned = *o.Assigned
		}
		fAssigned := frAssigned
		if fAssigned != "" {
			if err := r.SetFormParam("assigned", fAssigned); err != nil {
				return err
			}
		}

	}

	if o.AutoAddIps != nil {

		// form param auto_add_ips
		var frAutoAddIps string
		if o.AutoAddIps != nil {
			frAutoAddIps = *o.AutoAddIps
		}
		fAutoAddIps := frAutoAddIps
		if fAutoAddIps != "" {
			if err := r.SetFormParam("auto_add_ips", fAutoAddIps); err != nil {
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

	if o.CategoryID != nil {

		// form param category_id
		var frCategoryID string
		if o.CategoryID != nil {
			frCategoryID = *o.CategoryID
		}
		fCategoryID := frCategoryID
		if fCategoryID != "" {
			if err := r.SetFormParam("category_id", fCategoryID); err != nil {
				return err
			}
		}

	}

	if o.CustomerID != nil {

		// form param customer_id
		var frCustomerID string
		if o.CustomerID != nil {
			frCustomerID = *o.CustomerID
		}
		fCustomerID := frCustomerID
		if fCustomerID != "" {
			if err := r.SetFormParam("customer_id", fCustomerID); err != nil {
				return err
			}
		}

	}

	if o.Description != nil {

		// form param description
		var frDescription string
		if o.Description != nil {
			frDescription = *o.Description
		}
		fDescription := frDescription
		if fDescription != "" {
			if err := r.SetFormParam("description", fDescription); err != nil {
				return err
			}
		}

	}

	if o.Gateway != nil {

		// form param gateway
		var frGateway string
		if o.Gateway != nil {
			frGateway = *o.Gateway
		}
		fGateway := frGateway
		if fGateway != "" {
			if err := r.SetFormParam("gateway", fGateway); err != nil {
				return err
			}
		}

	}

	// form param mask_bits
	frMaskBits := o.MaskBits
	fMaskBits := frMaskBits
	if fMaskBits != "" {
		if err := r.SetFormParam("mask_bits", fMaskBits); err != nil {
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

	// form param network
	frNetwork := o.Network
	fNetwork := frNetwork
	if fNetwork != "" {
		if err := r.SetFormParam("network", fNetwork); err != nil {
			return err
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

	if o.ParentSubnetID != nil {

		// form param parent_subnet_id
		var frParentSubnetID string
		if o.ParentSubnetID != nil {
			frParentSubnetID = *o.ParentSubnetID
		}
		fParentSubnetID := frParentSubnetID
		if fParentSubnetID != "" {
			if err := r.SetFormParam("parent_subnet_id", fParentSubnetID); err != nil {
				return err
			}
		}

	}

	if o.ParentVlanID != nil {

		// form param parent_vlan_id
		var frParentVlanID string
		if o.ParentVlanID != nil {
			frParentVlanID = *o.ParentVlanID
		}
		fParentVlanID := frParentVlanID
		if fParentVlanID != "" {
			if err := r.SetFormParam("parent_vlan_id", fParentVlanID); err != nil {
				return err
			}
		}

	}

	if o.RangeBegin != nil {

		// form param range_begin
		var frRangeBegin string
		if o.RangeBegin != nil {
			frRangeBegin = *o.RangeBegin
		}
		fRangeBegin := frRangeBegin
		if fRangeBegin != "" {
			if err := r.SetFormParam("range_begin", fRangeBegin); err != nil {
				return err
			}
		}

	}

	if o.RangeEnd != nil {

		// form param range_end
		var frRangeEnd string
		if o.RangeEnd != nil {
			frRangeEnd = *o.RangeEnd
		}
		fRangeEnd := frRangeEnd
		if fRangeEnd != "" {
			if err := r.SetFormParam("range_end", fRangeEnd); err != nil {
				return err
			}
		}

	}

	if o.ServiceLevel != nil {

		// form param service_level
		var frServiceLevel string
		if o.ServiceLevel != nil {
			frServiceLevel = *o.ServiceLevel
		}
		fServiceLevel := frServiceLevel
		if fServiceLevel != "" {
			if err := r.SetFormParam("service_level", fServiceLevel); err != nil {
				return err
			}
		}

	}

	if o.VrfGroup != nil {

		// form param vrf_group
		var frVrfGroup string
		if o.VrfGroup != nil {
			frVrfGroup = *o.VrfGroup
		}
		fVrfGroup := frVrfGroup
		if fVrfGroup != "" {
			if err := r.SetFormParam("vrf_group", fVrfGroup); err != nil {
				return err
			}
		}

	}

	if o.VrfGroupID != nil {

		// form param vrf_group_id
		var frVrfGroupID string
		if o.VrfGroupID != nil {
			frVrfGroupID = *o.VrfGroupID
		}
		fVrfGroupID := frVrfGroupID
		if fVrfGroupID != "" {
			if err := r.SetFormParam("vrf_group_id", fVrfGroupID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}