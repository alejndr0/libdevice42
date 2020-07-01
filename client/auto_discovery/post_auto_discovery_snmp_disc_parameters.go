// Code generated by go-swagger; DO NOT EDIT.

package auto_discovery

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

// NewPostAutoDiscoverySnmpDiscParams creates a new PostAutoDiscoverySnmpDiscParams object
// with the default values initialized.
func NewPostAutoDiscoverySnmpDiscParams() *PostAutoDiscoverySnmpDiscParams {
	var (
		snmpVersionDefault = string("v2c")
	)
	return &PostAutoDiscoverySnmpDiscParams{
		SnmpVersion: &snmpVersionDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAutoDiscoverySnmpDiscParamsWithTimeout creates a new PostAutoDiscoverySnmpDiscParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAutoDiscoverySnmpDiscParamsWithTimeout(timeout time.Duration) *PostAutoDiscoverySnmpDiscParams {
	var (
		snmpVersionDefault = string("v2c")
	)
	return &PostAutoDiscoverySnmpDiscParams{
		SnmpVersion: &snmpVersionDefault,

		timeout: timeout,
	}
}

// NewPostAutoDiscoverySnmpDiscParamsWithContext creates a new PostAutoDiscoverySnmpDiscParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAutoDiscoverySnmpDiscParamsWithContext(ctx context.Context) *PostAutoDiscoverySnmpDiscParams {
	var (
		snmpVersionDefault = string("v2c")
	)
	return &PostAutoDiscoverySnmpDiscParams{
		SnmpVersion: &snmpVersionDefault,

		Context: ctx,
	}
}

// NewPostAutoDiscoverySnmpDiscParamsWithHTTPClient creates a new PostAutoDiscoverySnmpDiscParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAutoDiscoverySnmpDiscParamsWithHTTPClient(client *http.Client) *PostAutoDiscoverySnmpDiscParams {
	var (
		snmpVersionDefault = string("v2c")
	)
	return &PostAutoDiscoverySnmpDiscParams{
		SnmpVersion: &snmpVersionDefault,
		HTTPClient:  client,
	}
}

