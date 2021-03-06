// Code generated by go-swagger; DO NOT EDIT.

package racks

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

// NewPostRacksParams creates a new PostRacksParams object
// with the default values initialized.
func NewPostRacksParams() *PostRacksParams {
	var ()
	return &PostRacksParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRacksParamsWithTimeout creates a new PostRacksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRacksParamsWithTimeout(timeout time.Duration) *PostRacksParams {
	var ()
	return &PostRacksParams{

		timeout: timeout,
	}
}

// NewPostRacksParamsWithContext creates a new PostRacksParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRacksParamsWithContext(ctx context.Context) *PostRacksParams {
	var ()
	return &PostRacksParams{

		Context: ctx,
	}
}

// NewPostRacksParamsWithHTTPClient creates a new PostRacksParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostRacksParamsWithHTTPClient(client *http.Client) *PostRacksParams {
	var ()
	return &PostRacksParams{
		HTTPClient: client,
	}
}

/*PostRacksParams contains all the parameters to send to the API endpoint
for the post racks operation typically these are written to a http.Request
*/
type PostRacksParams struct {

	/*Building
	  Name of building - Used when there are non-unique room names.

	*/
	Building *string
	/*ColSize
	  how many racks wide the rack is

	*/
	ColSize *string
	/*FirstNumber
	  default 0, add to change.

	*/
	FirstNumber *string
	/*Groups
	  If multitenancy is on, admin groups that have access to this object are specified here, e.g. Prod_East:no,Corp:yes specifies that the admin groups for this object are Prod_East with view only permission and Corp with change permission. If this parameter is present with no value, all groups are deleted.

	*/
	Groups *string
	/*Manufacturer
	  name of the hardware/software manufacturer.

	*/
	Manufacturer *string
	/*Name
	  Rack name - must be unique within a room

	*/
	Name *string
	/*NewName
	  Use to change name of object.

	*/
	NewName *string
	/*Notes
	  Any additional notes

	*/
	Notes *string
	/*NumberingStartFromBottom
	  default is yes, no to change, otherwise ignored.

	*/
	NumberingStartFromBottom *string
	/*Orientation
	  orientation of the rack in room layout view. Possible values: right, left, up or down.

	*/
	Orientation *string
	/*RackID
	  To change a rack using ID. This has highest priority to change a rack.

	*/
	RackID *int64
	/*Room
	  Name of room - Required if changing a rack without rack_id.

	*/
	Room *string
	/*RoomID
	  Room ID if Room name is not unique

	*/
	RoomID *string
	/*Row
	  this row field is for the name of the rows, and not related to the grid positioning of the rack

	*/
	Row *string
	/*RowSize
	  How many rows long the rack is

	*/
	RowSize *string
	/*Size
	  In U

	*/
	Size *int64
	/*StartCol
	  Starting column for the rack, for grid positioning

	*/
	StartCol *string
	/*StartRow
	  Starting row for rack, for grid positioning

	*/
	StartRow *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post racks params
func (o *PostRacksParams) WithTimeout(timeout time.Duration) *PostRacksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post racks params
func (o *PostRacksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post racks params
func (o *PostRacksParams) WithContext(ctx context.Context) *PostRacksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post racks params
func (o *PostRacksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post racks params
func (o *PostRacksParams) WithHTTPClient(client *http.Client) *PostRacksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post racks params
func (o *PostRacksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBuilding adds the building to the post racks params
func (o *PostRacksParams) WithBuilding(building *string) *PostRacksParams {
	o.SetBuilding(building)
	return o
}

// SetBuilding adds the building to the post racks params
func (o *PostRacksParams) SetBuilding(building *string) {
	o.Building = building
}

// WithColSize adds the colSize to the post racks params
func (o *PostRacksParams) WithColSize(colSize *string) *PostRacksParams {
	o.SetColSize(colSize)
	return o
}

// SetColSize adds the colSize to the post racks params
func (o *PostRacksParams) SetColSize(colSize *string) {
	o.ColSize = colSize
}

// WithFirstNumber adds the firstNumber to the post racks params
func (o *PostRacksParams) WithFirstNumber(firstNumber *string) *PostRacksParams {
	o.SetFirstNumber(firstNumber)
	return o
}

// SetFirstNumber adds the firstNumber to the post racks params
func (o *PostRacksParams) SetFirstNumber(firstNumber *string) {
	o.FirstNumber = firstNumber
}

// WithGroups adds the groups to the post racks params
func (o *PostRacksParams) WithGroups(groups *string) *PostRacksParams {
	o.SetGroups(groups)
	return o
}

// SetGroups adds the groups to the post racks params
func (o *PostRacksParams) SetGroups(groups *string) {
	o.Groups = groups
}

// WithManufacturer adds the manufacturer to the post racks params
func (o *PostRacksParams) WithManufacturer(manufacturer *string) *PostRacksParams {
	o.SetManufacturer(manufacturer)
	return o
}

// SetManufacturer adds the manufacturer to the post racks params
func (o *PostRacksParams) SetManufacturer(manufacturer *string) {
	o.Manufacturer = manufacturer
}

// WithName adds the name to the post racks params
func (o *PostRacksParams) WithName(name *string) *PostRacksParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the post racks params
func (o *PostRacksParams) SetName(name *string) {
	o.Name = name
}

// WithNewName adds the newName to the post racks params
func (o *PostRacksParams) WithNewName(newName *string) *PostRacksParams {
	o.SetNewName(newName)
	return o
}

// SetNewName adds the newName to the post racks params
func (o *PostRacksParams) SetNewName(newName *string) {
	o.NewName = newName
}

// WithNotes adds the notes to the post racks params
func (o *PostRacksParams) WithNotes(notes *string) *PostRacksParams {
	o.SetNotes(notes)
	return o
}

// SetNotes adds the notes to the post racks params
func (o *PostRacksParams) SetNotes(notes *string) {
	o.Notes = notes
}

// WithNumberingStartFromBottom adds the numberingStartFromBottom to the post racks params
func (o *PostRacksParams) WithNumberingStartFromBottom(numberingStartFromBottom *string) *PostRacksParams {
	o.SetNumberingStartFromBottom(numberingStartFromBottom)
	return o
}

// SetNumberingStartFromBottom adds the numberingStartFromBottom to the post racks params
func (o *PostRacksParams) SetNumberingStartFromBottom(numberingStartFromBottom *string) {
	o.NumberingStartFromBottom = numberingStartFromBottom
}

// WithOrientation adds the orientation to the post racks params
func (o *PostRacksParams) WithOrientation(orientation *string) *PostRacksParams {
	o.SetOrientation(orientation)
	return o
}

// SetOrientation adds the orientation to the post racks params
func (o *PostRacksParams) SetOrientation(orientation *string) {
	o.Orientation = orientation
}

// WithRackID adds the rackID to the post racks params
func (o *PostRacksParams) WithRackID(rackID *int64) *PostRacksParams {
	o.SetRackID(rackID)
	return o
}

// SetRackID adds the rackId to the post racks params
func (o *PostRacksParams) SetRackID(rackID *int64) {
	o.RackID = rackID
}

// WithRoom adds the room to the post racks params
func (o *PostRacksParams) WithRoom(room *string) *PostRacksParams {
	o.SetRoom(room)
	return o
}

// SetRoom adds the room to the post racks params
func (o *PostRacksParams) SetRoom(room *string) {
	o.Room = room
}

// WithRoomID adds the roomID to the post racks params
func (o *PostRacksParams) WithRoomID(roomID *string) *PostRacksParams {
	o.SetRoomID(roomID)
	return o
}

// SetRoomID adds the roomId to the post racks params
func (o *PostRacksParams) SetRoomID(roomID *string) {
	o.RoomID = roomID
}

// WithRow adds the row to the post racks params
func (o *PostRacksParams) WithRow(row *string) *PostRacksParams {
	o.SetRow(row)
	return o
}

// SetRow adds the row to the post racks params
func (o *PostRacksParams) SetRow(row *string) {
	o.Row = row
}

// WithRowSize adds the rowSize to the post racks params
func (o *PostRacksParams) WithRowSize(rowSize *string) *PostRacksParams {
	o.SetRowSize(rowSize)
	return o
}

// SetRowSize adds the rowSize to the post racks params
func (o *PostRacksParams) SetRowSize(rowSize *string) {
	o.RowSize = rowSize
}

// WithSize adds the size to the post racks params
func (o *PostRacksParams) WithSize(size *int64) *PostRacksParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the post racks params
func (o *PostRacksParams) SetSize(size *int64) {
	o.Size = size
}

// WithStartCol adds the startCol to the post racks params
func (o *PostRacksParams) WithStartCol(startCol *string) *PostRacksParams {
	o.SetStartCol(startCol)
	return o
}

// SetStartCol adds the startCol to the post racks params
func (o *PostRacksParams) SetStartCol(startCol *string) {
	o.StartCol = startCol
}

// WithStartRow adds the startRow to the post racks params
func (o *PostRacksParams) WithStartRow(startRow *string) *PostRacksParams {
	o.SetStartRow(startRow)
	return o
}

// SetStartRow adds the startRow to the post racks params
func (o *PostRacksParams) SetStartRow(startRow *string) {
	o.StartRow = startRow
}

// WriteToRequest writes these params to a swagger request
func (o *PostRacksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.ColSize != nil {

		// form param col_size
		var frColSize string
		if o.ColSize != nil {
			frColSize = *o.ColSize
		}
		fColSize := frColSize
		if fColSize != "" {
			if err := r.SetFormParam("col_size", fColSize); err != nil {
				return err
			}
		}

	}

	if o.FirstNumber != nil {

		// form param first_number
		var frFirstNumber string
		if o.FirstNumber != nil {
			frFirstNumber = *o.FirstNumber
		}
		fFirstNumber := frFirstNumber
		if fFirstNumber != "" {
			if err := r.SetFormParam("first_number", fFirstNumber); err != nil {
				return err
			}
		}

	}

	if o.Groups != nil {

		// form param groups
		var frGroups string
		if o.Groups != nil {
			frGroups = *o.Groups
		}
		fGroups := frGroups
		if fGroups != "" {
			if err := r.SetFormParam("groups", fGroups); err != nil {
				return err
			}
		}

	}

	if o.Manufacturer != nil {

		// form param manufacturer
		var frManufacturer string
		if o.Manufacturer != nil {
			frManufacturer = *o.Manufacturer
		}
		fManufacturer := frManufacturer
		if fManufacturer != "" {
			if err := r.SetFormParam("manufacturer", fManufacturer); err != nil {
				return err
			}
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

	if o.NewName != nil {

		// form param new_name
		var frNewName string
		if o.NewName != nil {
			frNewName = *o.NewName
		}
		fNewName := frNewName
		if fNewName != "" {
			if err := r.SetFormParam("new_name", fNewName); err != nil {
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

	if o.NumberingStartFromBottom != nil {

		// form param numbering_start_from_bottom
		var frNumberingStartFromBottom string
		if o.NumberingStartFromBottom != nil {
			frNumberingStartFromBottom = *o.NumberingStartFromBottom
		}
		fNumberingStartFromBottom := frNumberingStartFromBottom
		if fNumberingStartFromBottom != "" {
			if err := r.SetFormParam("numbering_start_from_bottom", fNumberingStartFromBottom); err != nil {
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

	if o.RackID != nil {

		// form param rack_id
		var frRackID int64
		if o.RackID != nil {
			frRackID = *o.RackID
		}
		fRackID := swag.FormatInt64(frRackID)
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

	if o.RoomID != nil {

		// form param room_id
		var frRoomID string
		if o.RoomID != nil {
			frRoomID = *o.RoomID
		}
		fRoomID := frRoomID
		if fRoomID != "" {
			if err := r.SetFormParam("room_id", fRoomID); err != nil {
				return err
			}
		}

	}

	if o.Row != nil {

		// form param row
		var frRow string
		if o.Row != nil {
			frRow = *o.Row
		}
		fRow := frRow
		if fRow != "" {
			if err := r.SetFormParam("row", fRow); err != nil {
				return err
			}
		}

	}

	if o.RowSize != nil {

		// form param row_size
		var frRowSize string
		if o.RowSize != nil {
			frRowSize = *o.RowSize
		}
		fRowSize := frRowSize
		if fRowSize != "" {
			if err := r.SetFormParam("row_size", fRowSize); err != nil {
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

	if o.StartCol != nil {

		// form param start_col
		var frStartCol string
		if o.StartCol != nil {
			frStartCol = *o.StartCol
		}
		fStartCol := frStartCol
		if fStartCol != "" {
			if err := r.SetFormParam("start_col", fStartCol); err != nil {
				return err
			}
		}

	}

	if o.StartRow != nil {

		// form param start_row
		var frStartRow string
		if o.StartRow != nil {
			frStartRow = *o.StartRow
		}
		fStartRow := frStartRow
		if fStartRow != "" {
			if err := r.SetFormParam("start_row", fStartRow); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
