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

// NewGetLolPlatformConfigV1InitialConfigurationCompleteParams creates a new GetLolPlatformConfigV1InitialConfigurationCompleteParams object
// with the default values initialized.
func NewGetLolPlatformConfigV1InitialConfigurationCompleteParams() *GetLolPlatformConfigV1InitialConfigurationCompleteParams {

	return &GetLolPlatformConfigV1InitialConfigurationCompleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolPlatformConfigV1InitialConfigurationCompleteParamsWithTimeout creates a new GetLolPlatformConfigV1InitialConfigurationCompleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolPlatformConfigV1InitialConfigurationCompleteParamsWithTimeout(timeout time.Duration) *GetLolPlatformConfigV1InitialConfigurationCompleteParams {

	return &GetLolPlatformConfigV1InitialConfigurationCompleteParams{

		timeout: timeout,
	}
}

// NewGetLolPlatformConfigV1InitialConfigurationCompleteParamsWithContext creates a new GetLolPlatformConfigV1InitialConfigurationCompleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolPlatformConfigV1InitialConfigurationCompleteParamsWithContext(ctx context.Context) *GetLolPlatformConfigV1InitialConfigurationCompleteParams {

	return &GetLolPlatformConfigV1InitialConfigurationCompleteParams{

		Context: ctx,
	}
}

// NewGetLolPlatformConfigV1InitialConfigurationCompleteParamsWithHTTPClient creates a new GetLolPlatformConfigV1InitialConfigurationCompleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolPlatformConfigV1InitialConfigurationCompleteParamsWithHTTPClient(client *http.Client) *GetLolPlatformConfigV1InitialConfigurationCompleteParams {

	return &GetLolPlatformConfigV1InitialConfigurationCompleteParams{
		HTTPClient: client,
	}
}

/*GetLolPlatformConfigV1InitialConfigurationCompleteParams contains all the parameters to send to the API endpoint
for the get lol platform config v1 initial configuration complete operation typically these are written to a http.Request
*/
type GetLolPlatformConfigV1InitialConfigurationCompleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol platform config v1 initial configuration complete params
func (o *GetLolPlatformConfigV1InitialConfigurationCompleteParams) WithTimeout(timeout time.Duration) *GetLolPlatformConfigV1InitialConfigurationCompleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol platform config v1 initial configuration complete params
func (o *GetLolPlatformConfigV1InitialConfigurationCompleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol platform config v1 initial configuration complete params
func (o *GetLolPlatformConfigV1InitialConfigurationCompleteParams) WithContext(ctx context.Context) *GetLolPlatformConfigV1InitialConfigurationCompleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol platform config v1 initial configuration complete params
func (o *GetLolPlatformConfigV1InitialConfigurationCompleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol platform config v1 initial configuration complete params
func (o *GetLolPlatformConfigV1InitialConfigurationCompleteParams) WithHTTPClient(client *http.Client) *GetLolPlatformConfigV1InitialConfigurationCompleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol platform config v1 initial configuration complete params
func (o *GetLolPlatformConfigV1InitialConfigurationCompleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolPlatformConfigV1InitialConfigurationCompleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
