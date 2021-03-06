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

// NewPutAssetsParams creates a new PutAssetsParams object
// with the default values initialized.
func NewPutAssetsParams() *PutAssetsParams {
	var ()
	return &PutAssetsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutAssetsParamsWithTimeout creates a new PutAssetsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutAssetsParamsWithTimeout(timeout time.Duration) *PutAssetsParams {
	var ()
	return &PutAssetsParams{

		timeout: timeout,
	}
}

// NewPutAssetsParamsWithContext creates a new PutAssetsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutAssetsParamsWithContext(ctx context.Context) *PutAssetsParams {
	var ()
	return &PutAssetsParams{

		Context: ctx,
	}
}

// NewPutAssetsParamsWithHTTPClient creates a new PutAssetsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutAssetsParamsWithHTTPClient(client *http.Client) *PutAssetsParams {
	var ()
	return &PutAssetsParams{
		HTTPClient: client,
	}
}

/*PutAssetsParams contains all the parameters to send to the API endpoint
for the put assets operation typically these are written to a http.Request
*/
type PutAssetsParams struct {

	/*Asset
	  Name of asset. Required if asset_id not provided.

	*/
	Asset *string
	/*AssetID
	  Asset ID

	*/
	AssetID *int64
	/*AssetNo*/
	AssetNo *string
	/*BackImage
	  name of the back image file. Use instead of back_image_id.

	*/
	BackImage *string
	/*BackImageID
	  back image file id. You can see these from Tools > Import > Hardware Import for now.

	*/
	BackImageID *string
	/*Building*/
	Building *string
	/*Category
	  name of the category

	*/
	Category *string
	/*Customer
	  Name of existing customer.

	*/
	Customer *string
	/*Depth
	  Half depth by default. full to override. Optional.

	*/
	Depth *string
	/*DeviceID
	  ID of the related device

	*/
	DeviceID *int64
	/*Imgfile
	  name of the image file (Added in v5.8.2). Use instead of imgfile_id

	*/
	Imgfile *string
	/*ImgfileID
	  image file id. You can see these from Tools > Import > Hardware Import for now.

	*/
	ImgfileID *string
	/*Location
	  Location/region of instance deployment

	*/
	Location *string
	/*ModuleHost
	  Name of the Module host. Must be unique asset name for this to work. (use instead of ID, Added in v5.8.2)

	*/
	ModuleHost *string
	/*ModuleHostID*/
	ModuleHostID *string
	/*Notes
	  Any additional notes

	*/
	Notes *string
	/*NumberingStartFrom
	  This is starting # for patch panel ports. Defaults to 1 if not entered.

	*/
	NumberingStartFrom *string
	/*Orientation
	  Back if back facing. Otherwise ignored

	*/
	Orientation *string
	/*PatchPanelModel
	  Name of the patch panel model (use instead of ID, Added in v5.8.2)

	*/
	PatchPanelModel *string
	/*PatchPanelModelID
	  Patch Panel Model ID or UI Tools > Export > Patch Panel Model

	*/
	PatchPanelModelID *string
	/*PatchPanelModuleModel
	  Name of the patch panel module model (use instead of ID, Added in v5.8.2)

	*/
	PatchPanelModuleModel *string
	/*PatchPanelModuleModelID*/
	PatchPanelModuleModelID *string
	/*RackID
	  ID of existing rack to add asset to.

	*/
	RackID *string
	/*Room
	  Used to identify a unique rack or to place asset in existing room.

	*/
	Room *string
	/*SerialNo*/
	SerialNo *string
	/*ServiceLevel
	  In Service, Spare, Not in Service are pre-defined - or choose your own.

	*/
	ServiceLevel *string
	/*Size
	  Required if adding asset to rack. in U.

	*/
	Size *int64
	/*SlotNo*/
	SlotNo *string
	/*StartAt
	  Required if adding to rack. U Start location.

	*/
	StartAt *string
	/*Tags
	  add tags (comma separated)

	*/
	Tags *string
	/*TagsRemove
	  remove tags (comma separated)

	*/
	TagsRemove *string
	/*Type
	  Used to change/add the type of asset

	*/
	Type *string
	/*Vendor
	  Name of existing vendor

	*/
	Vendor *string
	/*Where
	  Location in a rack. Note: If mounted a size must be provided or available from the hardware model.

	*/
	Where *string
	/*XPos
	  A number between 0 and 2520 representing the position within the u slot in increments of 252, which is equal to 1/10th of the width of the rack. 0 will place a device flush left, 1260 will place the left side of a device in center.

	*/
	XPos *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put assets params
func (o *PutAssetsParams) WithTimeout(timeout time.Duration) *PutAssetsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put assets params
func (o *PutAssetsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put assets params
func (o *PutAssetsParams) WithContext(ctx context.Context) *PutAssetsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put assets params
func (o *PutAssetsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put assets params
func (o *PutAssetsParams) WithHTTPClient(client *http.Client) *PutAssetsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put assets params
func (o *PutAssetsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAsset adds the asset to the put assets params
func (o *PutAssetsParams) WithAsset(asset *string) *PutAssetsParams {
	o.SetAsset(asset)
	return o
}

// SetAsset adds the asset to the put assets params
func (o *PutAssetsParams) SetAsset(asset *string) {
	o.Asset = asset
}

// WithAssetID adds the assetID to the put assets params
func (o *PutAssetsParams) WithAssetID(assetID *int64) *PutAssetsParams {
	o.SetAssetID(assetID)
	return o
}

// SetAssetID adds the assetId to the put assets params
func (o *PutAssetsParams) SetAssetID(assetID *int64) {
	o.AssetID = assetID
}

// WithAssetNo adds the assetNo to the put assets params
func (o *PutAssetsParams) WithAssetNo(assetNo *string) *PutAssetsParams {
	o.SetAssetNo(assetNo)
	return o
}

// SetAssetNo adds the assetNo to the put assets params
func (o *PutAssetsParams) SetAssetNo(assetNo *string) {
	o.AssetNo = assetNo
}

// WithBackImage adds the backImage to the put assets params
func (o *PutAssetsParams) WithBackImage(backImage *string) *PutAssetsParams {
	o.SetBackImage(backImage)
	return o
}

// SetBackImage adds the backImage to the put assets params
func (o *PutAssetsParams) SetBackImage(backImage *string) {
	o.BackImage = backImage
}

// WithBackImageID adds the backImageID to the put assets params
func (o *PutAssetsParams) WithBackImageID(backImageID *string) *PutAssetsParams {
	o.SetBackImageID(backImageID)
	return o
}

// SetBackImageID adds the backImageId to the put assets params
func (o *PutAssetsParams) SetBackImageID(backImageID *string) {
	o.BackImageID = backImageID
}

// WithBuilding adds the building to the put assets params
func (o *PutAssetsParams) WithBuilding(building *string) *PutAssetsParams {
	o.SetBuilding(building)
	return o
}

// SetBuilding adds the building to the put assets params
func (o *PutAssetsParams) SetBuilding(building *string) {
	o.Building = building
}

// WithCategory adds the category to the put assets params
func (o *PutAssetsParams) WithCategory(category *string) *PutAssetsParams {
	o.SetCategory(category)
	return o
}

// SetCategory adds the category to the put assets params
func (o *PutAssetsParams) SetCategory(category *string) {
	o.Category = category
}

// WithCustomer adds the customer to the put assets params
func (o *PutAssetsParams) WithCustomer(customer *string) *PutAssetsParams {
	o.SetCustomer(customer)
	return o
}

// SetCustomer adds the customer to the put assets params
func (o *PutAssetsParams) SetCustomer(customer *string) {
	o.Customer = customer
}

// WithDepth adds the depth to the put assets params
func (o *PutAssetsParams) WithDepth(depth *string) *PutAssetsParams {
	o.SetDepth(depth)
	return o
}

// SetDepth adds the depth to the put assets params
func (o *PutAssetsParams) SetDepth(depth *string) {
	o.Depth = depth
}

// WithDeviceID adds the deviceID to the put assets params
func (o *PutAssetsParams) WithDeviceID(deviceID *int64) *PutAssetsParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the put assets params
func (o *PutAssetsParams) SetDeviceID(deviceID *int64) {
	o.DeviceID = deviceID
}

// WithImgfile adds the imgfile to the put assets params
func (o *PutAssetsParams) WithImgfile(imgfile *string) *PutAssetsParams {
	o.SetImgfile(imgfile)
	return o
}

// SetImgfile adds the imgfile to the put assets params
func (o *PutAssetsParams) SetImgfile(imgfile *string) {
	o.Imgfile = imgfile
}

// WithImgfileID adds the imgfileID to the put assets params
func (o *PutAssetsParams) WithImgfileID(imgfileID *string) *PutAssetsParams {
	o.SetImgfileID(imgfileID)
	return o
}

// SetImgfileID adds the imgfileId to the put assets params
func (o *PutAssetsParams) SetImgfileID(imgfileID *string) {
	o.ImgfileID = imgfileID
}

// WithLocation adds the location to the put assets params
func (o *PutAssetsParams) WithLocation(location *string) *PutAssetsParams {
	o.SetLocation(location)
	return o
}

// SetLocation adds the location to the put assets params
func (o *PutAssetsParams) SetLocation(location *string) {
	o.Location = location
}

// WithModuleHost adds the moduleHost to the put assets params
func (o *PutAssetsParams) WithModuleHost(moduleHost *string) *PutAssetsParams {
	o.SetModuleHost(moduleHost)
	return o
}

// SetModuleHost adds the moduleHost to the put assets params
func (o *PutAssetsParams) SetModuleHost(moduleHost *string) {
	o.ModuleHost = moduleHost
}

// WithModuleHostID adds the moduleHostID to the put assets params
func (o *PutAssetsParams) WithModuleHostID(moduleHostID *string) *PutAssetsParams {
	o.SetModuleHostID(moduleHostID)
	return o
}

// SetModuleHostID adds the moduleHostId to the put assets params
func (o *PutAssetsParams) SetModuleHostID(moduleHostID *string) {
	o.ModuleHostID = moduleHostID
}

// WithNotes adds the notes to the put assets params
func (o *PutAssetsParams) WithNotes(notes *string) *PutAssetsParams {
	o.SetNotes(notes)
	return o
}

// SetNotes adds the notes to the put assets params
func (o *PutAssetsParams) SetNotes(notes *string) {
	o.Notes = notes
}

// WithNumberingStartFrom adds the numberingStartFrom to the put assets params
func (o *PutAssetsParams) WithNumberingStartFrom(numberingStartFrom *string) *PutAssetsParams {
	o.SetNumberingStartFrom(numberingStartFrom)
	return o
}

// SetNumberingStartFrom adds the numberingStartFrom to the put assets params
func (o *PutAssetsParams) SetNumberingStartFrom(numberingStartFrom *string) {
	o.NumberingStartFrom = numberingStartFrom
}

// WithOrientation adds the orientation to the put assets params
func (o *PutAssetsParams) WithOrientation(orientation *string) *PutAssetsParams {
	o.SetOrientation(orientation)
	return o
}

// SetOrientation adds the orientation to the put assets params
func (o *PutAssetsParams) SetOrientation(orientation *string) {
	o.Orientation = orientation
}

// WithPatchPanelModel adds the patchPanelModel to the put assets params
func (o *PutAssetsParams) WithPatchPanelModel(patchPanelModel *string) *PutAssetsParams {
	o.SetPatchPanelModel(patchPanelModel)
	return o
}

// SetPatchPanelModel adds the patchPanelModel to the put assets params
func (o *PutAssetsParams) SetPatchPanelModel(patchPanelModel *string) {
	o.PatchPanelModel = patchPanelModel
}

// WithPatchPanelModelID adds the patchPanelModelID to the put assets params
func (o *PutAssetsParams) WithPatchPanelModelID(patchPanelModelID *string) *PutAssetsParams {
	o.SetPatchPanelModelID(patchPanelModelID)
	return o
}

// SetPatchPanelModelID adds the patchPanelModelId to the put assets params
func (o *PutAssetsParams) SetPatchPanelModelID(patchPanelModelID *string) {
	o.PatchPanelModelID = patchPanelModelID
}

// WithPatchPanelModuleModel adds the patchPanelModuleModel to the put assets params
func (o *PutAssetsParams) WithPatchPanelModuleModel(patchPanelModuleModel *string) *PutAssetsParams {
	o.SetPatchPanelModuleModel(patchPanelModuleModel)
	return o
}

// SetPatchPanelModuleModel adds the patchPanelModuleModel to the put assets params
func (o *PutAssetsParams) SetPatchPanelModuleModel(patchPanelModuleModel *string) {
	o.PatchPanelModuleModel = patchPanelModuleModel
}

// WithPatchPanelModuleModelID adds the patchPanelModuleModelID to the put assets params
func (o *PutAssetsParams) WithPatchPanelModuleModelID(patchPanelModuleModelID *string) *PutAssetsParams {
	o.SetPatchPanelModuleModelID(patchPanelModuleModelID)
	return o
}

// SetPatchPanelModuleModelID adds the patchPanelModuleModelId to the put assets params
func (o *PutAssetsParams) SetPatchPanelModuleModelID(patchPanelModuleModelID *string) {
	o.PatchPanelModuleModelID = patchPanelModuleModelID
}

// WithRackID adds the rackID to the put assets params
func (o *PutAssetsParams) WithRackID(rackID *string) *PutAssetsParams {
	o.SetRackID(rackID)
	return o
}

// SetRackID adds the rackId to the put assets params
func (o *PutAssetsParams) SetRackID(rackID *string) {
	o.RackID = rackID
}

// WithRoom adds the room to the put assets params
func (o *PutAssetsParams) WithRoom(room *string) *PutAssetsParams {
	o.SetRoom(room)
	return o
}

// SetRoom adds the room to the put assets params
func (o *PutAssetsParams) SetRoom(room *string) {
	o.Room = room
}

// WithSerialNo adds the serialNo to the put assets params
func (o *PutAssetsParams) WithSerialNo(serialNo *string) *PutAssetsParams {
	o.SetSerialNo(serialNo)
	return o
}

// SetSerialNo adds the serialNo to the put assets params
func (o *PutAssetsParams) SetSerialNo(serialNo *string) {
	o.SerialNo = serialNo
}

// WithServiceLevel adds the serviceLevel to the put assets params
func (o *PutAssetsParams) WithServiceLevel(serviceLevel *string) *PutAssetsParams {
	o.SetServiceLevel(serviceLevel)
	return o
}

// SetServiceLevel adds the serviceLevel to the put assets params
func (o *PutAssetsParams) SetServiceLevel(serviceLevel *string) {
	o.ServiceLevel = serviceLevel
}

// WithSize adds the size to the put assets params
func (o *PutAssetsParams) WithSize(size *int64) *PutAssetsParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the put assets params
func (o *PutAssetsParams) SetSize(size *int64) {
	o.Size = size
}

// WithSlotNo adds the slotNo to the put assets params
func (o *PutAssetsParams) WithSlotNo(slotNo *string) *PutAssetsParams {
	o.SetSlotNo(slotNo)
	return o
}

// SetSlotNo adds the slotNo to the put assets params
func (o *PutAssetsParams) SetSlotNo(slotNo *string) {
	o.SlotNo = slotNo
}

// WithStartAt adds the startAt to the put assets params
func (o *PutAssetsParams) WithStartAt(startAt *string) *PutAssetsParams {
	o.SetStartAt(startAt)
	return o
}

// SetStartAt adds the startAt to the put assets params
func (o *PutAssetsParams) SetStartAt(startAt *string) {
	o.StartAt = startAt
}

// WithTags adds the tags to the put assets params
func (o *PutAssetsParams) WithTags(tags *string) *PutAssetsParams {
	o.SetTags(tags)
	return o
}

// SetTags adds the tags to the put assets params
func (o *PutAssetsParams) SetTags(tags *string) {
	o.Tags = tags
}

// WithTagsRemove adds the tagsRemove to the put assets params
func (o *PutAssetsParams) WithTagsRemove(tagsRemove *string) *PutAssetsParams {
	o.SetTagsRemove(tagsRemove)
	return o
}

// SetTagsRemove adds the tagsRemove to the put assets params
func (o *PutAssetsParams) SetTagsRemove(tagsRemove *string) {
	o.TagsRemove = tagsRemove
}

// WithType adds the typeVar to the put assets params
func (o *PutAssetsParams) WithType(typeVar *string) *PutAssetsParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the put assets params
func (o *PutAssetsParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WithVendor adds the vendor to the put assets params
func (o *PutAssetsParams) WithVendor(vendor *string) *PutAssetsParams {
	o.SetVendor(vendor)
	return o
}

// SetVendor adds the vendor to the put assets params
func (o *PutAssetsParams) SetVendor(vendor *string) {
	o.Vendor = vendor
}

// WithWhere adds the where to the put assets params
func (o *PutAssetsParams) WithWhere(where *string) *PutAssetsParams {
	o.SetWhere(where)
	return o
}

// SetWhere adds the where to the put assets params
func (o *PutAssetsParams) SetWhere(where *string) {
	o.Where = where
}

// WithXPos adds the xPos to the put assets params
func (o *PutAssetsParams) WithXPos(xPos *int64) *PutAssetsParams {
	o.SetXPos(xPos)
	return o
}

// SetXPos adds the xPos to the put assets params
func (o *PutAssetsParams) SetXPos(xPos *int64) {
	o.XPos = xPos
}

// WriteToRequest writes these params to a swagger request
func (o *PutAssetsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.AssetID != nil {

		// form param asset_id
		var frAssetID int64
		if o.AssetID != nil {
			frAssetID = *o.AssetID
		}
		fAssetID := swag.FormatInt64(frAssetID)
		if fAssetID != "" {
			if err := r.SetFormParam("asset_id", fAssetID); err != nil {
				return err
			}
		}

	}

	if o.AssetNo != nil {

		// form param asset_no
		var frAssetNo string
		if o.AssetNo != nil {
			frAssetNo = *o.AssetNo
		}
		fAssetNo := frAssetNo
		if fAssetNo != "" {
			if err := r.SetFormParam("asset_no", fAssetNo); err != nil {
				return err
			}
		}

	}

	if o.BackImage != nil {

		// form param back_image
		var frBackImage string
		if o.BackImage != nil {
			frBackImage = *o.BackImage
		}
		fBackImage := frBackImage
		if fBackImage != "" {
			if err := r.SetFormParam("back_image", fBackImage); err != nil {
				return err
			}
		}

	}

	if o.BackImageID != nil {

		// form param back_image_id
		var frBackImageID string
		if o.BackImageID != nil {
			frBackImageID = *o.BackImageID
		}
		fBackImageID := frBackImageID
		if fBackImageID != "" {
			if err := r.SetFormParam("back_image_id", fBackImageID); err != nil {
				return err
			}
		}

	}

	if o.Building != nil {

		// form param building
		var frBuilding string
		if o.Building != nil {
			frBuilding = *o.Building
		}
		fBuilding := frBuilding
		if fBuilding != "" {
			if err := r.SetFormParam("building", fBuilding); err != nil {
				return err
			}
		}

	}

	if o.Category != nil {

		// query param category
		var qrCategory string
		if o.Category != nil {
			qrCategory = *o.Category
		}
		qCategory := qrCategory
		if qCategory != "" {
			if err := r.SetQueryParam("category", qCategory); err != nil {
				return err
			}
		}

	}

	if o.Customer != nil {

		// form param customer
		var frCustomer string
		if o.Customer != nil {
			frCustomer = *o.Customer
		}
		fCustomer := frCustomer
		if fCustomer != "" {
			if err := r.SetFormParam("customer", fCustomer); err != nil {
				return err
			}
		}

	}

	if o.Depth != nil {

		// form param depth
		var frDepth string
		if o.Depth != nil {
			frDepth = *o.Depth
		}
		fDepth := frDepth
		if fDepth != "" {
			if err := r.SetFormParam("depth", fDepth); err != nil {
				return err
			}
		}

	}

	if o.DeviceID != nil {

		// form param device_id
		var frDeviceID int64
		if o.DeviceID != nil {
			frDeviceID = *o.DeviceID
		}
		fDeviceID := swag.FormatInt64(frDeviceID)
		if fDeviceID != "" {
			if err := r.SetFormParam("device_id", fDeviceID); err != nil {
				return err
			}
		}

	}

	if o.Imgfile != nil {

		// form param imgfile
		var frImgfile string
		if o.Imgfile != nil {
			frImgfile = *o.Imgfile
		}
		fImgfile := frImgfile
		if fImgfile != "" {
			if err := r.SetFormParam("imgfile", fImgfile); err != nil {
				return err
			}
		}

	}

	if o.ImgfileID != nil {

		// form param imgfile_id
		var frImgfileID string
		if o.ImgfileID != nil {
			frImgfileID = *o.ImgfileID
		}
		fImgfileID := frImgfileID
		if fImgfileID != "" {
			if err := r.SetFormParam("imgfile_id", fImgfileID); err != nil {
				return err
			}
		}

	}

	if o.Location != nil {

		// form param location
		var frLocation string
		if o.Location != nil {
			frLocation = *o.Location
		}
		fLocation := frLocation
		if fLocation != "" {
			if err := r.SetFormParam("location", fLocation); err != nil {
				return err
			}
		}

	}

	if o.ModuleHost != nil {

		// form param module_host
		var frModuleHost string
		if o.ModuleHost != nil {
			frModuleHost = *o.ModuleHost
		}
		fModuleHost := frModuleHost
		if fModuleHost != "" {
			if err := r.SetFormParam("module_host", fModuleHost); err != nil {
				return err
			}
		}

	}

	if o.ModuleHostID != nil {

		// form param module_host_id
		var frModuleHostID string
		if o.ModuleHostID != nil {
			frModuleHostID = *o.ModuleHostID
		}
		fModuleHostID := frModuleHostID
		if fModuleHostID != "" {
			if err := r.SetFormParam("module_host_id", fModuleHostID); err != nil {
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

	if o.NumberingStartFrom != nil {

		// form param numbering_start_from
		var frNumberingStartFrom string
		if o.NumberingStartFrom != nil {
			frNumberingStartFrom = *o.NumberingStartFrom
		}
		fNumberingStartFrom := frNumberingStartFrom
		if fNumberingStartFrom != "" {
			if err := r.SetFormParam("numbering_start_from", fNumberingStartFrom); err != nil {
				return err
			}
		}

	}

	if o.Orientation != nil {

		// form param orientation
		var frOrientation string
		if o.Orientation != nil {
			frOrientation = *o.Orientation
		}
		fOrientation := frOrientation
		if fOrientation != "" {
			if err := r.SetFormParam("orientation", fOrientation); err != nil {
				return err
			}
		}

	}

	if o.PatchPanelModel != nil {

		// form param patch_panel_model
		var frPatchPanelModel string
		if o.PatchPanelModel != nil {
			frPatchPanelModel = *o.PatchPanelModel
		}
		fPatchPanelModel := frPatchPanelModel
		if fPatchPanelModel != "" {
			if err := r.SetFormParam("patch_panel_model", fPatchPanelModel); err != nil {
				return err
			}
		}

	}

	if o.PatchPanelModelID != nil {

		// form param patch_panel_model_id
		var frPatchPanelModelID string
		if o.PatchPanelModelID != nil {
			frPatchPanelModelID = *o.PatchPanelModelID
		}
		fPatchPanelModelID := frPatchPanelModelID
		if fPatchPanelModelID != "" {
			if err := r.SetFormParam("patch_panel_model_id", fPatchPanelModelID); err != nil {
				return err
			}
		}

	}

	if o.PatchPanelModuleModel != nil {

		// form param patch_panel_module_model
		var frPatchPanelModuleModel string
		if o.PatchPanelModuleModel != nil {
			frPatchPanelModuleModel = *o.PatchPanelModuleModel
		}
		fPatchPanelModuleModel := frPatchPanelModuleModel
		if fPatchPanelModuleModel != "" {
			if err := r.SetFormParam("patch_panel_module_model", fPatchPanelModuleModel); err != nil {
				return err
			}
		}

	}

	if o.PatchPanelModuleModelID != nil {

		// form param patch_panel_module_model_id
		var frPatchPanelModuleModelID string
		if o.PatchPanelModuleModelID != nil {
			frPatchPanelModuleModelID = *o.PatchPanelModuleModelID
		}
		fPatchPanelModuleModelID := frPatchPanelModuleModelID
		if fPatchPanelModuleModelID != "" {
			if err := r.SetFormParam("patch_panel_module_model_id", fPatchPanelModuleModelID); err != nil {
				return err
			}
		}

	}

	if o.RackID != nil {

		// form param rack_id
		var frRackID string
		if o.RackID != nil {
			frRackID = *o.RackID
		}
		fRackID := frRackID
		if fRackID != "" {
			if err := r.SetFormParam("rack_id", fRackID); err != nil {
				return err
			}
		}

	}

	if o.Room != nil {

		// form param room
		var frRoom string
		if o.Room != nil {
			frRoom = *o.Room
		}
		fRoom := frRoom
		if fRoom != "" {
			if err := r.SetFormParam("room", fRoom); err != nil {
				return err
			}
		}

	}

	if o.SerialNo != nil {

		// form param serial_no
		var frSerialNo string
		if o.SerialNo != nil {
			frSerialNo = *o.SerialNo
		}
		fSerialNo := frSerialNo
		if fSerialNo != "" {
			if err := r.SetFormParam("serial_no", fSerialNo); err != nil {
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

	if o.Size != nil {

		// form param size
		var frSize int64
		if o.Size != nil {
			frSize = *o.Size
		}
		fSize := swag.FormatInt64(frSize)
		if fSize != "" {
			if err := r.SetFormParam("size", fSize); err != nil {
				return err
			}
		}

	}

	if o.SlotNo != nil {

		// form param slot_no
		var frSlotNo string
		if o.SlotNo != nil {
			frSlotNo = *o.SlotNo
		}
		fSlotNo := frSlotNo
		if fSlotNo != "" {
			if err := r.SetFormParam("slot_no", fSlotNo); err != nil {
				return err
			}
		}

	}

	if o.StartAt != nil {

		// form param start_at
		var frStartAt string
		if o.StartAt != nil {
			frStartAt = *o.StartAt
		}
		fStartAt := frStartAt
		if fStartAt != "" {
			if err := r.SetFormParam("start_at", fStartAt); err != nil {
				return err
			}
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

	if o.Where != nil {

		// form param where
		var frWhere string
		if o.Where != nil {
			frWhere = *o.Where
		}
		fWhere := frWhere
		if fWhere != "" {
			if err := r.SetFormParam("where", fWhere); err != nil {
				return err
			}
		}

	}

	if o.XPos != nil {

		// form param x_pos
		var frXPos int64
		if o.XPos != nil {
			frXPos = *o.XPos
		}
		fXPos := swag.FormatInt64(frXPos)
		if fXPos != "" {
			if err := r.SetFormParam("x_pos", fXPos); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
