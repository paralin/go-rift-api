// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/paralin/go-rift-api/models"
)

// NewPutLolPlayerPreferencesV1PreferenceParams creates a new PutLolPlayerPreferencesV1PreferenceParams object
// with the default values initialized.
func NewPutLolPlayerPreferencesV1PreferenceParams() *PutLolPlayerPreferencesV1PreferenceParams {
	var ()
	return &PutLolPlayerPreferencesV1PreferenceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutLolPlayerPreferencesV1PreferenceParamsWithTimeout creates a new PutLolPlayerPreferencesV1PreferenceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutLolPlayerPreferencesV1PreferenceParamsWithTimeout(timeout time.Duration) *PutLolPlayerPreferencesV1PreferenceParams {
	var ()
	return &PutLolPlayerPreferencesV1PreferenceParams{

		timeout: timeout,
	}
}

// NewPutLolPlayerPreferencesV1PreferenceParamsWithContext creates a new PutLolPlayerPreferencesV1PreferenceParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutLolPlayerPreferencesV1PreferenceParamsWithContext(ctx context.Context) *PutLolPlayerPreferencesV1PreferenceParams {
	var ()
	return &PutLolPlayerPreferencesV1PreferenceParams{

		Context: ctx,
	}
}

// NewPutLolPlayerPreferencesV1PreferenceParamsWithHTTPClient creates a new PutLolPlayerPreferencesV1PreferenceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutLolPlayerPreferencesV1PreferenceParamsWithHTTPClient(client *http.Client) *PutLolPlayerPreferencesV1PreferenceParams {
	var ()
	return &PutLolPlayerPreferencesV1PreferenceParams{
		HTTPClient: client,
	}
}

/*PutLolPlayerPreferencesV1PreferenceParams contains all the parameters to send to the API endpoint
for the put lol player preferences v1 preference operation typically these are written to a http.Request
*/
type PutLolPlayerPreferencesV1PreferenceParams struct {

	/*Preferences*/
	Preferences *models.LolPlayerPreferencesPlayerPreferences

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put lol player preferences v1 preference params
func (o *PutLolPlayerPreferencesV1PreferenceParams) WithTimeout(timeout time.Duration) *PutLolPlayerPreferencesV1PreferenceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put lol player preferences v1 preference params
func (o *PutLolPlayerPreferencesV1PreferenceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put lol player preferences v1 preference params
func (o *PutLolPlayerPreferencesV1PreferenceParams) WithContext(ctx context.Context) *PutLolPlayerPreferencesV1PreferenceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put lol player preferences v1 preference params
func (o *PutLolPlayerPreferencesV1PreferenceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put lol player preferences v1 preference params
func (o *PutLolPlayerPreferencesV1PreferenceParams) WithHTTPClient(client *http.Client) *PutLolPlayerPreferencesV1PreferenceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put lol player preferences v1 preference params
func (o *PutLolPlayerPreferencesV1PreferenceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPreferences adds the preferences to the put lol player preferences v1 preference params
func (o *PutLolPlayerPreferencesV1PreferenceParams) WithPreferences(preferences *models.LolPlayerPreferencesPlayerPreferences) *PutLolPlayerPreferencesV1PreferenceParams {
	o.SetPreferences(preferences)
	return o
}

// SetPreferences adds the preferences to the put lol player preferences v1 preference params
func (o *PutLolPlayerPreferencesV1PreferenceParams) SetPreferences(preferences *models.LolPlayerPreferencesPlayerPreferences) {
	o.Preferences = preferences
}

// WriteToRequest writes these params to a swagger request
func (o *PutLolPlayerPreferencesV1PreferenceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Preferences != nil {
		if err := r.SetBodyParam(o.Preferences); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
