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
)

// NewPatchLolGameSettingsV1InputSettingsParams creates a new PatchLolGameSettingsV1InputSettingsParams object
// with the default values initialized.
func NewPatchLolGameSettingsV1InputSettingsParams() *PatchLolGameSettingsV1InputSettingsParams {
	var ()
	return &PatchLolGameSettingsV1InputSettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchLolGameSettingsV1InputSettingsParamsWithTimeout creates a new PatchLolGameSettingsV1InputSettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchLolGameSettingsV1InputSettingsParamsWithTimeout(timeout time.Duration) *PatchLolGameSettingsV1InputSettingsParams {
	var ()
	return &PatchLolGameSettingsV1InputSettingsParams{

		timeout: timeout,
	}
}

// NewPatchLolGameSettingsV1InputSettingsParamsWithContext creates a new PatchLolGameSettingsV1InputSettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchLolGameSettingsV1InputSettingsParamsWithContext(ctx context.Context) *PatchLolGameSettingsV1InputSettingsParams {
	var ()
	return &PatchLolGameSettingsV1InputSettingsParams{

		Context: ctx,
	}
}

// NewPatchLolGameSettingsV1InputSettingsParamsWithHTTPClient creates a new PatchLolGameSettingsV1InputSettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchLolGameSettingsV1InputSettingsParamsWithHTTPClient(client *http.Client) *PatchLolGameSettingsV1InputSettingsParams {
	var ()
	return &PatchLolGameSettingsV1InputSettingsParams{
		HTTPClient: client,
	}
}

/*PatchLolGameSettingsV1InputSettingsParams contains all the parameters to send to the API endpoint
for the patch lol game settings v1 input settings operation typically these are written to a http.Request
*/
type PatchLolGameSettingsV1InputSettingsParams struct {

	/*SettingsResource*/
	SettingsResource interface{}

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch lol game settings v1 input settings params
func (o *PatchLolGameSettingsV1InputSettingsParams) WithTimeout(timeout time.Duration) *PatchLolGameSettingsV1InputSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch lol game settings v1 input settings params
func (o *PatchLolGameSettingsV1InputSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch lol game settings v1 input settings params
func (o *PatchLolGameSettingsV1InputSettingsParams) WithContext(ctx context.Context) *PatchLolGameSettingsV1InputSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch lol game settings v1 input settings params
func (o *PatchLolGameSettingsV1InputSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch lol game settings v1 input settings params
func (o *PatchLolGameSettingsV1InputSettingsParams) WithHTTPClient(client *http.Client) *PatchLolGameSettingsV1InputSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch lol game settings v1 input settings params
func (o *PatchLolGameSettingsV1InputSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSettingsResource adds the settingsResource to the patch lol game settings v1 input settings params
func (o *PatchLolGameSettingsV1InputSettingsParams) WithSettingsResource(settingsResource interface{}) *PatchLolGameSettingsV1InputSettingsParams {
	o.SetSettingsResource(settingsResource)
	return o
}

// SetSettingsResource adds the settingsResource to the patch lol game settings v1 input settings params
func (o *PatchLolGameSettingsV1InputSettingsParams) SetSettingsResource(settingsResource interface{}) {
	o.SettingsResource = settingsResource
}

// WriteToRequest writes these params to a swagger request
func (o *PatchLolGameSettingsV1InputSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.SettingsResource != nil {
		if err := r.SetBodyParam(o.SettingsResource); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
