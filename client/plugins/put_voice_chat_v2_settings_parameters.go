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

// NewPutVoiceChatV2SettingsParams creates a new PutVoiceChatV2SettingsParams object
// with the default values initialized.
func NewPutVoiceChatV2SettingsParams() *PutVoiceChatV2SettingsParams {
	var ()
	return &PutVoiceChatV2SettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutVoiceChatV2SettingsParamsWithTimeout creates a new PutVoiceChatV2SettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutVoiceChatV2SettingsParamsWithTimeout(timeout time.Duration) *PutVoiceChatV2SettingsParams {
	var ()
	return &PutVoiceChatV2SettingsParams{

		timeout: timeout,
	}
}

// NewPutVoiceChatV2SettingsParamsWithContext creates a new PutVoiceChatV2SettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutVoiceChatV2SettingsParamsWithContext(ctx context.Context) *PutVoiceChatV2SettingsParams {
	var ()
	return &PutVoiceChatV2SettingsParams{

		Context: ctx,
	}
}

// NewPutVoiceChatV2SettingsParamsWithHTTPClient creates a new PutVoiceChatV2SettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutVoiceChatV2SettingsParamsWithHTTPClient(client *http.Client) *PutVoiceChatV2SettingsParams {
	var ()
	return &PutVoiceChatV2SettingsParams{
		HTTPClient: client,
	}
}

/*PutVoiceChatV2SettingsParams contains all the parameters to send to the API endpoint
for the put voice chat v2 settings operation typically these are written to a http.Request
*/
type PutVoiceChatV2SettingsParams struct {

	/*Settings*/
	Settings *models.VoiceChatUpdateSettingsResource

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put voice chat v2 settings params
func (o *PutVoiceChatV2SettingsParams) WithTimeout(timeout time.Duration) *PutVoiceChatV2SettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put voice chat v2 settings params
func (o *PutVoiceChatV2SettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put voice chat v2 settings params
func (o *PutVoiceChatV2SettingsParams) WithContext(ctx context.Context) *PutVoiceChatV2SettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put voice chat v2 settings params
func (o *PutVoiceChatV2SettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put voice chat v2 settings params
func (o *PutVoiceChatV2SettingsParams) WithHTTPClient(client *http.Client) *PutVoiceChatV2SettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put voice chat v2 settings params
func (o *PutVoiceChatV2SettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSettings adds the settings to the put voice chat v2 settings params
func (o *PutVoiceChatV2SettingsParams) WithSettings(settings *models.VoiceChatUpdateSettingsResource) *PutVoiceChatV2SettingsParams {
	o.SetSettings(settings)
	return o
}

// SetSettings adds the settings to the put voice chat v2 settings params
func (o *PutVoiceChatV2SettingsParams) SetSettings(settings *models.VoiceChatUpdateSettingsResource) {
	o.Settings = settings
}

// WriteToRequest writes these params to a swagger request
func (o *PutVoiceChatV2SettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Settings != nil {
		if err := r.SetBodyParam(o.Settings); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