/*PostAutoDiscoverySnmpDiscParams contains all the parameters to send to the API endpoint
for the post auto discovery snmp disc operation typically these are written to a http.Request
*/
type PostAutoDiscoverySnmpDiscParams struct {

	/*ClearExistingSchedule*/
	ClearExistingSchedule *string
	/*DebugLevel*/
	DebugLevel *string
	/*EndIPAddress
	  End IP address

	*/
	EndIPAddress *string
	/*Groups
	  name of one or more groups separated by commas

	*/
	Groups *string
	/*Ipaddress
	  IP address. Required if new

	*/
	Ipaddress *string
	/*Name
	  name of the job

	*/
	Name string
	/*ScheduleDays
	  Comma separated days of week, where Monday = 0. e.g. 0,1,2 wil set the job for Mon, Tue and Wed. For multiple schedules, separate with a slash (/).

	*/
	ScheduleDays *string
	/*ScheduleTime
	  Time in HH:MM format if you want to schedule the job. Note: Must be formatted as text NOT date. For multiple schedules, separate with a slash (/).

	*/
	ScheduleTime *string
	/*ServiceLevel
	  Must already exist

	*/
	ServiceLevel *string
	/*SnmpPort
	  snmp port (integer only) (added in v10.4.0)

	*/
	SnmpPort *int64
	/*SnmpString
	  required, if new

	*/
	SnmpString *string
	/*SnmpStringID
	  The id of the password for the community string

	*/
	SnmpStringID *string
	/*SnmpVersion*/
	SnmpVersion *string
	/*Snmpv3AuthMode*/
	Snmpv3AuthMode *string
	/*Snmpv3AuthPassword
	  password (added in v10.4.0)

	*/
	Snmpv3AuthPassword *string
	/*Snmpv3AuthPasswordID
	  The id of the password for the auth password

	*/
	Snmpv3AuthPasswordID *string
	/*Snmpv3AuthProtocol*/
	Snmpv3AuthProtocol *string
	/*Snmpv3Context*/
	Snmpv3Context *string
	/*Snmpv3PrivacyProtocol*/
	Snmpv3PrivacyProtocol *string
	/*Snmpv3PrivacyProtocolPassword
	  password (added in v10.4.0)

	*/
	Snmpv3PrivacyProtocolPassword *string
	/*Snmpv3PrivacyProtocolPasswordID
	  The id of the password for the privacy protocol password

	*/
	Snmpv3PrivacyProtocolPasswordID *string
	/*Snmpv3User
	  name of snmp v3 user (added in v10.4.0)

	*/
	Snmpv3User *string
	/*StripDomainName*/
	StripDomainName *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post auto discovery snmp disc params
func (o *PostAutoDiscoverySnmpDiscParams) WithTimeout(timeout time.Duration) *PostAutoDiscoverySnmpDiscParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post auto discovery snmp disc params
func (o *PostAutoDiscoverySnmpDiscParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post auto discovery snmp disc params
func (o *PostAutoDiscoverySnmpDiscParams) WithContext(ctx context.Context) *PostAutoDiscoverySnmpDiscParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post auto discovery snmp disc params
func (o *PostAutoDiscoverySnmpDiscParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post auto discovery snmp disc params
func (o *PostAutoDiscoverySnmpDiscParams) WithHTTPClient(client *http.Client) *PostAutoDiscoverySnmpDiscParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post auto discovery snmp disc params
func (o *PostAutoDiscoverySnmpDiscParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClearExistingSchedule adds the clearExistingSchedule to the post auto discovery snmp disc params
func (o *PostAutoDiscoverySnmpDiscParams) WithClearExistingSchedule(clearExistingSchedule *string) *PostAutoDiscoverySnmpDiscParams {
	o.SetClearExistingSchedule(clearExistingSchedule)
	return o
}

// SetClearExistingSchedule adds the clearExistingSchedule to the post auto discovery snmp disc params
func (o *PostAutoDiscoverySnmpDiscParams) SetClearExistingSchedule(clearExistingSchedule *string) {
	o.ClearExistingSchedule = clearExistingSchedule
}

// WithDebugLevel adds the debugLevel to the post auto discovery snmp disc params
func (o *PostAutoDiscoverySnmpDiscParams) WithDebugLevel(debugLevel *string) *PostAutoDiscoverySnmpDiscParams {
	o.SetDebugLevel(debugLevel)
	return o
}

// SetDebugLevel adds the debugLevel to the post auto discovery snmp disc params
func (o *PostAutoDiscoverySnmpDiscParams) SetDebugLevel(debugLevel *string) {
	o.DebugLevel = debugLevel
}

// WithEndIPAddress adds the endIPAddress to the post auto discovery snmp disc params
func (o *PostAutoDiscoverySnmpDiscParams) WithEndIPAddress(endIPAddress *string) *PostAutoDiscoverySnmpDiscParams {
	o.SetEndIPAddress(endIPAddress)
	return o
}

// SetEndIPAddress adds the endIpAddress to the post auto discovery snmp disc params
func (o *PostAutoDiscoverySnmpDiscParams) SetEndIPAddress(endIPAddress *string) {
	o.EndIPAddress = endIPAddress
}

// WithGroups adds the groups to the post auto discovery snmp disc params
func (o *PostAutoDiscoverySnmpDiscParams) WithGroups(groups *string) *PostAutoDiscoverySnmpDiscParams {
	o.SetGroups(groups)
	return o
}

// SetGroups adds the groups to the post auto discovery snmp disc params
func (o *PostAutoDiscoverySnmpDiscParams) SetGroups(groups *string) {
	o.Groups = groups
}

// WithIpaddress adds the ipaddress to the post auto discovery snmp disc params
func (o *PostAutoDiscoverySnmpDiscParams) WithIpaddress(ipaddress *string) *PostAutoDiscoverySnmpDiscParams {
	o.SetIpaddress(ipaddress)
	return o
}

// SetIpaddress adds the ipaddress to the post auto discovery snmp disc params
func (o *PostAutoDiscoverySnmpDiscParams) SetIpaddress(ipaddress *string) {
	o.Ipaddress = ipaddress
}

// WithName adds the name to the post auto discovery snmp disc params
func (o *PostAutoDiscoverySnmpDiscParams) WithName(name string) *PostAutoDiscoverySnmpDiscParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the post auto discovery snmp disc params
func (o *PostAutoDiscoverySnmpDiscParams) SetName(name string) {
	o.Name = name
}

// WithScheduleDays adds the scheduleDays to the post auto discovery snmp disc params
func (o *PostAutoDiscoverySnmpDiscParams) WithScheduleDays(scheduleDays *string) *PostAutoDiscoverySnmpDiscParams {
	o.SetScheduleDays(scheduleDays)
	return o
}

// SetScheduleDays adds the scheduleDays to the post auto discovery snmp disc params
func (o *PostAutoDiscoverySnmpDiscParams) SetScheduleDays(scheduleDays *string) {
	o.ScheduleDays = scheduleDays
}

// WithScheduleTime adds the scheduleTime to the post auto discovery snmp disc params
func (o *PostAutoDiscoverySnmpDiscParams) WithScheduleTime(scheduleTime *string) *PostAutoDiscoverySnmpDiscParams {
	o.SetScheduleTime(scheduleTime)
	return o
}

// SetScheduleTime adds the scheduleTime to the post auto discovery snmp disc params
func (o *PostAutoDiscoverySnmpDiscParams) SetScheduleTime(scheduleTime *string) {
	o.ScheduleTime = scheduleTime
}

// WithServiceLevel adds the serviceLevel to the post auto discovery snmp disc params
func (o *PostAutoDiscoverySnmpDiscParams) WithServiceLevel(serviceLevel *string) *PostAutoDiscoverySnmpDiscParams {
	o.SetServiceLevel(serviceLevel)
	return o
}

// SetServiceLevel adds the serviceLevel to the post auto discovery snmp disc params
func (o *PostAutoDiscoverySnmpDiscParams) SetServiceLevel(serviceLevel *string) {
	o.ServiceLevel = serviceLevel
}

// WithSnmpPort adds the snmpPort to the post auto discovery snmp disc params
func (o *PostAutoDiscoverySnmpDiscParams) WithSnmpPort(snmpPort *int64) *PostAutoDiscoverySnmpDiscParams {
	o.SetSnmpPort(snmpPort)
	return o
}

// SetSnmpPort adds the snmpPort to the post auto discovery snmp disc params
func (o *PostAutoDiscoverySnmpDiscParams) SetSnmpPort(snmpPort *int64) {
	o.SnmpPort = snmpPort
}

// WithSnmpString adds the snmpString to the post auto discovery snmp disc params
func (o *PostAutoDiscoverySnmpDiscParams) WithSnmpString(snmpString *string) *PostAutoDiscoverySnmpDiscParams {
	o.SetSnmpString(snmpString)
	return o
}

// SetSnmpString adds the snmpString to the post auto discovery snmp disc params
func (o *PostAutoDiscoverySnmpDiscParams) SetSnmpString(snmpString *string) {
	o.SnmpString = snmpString
}

// WithSnmpStringID adds the snmpStringID to the post auto discovery snmp disc params
func (o *PostAutoDiscoverySnmpDiscParams) WithSnmpStringID(snmpStringID *string) *PostAutoDiscoverySnmpDiscParams {
	o.SetSnmpStringID(snmpStringID)
	return o
}

// SetSnmpStringID adds the snmpStringId to the post auto discovery snmp disc params
func (o *PostAutoDiscoverySnmpDiscParams) SetSnmpStringID(snmpStringID *string) {
	o.SnmpStringID = snmpStringID
}

// WithSnmpVersion adds the snmpVersion to the post auto discovery snmp disc params
func (o *PostAutoDiscoverySnmpDiscParams) WithSnmpVersion(snmpVersion *string) *PostAutoDiscoverySnmpDiscParams {
	o.SetSnmpVersion(snmpVersion)
	return o
}

// SetSnmpVersion adds the snmpVersion to the post auto discovery snmp disc params
func (o *PostAutoDiscoverySnmpDiscParams) SetSnmpVersion(snmpVersion *string) {
	o.SnmpVersion = snmpVersion
}

// WithSnmpv3AuthMode adds the snmpv3AuthMode to the post auto discovery snmp disc params
func (o *PostAutoDiscoverySnmpDiscParams) WithSnmpv3AuthMode(snmpv3AuthMode *string) *PostAutoDiscoverySnmpDiscParams {
	o.SetSnmpv3AuthMode(snmpv3AuthMode)
	return o
}

// SetSnmpv3AuthMode adds the snmpv3AuthMode to the post auto discovery snmp disc params
func (o *PostAutoDiscoverySnmpDiscParams) SetSnmpv3AuthMode(snmpv3AuthMode *string) {
	o.Snmpv3AuthMode = snmpv3AuthMode
}

// WithSnmpv3AuthPassword adds the snmpv3AuthPassword to the post auto discovery snmp disc params
func (o *PostAutoDiscoverySnmpDiscParams) WithSnmpv3AuthPassword(snmpv3AuthPassword *string) *PostAutoDiscoverySnmpDiscParams {
	o.SetSnmpv3AuthPassword(snmpv3AuthPassword)
	return o
}

// SetSnmpv3AuthPassword adds the snmpv3AuthPassword to the post auto discovery snmp disc params
func (o *PostAutoDiscoverySnmpDiscParams) SetSnmpv3AuthPassword(snmpv3AuthPassword *string) {
	o.Snmpv3AuthPassword = snmpv3AuthPassword
}

// WithSnmpv3AuthPasswordID adds the snmpv3AuthPasswordID to the post auto discovery snmp disc params
func (o *PostAutoDiscoverySnmpDiscParams) WithSnmpv3AuthPasswordID(snmpv3AuthPasswordID *string) *PostAutoDiscoverySnmpDiscParams {
	o.SetSnmpv3AuthPasswordID(snmpv3AuthPasswordID)
	return o
}

// SetSnmpv3AuthPasswordID adds the snmpv3AuthPasswordId to the post auto discovery snmp disc params
func (o *PostAutoDiscoverySnmpDiscParams) SetSnmpv3AuthPasswordID(snmpv3AuthPasswordID *string) {
	o.Snmpv3AuthPasswordID = snmpv3AuthPasswordID
}

// WithSnmpv3AuthProtocol adds the snmpv3AuthProtocol to the post auto discovery snmp disc params
func (o *PostAutoDiscoverySnmpDiscParams) WithSnmpv3AuthProtocol(snmpv3AuthProtocol *string) *PostAutoDiscoverySnmpDiscParams {
	o.SetSnmpv3AuthProtocol(snmpv3AuthProtocol)
	return o
}

// SetSnmpv3AuthProtocol adds the snmpv3AuthProtocol to the post auto discovery snmp disc params
func (o *PostAutoDiscoverySnmpDiscParams) SetSnmpv3AuthProtocol(snmpv3AuthProtocol *string) {
	o.Snmpv3AuthProtocol = snmpv3AuthProtocol
}

// WithSnmpv3Context adds the snmpv3Context to the post auto discovery snmp disc params
func (o *PostAutoDiscoverySnmpDiscParams) WithSnmpv3Context(snmpv3Context *string) *PostAutoDiscoverySnmpDiscParams {
	o.SetSnmpv3Context(snmpv3Context)
	return o
}

// SetSnmpv3Context adds the snmpv3Context to the post auto discovery snmp disc params
func (o *PostAutoDiscoverySnmpDiscParams) SetSnmpv3Context(snmpv3Context *string) {
	o.Snmpv3Context = snmpv3Context
}

// WithSnmpv3PrivacyProtocol adds the snmpv3PrivacyProtocol to the post auto discovery snmp disc params
func (o *PostAutoDiscoverySnmpDiscParams) WithSnmpv3PrivacyProtocol(snmpv3PrivacyProtocol *string) *PostAutoDiscoverySnmpDiscParams {
	o.SetSnmpv3PrivacyProtocol(snmpv3PrivacyProtocol)
	return o
}

// SetSnmpv3PrivacyProtocol adds the snmpv3PrivacyProtocol to the post auto discovery snmp disc params
func (o *PostAutoDiscoverySnmpDiscParams) SetSnmpv3PrivacyProtocol(snmpv3PrivacyProtocol *string) {
	o.Snmpv3PrivacyProtocol = snmpv3PrivacyProtocol
}

// WithSnmpv3PrivacyProtocolPassword adds the snmpv3PrivacyProtocolPassword to the post auto discovery snmp disc params
func (o *PostAutoDiscoverySnmpDiscParams) WithSnmpv3PrivacyProtocolPassword(snmpv3PrivacyProtocolPassword *string) *PostAutoDiscoverySnmpDiscParams {
	o.SetSnmpv3PrivacyProtocolPassword(snmpv3PrivacyProtocolPassword)
	return o
}

// SetSnmpv3PrivacyProtocolPassword adds the snmpv3PrivacyProtocolPassword to the post auto discovery snmp disc params
func (o *PostAutoDiscoverySnmpDiscParams) SetSnmpv3PrivacyProtocolPassword(snmpv3PrivacyProtocolPassword *string) {
	o.Snmpv3PrivacyProtocolPassword = snmpv3PrivacyProtocolPassword
}

// WithSnmpv3PrivacyProtocolPasswordID adds the snmpv3PrivacyProtocolPasswordID to the post auto discovery snmp disc params
func (o *PostAutoDiscoverySnmpDiscParams) WithSnmpv3PrivacyProtocolPasswordID(snmpv3PrivacyProtocolPasswordID *string) *PostAutoDiscoverySnmpDiscParams {
	o.SetSnmpv3PrivacyProtocolPasswordID(snmpv3PrivacyProtocolPasswordID)
	return o
}

// SetSnmpv3PrivacyProtocolPasswordID adds the snmpv3PrivacyProtocolPasswordId to the post auto discovery snmp disc params
func (o *PostAutoDiscoverySnmpDiscParams) SetSnmpv3PrivacyProtocolPasswordID(snmpv3PrivacyProtocolPasswordID *string) {
	o.Snmpv3PrivacyProtocolPasswordID = snmpv3PrivacyProtocolPasswordID
}

// WithSnmpv3User adds the snmpv3User to the post auto discovery snmp disc params
func (o *PostAutoDiscoverySnmpDiscParams) WithSnmpv3User(snmpv3User *string) *PostAutoDiscoverySnmpDiscParams {
	o.SetSnmpv3User(snmpv3User)
	return o
}

// SetSnmpv3User adds the snmpv3User to the post auto discovery snmp disc params
func (o *PostAutoDiscoverySnmpDiscParams) SetSnmpv3User(snmpv3User *string) {
	o.Snmpv3User = snmpv3User
}

// WithStripDomainName adds the stripDomainName to the post auto discovery snmp disc params
func (o *PostAutoDiscoverySnmpDiscParams) WithStripDomainName(stripDomainName *string) *PostAutoDiscoverySnmpDiscParams {
	o.SetStripDomainName(stripDomainName)
	return o
}

// SetStripDomainName adds the stripDomainName to the post auto discovery snmp disc params
func (o *PostAutoDiscoverySnmpDiscParams) SetStripDomainName(stripDomainName *string) {
	o.StripDomainName = stripDomainName
}

// WriteToRequest writes these params to a swagger request
func (o *PostAutoDiscoverySnmpDiscParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ClearExistingSchedule != nil {

		// form param clear_existing_schedule
		var frClearExistingSchedule string
		if o.ClearExistingSchedule != nil {
			frClearExistingSchedule = *o.ClearExistingSchedule
		}
		fClearExistingSchedule := frClearExistingSchedule
		if fClearExistingSchedule != "" {
			if err := r.SetFormParam("clear_existing_schedule", fClearExistingSchedule); err != nil {
				return err
			}
		}

	}

	if o.DebugLevel != nil {

		// form param debug_level
		var frDebugLevel string
		if o.DebugLevel != nil {
			frDebugLevel = *o.DebugLevel
		}
		fDebugLevel := frDebugLevel
		if fDebugLevel != "" {
			if err := r.SetFormParam("debug_level", fDebugLevel); err != nil {
				return err
			}
		}

	}

	if o.EndIPAddress != nil {

		// form param end_ip_address
		var frEndIPAddress string
		if o.EndIPAddress != nil {
			frEndIPAddress = *o.EndIPAddress
		}
		fEndIPAddress := frEndIPAddress
		if fEndIPAddress != "" {
			if err := r.SetFormParam("end_ip_address", fEndIPAddress); err != nil {
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

	if o.Ipaddress != nil {

		// form param ipaddress
		var frIpaddress string
		if o.Ipaddress != nil {
			frIpaddress = *o.Ipaddress
		}
		fIpaddress := frIpaddress
		if fIpaddress != "" {
			if err := r.SetFormParam("ipaddress", fIpaddress); err != nil {
				return err
			}
		}

	}

	// form param name
	frName := o.Name
	fName := frName
	if fName != "" {
		if err := r.SetFormParam("name", fName); err != nil {
			return err
		}
	}

	if o.ScheduleDays != nil {

		// form param schedule_days
		var frScheduleDays string
		if o.ScheduleDays != nil {
			frScheduleDays = *o.ScheduleDays
		}
		fScheduleDays := frScheduleDays
		if fScheduleDays != "" {
			if err := r.SetFormParam("schedule_days", fScheduleDays); err != nil {
				return err
			}
		}

	}

	if o.ScheduleTime != nil {

		// form param schedule_time
		var frScheduleTime string
		if o.ScheduleTime != nil {
			frScheduleTime = *o.ScheduleTime
		}
		fScheduleTime := frScheduleTime
		if fScheduleTime != "" {
			if err := r.SetFormParam("schedule_time", fScheduleTime); err != nil {
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

	if o.SnmpPort != nil {

		// form param snmp_port
		var frSnmpPort int64
		if o.SnmpPort != nil {
			frSnmpPort = *o.SnmpPort
		}
		fSnmpPort := swag.FormatInt64(frSnmpPort)
		if fSnmpPort != "" {
			if err := r.SetFormParam("snmp_port", fSnmpPort); err != nil {
				return err
			}
		}

	}

	if o.SnmpString != nil {

		// form param snmp_string
		var frSnmpString string
		if o.SnmpString != nil {
			frSnmpString = *o.SnmpString
		}
		fSnmpString := frSnmpString
		if fSnmpString != "" {
			if err := r.SetFormParam("snmp_string", fSnmpString); err != nil {
				return err
			}
		}

	}

	if o.SnmpStringID != nil {

		// form param snmp_string_id
		var frSnmpStringID string
		if o.SnmpStringID != nil {
			frSnmpStringID = *o.SnmpStringID
		}
		fSnmpStringID := frSnmpStringID
		if fSnmpStringID != "" {
			if err := r.SetFormParam("snmp_string_id", fSnmpStringID); err != nil {
				return err
			}
		}

	}

	if o.SnmpVersion != nil {

		// form param snmp_version
		var frSnmpVersion string
		if o.SnmpVersion != nil {
			frSnmpVersion = *o.SnmpVersion
		}
		fSnmpVersion := frSnmpVersion
		if fSnmpVersion != "" {
			if err := r.SetFormParam("snmp_version", fSnmpVersion); err != nil {
				return err
			}
		}

	}

	if o.Snmpv3AuthMode != nil {

		// form param snmpv3_auth_mode
		var frSnmpv3AuthMode string
		if o.Snmpv3AuthMode != nil {
			frSnmpv3AuthMode = *o.Snmpv3AuthMode
		}
		fSnmpv3AuthMode := frSnmpv3AuthMode
		if fSnmpv3AuthMode != "" {
			if err := r.SetFormParam("snmpv3_auth_mode", fSnmpv3AuthMode); err != nil {
				return err
			}
		}

	}

	if o.Snmpv3AuthPassword != nil {

		// form param snmpv3_auth_password
		var frSnmpv3AuthPassword string
		if o.Snmpv3AuthPassword != nil {
			frSnmpv3AuthPassword = *o.Snmpv3AuthPassword
		}
		fSnmpv3AuthPassword := frSnmpv3AuthPassword
		if fSnmpv3AuthPassword != "" {
			if err := r.SetFormParam("snmpv3_auth_password", fSnmpv3AuthPassword); err != nil {
				return err
			}
		}

	}

	if o.Snmpv3AuthPasswordID != nil {

		// form param snmpv3_auth_password_id
		var frSnmpv3AuthPasswordID string
		if o.Snmpv3AuthPasswordID != nil {
			frSnmpv3AuthPasswordID = *o.Snmpv3AuthPasswordID
		}
		fSnmpv3AuthPasswordID := frSnmpv3AuthPasswordID
		if fSnmpv3AuthPasswordID != "" {
			if err := r.SetFormParam("snmpv3_auth_password_id", fSnmpv3AuthPasswordID); err != nil {
				return err
			}
		}

	}

	if o.Snmpv3AuthProtocol != nil {

		// form param snmpv3_auth_protocol
		var frSnmpv3AuthProtocol string
		if o.Snmpv3AuthProtocol != nil {
			frSnmpv3AuthProtocol = *o.Snmpv3AuthProtocol
		}
		fSnmpv3AuthProtocol := frSnmpv3AuthProtocol
		if fSnmpv3AuthProtocol != "" {
			if err := r.SetFormParam("snmpv3_auth_protocol", fSnmpv3AuthProtocol); err != nil {
				return err
			}
		}

	}

	if o.Snmpv3Context != nil {

		// form param snmpv3_context
		var frSnmpv3Context string
		if o.Snmpv3Context != nil {
			frSnmpv3Context = *o.Snmpv3Context
		}
		fSnmpv3Context := frSnmpv3Context
		if fSnmpv3Context != "" {
			if err := r.SetFormParam("snmpv3_context", fSnmpv3Context); err != nil {
				return err
			}
		}

	}

	if o.Snmpv3PrivacyProtocol != nil {

		// form param snmpv3_privacy_protocol
		var frSnmpv3PrivacyProtocol string
		if o.Snmpv3PrivacyProtocol != nil {
			frSnmpv3PrivacyProtocol = *o.Snmpv3PrivacyProtocol
		}
		fSnmpv3PrivacyProtocol := frSnmpv3PrivacyProtocol
		if fSnmpv3PrivacyProtocol != "" {
			if err := r.SetFormParam("snmpv3_privacy_protocol", fSnmpv3PrivacyProtocol); err != nil {
				return err
			}
		}

	}

	if o.Snmpv3PrivacyProtocolPassword != nil {

		// form param snmpv3_privacy_protocol_password
		var frSnmpv3PrivacyProtocolPassword string
		if o.Snmpv3PrivacyProtocolPassword != nil {
			frSnmpv3PrivacyProtocolPassword = *o.Snmpv3PrivacyProtocolPassword
		}
		fSnmpv3PrivacyProtocolPassword := frSnmpv3PrivacyProtocolPassword
		if fSnmpv3PrivacyProtocolPassword != "" {
			if err := r.SetFormParam("snmpv3_privacy_protocol_password", fSnmpv3PrivacyProtocolPassword); err != nil {
				return err
			}
		}

	}

	if o.Snmpv3PrivacyProtocolPasswordID != nil {

		// form param snmpv3_privacy_protocol_password_id
		var frSnmpv3PrivacyProtocolPasswordID string
		if o.Snmpv3PrivacyProtocolPasswordID != nil {
			frSnmpv3PrivacyProtocolPasswordID = *o.Snmpv3PrivacyProtocolPasswordID
		}
		fSnmpv3PrivacyProtocolPasswordID := frSnmpv3PrivacyProtocolPasswordID
		if fSnmpv3PrivacyProtocolPasswordID != "" {
			if err := r.SetFormParam("snmpv3_privacy_protocol_password_id", fSnmpv3PrivacyProtocolPasswordID); err != nil {
				return err
			}
		}

	}

	if o.Snmpv3User != nil {

		// form param snmpv3_user
		var frSnmpv3User string
		if o.Snmpv3User != nil {
			frSnmpv3User = *o.Snmpv3User
		}
		fSnmpv3User := frSnmpv3User
		if fSnmpv3User != "" {
			if err := r.SetFormParam("snmpv3_user", fSnmpv3User); err != nil {
				return err
			}
		}

	}

	if o.StripDomainName != nil {

		// form param strip_domain_name
		var frStripDomainName string
		if o.StripDomainName != nil {
			frStripDomainName = *o.StripDomainName
		}
		fStripDomainName := frStripDomainName
		if fStripDomainName != "" {
			if err := r.SetFormParam("strip_domain_name", fStripDomainName); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}