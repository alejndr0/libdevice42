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

// NewPostUpdateServicePortsParams creates a new PostUpdateServicePortsParams object
// with the default values initialized.
func NewPostUpdateServicePortsParams() *PostUpdateServicePortsParams {
	var ()
	return &PostUpdateServicePortsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostUpdateServicePortsParamsWithTimeout creates a new PostUpdateServicePortsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostUpdateServicePortsParamsWithTimeout(timeout time.Duration) *PostUpdateServicePortsParams {
	var ()
	return &PostUpdateServicePortsParams{

		timeout: timeout,
	}
}

// NewPostUpdateServicePortsParamsWithContext creates a new PostUpdateServicePortsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostUpdateServicePortsParamsWithContext(ctx context.Context) *PostUpdateServicePortsParams {
	var ()
	return &PostUpdateServicePortsParams{

		Context: ctx,
	}
}

// NewPostUpdateServicePortsParamsWithHTTPClient creates a new PostUpdateServicePortsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostUpdateServicePortsParamsWithHTTPClient(client *http.Client) *PostUpdateServicePortsParams {
	var ()
	return &PostUpdateServicePortsParams{
		HTTPClient: client,
	}
}

/*PostUpdateServicePortsParams contains all the parameters to send to the API endpoint
for the post update service ports operation typically these are written to a http.Request
*/
type PostUpdateServicePortsParams struct {

	/*CountInLicensing
	  Whether or not to count OS in licensing

	*/
	CountInLicensing *string
	/*Device
	  The name of the device where this software is installed

	*/
	Device string
	/*InstallDate
	  The date that the software was installed

	*/
	InstallDate *string
	/*LicenseKey
	  OS license key

	*/
	LicenseKey *string
	/*LicenseKeyCount
	  The number of licenses this software key supports

	*/
	LicenseKeyCount *string
	/*LicenseUseCount
	  the number of licenses that are in use for this software instance

	*/
	LicenseUseCount *string
	/*Software
	  the name of the software

	*/
	Software string
	/*User
	  The user assigned to this software

	*/
	User *string
	/*Vendor
	  The vendor that created the server, linked to Organization

	*/
	Vendor *string
	/*Version
	  The version number of the software

	*/
	Version *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post update service ports params
func (o *PostUpdateServicePortsParams) WithTimeout(timeout time.Duration) *PostUpdateServicePortsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post update service ports params
func (o *PostUpdateServicePortsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post update service ports params
func (o *PostUpdateServicePortsParams) WithContext(ctx context.Context) *PostUpdateServicePortsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post update service ports params
func (o *PostUpdateServicePortsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post update service ports params
func (o *PostUpdateServicePortsParams) WithHTTPClient(client *http.Client) *PostUpdateServicePortsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post update service ports params
func (o *PostUpdateServicePortsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCountInLicensing adds the countInLicensing to the post update service ports params
func (o *PostUpdateServicePortsParams) WithCountInLicensing(countInLicensing *string) *PostUpdateServicePortsParams {
	o.SetCountInLicensing(countInLicensing)
	return o
}

// SetCountInLicensing adds the countInLicensing to the post update service ports params
func (o *PostUpdateServicePortsParams) SetCountInLicensing(countInLicensing *string) {
	o.CountInLicensing = countInLicensing
}

// WithDevice adds the device to the post update service ports params
func (o *PostUpdateServicePortsParams) WithDevice(device string) *PostUpdateServicePortsParams {
	o.SetDevice(device)
	return o
}

// SetDevice adds the device to the post update service ports params
func (o *PostUpdateServicePortsParams) SetDevice(device string) {
	o.Device = device
}

// WithInstallDate adds the installDate to the post update service ports params
func (o *PostUpdateServicePortsParams) WithInstallDate(installDate *string) *PostUpdateServicePortsParams {
	o.SetInstallDate(installDate)
	return o
}

// SetInstallDate adds the installDate to the post update service ports params
func (o *PostUpdateServicePortsParams) SetInstallDate(installDate *string) {
	o.InstallDate = installDate
}

// WithLicenseKey adds the licenseKey to the post update service ports params
func (o *PostUpdateServicePortsParams) WithLicenseKey(licenseKey *string) *PostUpdateServicePortsParams {
	o.SetLicenseKey(licenseKey)
	return o
}

// SetLicenseKey adds the licenseKey to the post update service ports params
func (o *PostUpdateServicePortsParams) SetLicenseKey(licenseKey *string) {
	o.LicenseKey = licenseKey
}

// WithLicenseKeyCount adds the licenseKeyCount to the post update service ports params
func (o *PostUpdateServicePortsParams) WithLicenseKeyCount(licenseKeyCount *string) *PostUpdateServicePortsParams {
	o.SetLicenseKeyCount(licenseKeyCount)
	return o
}

// SetLicenseKeyCount adds the licenseKeyCount to the post update service ports params
func (o *PostUpdateServicePortsParams) SetLicenseKeyCount(licenseKeyCount *string) {
	o.LicenseKeyCount = licenseKeyCount
}

// WithLicenseUseCount adds the licenseUseCount to the post update service ports params
func (o *PostUpdateServicePortsParams) WithLicenseUseCount(licenseUseCount *string) *PostUpdateServicePortsParams {
	o.SetLicenseUseCount(licenseUseCount)
	return o
}

// SetLicenseUseCount adds the licenseUseCount to the post update service ports params
func (o *PostUpdateServicePortsParams) SetLicenseUseCount(licenseUseCount *string) {
	o.LicenseUseCount = licenseUseCount
}

// WithSoftware adds the software to the post update service ports params
func (o *PostUpdateServicePortsParams) WithSoftware(software string) *PostUpdateServicePortsParams {
	o.SetSoftware(software)
	return o
}

// SetSoftware adds the software to the post update service ports params
func (o *PostUpdateServicePortsParams) SetSoftware(software string) {
	o.Software = software
}

// WithUser adds the user to the post update service ports params
func (o *PostUpdateServicePortsParams) WithUser(user *string) *PostUpdateServicePortsParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the post update service ports params
func (o *PostUpdateServicePortsParams) SetUser(user *string) {
	o.User = user
}

// WithVendor adds the vendor to the post update service ports params
func (o *PostUpdateServicePortsParams) WithVendor(vendor *string) *PostUpdateServicePortsParams {
	o.SetVendor(vendor)
	return o
}

// SetVendor adds the vendor to the post update service ports params
func (o *PostUpdateServicePortsParams) SetVendor(vendor *string) {
	o.Vendor = vendor
}

// WithVersion adds the version to the post update service ports params
func (o *PostUpdateServicePortsParams) WithVersion(version *string) *PostUpdateServicePortsParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the post update service ports params
func (o *PostUpdateServicePortsParams) SetVersion(version *string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *PostUpdateServicePortsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CountInLicensing != nil {

		// form param count_in_licensing
		var frCountInLicensing string
		if o.CountInLicensing != nil {
			frCountInLicensing = *o.CountInLicensing
		}
		fCountInLicensing := frCountInLicensing
		if fCountInLicensing != "" {
			if err := r.SetFormParam("count_in_licensing", fCountInLicensing); err != nil {
				return err
			}
		}

	}

	// form param device
	frDevice := o.Device
	fDevice := frDevice
	if fDevice != "" {
		if err := r.SetFormParam("device", fDevice); err != nil {
			return err
		}
	}

	if o.InstallDate != nil {

		// form param install_date
		var frInstallDate string
		if o.InstallDate != nil {
			frInstallDate = *o.InstallDate
		}
		fInstallDate := frInstallDate
		if fInstallDate != "" {
			if err := r.SetFormParam("install_date", fInstallDate); err != nil {
				return err
			}
		}

	}

	if o.LicenseKey != nil {

		// form param license_key
		var frLicenseKey string
		if o.LicenseKey != nil {
			frLicenseKey = *o.LicenseKey
		}
		fLicenseKey := frLicenseKey
		if fLicenseKey != "" {
			if err := r.SetFormParam("license_key", fLicenseKey); err != nil {
				return err
			}
		}

	}

	if o.LicenseKeyCount != nil {

		// form param license_key_count
		var frLicenseKeyCount string
		if o.LicenseKeyCount != nil {
			frLicenseKeyCount = *o.LicenseKeyCount
		}
		fLicenseKeyCount := frLicenseKeyCount
		if fLicenseKeyCount != "" {
			if err := r.SetFormParam("license_key_count", fLicenseKeyCount); err != nil {
				return err
			}
		}

	}

	if o.LicenseUseCount != nil {

		// form param license_use_count
		var frLicenseUseCount string
		if o.LicenseUseCount != nil {
			frLicenseUseCount = *o.LicenseUseCount
		}
		fLicenseUseCount := frLicenseUseCount
		if fLicenseUseCount != "" {
			if err := r.SetFormParam("license_use_count", fLicenseUseCount); err != nil {
				return err
			}
		}

	}

	// form param software
	frSoftware := o.Software
	fSoftware := frSoftware
	if fSoftware != "" {
		if err := r.SetFormParam("software", fSoftware); err != nil {
			return err
		}
	}

	if o.User != nil {

		// form param user
		var frUser string
		if o.User != nil {
			frUser = *o.User
		}
		fUser := frUser
		if fUser != "" {
			if err := r.SetFormParam("user", fUser); err != nil {
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

	if o.Version != nil {

		// form param version
		var frVersion string
		if o.Version != nil {
			frVersion = *o.Version
		}
		fVersion := frVersion
		if fVersion != "" {
			if err := r.SetFormParam("version", fVersion); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}