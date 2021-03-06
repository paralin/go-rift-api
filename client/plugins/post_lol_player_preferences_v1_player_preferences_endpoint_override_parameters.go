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

// NewPostLolPlayerPreferencesV1PlayerPreferencesEndpointOverrideParams creates a new PostLolPlayerPreferencesV1PlayerPreferencesEndpointOverrideParams object
// with the default values initialized.
func NewPostLolPlayerPreferencesV1PlayerPreferencesEndpointOverrideParams() *PostLolPlayerPreferencesV1PlayerPreferencesEndpointOverrideParams {
	var ()
	return &PostLolPlayerPreferencesV1PlayerPreferencesEndpointOverrideParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolPlayerPreferencesV1PlayerPreferencesEndpointOverrideParamsWithTimeout creates a new PostLolPlayerPreferencesV1PlayerPreferencesEndpointOverrideParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolPlayerPreferencesV1PlayerPreferencesEndpointOverrideParamsWithTimeout(timeout time.Duration) *PostLolPlayerPreferencesV1PlayerPreferencesEndpointOverrideParams {
	var ()
	return &PostLolPlayerPreferencesV1PlayerPreferencesEndpointOverrideParams{

		timeout: timeout,
	}
}

// NewPostLolPlayerPreferencesV1PlayerPreferencesEndpointOverrideParamsWithContext creates a new PostLolPlayerPreferencesV1PlayerPreferencesEndpointOverrideParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolPlayerPreferencesV1PlayerPreferencesEndpointOverrideParamsWithContext(ctx context.Context) *PostLolPlayerPreferencesV1PlayerPreferencesEndpointOverrideParams {
	var ()
	return &PostLolPlayerPreferencesV1PlayerPreferencesEndpointOverrideParams{

		Context: ctx,
	}
}

// NewPostLolPlayerPreferencesV1PlayerPreferencesEndpointOverrideParamsWithHTTPClient creates a new PostLolPlayerPreferencesV1PlayerPreferencesEndpointOverrideParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolPlayerPreferencesV1PlayerPreferencesEndpointOverrideParamsWithHTTPClient(client *http.Client) *PostLolPlayerPreferencesV1PlayerPreferencesEndpointOverrideParams {
	var ()
	return &PostLolPlayerPreferencesV1PlayerPreferencesEndpointOverrideParams{
		HTTPClient: client,
	}
}

/*PostLolPlayerPreferencesV1PlayerPreferencesEndpointOverrideParams contains all the parameters to send to the API endpoint
for the post lol player preferences v1 player preferences endpoint override operation typically these are written to a http.Request
*/
type PostLolPlayerPreferencesV1PlayerPreferencesEndpointOverrideParams struct {

	/*Preferences*/
	Preferences *models.LolPlayerPreferencesPlayerPreferencesEndpoint

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol player preferences v1 player preferences endpoint override params
func (o *PostLolPlayerPreferencesV1PlayerPreferencesEndpointOverrideParams) WithTimeout(timeout time.Duration) *PostLolPlayerPreferencesV1PlayerPreferencesEndpointOverrideParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol player preferences v1 player preferences endpoint override params
func (o *PostLolPlayerPreferencesV1PlayerPreferencesEndpointOverrideParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol player preferences v1 player preferences endpoint override params
func (o *PostLolPlayerPreferencesV1PlayerPreferencesEndpointOverrideParams) WithContext(ctx context.Context) *PostLolPlayerPreferencesV1PlayerPreferencesEndpointOverrideParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol player preferences v1 player preferences endpoint override params
func (o *PostLolPlayerPreferencesV1PlayerPreferencesEndpointOverrideParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol player preferences v1 player preferences endpoint override params
func (o *PostLolPlayerPreferencesV1PlayerPreferencesEndpointOverrideParams) WithHTTPClient(client *http.Client) *PostLolPlayerPreferencesV1PlayerPreferencesEndpointOverrideParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol player preferences v1 player preferences endpoint override params
func (o *PostLolPlayerPreferencesV1PlayerPreferencesEndpointOverrideParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPreferences adds the preferences to the post lol player preferences v1 player preferences endpoint override params
func (o *PostLolPlayerPreferencesV1PlayerPreferencesEndpointOverrideParams) WithPreferences(preferences *models.LolPlayerPreferencesPlayerPreferencesEndpoint) *PostLolPlayerPreferencesV1PlayerPreferencesEndpointOverrideParams {
	o.SetPreferences(preferences)
	return o
}

// SetPreferences adds the preferences to the post lol player preferences v1 player preferences endpoint override params
func (o *PostLolPlayerPreferencesV1PlayerPreferencesEndpointOverrideParams) SetPreferences(preferences *models.LolPlayerPreferencesPlayerPreferencesEndpoint) {
	o.Preferences = preferences
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolPlayerPreferencesV1PlayerPreferencesEndpointOverrideParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
